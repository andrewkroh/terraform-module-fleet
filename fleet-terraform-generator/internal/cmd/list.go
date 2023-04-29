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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/elastic/terraform-module-fleet/fleet-terraform-generator/internal/module"
)

const (
	packagesDirectoryFlagName = "packages-dir"
)

func ListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all Fleet packages, policy templates, data streams, and inputs.",
		RunE:  listRunE,
	}
	cmd.Flags().String(packagesDirectoryFlagName, "", "Directory containing Fleet packages.")
	must(cmd.MarkFlagRequired(packagesDirectoryFlagName))
	return cmd
}

func listRunE(cmd *cobra.Command, args []string) error {
	pkgsDir, err := cmd.Flags().GetString(packagesDirectoryFlagName)
	if err != nil {
		return err
	}

	moduleSpecifiers, err := module.List(pkgsDir)
	if err != nil {
		return err
	}

	moduleSpecifiers, err = moduleSpecifiers.Filter(args...)
	if err != nil {
		return err
	}

	for _, s := range moduleSpecifiers {
		fmt.Fprintln(cmd.OutOrStdout(), s.String())
	}
	return nil
}
