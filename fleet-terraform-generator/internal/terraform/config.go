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

package terraform

import "encoding/json"

// https://developer.hashicorp.com/terraform/language/syntax/json

type File struct {
	Comment       string                  `json:"//,omitempty"`       // File level comment.
	Variables     map[string]Variable     `json:"variable,omitempty"` // Name to variable definition.
	ResourceTypes map[string]ResourceType `json:"resource,omitempty"` // Resource type to name.
	Outputs       map[string]Output       `json:"output,omitempty"`   // Name to output variable definition.
	Modules       map[string]Module       `json:"module,omitempty"`   // Name to module definition.
}

type Variable struct {
	Type        string         `json:"type,omitempty"`
	Description string         `json:"description,omitempty"`
	Default     *NullableValue `json:"default,omitempty"`
	Sensitive   *bool          `json:"sensitive,omitempty"`
	Nullable    *bool          `json:"nullable,omitempty"`
	Validation  *Validation    `json:"validation,omitempty"`
}

// NullableValue is type that can be used in conjunction with encoding/json omitempty
// to obtain the ability to represent a field as non-present, null, or non-null.
type NullableValue struct {
	Value any
}

func (v *NullableValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

// Need a type that can render as JSON null,

type Validation struct {
	Condition    string `json:"condition"`
	ErrorMessage string `json:"error_message"`
}

type ResourceType map[string]Resource // Name to resource definition.

type Resource map[string]any

type Output struct {
	Description string `json:"description,omitempty"`
	Sensitive   *bool  `json:"sensitive,omitempty"`
	Value       string `json:"value"`
}

type Module struct {
	Source string `json:"source"`
	Params map[string]any
}

func (m Module) MarshalJSON() ([]byte, error) {
	// hack for https://github.com/golang/go/issues/6213
	clone := make(map[string]any, len(m.Params)+1)
	for k, v := range m.Params {
		clone[k] = v
	}
	clone["source"] = m.Source
	return json.Marshal(clone)
}

type ModuleParams map[string]any

func Ptr[T any](v T) *T {
	return &v
}
