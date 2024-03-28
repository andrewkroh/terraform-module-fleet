<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_elasticstack"></a> [elasticstack](#requirement\_elasticstack) | >= 0.11.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_elasticstack"></a> [elasticstack](#provider\_elasticstack) | >= 0.11.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [elasticstack_fleet_agent_policy.agent_policy](https://registry.terraform.io/providers/elastic/elasticstack/latest/docs/resources/fleet_agent_policy) | resource |
| [elasticstack_fleet_enrollment_tokens.enrollment_token](https://registry.terraform.io/providers/elastic/elasticstack/latest/docs/data-sources/fleet_enrollment_tokens) | data source |

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
| <a name="input_skip_destroy"></a> [skip\_destroy](#input\_skip\_destroy) | If true, the agent policy will not be deleted at destroy time, and instead just removed from the Terraform state. | `bool` | `false` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_enrollment_token"></a> [enrollment\_token](#output\_enrollment\_token) | Token that can be used to enroll Agents into the policy. |
| <a name="output_id"></a> [id](#output\_id) | n/a |
<!-- END_TF_DOCS -->