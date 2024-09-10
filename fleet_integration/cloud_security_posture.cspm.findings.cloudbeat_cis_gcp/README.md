<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_fleet_package_policy"></a> [fleet\_package\_policy](#module\_fleet\_package\_policy) | ../../fleet_package_policy | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cloud_shell_url"></a> [cloud\_shell\_url](#input\_cloud\_shell\_url) | A URL to CloudShell for creating a new deployment | `string` | `"https://shell.cloud.google.com/cloudshell/?ephemeral=true&cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Felastic%2Fcloudbeat&cloudshell_git_branch=8.15&cloudshell_workspace=deploy%2Fdeployment-manager&show=terminal"` | no |
| <a name="input_condition"></a> [condition](#input\_condition) | Condition to filter when to collect this input. See [Dynamic Input Configuration](https://www.elastic.co/guide/en/fleet/current/dynamic-input-configuration.html) for details. | `string` | `null` | no |
| <a name="input_deployment"></a> [deployment](#input\_deployment) | Chosen deployment type (aws/gcp/azure/eks/k8s) | `string` | n/a | yes |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the cloud\_security\_posture package to use. | `string` | `"1.11.0-preview07"` | no |
| <a name="input_gcp_account_type"></a> [gcp\_account\_type](#input\_gcp\_account\_type) | n/a | `string` | n/a | yes |
| <a name="input_gcp_credentials_file"></a> [gcp\_credentials\_file](#input\_gcp\_credentials\_file) | n/a | `string` | `null` | no |
| <a name="input_gcp_credentials_json"></a> [gcp\_credentials\_json](#input\_gcp\_credentials\_json) | n/a | `string` | `null` | no |
| <a name="input_gcp_credentials_type"></a> [gcp\_credentials\_type](#input\_gcp\_credentials\_type) | n/a | `string` | `null` | no |
| <a name="input_gcp_organization_id"></a> [gcp\_organization\_id](#input\_gcp\_organization\_id) | n/a | `string` | `null` | no |
| <a name="input_gcp_project_id"></a> [gcp\_project\_id](#input\_gcp\_project\_id) | n/a | `string` | `null` | no |
| <a name="input_posture"></a> [posture](#input\_posture) | Chosen posture type (cspm/kspm) | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->