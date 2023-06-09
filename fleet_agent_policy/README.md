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
| [restapi_object.agent_policy](https://registry.terraform.io/providers/Mastercard/restapi/latest/docs/resources/object) | resource |
| [restapi_object.enrollment_token](https://registry.terraform.io/providers/Mastercard/restapi/latest/docs/data-sources/object) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_data_output_id"></a> [data\_output\_id](#input\_data\_output\_id) | ID of the Fleet data output. Use this to choose a non-default output. | `string` | `null` | no |
| <a name="input_description"></a> [description](#input\_description) | Description of Agent policy. | `string` | n/a | yes |
| <a name="input_fleet_server_host_id"></a> [fleet\_server\_host\_id](#input\_fleet\_server\_host\_id) | ID of the Fleet server host. Use this to choose a non-default Fleet server address. | `string` | `null` | no |
| <a name="input_inactivity_timeout_sec"></a> [inactivity\_timeout\_sec](#input\_inactivity\_timeout\_sec) | An agent will automatically change to inactive status and be filtered out | `number` | `3600` | no |
| <a name="input_monitor_logs"></a> [monitor\_logs](#input\_monitor\_logs) | n/a | `bool` | `true` | no |
| <a name="input_monitor_metrics"></a> [monitor\_metrics](#input\_monitor\_metrics) | n/a | `bool` | `true` | no |
| <a name="input_monitoring_output_id"></a> [monitoring\_output\_id](#input\_monitoring\_output\_id) | ID of the Fleet monitoring output. Use this to choose a non-default output. | `string` | `null` | no |
| <a name="input_name"></a> [name](#input\_name) | Name of Agent policy. | `string` | n/a | yes |
| <a name="input_namespace"></a> [namespace](#input\_namespace) | Namespace for the policy. Controls namespace of monitoring data. | `string` | `"default"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_enrollment_token"></a> [enrollment\_token](#output\_enrollment\_token) | Token that can be used to enroll Agents into the policy. |
| <a name="output_id"></a> [id](#output\_id) | n/a |
<!-- END_TF_DOCS -->