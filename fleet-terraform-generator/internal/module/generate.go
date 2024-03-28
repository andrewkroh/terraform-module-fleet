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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/andrewkroh/go-fleetpkg"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/elastic/terraform-module-fleet/fleet-terraform-generator/internal/terraform"
)

type Terraform struct {
	Name        string
	PackageType string // Package type (input or integration).
	File        *terraform.File
}

// Generate generates a Terraform module.
func Generate(path, policyTemplateName, dataStreamName, inputName string, ignoreVariableShadowing bool) (*Terraform, error) {
	// Read in the package metadata.
	pkg, err := fleetpkg.Read(path)
	if err != nil {
		return nil, err
	}

	var (
		manifest            = pkg.Manifest
		policyTemplate      *fleetpkg.PolicyTemplate
		policyTemplateInput *fleetpkg.Input
		dataStream          *fleetpkg.DataStream
		stream              *fleetpkg.Stream
	)

	// Policy template.

	if policyTemplateName == "" {
		policyTemplate = &pkg.Manifest.PolicyTemplates[0]
		policyTemplateName = policyTemplate.Name
	} else {
		for i, pt := range pkg.Manifest.PolicyTemplates {
			if pt.Name == policyTemplateName {
				policyTemplate = &pkg.Manifest.PolicyTemplates[i]
				break
			}
		}
		if policyTemplate == nil {
			return nil, fmt.Errorf("policy template %q not found", policyTemplateName)
		}
	}

	if pkg.Manifest.Type != "input" {
		for i, input := range policyTemplate.Inputs {
			if input.Type == inputName {
				policyTemplateInput = &policyTemplate.Inputs[i]
				break
			}
		}
		if policyTemplateInput == nil {
			return nil, fmt.Errorf("input %q was not found within policy template %q", inputName, policyTemplateName)
		}

		// Data stream.
		{
			ds, found := pkg.DataStreams[dataStreamName]
			if !found {
				return nil, fmt.Errorf("data stream %q was not found in the package", dataStreamName)
			}
			dataStream = ds
		}
		// Input type.
		{
			for i, s := range dataStream.Manifest.Streams {
				if s.Input == inputName {
					stream = &dataStream.Manifest.Streams[i]
					break
				}
			}
			if stream == nil {
				return nil, fmt.Errorf("input type %q was not found in data stream %q", inputName, dataStreamName)
			}
		}
	} else {
		// Inject data_stream.dataset for input packages.
		injectDataStreamDatasetVar(policyTemplate)
	}

	tfVariables := map[string]moduleVariable{
		"fleet_agent_policy_id": {
			Terraform: terraform.Variable{
				Type:        "string",
				Description: "Agent policy ID to add the package policy to.",
			},
		},
		"fleet_data_stream_namespace": {
			Terraform: terraform.Variable{
				Type:        "string",
				Description: "Namespace to use for the data stream.",
				Default:     &terraform.NullableValue{Value: "default"},
			},
		},
		"fleet_package_policy_description": {
			Terraform: terraform.Variable{
				Type:        "string",
				Description: "Description to use for the package policy.",
				Default:     &terraform.NullableValue{Value: ""},
			},
		},
		"fleet_package_policy_name_suffix": {
			Terraform: terraform.Variable{
				Type:        "string",
				Description: "Suffix to append to the end of the package policy name.",
				Default:     &terraform.NullableValue{Value: ""},
			},
		},
		"fleet_package_version": {
			Terraform: terraform.Variable{
				Type:        "string",
				Description: "Version of the " + pkg.Manifest.Name + " package to use.",
				Default:     &terraform.NullableValue{Value: pkg.Manifest.Version},
			},
		},
	}

	// Iterate over all variables in the package and create Terraform variables.
	packageLevelVarAssociations, err := addVariables(pkg.Manifest.Vars, tfVariables, ignoreVariableShadowing)
	if err != nil {
		return nil, fmt.Errorf("error adding package level variables: %w", err)
	}
	policyTemplateLevelVarAssociations, err := addVariables(policyTemplate.Vars, tfVariables, ignoreVariableShadowing)
	if err != nil {
		return nil, fmt.Errorf("error adding policy template level variables: %w", err)
	}
	var inputLevelVarAssociations map[string]string
	if policyTemplateInput != nil {
		inputLevelVarAssociations, err = addVariables(policyTemplateInput.Vars, tfVariables, ignoreVariableShadowing)
		if err != nil {
			return nil, fmt.Errorf("error adding input level variables: %w", err)
		}
	}
	var dataStreamVarAssociations map[string]string
	if stream != nil {
		dataStreamVarAssociations, err = addVariables(stream.Vars, tfVariables, ignoreVariableShadowing)
		if err != nil {
			return nil, fmt.Errorf("error adding data stream level variables: %w", err)
		}
	}

	packageLevelVarExpression, err := buildVariableExpression(packageLevelVarAssociations)
	if err != nil {
		return nil, err
	}
	inputLevelVarExpression, err := buildVariableExpression(inputLevelVarAssociations)
	if err != nil {
		return nil, err
	}
	// Empirically, it appears that input package policy template variables are treated
	// the same as data stream variables.
	dataStreamVarExpression, err := buildVariableExpression(dataStreamVarAssociations, policyTemplateLevelVarAssociations)
	if err != nil {
		return nil, err
	}

	dataStreamsForInput := []string{} // Declare empty slice.
	{
		// Get a list of data streams so that we can disable all the ones not being
		// used. This avoids validation errors for required variables.
		dataStreams := maps.Clone(pkg.DataStreams)

		// If the policy template declares specific data streams, then honor that list.
		if len(policyTemplate.DataStreams) > 0 {
			maps.DeleteFunc(dataStreams, func(s string, stream *fleetpkg.DataStream) bool {
				return !slices.Contains(policyTemplate.DataStreams, s)
			})
		}
		// Filter data streams to only include ones with the selected input type.
		for directoryName, ds := range dataStreams {
			for _, stream := range ds.Manifest.Streams {
				if stream.Input == inputName {
					dataStreamsForInput = append(dataStreamsForInput, datasetName(directoryName, ds))
				}
			}
		}
		sort.Strings(dataStreamsForInput)
	}

	// All "${policy_template.name}-${input.type}" combinations.
	allPolicyTemplateInputs := []string{} // Declare empty slice.
	{
		for _, p := range manifest.PolicyTemplates {
			for _, input := range p.Inputs {
				allPolicyTemplateInputs = append(allPolicyTemplateInputs, p.Name+"-"+input.Type)
			}
		}
	}

	packagePolicyName := manifest.Name + "-" + dataStreamName + "-${var.fleet_data_stream_namespace}${var.fleet_package_policy_name_suffix}"
	if dataStreamName == "" {
		packagePolicyName = manifest.Name + "-${var.fleet_data_stream_namespace}${var.fleet_package_policy_name_suffix}"
	}

	tf := &terraform.File{
		Comment:   "Generated by fleet-terraform-generator - DO NOT EDIT",
		Variables: moduleVariableToTerraformVariable(tfVariables),
		Modules: map[string]terraform.Module{
			"fleet_package_policy": {
				Source: "../../fleet_package_policy",
				Params: toMap(FleetPackagePolicyModule{
					AgentPolicyID:           "${var.fleet_agent_policy_id}",
					PackagePolicyName:       packagePolicyName,
					PackageName:             manifest.Name,
					PackageVersion:          "${var.fleet_package_version}",
					Namespace:               "${var.fleet_data_stream_namespace}",
					Description:             "${var.fleet_package_policy_description}",
					PolicyTemplate:          policyTemplate.Name,
					DataStream:              datasetName(dataStreamName, dataStream),
					InputType:               inputName,
					PackageVariablesJSON:    packageLevelVarExpression,
					InputVariablesJSON:      inputLevelVarExpression,
					DataStreamVariablesJSON: dataStreamVarExpression,
					AllDataStreams:          dataStreamsForInput,
					AllPolicyTemplateInputs: allPolicyTemplateInputs,
				}),
			},
		},
		Outputs: map[string]terraform.Output{
			"id": {
				Description: "Package policy ID",
				Value:       "${module.fleet_package_policy.id}",
			},
		},
	}

	pkgType := manifest.Type
	if pkgType == "" {
		pkgType = "integration"
	}

	tfModuleName := manifest.Name
	if pkgType == "integration" {
		tfModuleName = moduleName(manifest.Name, policyTemplateName, dataStreamName, inputName)
	}
	return &Terraform{
		Name:        tfModuleName,
		PackageType: pkgType,
		File:        tf,
	}, nil
}

type moduleVariable struct {
	Terraform terraform.Variable // Terraform variable definition.
	Fleet     *fleetpkg.Var      // Source of the Terraform variable in Fleet package (optional).
}

// addVariables adds the given Fleet package vars to the Terraform module variables in m.
// It returns an error if a variable with the given name already exists in the Terraform
// module. It returns an associations mapping that contains a mapping of Fleet variable
// names to Terraform variable names.
func addVariables(vars []fleetpkg.Var, m map[string]moduleVariable, ignoreShadowing bool) (associations map[string]string, err error) {
	associations = make(map[string]string, len(vars))
	for _, v := range vars {
		tfName, err := addVariable(v, m, ignoreShadowing)
		if err != nil {
			return nil, fmt.Errorf("error adding variable %q: %w", v.Name, err)
		}
		associations[v.Name] = tfName
	}

	return associations, nil
}

func addVariable(v fleetpkg.Var, m map[string]moduleVariable, ignoreShadowing bool) (tfName string, err error) {
	tfVar := terraform.Variable{
		Description: v.Description,
	}

	if tfVar.Type, err = dataType(v); err != nil {
		return "", err
	}
	if tfVar.Type == "string" && isSensitive(v) {
		tfVar.Sensitive = terraform.Ptr(true)
	}
	if v.Required != nil && *v.Required {
		tfVar.Nullable = terraform.Ptr(false)
	}
	if v.Default != nil {
		// Pass the default to Terraform.
		tfVar.Default = &terraform.NullableValue{Value: v.Default}
	} else if v.Required == nil || v.Required != nil && !*v.Required {
		tfVar.Default = &terraform.NullableValue{}
	}

	// Append yaml suffix to indicate to users that they must yamlencode() the value.
	name := strings.ReplaceAll(v.Name, ".", "_")
	if v.Type == "yaml" {
		name += "_yaml"
	}

	// 'providers' is a reserved word in Terraform.
	if name == "providers" {
		name = "providers_names"
	}

	// Don't allow variables shadowing. See https://github.com/elastic/package-spec/issues/421
	if !ignoreShadowing {
		if existing, found := m[name]; found {
			if existing.Fleet != nil {
				msg := fmt.Sprintf("duplicate variable %q found at both\n\t%s:%d\n\t%s:%d", name,
					existing.Fleet.Path(), existing.Fleet.Line(),
					v.Path(), v.Line())
				if diff := cmp.Diff(*existing.Fleet, v, cmpopts.IgnoreFields(fleetpkg.Var{}, "FileMetadata")); diff != "" {
					return "", fmt.Errorf(msg+"\ndiff:\n%s", diff)
				}
				return "", errors.New(msg)
			}
			return "", fmt.Errorf("duplicate variable named %q found at %s:%d", name, v.Path(), v.Line())
		}
	}
	m[name] = moduleVariable{Fleet: &v, Terraform: tfVar}
	return name, nil
}

func moduleVariableToTerraformVariable(in map[string]moduleVariable) map[string]terraform.Variable {
	out := make(map[string]terraform.Variable, len(in))
	for k, mv := range in {
		out[k] = mv.Terraform
	}
	return out
}

func isSensitive(v fleetpkg.Var) bool {
	name := strings.ToLower(v.Name)
	switch {
	case v.Type == "password",
		strings.Contains(name, "token") && !strings.Contains(name, "file"),
		strings.Contains(name, "api_key"),
		strings.Contains(name, "secret"):
		return true
	default:
		return false
	}
}

func dataType(v fleetpkg.Var) (string, error) {
	var tfType string
	switch v.Type {
	case "bool":
		tfType = "bool"
	case "integer":
		tfType = "number"
	case "password", "email", "select", "text", "textarea", "time_zone", "url", "yaml":
		tfType = "string"
	default:
		// package-spec controls the allow types.
		return "", fmt.Errorf("unknown fleet variable type %q", v.Type)
	}

	if v.Multi != nil && *v.Multi {
		tfType = "list(" + tfType + ")"
	}
	return tfType, nil
}

type FleetPackagePolicyModule struct {
	AgentPolicyID           string   `json:"agent_policy_id"`
	PackagePolicyName       string   `json:"package_policy_name,omitempty"`
	PackageName             string   `json:"package_name"`
	PackageVersion          string   `json:"package_version"`
	Namespace               string   `json:"namespace"`
	Description             string   `json:"description"`
	PolicyTemplate          string   `json:"policy_template"`
	DataStream              string   `json:"data_stream"`
	InputType               string   `json:"input_type"`
	PackageVariablesJSON    string   `json:"package_variables_json,omitempty"`
	InputVariablesJSON      string   `json:"input_variables_json,omitempty"`
	DataStreamVariablesJSON string   `json:"data_stream_variables_json,omitempty"`
	AllDataStreams          []string `json:"all_data_streams"`
	AllPolicyTemplateInputs []string `json:"all_policy_template_inputs"` // All policy_template + input combos.
}

func toMap(v any) map[string]any {
	buf := new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		panic(err)
	}

	var m map[string]any
	dec := json.NewDecoder(buf)
	dec.UseNumber()
	if err := dec.Decode(&m); err != nil {
		panic(err)
	}

	return m
}

var variableExpressionTemplate = template.Must(template.New("jsonencode").
	Option("missingkey=error").
	Funcs(template.FuncMap{
		"quoteIfNeeded": quoteIfNeeded,
	}).
	Parse(`${jsonencode({
{{- range $fleetVar, $tfVar := . }}
  {{ quoteIfNeeded $fleetVar }} = var.{{ $tfVar }}
{{- end }}
})}`))

func buildVariableExpression(associations ...map[string]string) (string, error) {
	allAssociations := joinMaps(associations...)

	if len(allAssociations) == 0 {
		return "", nil
	}

	buf := new(bytes.Buffer)
	if err := variableExpressionTemplate.Execute(buf, allAssociations); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func joinMaps(maps ...map[string]string) map[string]string {
	if len(maps) == 0 {
		return nil
	}
	if len(maps) == 1 {
		return maps[0]
	}

	out := map[string]string{}
	for _, m := range maps {
		for k, v := range m {
			if _, found := out[k]; found {
				panic("Multiple definitions for variable " + k)
			}
			out[k] = v
		}
	}

	return out
}

func moduleName(integration, policyTemplate, dataStream, input string) string {
	name := []string{integration}
	// Most policy templates are named the same as their integrations so
	// remove that to keep names shorter.
	if integration != policyTemplate {
		name = append(name, policyTemplate)
	}
	if policyTemplate != dataStream && dataStream != "" {
		name = append(name, dataStream)
	}
	name = append(name, strings.ReplaceAll(input, "/", "_"))
	return strings.Join(name, ".")
}

func datasetName(dataStreamDirName string, m *fleetpkg.DataStream) string {
	if m != nil {
		if _, dataset, found := strings.Cut(m.Manifest.Dataset, "."); found {
			return dataset
		}
	}

	return dataStreamDirName
}

// quoteIfNeeded surrounds a variable name with quotes if it contains dots.
// HCL used in Terraform does not allow names with dots.
func quoteIfNeeded(name string) string {
	if strings.Contains(name, ".") {
		return strconv.Quote(name)
	}
	return name
}

// injectDataStreamDatasetVar creates a data_stream.dataset variable in the
// given policy_template if it does not exist.
//
// Reference: https://github.com/andrewkroh/terraform-module-fleet/issues/26
func injectDataStreamDatasetVar(policy *fleetpkg.PolicyTemplate) {
	const datasetVarName = "data_stream.dataset"
	for _, v := range policy.Vars {
		if v.Name == datasetVarName {
			return
		}
	}

	policy.Vars = append(policy.Vars, fleetpkg.Var{
		Name:        datasetVarName,
		Description: "Set the name for your dataset. Once selected a dataset cannot be changed without creating a new integration policy. You can't use - in the name of a dataset and only valid characters for Elasticsearch index names are permitted.",
		Type:        "text",
		Title:       "Dataset Name",
		Default:     policy.Name,
		Required:    pointerTo(true),
	})
}

func pointerTo[T any](t T) *T {
	return &t
}
