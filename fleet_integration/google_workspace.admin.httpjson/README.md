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
| <a name="input_api_host"></a> [api\_host](#input\_api\_host) | The Google Workspace API Host. The path will be automatically set. | `string` | `"https://www.googleapis.com"` | no |
| <a name="input_delegated_account"></a> [delegated\_account](#input\_delegated\_account) | Email of the admin user used to access the API. | `string` | n/a | yes |
| <a name="input_enable_request_tracer"></a> [enable\_request\_tracer](#input\_enable\_request\_tracer) | The request tracer logs requests and responses to the agent's local file-system for debugging configurations. Enabling this request tracing compromises security and should only be used for debugging. See [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-httpjson.html#_request_tracer_filename) for details. | `bool` | `null` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the google\_workspace package to use. | `string` | `"2.14.0"` | no |
| <a name="input_http_client_timeout"></a> [http\_client\_timeout](#input\_http\_client\_timeout) | Duration of the time limit on HTTP requests. Valid time units are ns, us, ms, s, m, h. | `string` | `"60s"` | no |
| <a name="input_initial_interval"></a> [initial\_interval](#input\_initial\_interval) | How far back to pull the DHCP Lease events from Google Workspace. NOTE: Supported units for this parameter are h/m/s. | `string` | `"24h"` | no |
| <a name="input_interval"></a> [interval](#input\_interval) | Duration between requests to the API. Google Workspace defaults to a 2 hour polling interval because Google reports can go from some minutes up to 3 days of delay. For more details on this, you can read more at https://support.google.com/a/answer/7061566. NOTE: Supported units for this parameter are h/m/s. | `string` | `"2h"` | no |
| <a name="input_jwt_file"></a> [jwt\_file](#input\_jwt\_file) | Specifies the path to the JWT credentials file.<br>NOTE: Please use either JWT File or JWT JSON parameter. | `string` | `null` | no |
| <a name="input_jwt_json"></a> [jwt\_json](#input\_jwt\_json) | Raw contents of the JWT file. Useful when hosting a file along with the agent is not possible.<br>NOTE: Please use either JWT File or JWT JSON parameter. | `string` | `null` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original` | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded",<br>  "google_workspace-admin"<br>]</pre> | no |
| <a name="input_user_key"></a> [user\_key](#input\_user\_key) | Specifies the user key to fetch reports from. | `string` | `"all"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->