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

package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/elastic/terraform-module-fleet/fleet-terraform-generator/internal/module"
)

const (
	generatePackageFlagName        = "package"
	generatePolicyTemplateFlagName = "policy-template"
	generateDataStreamFlagName     = "data-stream"
	generateInputFlagName          = "input"
	generateOutputPathFlagName     = "out"
)

func GenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "List all Fleet packages, policy templates, data streams, and inputs.",
		RunE:  generateModuleRunE,
	}
	cmd.PersistentFlags().String(packagesDirectoryFlagName, "", "Directory containing Fleet packages.")
	must(cmd.MarkPersistentFlagRequired(packagesDirectoryFlagName))

	cmd.Flags().StringP(generatePackageFlagName, "p", "", "Package name.")
	must(cmd.MarkFlagRequired(generatePackageFlagName))

	cmd.Flags().StringP(generatePolicyTemplateFlagName, "t", "", "Policy template name.")

	cmd.Flags().StringP(generateDataStreamFlagName, "d", "", "Data stream name.")
	must(cmd.MarkFlagRequired(generateDataStreamFlagName))

	cmd.Flags().StringP(generateInputFlagName, "i", "", "Input name.")
	must(cmd.MarkFlagRequired(generateInputFlagName))

	cmd.PersistentFlags().String(generateOutputPathFlagName, "", "Output path. It creates a new sub-directory named based on the package, policy template, data stream, and input.")
	must(cmd.MarkPersistentFlagRequired(generateOutputPathFlagName))

	batchCmd := &cobra.Command{
		Use:   "batch",
		Short: "Generate a batch of Terraform modules as specified by glob or config file.",
		RunE:  generateBatchRunE,
	}
	cmd.AddCommand(batchCmd)

	return cmd
}

func generateModuleRunE(cmd *cobra.Command, _ []string) error {
	pkgsDir, err := cmd.Flags().GetString(packagesDirectoryFlagName)
	if err != nil {
		return err
	}
	packageName, err := cmd.Flags().GetString(generatePackageFlagName)
	if err != nil {
		return err
	}
	policyTemplateName, err := cmd.Flags().GetString(generatePolicyTemplateFlagName)
	if err != nil {
		return err
	}
	dataStreamName, err := cmd.Flags().GetString(generateDataStreamFlagName)
	if err != nil {
		return err
	}
	inputName, err := cmd.Flags().GetString(generateInputFlagName)
	if err != nil {
		return err
	}
	outputDir, err := cmd.Flags().GetString(generateOutputPathFlagName)
	if err != nil {
		return err
	}
	info, err := os.Stat(outputDir)
	if err != nil {
		return fmt.Errorf("failed checkout the output: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("output dir %q is not a directory", outputDir)
	}

	err = generateTerraformModuleToDisk(pkgsDir, packageName, policyTemplateName, dataStreamName, inputName, outputDir)
	if err != nil {
		s := module.Specifier{
			Integration:    packageName,
			PolicyTemplate: policyTemplateName,
			DataStream:     dataStreamName,
			Input:          inputName,
		}
		return fmt.Errorf("error generating module for %s: %w", s, err)
	}
	return nil
}

func generateBatchRunE(cmd *cobra.Command, args []string) error {
	pkgsDir, err := cmd.Flags().GetString(packagesDirectoryFlagName)
	if err != nil {
		return err
	}
	outputDir, err := cmd.Flags().GetString(generateOutputPathFlagName)
	if err != nil {
		return err
	}
	info, err := os.Stat(outputDir)
	if err != nil {
		return fmt.Errorf("failed checkout the output: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("output dir %q is not a directory", outputDir)
	}

	specs, err := module.List(pkgsDir)
	if err != nil {
		return err
	}

	specs, err = specs.Filter(args...)
	if err != nil {
		return err
	}

	var errs []error
	for _, s := range specs {
		if err = generateTerraformModuleToDisk(pkgsDir, s.Integration, s.PolicyTemplate, s.DataStream, s.Input, outputDir); err != nil {
			errs = append(errs, fmt.Errorf("error generating module for %s: %w", s.String(), err))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func generateTerraformModuleToDisk(pkgsDir, packageName, policyTemplateName, dataStreamName, inputName, outputDir string) error {
	tf, err := module.Generate(pkgsDir, packageName, policyTemplateName, dataStreamName, inputName)
	if err != nil {
		return err
	}

	outputPath := filepath.Join(outputDir, tf.Name, "module.tf.json")
	if err = os.Mkdir(filepath.Dir(outputPath), 0o700); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return fmt.Errorf("failed to create directory for module: %w", err)
		}
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err = enc.Encode(tf.File); err != nil {
		return err
	}

	if err = os.WriteFile(outputPath, buf.Bytes(), 0o600); err != nil {
		return err
	}
	return nil
}
