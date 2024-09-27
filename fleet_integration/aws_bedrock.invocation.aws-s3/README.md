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
| <a name="input_api_timeout"></a> [api\_timeout](#input\_api\_timeout) | The maximum duration of AWS API can take. The maximum is half of the visibility timeout value. | `string` | `null` | no |
| <a name="input_bucket_arn"></a> [bucket\_arn](#input\_bucket\_arn) | ARN of the AWS S3 bucket that will be polled for list operation. (Required when `queue_url` and `non_aws_bucket_name` are not set). | `string` | `null` | no |
| <a name="input_bucket_list_interval"></a> [bucket\_list\_interval](#input\_bucket\_list\_interval) | Time interval for polling listing of the S3 bucket. | `string` | `"120s"` | no |
| <a name="input_bucket_list_prefix"></a> [bucket\_list\_prefix](#input\_bucket\_list\_prefix) | Prefix to apply for the list request to the S3 bucket. | `string` | `null` | no |
| <a name="input_buffer_size"></a> [buffer\_size](#input\_buffer\_size) | The size in bytes of the buffer that each harvester uses when fetching a file. This only applies to non-JSON logs. | `string` | `null` | no |
| <a name="input_content_type"></a> [content\_type](#input\_content\_type) | A standard MIME type describing the format of the object data. This can be set to override the MIME type that was given to the object when it was uploaded. For example application/json. | `string` | `null` | no |
| <a name="input_credential_profile_name"></a> [credential\_profile\_name](#input\_credential\_profile\_name) | n/a | `string` | `null` | no |
| <a name="input_default_region"></a> [default\_region](#input\_default\_region) | Default region to use prior to connecting to region specific services/endpoints if no AWS region is set from environment variable, credentials or instance profile. If none of the above are set and no default region is set as well, `us-east-1` is used. A region, either from environment variable, credentials or instance profile or from this default region setting, needs to be set when using regions in non-regular AWS environments such as AWS China or US Government Isolated. | `string` | `""` | no |
| <a name="input_encoding"></a> [encoding](#input\_encoding) | The file encoding to use for reading data that contains international characters. This only applies to non-JSON logs. | `string` | `null` | no |
| <a name="input_endpoint"></a> [endpoint](#input\_endpoint) | URL of the entry point for an AWS web service | `string` | `""` | no |
| <a name="input_expand_event_list_from_field"></a> [expand\_event\_list\_from\_field](#input\_expand\_event\_list\_from\_field) | If the fileset using this input expects to receive multiple messages bundled under a specific field then the config option expand\_event\_list\_from\_field value can be assigned the name of the field. This setting will be able to split the messages under the group value into separate events. For example, CloudTrail logs are in JSON format and events are found under the JSON object "Records". | `string` | `null` | no |
| <a name="input_file_selectors_yaml"></a> [file\_selectors\_yaml](#input\_file\_selectors\_yaml) | If the SQS queue will have events that correspond to files that this integration shouldn’t process file\_selectors can be used to limit the files that are downloaded. This is a list of selectors which are made up of regex and expand\_event\_list\_from\_field options. The regex should match the S3 object key in the SQS message, and the optional expand\_event\_list\_from\_field is the same as the global setting. If file\_selectors is given, then any global expand\_event\_list\_from\_field value is ignored in favor of the ones specified in the file\_selectors. Regex syntax is the same as the Go language. Files that don’t match one of the regexes won’t be processed. content\_type, parsers, include\_s3\_metadata,max\_bytes, buffer\_size, and encoding may also be set for each file selector. | `list(string)` | `null` | no |
| <a name="input_fips_enabled"></a> [fips\_enabled](#input\_fips\_enabled) | Enabling this option changes the service name from `s3` to `s3-fips` for connecting to the correct service endpoint. | `bool` | `false` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the aws\_bedrock package to use. | `string` | `"0.9.1"` | no |
| <a name="input_include_s3_metadata"></a> [include\_s3\_metadata](#input\_include\_s3\_metadata) | This input can include S3 object metadata in the generated events for use in follow-on processing. You must specify the list of keys to include. By default none are included. If the key exists in the S3 response then it will be included in the event as aws.s3.metadata.<key> where the key name as been normalized to all lowercase. | `list(string)` | `null` | no |
| <a name="input_max_bytes"></a> [max\_bytes](#input\_max\_bytes) | The maximum number of bytes that a single log message can have. All bytes after max\_bytes are discarded and not sent. This setting is especially useful for multiline log messages, which can get large. This only applies to non-JSON logs. | `string` | `"10MiB"` | no |
| <a name="input_max_number_of_messages"></a> [max\_number\_of\_messages](#input\_max\_number\_of\_messages) | The maximum number of SQS messages that can be inflight at any time. | `number` | `5` | no |
| <a name="input_non_aws_bucket_name"></a> [non\_aws\_bucket\_name](#input\_non\_aws\_bucket\_name) | Name of the S3 bucket that will be polled for list operation. Required for 3rd party S3 compatible services. (Required when queue\_url and bucket\_arn are not set). | `string` | `null` | no |
| <a name="input_number_of_workers"></a> [number\_of\_workers](#input\_number\_of\_workers) | Number of workers that will process the S3 objects listed. (Required when `bucket_arn` is set). | `number` | `1` | no |
| <a name="input_path_style"></a> [path\_style](#input\_path\_style) | Enabling this option sets the bucket name as a path in the API call instead of a subdomain. When enabled https://<bucket-name>.s3.<region>.<provider>.com becomes https://s3.<region>.<provider>.com/<bucket-name>. This is only supported with 3rd party S3 providers.  AWS does not support path style. | `string` | `null` | no |
| <a name="input_preserve_duplicate_custom_fields"></a> [preserve\_duplicate\_custom\_fields](#input\_preserve\_duplicate\_custom\_fields) | Preserve Bedrock fields that were copied to Elastic Common Schema (ECS) fields. | `bool` | `false` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original` | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_provider_name"></a> [provider\_name](#input\_provider\_name) | Name of the 3rd party S3 bucket provider like backblaze or GCP. | `string` | `null` | no |
| <a name="input_proxy_url"></a> [proxy\_url](#input\_proxy\_url) | URL to proxy connections in the form of http\[s\]://<user>:<password>@<server name/ip>:<port> | `string` | `null` | no |
| <a name="input_queue_url"></a> [queue\_url](#input\_queue\_url) | URL of the AWS SQS queue that messages will be received from. | `string` | `null` | no |
| <a name="input_role_arn"></a> [role\_arn](#input\_role\_arn) | n/a | `string` | `null` | no |
| <a name="input_secret_access_key"></a> [secret\_access\_key](#input\_secret\_access\_key) | n/a | `string` | `null` | no |
| <a name="input_session_token"></a> [session\_token](#input\_session\_token) | n/a | `string` | `null` | no |
| <a name="input_shared_credential_file"></a> [shared\_credential\_file](#input\_shared\_credential\_file) | Directory of the shared credentials file | `string` | `null` | no |
| <a name="input_sqs_max_receive_count"></a> [sqs\_max\_receive\_count](#input\_sqs\_max\_receive\_count) | The maximum number of times a SQS message should be received (retried) before deleting it. This feature prevents poison-pill messages (messages that can be received but can’t be processed) from consuming resources. | `number` | `5` | no |
| <a name="input_sqs_wait_time"></a> [sqs\_wait\_time](#input\_sqs\_wait\_time) | The maximum duration that an SQS `ReceiveMessage` call should wait for a message to arrive in the queue before returning. The maximum value is `20s`. | `string` | `"20s"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded"<br>]</pre> | no |
| <a name="input_visibility_timeout"></a> [visibility\_timeout](#input\_visibility\_timeout) | The duration that the received messages are hidden from subsequent retrieve requests after being retrieved by a ReceiveMessage request.  The maximum is 12 hours. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->