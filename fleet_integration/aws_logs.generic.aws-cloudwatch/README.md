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
| <a name="input_access_key_id"></a> [access\_key\_id](#input\_access\_key\_id) | n/a | `string` | `null` | no |
| <a name="input_api_sleep"></a> [api\_sleep](#input\_api\_sleep) | This is used to sleep between AWS FilterLogEvents API calls inside the same collection period. `FilterLogEvents` API has a quota of 5 transactions per second (TPS)/account/Region. This value should only be adjusted when there are multiple Filebeats or multiple Filebeat inputs collecting logs from the same region and AWS account. | `string` | `"200ms"` | no |
| <a name="input_api_timeput"></a> [api\_timeput](#input\_api\_timeput) | The maximum duration of AWS API can take. If it exceeds the timeout, AWS API will be interrupted. | `string` | `"120s"` | no |
| <a name="input_credential_profile_name"></a> [credential\_profile\_name](#input\_credential\_profile\_name) | n/a | `string` | `null` | no |
| <a name="input_data_stream_dataset"></a> [data\_stream\_dataset](#input\_data\_stream\_dataset) | Set the name for your dataset. Changing the dataset will send the data to a different index. You can't use `-` in the name of a dataset and only valid characters for [Elasticsearch index names](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html). | `string` | `"aws_logs.generic"` | no |
| <a name="input_default_region"></a> [default\_region](#input\_default\_region) | Default region to use prior to connecting to region specific services/endpoints if no AWS region is set from environment variable, credentials or instance profile. If none of the above are set and no default region is set as well, `us-east-1` is used. A region, either from environment variable, credentials or instance profile or from this default region setting, needs to be set when using regions in non-regular AWS environments such as AWS China or US Government Isolated. | `string` | `""` | no |
| <a name="input_endpoint"></a> [endpoint](#input\_endpoint) | URL of the entry point for an AWS web service | `string` | `""` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the aws\_logs package to use. | `string` | `"1.1.0"` | no |
| <a name="input_latency"></a> [latency](#input\_latency) | The amount of time required for the logs to be available to CloudWatch Logs. Sample values, `1m` or `5m` — see Golang [time.ParseDuration](https://pkg.go.dev/time#ParseDuration) for more details. Latency translates the query's time range to consider the CloudWatch Logs latency. Example: `5m` means that the integration will query CloudWatch to search for logs available 5 minutes ago. | `string` | `null` | no |
| <a name="input_log_group_arn"></a> [log\_group\_arn](#input\_log\_group\_arn) | ARN of the log group to collect logs from. | `string` | `null` | no |
| <a name="input_log_group_name"></a> [log\_group\_name](#input\_log\_group\_name) | Name of the log group to collect logs from. `region_name` is required when `log_group_name` is given. | `string` | `null` | no |
| <a name="input_log_group_name_prefix"></a> [log\_group\_name\_prefix](#input\_log\_group\_name\_prefix) | The prefix for a group of log group names. `region_name` is required when `log_group_name_prefix` is given. `log_group_name` and `log_group_name_prefix` cannot be given at the same time. | `string` | `null` | no |
| <a name="input_log_stream_prefix"></a> [log\_stream\_prefix](#input\_log\_stream\_prefix) | A string to filter the results to include only log events from log streams that have names starting with this prefix. | `string` | `null` | no |
| <a name="input_log_streams"></a> [log\_streams](#input\_log\_streams) | A list of strings of log streams names that Filebeat collect log events from. | `list(string)` | `null` | no |
| <a name="input_number_of_workers"></a> [number\_of\_workers](#input\_number\_of\_workers) | The number of workers assigned to read from log groups. Each worker will read log events from one of the log groups matching `log_group_name_prefix`. For example, if `log_group_name_prefix` matches five log groups, then `number_of_workers` should be set to `5`. The default value is `1`. | `number` | `1` | no |
| <a name="input_pipeline"></a> [pipeline](#input\_pipeline) | The Ingest Node pipeline ID to be used by the integration. | `string` | `null` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original` | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_proxy_url"></a> [proxy\_url](#input\_proxy\_url) | URL to proxy connections in the form of http\[s\]://<user>:<password>@<server name/ip>:<port> | `string` | `null` | no |
| <a name="input_region_name"></a> [region\_name](#input\_region\_name) | Region that the specified log group or log group prefix belongs to. | `string` | `null` | no |
| <a name="input_role_arn"></a> [role\_arn](#input\_role\_arn) | n/a | `string` | `null` | no |
| <a name="input_scan_frequency"></a> [scan\_frequency](#input\_scan\_frequency) | This config parameter sets how often Filebeat checks for new log events from the specified log group. | `string` | `"1m"` | no |
| <a name="input_secret_access_key"></a> [secret\_access\_key](#input\_secret\_access\_key) | n/a | `string` | `null` | no |
| <a name="input_session_token"></a> [session\_token](#input\_session\_token) | n/a | `string` | `null` | no |
| <a name="input_shared_credential_file"></a> [shared\_credential\_file](#input\_shared\_credential\_file) | Directory of the shared credentials file | `string` | `null` | no |
| <a name="input_start_position"></a> [start\_position](#input\_start\_position) | Allows user to specify if this input should read log files from the beginning or from the end. | `string` | `"beginning"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded"<br>]</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->