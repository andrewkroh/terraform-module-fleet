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
| <a name="input_auth_key"></a> [auth\_key](#input\_auth\_key) | abuse.ch Auth Key. This is required for authentication for all API requests. See [documentation](https://abuse.ch/blog/community-first/) for details. | `string` | n/a | yes |
| <a name="input_enable_request_tracer"></a> [enable\_request\_tracer](#input\_enable\_request\_tracer) | The request tracer logs requests and responses to the agent's local file-system for debugging configurations. Enabling this request tracing compromises security and should only be used for debugging. Disabling the request tracer will delete any stored traces. See [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-cel.html#_resource_tracer_enable) for details. | `bool` | `false` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_force"></a> [fleet\_package\_policy\_force](#input\_fleet\_package\_policy\_force) | Force reinstallation of the package even if already installed. When true, bypasses "already installed" checks and triggers complete re-installation. This deletes and recreates Kibana assets (dashboards, visualizations), removes transforms and their destination indices, and overwrites ingest pipelines and templates. | `bool` | `true` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the ti\_abusech package to use. | `string` | `"3.6.0"` | no |
| <a name="input_http_client_timeout"></a> [http\_client\_timeout](#input\_http\_client\_timeout) | Duration before declaring that the HTTP client connection has timed out. Valid time units are ns, us, ms, s, m, h. | `string` | `"30s"` | no |
| <a name="input_initial_interval"></a> [initial\_interval](#input\_initial\_interval) | How far back to pull the threat indicators from ThreatFox in days. Can be any number between `1` to `7`. Example `5`. Default value is `7` i.e., 7 days. | `number` | `7` | no |
| <a name="input_interval"></a> [interval](#input\_interval) | Duration between requests to the ThreatFox API. Supported units for this parameter are h/m/s. Example `24h`. | `string` | `"24h"` | no |
| <a name="input_ioc_expiration_duration"></a> [ioc\_expiration\_duration](#input\_ioc\_expiration\_duration) | Indicator is expired after this duration since its last seen timestamp. Supported units for this parameter are d/h/m. Example `10d`. Default value is `90d` i.e., 90 days. | `string` | `"90d"` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original` | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_proxy_url"></a> [proxy\_url](#input\_proxy\_url) | URL to proxy connections in the form of http\[s\]://<user>:<password>@<server name/ip>:<port> | `string` | `null` | no |
| <a name="input_ssl_yaml"></a> [ssl\_yaml](#input\_ssl\_yaml) | SSL configuration options. See [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/configuration-ssl.html#ssl-common-config) for details. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded",<br>  "abusech-threatfox"<br>]</pre> | no |
| <a name="input_url"></a> [url](#input\_url) | Base URL of the abuse.ch ThreatFox API to collect threat indicators. | `string` | `"https://threatfox-api.abuse.ch/api/v1/"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->