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
| <a name="input_data_stream_dataset"></a> [data\_stream\_dataset](#input\_data\_stream\_dataset) | Dataset to write data to. Changing the dataset will send the data to a different index. You can't use `-` in the name of a dataset and only valid characters for [Elasticsearch index names](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html). | `string` | `"cel.cel"` | no |
| <a name="input_delete_redacted_fields"></a> [delete\_redacted\_fields](#input\_delete\_redacted\_fields) | The default behavior for field redaction is to replace characters with `*`s. If field value length or presence will<br>leak information, the fields can be deleted from logging by setting this configuration to true. | `bool` | `false` | no |
| <a name="input_enable_request_tracer"></a> [enable\_request\_tracer](#input\_enable\_request\_tracer) | The request tracer logs HTTP requests and responses to the agent's local file-system for debugging configurations. Enabling this request tracing compromises security and should only be used for debugging. See [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-cel.html#_resource_tracer_filename) for details. | `bool` | `null` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the cel package to use. | `string` | `"1.10.0"` | no |
| <a name="input_oauth_azure_resource"></a> [oauth\_azure\_resource](#input\_oauth\_azure\_resource) | Optional setting for the accessed WebAPI resource when using azure provider. | `string` | `null` | no |
| <a name="input_oauth_azure_tenant_id"></a> [oauth\_azure\_tenant\_id](#input\_oauth\_azure\_tenant\_id) | Optional setting used for authentication when using Azure provider. Since it is used in the process to generate the token\_url, it can’t be used in combination with it. | `string` | `null` | no |
| <a name="input_oauth_endpoint_params_yaml"></a> [oauth\_endpoint\_params\_yaml](#input\_oauth\_endpoint\_params\_yaml) | Set of values that will be sent on each resource to the token\_url. Each param key can have multiple values. Can be set for all providers except google. | `string` | `null` | no |
| <a name="input_oauth_google_credentials_file"></a> [oauth\_google\_credentials\_file](#input\_oauth\_google\_credentials\_file) | The full path to the credentials file for Google. | `string` | `null` | no |
| <a name="input_oauth_google_credentials_json"></a> [oauth\_google\_credentials\_json](#input\_oauth\_google\_credentials\_json) | Your Google credentials information as raw JSON. | `string` | `null` | no |
| <a name="input_oauth_google_delegated_account"></a> [oauth\_google\_delegated\_account](#input\_oauth\_google\_delegated\_account) | Email of the delegated account used to create the credentials (usually an admin). | `string` | `null` | no |
| <a name="input_oauth_google_jwt_file"></a> [oauth\_google\_jwt\_file](#input\_oauth\_google\_jwt\_file) | Full path to the JWT Account Key file for Google. | `string` | `null` | no |
| <a name="input_oauth_google_jwt_json"></a> [oauth\_google\_jwt\_json](#input\_oauth\_google\_jwt\_json) | Your Google JWT information as raw JSON. | `string` | `null` | no |
| <a name="input_oauth_id"></a> [oauth\_id](#input\_oauth\_id) | Client ID used for OAuth2 authentication | `string` | `null` | no |
| <a name="input_oauth_okta_jwt_file"></a> [oauth\_okta\_jwt\_file](#input\_oauth\_okta\_jwt\_file) | Full path to the JWT account private key file for Okta. | `string` | `null` | no |
| <a name="input_oauth_okta_jwt_json"></a> [oauth\_okta\_jwt\_json](#input\_oauth\_okta\_jwt\_json) | Your Okta JWT private key as raw JSON. | `string` | `null` | no |
| <a name="input_oauth_okta_jwt_pem"></a> [oauth\_okta\_jwt\_pem](#input\_oauth\_okta\_jwt\_pem) | Your Okta JWT private key encoded as a PEM block. | `string` | `null` | no |
| <a name="input_oauth_provider"></a> [oauth\_provider](#input\_oauth\_provider) | Used to configure supported oAuth2 providers. Each supported provider will require specific settings. It is not set by default. Supported providers are "azure" and "google". | `string` | `null` | no |
| <a name="input_oauth_scopes"></a> [oauth\_scopes](#input\_oauth\_scopes) | A list of scopes that will be resourceed during the oAuth2 flow. It is optional for all providers. | `list(string)` | `null` | no |
| <a name="input_oauth_secret"></a> [oauth\_secret](#input\_oauth\_secret) | Client secret used for OAuth2 authentication | `string` | `null` | no |
| <a name="input_oauth_token_url"></a> [oauth\_token\_url](#input\_oauth\_token\_url) | The URL endpoint that will be used to generate the tokens during the oAuth2 flow. It is required if no oauth\_custom variable is set or provider is not specified in oauth\_custom variable. | `string` | `null` | no |
| <a name="input_password"></a> [password](#input\_password) | The password to be used with Basic Auth headers | `string` | `null` | no |
| <a name="input_pipeline"></a> [pipeline](#input\_pipeline) | The Ingest Node pipeline ID to be used by the integration. | `string` | `null` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_program"></a> [program](#input\_program) | Program is the CEL program that is executed each polling period to get and transform the API data.<br>More information can be found in the [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-cel.html#_execution). | `string` | `"# // Fetch the agent's public IP every minute and note when the last request was made.\n# // It does not use the Resource URL configuration value.\n# bytes(get(\"https://api.ipify.org/?format=json\").Body).as(body, {\n#     \"events\": [body.decode_json().with({\n#         \"last_requested_at\": has(state.cursor) && has(state.cursor.last_requested_at) ?\n#             state.cursor.last_requested_at\n#         :\n#             now\n#     })],\n#     \"cursor\": {\"last_requested_at\": now}\n# })\n"` | no |
| <a name="input_redact_fields"></a> [redact\_fields](#input\_redact\_fields) | Fields to redact in debug logs. When logging at debug-level the input state and CEL evaluation state are included<br>in logs. This may leak secrets, so list sensitive state fields in this configuration. | `list(string)` | `null` | no |
| <a name="input_regexp_yaml"></a> [regexp\_yaml](#input\_regexp\_yaml) | Regexps is the set of regular expression to be made available to the program by name. The syntax used is [RE2](https://github.com/google/re2/wiki/Syntax).<br>More information can be found in the [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-cel.html#regexp-cel). | `string` | `"#products: '(?i)(Elasticsearch|Beats|Logstash|Kibana)'\n#solutions: '(?i)(Search|Observability|Security)'\n"` | no |
| <a name="input_resource_interval"></a> [resource\_interval](#input\_resource\_interval) | How often the API is polled, supports seconds, minutes and hours. | `string` | `"1m"` | no |
| <a name="input_resource_proxy_url"></a> [resource\_proxy\_url](#input\_resource\_proxy\_url) | This specifies proxy configuration in the form of `http[s]://<user>:<password>@<server name/ip>:<port>`. | `string` | `null` | no |
| <a name="input_resource_rate_limit_burst"></a> [resource\_rate\_limit\_burst](#input\_resource\_rate\_limit\_burst) | The maximum burst size. Burst is the maximum number of resource requests that can be made above the overall rate limit. | `string` | `null` | no |
| <a name="input_resource_rate_limit_limit"></a> [resource\_rate\_limit\_limit](#input\_resource\_rate\_limit\_limit) | The value of the response that specifies the total limit. | `string` | `null` | no |
| <a name="input_resource_redirect_forward_headers"></a> [resource\_redirect\_forward\_headers](#input\_resource\_redirect\_forward\_headers) | When set to true resource headers are forwarded in case of a redirect. Default is "false". | `bool` | `null` | no |
| <a name="input_resource_redirect_headers_ban_list"></a> [resource\_redirect\_headers\_ban\_list](#input\_resource\_redirect\_headers\_ban\_list) | When Redirect Forward Headers is set to true, all headers except the ones defined in this list will be forwarded. All headers are forwarded by default. | `list(string)` | `null` | no |
| <a name="input_resource_redirect_max_redirects"></a> [resource\_redirect\_max\_redirects](#input\_resource\_redirect\_max\_redirects) | The maximum number of redirects to follow for a resource. Default is "10". | `string` | `null` | no |
| <a name="input_resource_retry_max_attempts"></a> [resource\_retry\_max\_attempts](#input\_resource\_retry\_max\_attempts) | The maximum number of retries for the HTTP client. Default is "5". | `string` | `null` | no |
| <a name="input_resource_retry_wait_max"></a> [resource\_retry\_wait\_max](#input\_resource\_retry\_wait\_max) | The maximum time to wait before a retry is attempted. Default is "60s". | `string` | `null` | no |
| <a name="input_resource_retry_wait_min"></a> [resource\_retry\_wait\_min](#input\_resource\_retry\_wait\_min) | The minimum time to wait before a retry is attempted. Default is "1s". | `string` | `null` | no |
| <a name="input_resource_ssl_yaml"></a> [resource\_ssl\_yaml](#input\_resource\_ssl\_yaml) | i.e. certificate\_authorities, supported\_protocols, verification\_mode etc, more examples found in the [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/configuration-ssl.html#ssl-common-config) | `string` | `null` | no |
| <a name="input_resource_timeout"></a> [resource\_timeout](#input\_resource\_timeout) | Duration before declaring that the HTTP client connection has timed out. Valid time units are ns, us, ms, s, m, h. Default is "30"s. | `string` | `null` | no |
| <a name="input_resource_url"></a> [resource\_url](#input\_resource\_url) | i.e. scheme://host:port/path | `string` | `"https://server.example.com:8089/api"` | no |
| <a name="input_state_yaml"></a> [state\_yaml](#input\_state\_yaml) | State is the initial state to be provided to the program. If it has a cursor field, that field will be overwritten by any stored cursor, but will be available if no stored cursor exists.<br>More information can be found in the [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-cel.html#input-state-cel). | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded"<br>]</pre> | no |
| <a name="input_username"></a> [username](#input\_username) | The username to be used with Basic Auth headers | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->