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
| <a name="input_access_id"></a> [access\_id](#input\_access\_id) | Access ID of a ThreatConnect API User. | `string` | n/a | yes |
| <a name="input_batch_size"></a> [batch\_size](#input\_batch\_size) | Batch size for the response of the ThreatConnect API. The maximum supported batch size value is 10000. | `number` | `2000` | no |
| <a name="input_enable_request_tracer"></a> [enable\_request\_tracer](#input\_enable\_request\_tracer) | The request tracer logs requests and responses to the agent's local file-system for debugging configurations. Enabling this request tracing compromises security and should only be used for debugging. See [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-httpjson.html#_request_tracer_filename) for details. | `bool` | `null` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the ti\_threatconnect package to use. | `string` | `"1.8.0"` | no |
| <a name="input_http_client_timeout"></a> [http\_client\_timeout](#input\_http\_client\_timeout) | Duration before declaring that the HTTP client connection has timed out. Valid time units are ns, us, ms, s, m, h. | `string` | `"2m"` | no |
| <a name="input_include_attributes"></a> [include\_attributes](#input\_include\_attributes) | Elastic has a default async search response size of 10MB. Enabling this option may result in larger response size, which may cause the Kibana discover view to have issues. | `bool` | `false` | no |
| <a name="input_include_group_assoc"></a> [include\_group\_assoc](#input\_include\_group\_assoc) | Elastic has a default async search response size of 10MB. Enabling this option may result in larger response size, which may cause the Kibana discover view to have issues. | `bool` | `false` | no |
| <a name="input_include_indicator_assoc"></a> [include\_indicator\_assoc](#input\_include\_indicator\_assoc) | Elastic has a default async search response size of 10MB. Enabling this option may result in larger response size, which may cause the Kibana discover view to have issues. | `bool` | `false` | no |
| <a name="input_initial_interval"></a> [initial\_interval](#input\_initial\_interval) | How far back to pull Indicators and the groups associated with those indicators from ThreatConnect. Supported units for this parameter are h/m/s. | `string` | `"168h"` | no |
| <a name="input_interval"></a> [interval](#input\_interval) | Duration between requests to the ThreatConnect API. Supported units for this parameter are h/m/s. | `string` | `"24h"` | no |
| <a name="input_ioc_expiration_duration"></a> [ioc\_expiration\_duration](#input\_ioc\_expiration\_duration) | Enforces all IOCs to expire after this duration. This setting is required to avoid "orphaned" IOCs that never expire. Specify [Elasticsearch time units](https://www.elastic.co/guide/en/elasticsearch/reference/current/api-conventions.html#time-units) using only days, hours, or minutes (e.g., 10d), avoiding mixed time units. | `string` | `"90d"` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original`. | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_proxy_url"></a> [proxy\_url](#input\_proxy\_url) | URL to proxy connections in the form of http[s]://<user>:<password>@<server name/ip>:<port>. Please ensure your username and password are in URL encoded format. | `string` | `null` | no |
| <a name="input_secret_key"></a> [secret\_key](#input\_secret\_key) | Secret Key of a ThreatConnect API User. | `string` | n/a | yes |
| <a name="input_ssl_yaml"></a> [ssl\_yaml](#input\_ssl\_yaml) | SSL configuration options. See [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/configuration-ssl.html#ssl-common-config) for details. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded",<br>  "threatconnect-indicator"<br>]</pre> | no |
| <a name="input_tql"></a> [tql](#input\_tql) | Filter results based on query written in [TQL](https://knowledge.threatconnect.com/docs/constructing-query-expressions). | `string` | `null` | no |
| <a name="input_url"></a> [url](#input\_url) | Base URL of the ThreatConnect API. Default URL given is for ThreatConnect's Public Cloud instance. Note: Do not include trailing slash “/” character. | `string` | `"https://app.threatconnect.com"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->