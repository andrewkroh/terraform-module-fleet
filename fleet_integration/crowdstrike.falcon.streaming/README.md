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
| <a name="input_app_id"></a> [app\_id](#input\_app\_id) | This field specifies the `appId` parameter sent to the CrowdStrike API. See the CrowdStrike documentation for details. | `string` | n/a | yes |
| <a name="input_client_id"></a> [client\_id](#input\_client\_id) | Client ID for the CrowdStrike API. | `string` | n/a | yes |
| <a name="input_client_secret"></a> [client\_secret](#input\_client\_secret) | Client Secret for the CrowdStrike API. | `string` | n/a | yes |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_force"></a> [fleet\_package\_policy\_force](#input\_fleet\_package\_policy\_force) | Force reinstallation of the package even if already installed. When true, bypasses "already installed" checks and triggers complete re-installation. This deletes and recreates Kibana assets (dashboards, visualizations), removes transforms and their destination indices, and overwrites ingest pipelines and templates. | `bool` | `true` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_prerelease"></a> [fleet\_package\_prerelease](#input\_fleet\_package\_prerelease) | Include prerelease package versions when fleet\_package\_version is "latest". Required for packages below version 1.0.0. | `bool` | `false` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the crowdstrike package to use. Use "latest" to resolve the latest available version from the package registry at plan time. | `string` | `"4.4.0"` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original`. | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_proxy_headers_yaml"></a> [proxy\_headers\_yaml](#input\_proxy\_headers\_yaml) | This specifies the headers to be sent to the proxy server. | `string` | `null` | no |
| <a name="input_proxy_url"></a> [proxy\_url](#input\_proxy\_url) | URL to proxy connections in the form of http[s]://<user>:<password>@<server name/ip>:<port>. Ensure your username and password are in URL encoded format. | `string` | `null` | no |
| <a name="input_retry_infinite_retries"></a> [retry\_infinite\_retries](#input\_retry\_infinite\_retries) | Retry failed connections indefinitely instead of terminating after the maximum number of attempts is reached. When enabled, `Retry - Maximum Attempts` is ignored. Requires Elastic Agent 8.19.19, 9.3.8, 9.4.4, or later; on earlier agents retries are silently capped at 10. | `bool` | `false` | no |
| <a name="input_retry_max_attempts"></a> [retry\_max\_attempts](#input\_retry\_max\_attempts) | The maximum number of times a failed connection to the Event Streams API is retried before the input terminates and needs to be restarted. Failures such as a non-200 response or bad credentials count toward this limit. Enable `Retry - Infinite Retries` to never terminate on repeated failures. On Elastic Agent 8.19.19, 9.3.8, 9.4.4, or later, transient connection-level failures (empty responses, network errors, timeouts) are retried indefinitely and do not count toward this limit; values above 10 (and `Retry - Infinite Retries`) also require those agent versions and are silently capped at 10 on earlier agents. | `number` | `5` | no |
| <a name="input_retry_wait_max"></a> [retry\_wait\_max](#input\_retry\_wait\_max) | The maximum time to wait before retrying a failed connection, as a duration string (for example `30s`). Back-off grows exponentially between the minimum and maximum wait. | `string` | `"30s"` | no |
| <a name="input_retry_wait_min"></a> [retry\_wait\_min](#input\_retry\_wait\_min) | The minimum time to wait before retrying a failed connection, as a duration string (for example `1s` or `500ms`). Must be less than or equal to `Retry - Maximum Wait`. | `string` | `"1s"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded",<br>  "crowdstrike-falcon"<br>]</pre> | no |
| <a name="input_token_url"></a> [token\_url](#input\_token\_url) | CrowdStrike API token URL. | `string` | `"https://api.crowdstrike.com/oauth2/token"` | no |
| <a name="input_url"></a> [url](#input\_url) | Base URL of the CrowdStrike API. Defaults to https://api.crowdstrike.com. | `string` | `"https://api.crowdstrike.com"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->