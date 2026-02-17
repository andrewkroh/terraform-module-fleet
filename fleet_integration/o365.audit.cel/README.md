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
| <a name="input_azure_tenant_id"></a> [azure\_tenant\_id](#input\_azure\_tenant\_id) | Directory (tenant) ID | `string` | n/a | yes |
| <a name="input_batch_interval"></a> [batch\_interval](#input\_batch\_interval) | Interval for each API request. The default fetches a single hour of events for each request. This value may not be more than 24h. Supports following suffixes - "h" (hour), "m" (minute), "s" (second), "ms" (millisecond), "us" (microsecond), and "ns" (nanosecond) | `string` | `"1h"` | no |
| <a name="input_client_id"></a> [client\_id](#input\_client\_id) | Client ID used for OAuth 2.0 authentication | `string` | n/a | yes |
| <a name="input_client_secret"></a> [client\_secret](#input\_client\_secret) | Client secret used for OAuth 2.0 authentication | `string` | n/a | yes |
| <a name="input_content_types"></a> [content\_types](#input\_content\_types) | Comma separated list of content types to fetch from Management API.<br>Supported content types are - `Audit.AzureActiveDirectory, Audit.Exchange, Audit.SharePoint, Audit.General, DLP.All`.<br><br>More information can be found in the [documentation](https://learn.microsoft.com/en-us/office/office-365-management-api/office-365-management-activity-api-reference#working-with-the-office-365-management-activity-api). | `string` | `"Audit.AzureActiveDirectory, Audit.Exchange, Audit.SharePoint, Audit.General, DLP.All"` | no |
| <a name="input_enable_request_tracer"></a> [enable\_request\_tracer](#input\_enable\_request\_tracer) | The request tracer logs requests and responses to the agent's local file-system for debugging configurations. Enabling this request tracing compromises security and should only be used for debugging. Disabling the request tracer will delete any stored traces. See [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-cel.html#_resource_tracer_enable) for details. | `bool` | `false` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_force"></a> [fleet\_package\_policy\_force](#input\_fleet\_package\_policy\_force) | Force reinstallation of the package even if already installed. When true, bypasses "already installed" checks and triggers complete re-installation. This deletes and recreates Kibana assets (dashboards, visualizations), removes transforms and their destination indices, and overwrites ingest pipelines and templates. | `bool` | `true` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the o365 package to use. | `string` | `"3.4.0"` | no |
| <a name="input_initial_interval"></a> [initial\_interval](#input\_initial\_interval) | Initial interval for the first API call. Default starts fetching events from 167h55m, i.e., 7 days ago, and must not go further back than that. Supports following suffixes - "h" (hour), "m" (minute), "s" (second), "ms" (millisecond), "us" (microsecond), and "ns" (nanosecond) | `string` | `"167h55m"` | no |
| <a name="input_interval"></a> [interval](#input\_interval) | How often the API is polled, supports seconds, minutes and hours. | `string` | `"3m"` | no |
| <a name="input_max_executions"></a> [max\_executions](#input\_max\_executions) | Maximum Executions Per Interval is the maximum number of executions that can be performed without waiting for the interval time. | `number` | `10000` | no |
| <a name="input_maximum_age"></a> [maximum\_age](#input\_maximum\_age) | A hard maximum age limit for data that can be requested. It defaults to 5 mins less than the API's documented limit but may be shortened as a workaround for errors related to expired data. Supports following suffixes - "h" (hour), "m" (minute), "s" (second), "ms" (millisecond), "us" (microsecond), and "ns" (nanosecond) | `string` | `"167h55m"` | no |
| <a name="input_oauth_endpoint_params_yaml"></a> [oauth\_endpoint\_params\_yaml](#input\_oauth\_endpoint\_params\_yaml) | Endpoint Params used for OAuth2 authentication as YAML. See [documentation](https://www.elastic.co/docs/reference/beats/filebeat/filebeat-input-cel#_auth_oauth2_endpoint_params) for details. | `string` | `"grant_type: client_credentials\n"` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original` | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_resource_proxy_url"></a> [resource\_proxy\_url](#input\_resource\_proxy\_url) | This specifies proxy configuration in the form of `http[s]://<user>:<password>@<server name/ip>:<port>`. | `string` | `null` | no |
| <a name="input_resource_rate_limit_burst"></a> [resource\_rate\_limit\_burst](#input\_resource\_rate\_limit\_burst) | The maximum burst size. Burst is the maximum number of resource requests that can be made above the overall rate limit. | `string` | `null` | no |
| <a name="input_resource_rate_limit_limit"></a> [resource\_rate\_limit\_limit](#input\_resource\_rate\_limit\_limit) | Value of the response that specifies the total limit. | `string` | `null` | no |
| <a name="input_resource_redirect_forward_headers"></a> [resource\_redirect\_forward\_headers](#input\_resource\_redirect\_forward\_headers) | If set to true resource headers are forwarded in case of a redirect. Default is "false". | `bool` | `null` | no |
| <a name="input_resource_redirect_headers_ban_list"></a> [resource\_redirect\_headers\_ban\_list](#input\_resource\_redirect\_headers\_ban\_list) | If "Redirect Forward Headers" is set to true, all headers except the ones defined in this list will be forwarded. All headers are forwarded by default. | `list(string)` | `null` | no |
| <a name="input_resource_redirect_max_redirects"></a> [resource\_redirect\_max\_redirects](#input\_resource\_redirect\_max\_redirects) | Maximum number of redirects to follow for a resource. Default is "10". | `string` | `null` | no |
| <a name="input_resource_retry_max_attempts"></a> [resource\_retry\_max\_attempts](#input\_resource\_retry\_max\_attempts) | Maximum number of retries for the HTTP client. Default is "5". | `string` | `null` | no |
| <a name="input_resource_retry_wait_max"></a> [resource\_retry\_wait\_max](#input\_resource\_retry\_wait\_max) | Maximum time to wait before a retry is attempted. Default is "60s". | `string` | `null` | no |
| <a name="input_resource_retry_wait_min"></a> [resource\_retry\_wait\_min](#input\_resource\_retry\_wait\_min) | Minimum time to wait before a retry is attempted. Default is "1s". | `string` | `null` | no |
| <a name="input_resource_ssl_yaml"></a> [resource\_ssl\_yaml](#input\_resource\_ssl\_yaml) | i.e. certificate\_authorities, supported\_protocols, verification\_mode etc, more examples found in the [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/configuration-ssl.html#ssl-common-config) | `string` | `null` | no |
| <a name="input_resource_timeout"></a> [resource\_timeout](#input\_resource\_timeout) | Duration before declaring that the HTTP client connection has timed out. Valid time units are ns, us, ms, s, m, h. Default is "30"s. | `string` | `"60s"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded",<br>  "o365-cel"<br>]</pre> | no |
| <a name="input_token_scopes"></a> [token\_scopes](#input\_token\_scopes) | Scopes for OAuth 2.0 token. | `list(string)` | <pre>[<br>  "https://manage.office.com/.default"<br>]</pre> | no |
| <a name="input_token_url"></a> [token\_url](#input\_token\_url) | Base URL endpoint that will be used to generate the tokens during the OAuth 2.0 flow. If not provided, above `Azure Tenant ID` will be used for OAuth 2.0 token generation. Default value - `https://login.microsoftonline.com` | `string` | `"https://login.microsoftonline.com"` | no |
| <a name="input_url"></a> [url](#input\_url) | n/a | `string` | `"https://manage.office.com"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->