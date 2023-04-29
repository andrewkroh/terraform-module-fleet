<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_restapi"></a> [restapi](#requirement\_restapi) | >= 1.18.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_restapi"></a> [restapi](#provider\_restapi) | >= 1.18.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [restapi_object.package_policy](https://registry.terraform.io/providers/Mastercard/restapi/latest/docs/resources/object) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_agent_policy_id"></a> [agent\_policy\_id](#input\_agent\_policy\_id) | Fleet Agent policy ID | `string` | n/a | yes |
| <a name="input_all_data_streams"></a> [all\_data\_streams](#input\_all\_data\_streams) | n/a | `list(string)` | `[]` | no |
| <a name="input_data_stream"></a> [data\_stream](#input\_data\_stream) | Name of the data\_stream within the integration (e.g. "log"). | `any` | n/a | yes |
| <a name="input_data_stream_variables_json"></a> [data\_stream\_variables\_json](#input\_data\_stream\_variables\_json) | JSON encoded data stream specific variables. | `string` | `"{}"` | no |
| <a name="input_input_type"></a> [input\_type](#input\_input\_type) | Input type. | `string` | n/a | yes |
| <a name="input_input_variables_json"></a> [input\_variables\_json](#input\_input\_variables\_json) | JSON encoded input level variables that are shared by each stream. | `string` | `"{}"` | no |
| <a name="input_namespace"></a> [namespace](#input\_namespace) | Namespace to apply to data streams. | `string` | `"default"` | no |
| <a name="input_package_name"></a> [package\_name](#input\_package\_name) | Package name. | `string` | n/a | yes |
| <a name="input_package_policy_name"></a> [package\_policy\_name](#input\_package\_policy\_name) | Name of the package policy. | `any` | `null` | no |
| <a name="input_package_version"></a> [package\_version](#input\_package\_version) | Package version. | `string` | n/a | yes |
| <a name="input_policy_template"></a> [policy\_template](#input\_policy\_template) | Name of the policy template containing the input (see the integration's manifest.yml). | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->