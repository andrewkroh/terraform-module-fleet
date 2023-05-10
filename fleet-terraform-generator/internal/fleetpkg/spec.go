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

package fleetpkg

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v3"
)

type FileMetadata struct {
	File   string `json:"-"` // File from which field was read.
	Line   int    `json:"-"` // Line from which field was read.
	Column int    `json:"-"` // Column from which field was read.
}

type Integration struct {
	Manifest    Manifest               `json:"manifest"`
	DataStreams map[string]*DataStream `json:"data_streams"`
}

type DataStream struct {
	Manifest DataStreamManifest `json:"manifest"`
}

type Manifest struct {
	Name            string           `json:"name" yaml:"name"`
	Title           string           `json:"title" yaml:"title"`
	Version         string           `json:"version" yaml:"version"`
	Release         string           `json:"release" yaml:"release"`
	Description     string           `json:"description" yaml:"description"`
	Type            string           `json:"type" yaml:"type"`
	Icons           []Icons          `json:"icons" yaml:"icons,omitempty"`
	FormatVersion   string           `json:"format_version" yaml:"format_version"`
	License         string           `json:"license" yaml:"license,omitempty"`
	Categories      []string         `json:"categories" yaml:"categories,omitempty"`
	Conditions      Conditions       `json:"conditions" yaml:"conditions"`
	Screenshots     []Screenshots    `json:"screenshots" yaml:"screenshots,omitempty"`
	Source          Source           `json:"source"`
	Vars            []Var            `json:"vars" yaml:"vars,omitempty"`
	PolicyTemplates []PolicyTemplate `json:"policy_templates" yaml:"policy_templates,omitempty"`
	Owner           Owner            `json:"owner" yaml:"owner"`
}

type Source struct {
	License string `json:"license,omitempty"`
}

type Icons struct {
	Src      string `json:"src"`
	Title    string `json:"title"`
	Size     string `json:"size"`
	Type     string `json:"type"`
	DarkMode bool   `json:"dark_mode" yaml:"dark_mode"`
}

type Conditions struct {
	KibanaVersion       string `json:"kibana.version" yaml:"kibana.version"`
	ElasticSubscription string `json:"elastic.subscription" yaml:"elastic.subscription"`
	Elastic             struct {
		Subscription string `json:"subscription"`
	} `json:"elastic"`
	Kibana struct {
		Version string `json:"version"`
	} `json:"kibana"`
}

type Screenshots struct {
	Src   string `json:"src"`
	Title string `json:"title"`
	Size  string `json:"size"`
	Type  string `json:"type"`
}

type Var struct {
	Name        string   `json:"name"`
	Default     any      `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type"`
	Title       string   `json:"title"`
	Multi       bool     `json:"multi"`
	Required    bool     `json:"required"`
	ShowUser    bool     `json:"show_user" yaml:"show_user"`
	Options     []Option `json:"options"` // List of options for 'type: select'.

	Meta FileMetadata `json:"-"`
}

func (f *Var) UnmarshalYAML(value *yaml.Node) error {
	// Prevent recursion by creating a new type that does not implement Unmarshaler.
	type notVar Var
	x := (*notVar)(f)

	if err := value.Decode(&x); err != nil {
		return err
	}
	f.Meta.Line = value.Line
	f.Meta.Column = value.Column
	return nil
}

type Option struct {
	Value string `json:"value"`
	Text  string `json:"text"`
}

type Input struct {
	Type         string `json:"type"`
	Title        string `json:"title"`
	Description  string `json:"description" yaml:"description,omitempty"`
	InputGroup   string `json:"input_group" yaml:"input_group,omitempty"`
	TemplatePath string `json:"template_path" yaml:"template_path"`
	Multi        bool   `json:"multi"`
	Vars         []Var  `json:"vars"`
}

type PolicyTemplate struct {
	Name         string        `json:"name"`
	Title        string        `json:"title"`
	Categories   []string      `json:"categories" yaml:"categories,omitempty"`
	Description  string        `json:"description"`
	DataStreams  []string      `json:"data_streams" yaml:"data_streams,omitempty"`
	Inputs       []Input       `json:"inputs" yaml:"inputs,omitempty"`
	Icons        []Icons       `json:"icons" yaml:"icons,omitempty"`
	Screenshots  []Screenshots `json:"screenshots" yaml:"screenshots,omitempty"`
	Multiple     bool          `json:"multiple" yaml:"multiple"`
	Type         string        `json:"type" yaml:"type"` // Type of data stream.
	Input        string        `json:"input" yaml:"input"`
	TemplatePath string        `json:"template_path" yaml:"template_path"`
	Vars         []Var         `json:"vars"` // Policy template level variables.
}

type Owner struct {
	Github string `json:"github"`
}

type DataStreamManifest struct {
	Dataset         string
	DatasetIsPrefix bool           `json:"dataset_is_prefix" yaml:"dataset_is_prefix"`
	ILMPolicy       string         `json:"ilm_policy" yaml:"ilm_policy"`
	Release         string         `json:"release"`
	Title           string         `json:"title"`
	Type            string         `json:"type"`
	Streams         []Stream       `json:"streams"`
	Elasticsearch   map[string]any `json:"elasticsearch"`
}

type Stream struct {
	Input        string `json:"input"`
	Description  string `json:"description"`
	Title        string `json:"title"`
	TemplatePath string `json:"template_path" yaml:"template_path"`
	Vars         []Var  `json:"vars"`
	Enabled      bool   `json:"enabled"`
}

func Load(path string) (*Integration, error) {
	var manifest Manifest
	if err := readYAML(filepath.Join(path, "manifest.yml"), &manifest, true); err != nil {
		return nil, err
	}

	integration := &Integration{
		Manifest:    manifest,
		DataStreams: map[string]*DataStream{},
	}

	dataStreams, err := filepath.Glob(filepath.Join(path, "data_stream/*/manifest.yml"))
	if err != nil {
		return nil, err
	}
	for _, f := range dataStreams {
		var dataStreamManifest DataStreamManifest
		if err := readYAML(f, &dataStreamManifest, true); err != nil {
			return nil, err
		}
		integration.DataStreams[filepath.Base(filepath.Dir(f))] = &DataStream{
			Manifest: dataStreamManifest,
		}
	}

	return integration, nil
}

func readYAML(path string, v any, strict bool) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	dec.KnownFields(strict)

	if err := dec.Decode(v); err != nil {
		return fmt.Errorf("failed reading %s: %w", path, err)
	}
	annotateFileMetadata(path, v)
	return nil
}

// annotateFileMetadata sets the file name on any types that contains FileMetadata.
func annotateFileMetadata(file string, v any) {
	fileAnnotator{Name: file}.Annotate(reflect.ValueOf(v))
}

type fileAnnotator struct {
	Name string
}

func (a fileAnnotator) Annotate(val reflect.Value) {
	// Need an addressable value in order to edit the metadata value.
	if val.CanAddr() {
		if m, ok := val.Addr().Interface().(*FileMetadata); ok {
			m.File = a.Name
			return
		}
	}

	switch val.Kind() {
	case reflect.Pointer:
		a.Annotate(val.Elem())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			valueField := val.Field(i)
			a.Annotate(valueField)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			a.Annotate(val.Index(i))
		}
	case reflect.Map:
		itr := val.MapRange()
		for itr.Next() {
			// NOTE: This can only edit the map value if it is addressable (aka a pointer).
			a.Annotate(itr.Value())
		}
	}
}
