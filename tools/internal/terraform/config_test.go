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

package terraform_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/terraform-module-fleet/tools/internal/terraform"
)

func TestMarshalJSON(t *testing.T) {
	data := terraform.File{
		Variables: map[string]terraform.Variable{
			"foo": {
				Type:        "string",
				Description: "Configures the bar.",
			},
		},
		Modules: map[string]terraform.Module{
			"package_policy_github_issues": {
				Source: "../../fleet_package_policy",
				Params: map[string]any{
					"agent_policy_id": "ebe2efde-b965-4df3-9d2c-1dc8b808fe72",
				},
			},
		},
		Outputs: map[string]terraform.Output{
			"agent_policy_id": {
				Description: "Agent policy ID",
				Sensitive:   terraform.Ptr(true),
				Value:       "${module.fleet_agent_policy.id}",
			},
		},
	}

	expected := `
{
  "variable": {
    "foo": {
      "type": "string",
      "description": "Configures the bar."
    }
  },
  "output": {
    "agent_policy_id": {
      "description": "Agent policy ID",
      "sensitive": true,
      "value": "${module.fleet_agent_policy.id}"
    }
  },
  "module": {
    "package_policy_github_issues": {
      "agent_policy_id": "ebe2efde-b965-4df3-9d2c-1dc8b808fe72",
      "source": "../../fleet_package_policy"
    }
  }
}

`
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(t, expected, buf.String())
}
