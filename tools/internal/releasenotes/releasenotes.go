// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package releasenotes

import (
	"bytes"
	"cmp"
	"embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/elastic/terraform-module-fleet/tools/internal/terraform"
)

var (
	ErrNoChanges                  = errors.New("no changes detected")
	errIntegrationsCommitNotFound = errors.New("elastic/integrations commit not found in git log")
)

var flagset = flag.NewFlagSet("", flag.ContinueOnError)

// Parameters
var (
	previousCommit string
	releaseCommit  string
)

func init() {
	flagset.StringVar(&previousCommit, "p", "", "Previous release commit. Defaults to most recent tag.")
	flagset.StringVar(&releaseCommit, "r", "", "This release commit. Defaults to latest commit on current branch.")
}

func Main() error {
	if err := flagset.Parse(os.Args[1:]); err != nil {
		return err
	}

	gitDir, err := detectDotGit()
	if err != nil {
		return err
	}

	fs := memfs.New()
	memStore := memory.NewStorage()
	r, err := git.Clone(memStore, fs, &git.CloneOptions{
		URL: gitDir,
	})
	if err != nil {
		return err
	}

	if previousCommit == "" {
		tag, err := latestGitTag(r)
		if err != nil {
			return err
		}
		previousCommit = tag.Name().String()
	}
	log.Println("Previous release:", previousCommit)

	if releaseCommit == "" {
		current, err := r.Head()
		if err != nil {
			return err
		}
		releaseCommit = current.Hash().String()
	}
	log.Println("Latest commit:", releaseCommit)

	before, err := readMetadata(r, plumbing.Revision(previousCommit))
	if err != nil {
		return fmt.Errorf("failed to read module metadata from previous release tag %q: %w", previousCommit, err)
	}
	after, err := readMetadata(r, plumbing.Revision(releaseCommit))
	if err != nil {
		return fmt.Errorf("failed to read metadata from HEAD: %w", err)
	}
	integrationsCommit, err := detectLatestElasticIntegrations(r)
	if err != nil {
		return fmt.Errorf("failed to determine last sync with elastic/integrations: %w", err)
	}

	data := diffMetadata(before, after)
	data.ElasticIntegrationsCommit = integrationsCommit

	if len(data.Inputs) == 0 && len(data.Integrations) == 0 {
		return ErrNoChanges
	}

	tmpl := htmlTmpl.Lookup("release-notes.md.gotmpl")
	notes, err := renderTemplate(tmpl, data)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", notes)
	return nil
}

// detectLatestElasticIntegrations iterates through git history to find the
// commit ID of the elastic/integrations project that was used to generate
// modules. The commit message will look like this
//
//	Update to elastic/integrations@6805e4daec4bab0c3bff57d1464ba35695cfa450
func detectLatestElasticIntegrations(r *git.Repository) (string, error) {
	logIter, err := r.Log(&git.LogOptions{})
	if err != nil {
		return "", err
	}

	var commit string
	err = logIter.ForEach(func(c *object.Commit) error {
		commit = parseElasticIntegrationsCommit(c.Message)
		if commit != "" {
			return storer.ErrStop
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if commit == "" {
		return "", errIntegrationsCommitNotFound
	}

	return commit, nil
}

var updateCommitRegex = regexp.MustCompile(`(?m)Update to elastic/integrations@([a-f0-9]+)`)

func parseElasticIntegrationsCommit(message string) (commit string) {
	matches := updateCommitRegex.FindStringSubmatch(message)
	if len(matches) == 2 {
		return matches[1]
	}
	return ""
}

type Metadata struct {
	Inputs       []terraform.File
	Integrations []terraform.File
}

func readMetadata(r *git.Repository, from plumbing.Revision) (*Metadata, error) {
	w, err := r.Worktree()
	if err != nil {
		return nil, err
	}

	ref, err := r.ResolveRevision(from)
	if err != nil {
		return nil, err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Hash: *ref,
	})
	if err != nil {
		return nil, err
	}

	tfModules, err := util.Glob(w.Filesystem, "fleet_in*/*/module.tf.json")
	if err != nil {
		return nil, err
	}

	var m Metadata
	for _, path := range tfModules {
		f, err := w.Filesystem.Open(path)
		if err != nil {
			return nil, err
		}

		dec := json.NewDecoder(f)
		dec.DisallowUnknownFields()
		dec.UseNumber()

		var tfFile terraform.File
		if err = dec.Decode(&tfFile); err != nil {
			f.Close()
			return nil, err
		}

		switch {
		case strings.HasPrefix(path, "fleet_input"):
			m.Inputs = append(m.Inputs, tfFile)
		case strings.HasPrefix(path, "fleet_integration"):
			m.Integrations = append(m.Integrations, tfFile)
		default:
			f.Close()
			return nil, fmt.Errorf("unexpected path prefix in %q", path)
		}

		f.Close()
	}

	return &m, nil
}

//go:embed templates/*.gotmpl
var templates embed.FS

var templateFuncs = template.FuncMap{}

var htmlTmpl = template.Must(template.New("github").
	Option("missingkey=error").
	Funcs(templateFuncs).
	ParseFS(templates, "*/*.gotmpl"))

func renderTemplate(tmpl *template.Template, data *TemplateData) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.Bytes(), nil
}

type TemplateData struct {
	Inputs                    []ModuleChange
	Integrations              []ModuleChange
	ElasticIntegrationsCommit string
}

type ChangeType string

var (
	Update ChangeType = "update"
	Add    ChangeType = "add"
	Remove ChangeType = "remove"
)

type ModuleChange struct {
	Change     ChangeType
	Package    string
	Version    string
	OldVersion string
}

func diffMetadata(before, after *Metadata) *TemplateData {
	var data TemplateData
	data.Inputs = diffTerraformFiles(before.Inputs, after.Inputs)
	data.Integrations = diffTerraformFiles(before.Integrations, after.Integrations)
	return &data
}

func diffTerraformFiles(before, after []terraform.File) []ModuleChange {
	oldPackages := map[string]string{}

	for _, f := range before {
		m := f.Modules["fleet_package_policy"]
		name := m.Params["package_name"].(string)

		v, ok := f.Variables["fleet_package_version"]
		if ok && v.Default != nil {
			oldPackages[name] = v.Default.Value.(string)
		}
	}

	newPackages := map[string]string{}

	for _, f := range after {
		m := f.Modules["fleet_package_policy"]
		name := m.Params["package_name"].(string)

		v, ok := f.Variables["fleet_package_version"]
		if ok && v.Default != nil {
			newPackages[name] = v.Default.Value.(string)
		}
	}

	var changes []ModuleChange
	for pkg, newVer := range newPackages {
		if oldVer, found := oldPackages[pkg]; found {
			if oldVer != newVer {
				changes = append(changes, ModuleChange{
					Change:     Update,
					Package:    pkg,
					Version:    newVer,
					OldVersion: oldVer,
				})
			}
		} else {
			changes = append(changes, ModuleChange{
				Change:  Add,
				Package: pkg,
				Version: newVer,
			})
		}
		delete(oldPackages, pkg)
	}
	for pkg := range oldPackages {
		changes = append(changes, ModuleChange{
			Change:  Remove,
			Package: pkg,
		})
	}

	slices.SortFunc(changes, func(a, b ModuleChange) int {
		return cmp.Compare(a.Package, b.Package)
	})

	return changes
}

func detectDotGit() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		gitDir := filepath.Join(path, ".git")
		if _, err := os.Stat(gitDir); err == nil {
			return gitDir, nil
		}

		nextPath := filepath.Dir(path)
		if nextPath == path || nextPath == "" {
			break
		}
		path = nextPath
	}

	return "", errors.New(".git directory not found")
}

func latestGitTag(r *git.Repository) (*plumbing.Reference, error) {
	tags, err := r.Tags()
	if err != nil {
		return nil, err
	}

	var latestTag *plumbing.Reference
	var latestTagTime time.Time
	err = tags.ForEach(func(ref *plumbing.Reference) error {
		commit, err := r.CommitObject(ref.Hash())
		if err != nil {
			return err
		}

		// Is this commit later?
		if commit.Author.When.After(latestTagTime) {
			latestTag = ref
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if latestTag == nil {
		return nil, errors.New("no git tags found")
	}
	return latestTag, nil
}
