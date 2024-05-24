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

package module

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"github.com/andrewkroh/go-fleetpkg"

	"golang.org/x/exp/maps"
)

type Specifier struct {
	Integration    string
	PolicyTemplate string
	DataStream     string
	Input          string

	Path string
}

func (s Specifier) String() string {
	parts := []string{
		s.Integration,
		s.PolicyTemplate,
	}
	// Input type packages don't have data streams.
	if s.DataStream != "" {
		parts = append(parts, s.DataStream)
	}
	parts = append(parts, strings.ReplaceAll(s.Input, "/", "_"))
	return strings.Join(parts, "/")
}

type Specifiers []Specifier

func (l Specifiers) Filter(globs ...string) (Specifiers, error) {
	if len(globs) == 0 {
		return l, nil
	}

	set := map[Specifier]struct{}{}
	for _, glob := range globs {
		foundMatch := false
		for _, specifier := range l {
			match, err := filepath.Match(glob, specifier.String())
			if err != nil {
				return nil, err
			}
			if !match {
				continue
			}
			set[specifier] = struct{}{}
			foundMatch = true
		}
		if !foundMatch {
			return nil, fmt.Errorf("package specifier %q does not match any known package", glob)
		}
	}

	out := maps.Keys(set)
	sortSpecifiers(out)
	return out, nil
}

func List(dir string, continueOnError bool) (Specifiers, error) {
	var result []Specifier

	// Generate the product of policy_template x data_stream x input.
	err := walkPackages(dir, func(pkg *fleetpkg.Integration, err error) error {
		if err != nil {
			if continueOnError {
				log.Println("[WARN] Ignoring package:", err)
				return nil
			}
			return err
		}

		if pkg.Manifest.Type == "input" {
			for _, pt := range pkg.Manifest.PolicyTemplates {
				result = append(result, Specifier{
					Integration:    pkg.Manifest.Name,
					PolicyTemplate: pt.Name,
					Input:          pt.Input,
					Path:           pkg.Path(),
				})
			}
			return nil
		}

		for _, pt := range pkg.Manifest.PolicyTemplates {
			// Some policy templates do not list the associated data streams so I assume
			// this means consider all data streams in the package.
			if len(pt.DataStreams) == 0 {
				pt.DataStreams = maps.Keys(pkg.DataStreams)
			}

			for _, dsName := range pt.DataStreams {
				ds, found := pkg.DataStreams[dsName]
				if !found {
					log.Printf("WARN: In %s/%s data stream %s was not found.", pkg.Manifest.Name, pt.Name, dsName)
					continue
				}

				for _, ptInput := range pt.Inputs {
					// Does the data stream contain the input type defined in the policy template?
					for _, stream := range ds.Manifest.Streams {
						if stream.Input == ptInput.Type {
							result = append(result, Specifier{
								Integration:    pkg.Manifest.Name,
								PolicyTemplate: pt.Name,
								DataStream:     dsName,
								Input:          ptInput.Type,
								Path:           pkg.Path(),
							})
							break
						}
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	sortSpecifiers(result)
	return result, nil
}

func walkPackages(dir string, walk func(pkg *fleetpkg.Integration, err error) error) error {
	allPackages, err := filepath.Glob(filepath.Join(dir, "*/manifest.yml"))
	if err != nil {
		return err
	}

	for _, manifestPath := range allPackages {
		integration, err := fleetpkg.Read(filepath.Dir(manifestPath))
		if err = walk(integration, err); err != nil {
			return err
		}
	}

	return nil
}

func sortSpecifiers(l []Specifier) {
	sort.Slice(l, func(i, j int) bool {
		if l[i].Integration != l[j].Integration {
			return l[i].Integration < l[j].Integration
		}
		if l[i].PolicyTemplate != l[j].PolicyTemplate {
			return l[i].PolicyTemplate < l[j].PolicyTemplate
		}
		if l[i].DataStream != l[j].DataStream {
			return l[i].DataStream < l[j].DataStream
		}
		return l[i].Input < l[j].Input
	})
}
