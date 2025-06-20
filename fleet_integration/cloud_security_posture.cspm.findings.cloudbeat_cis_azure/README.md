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
| <a name="input_arm_template_url"></a> [arm\_template\_url](#input\_arm\_template\_url) | A URL to the ARM Template for creating a new deployment | `string` | `"https://portal.azure.com/#create/Microsoft.Template/uri/https%3A%2F%2Fraw.githubusercontent.com%2Felastic%2Fcloudbeat%2F8.18%2Fdeploy%2Fazure%2FARM-for-ACCOUNT_TYPE.json"` | no |
| <a name="input_azure_account_type"></a> [azure\_account\_type](#input\_azure\_account\_type) | n/a | `string` | `null` | no |
| <a name="input_azure_credentials_client_certificate_password"></a> [azure\_credentials\_client\_certificate\_password](#input\_azure\_credentials\_client\_certificate\_password) | n/a | `string` | `null` | no |
| <a name="input_azure_credentials_client_certificate_path"></a> [azure\_credentials\_client\_certificate\_path](#input\_azure\_credentials\_client\_certificate\_path) | n/a | `string` | `null` | no |
| <a name="input_azure_credentials_client_id"></a> [azure\_credentials\_client\_id](#input\_azure\_credentials\_client\_id) | n/a | `string` | `null` | no |
| <a name="input_azure_credentials_client_secret"></a> [azure\_credentials\_client\_secret](#input\_azure\_credentials\_client\_secret) | n/a | `string` | `null` | no |
| <a name="input_azure_credentials_tenant_id"></a> [azure\_credentials\_tenant\_id](#input\_azure\_credentials\_tenant\_id) | n/a | `string` | `null` | no |
| <a name="input_azure_credentials_type"></a> [azure\_credentials\_type](#input\_azure\_credentials\_type) | n/a | `string` | `null` | no |
| <a name="input_condition"></a> [condition](#input\_condition) | Condition to filter when to collect this input. See [Dynamic Input Configuration](https://www.elastic.co/guide/en/fleet/current/dynamic-input-configuration.html) for details. | `string` | `null` | no |
| <a name="input_deployment"></a> [deployment](#input\_deployment) | Chosen deployment type (aws/gcp/azure/eks/k8s) | `string` | n/a | yes |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the cloud\_security\_posture package to use. | `string` | `"2.0.0-preview04"` | no |
| <a name="input_posture"></a> [posture](#input\_posture) | Chosen posture type (cspm/kspm) | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->