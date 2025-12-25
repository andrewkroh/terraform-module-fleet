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
| <a name="input_access_token"></a> [access\_token](#input\_access\_token) | the GitHub Personal Access Token. | `string` | n/a | yes |
| <a name="input_api_url"></a> [api\_url](#input\_api\_url) | The API URL without the path. | `string` | `"https://api.github.com"` | no |
| <a name="input_enable_request_tracer"></a> [enable\_request\_tracer](#input\_enable\_request\_tracer) | The request tracer logs requests and responses to the agent's local file-system for debugging configurations. Enabling this request tracing compromises security and should only be used for debugging. See [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-httpjson.html#_request_tracer_filename) for details. | `bool` | `null` | no |
| <a name="input_filter"></a> [filter](#input\_filter) | Indicates which sorts of issues to return. <br>Can be one of - `assigned`, `created`, `mentioned`, `subscribed`, `repos`, `all`. <br>`assigned` means issues assigned to you. `created` means issues created by you. `mentioned` means issues mentioning you. `subscribed` means issues you're subscribed to updates for. `all` or repos means all issues you can see, regardless of participation or creation. | `string` | `"all"` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the github package to use. | `string` | `"2.19.0"` | no |
| <a name="input_http_client_timeout"></a> [http\_client\_timeout](#input\_http\_client\_timeout) | Duration before declaring that the HTTP client connection has timed out. Valid time units are ns, us, ms, s, m, h. | `string` | `"60s"` | no |
| <a name="input_interval"></a> [interval](#input\_interval) | Interval at which the alerts will be pulled. The value must be between 2m and 1h. Supported units for this parameter are h/m/s. | `string` | `"10m"` | no |
| <a name="input_labels"></a> [labels](#input\_labels) | A list of comma separated label names. Example - bug,ui,@high | `string` | `null` | no |
| <a name="input_owner"></a> [owner](#input\_owner) | The owner of GitHub Repository. If repository belongs to an organization, owner is name of the organization | `string` | n/a | yes |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original` | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. <br>This executes in the agent before the logs are parsed. <br>See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_proxy_url"></a> [proxy\_url](#input\_proxy\_url) | URL to proxy connections in the form of http\[s\]://<user>:<password>@<server name/ip>:<port> | `string` | `null` | no |
| <a name="input_repo"></a> [repo](#input\_repo) | The GitHub Repository. If not provided, alerts for all the repositories of the owner will be ingested | `string` | `null` | no |
| <a name="input_since"></a> [since](#input\_since) | Only show notifications updated after the given time are returned. This is a timestamp in ISO 8601 format - `YYYY-MM-DDTHH:MM:SSZ`. | `string` | `null` | no |
| <a name="input_ssl_yaml"></a> [ssl\_yaml](#input\_ssl\_yaml) | SSL configuration options. See [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/configuration-ssl.html#ssl-common-config) for details. | `string` | `null` | no |
| <a name="input_state"></a> [state](#input\_state) | Indicates the state of the issues to return. | `string` | `"all"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded",<br>  "github-issues"<br>]</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->