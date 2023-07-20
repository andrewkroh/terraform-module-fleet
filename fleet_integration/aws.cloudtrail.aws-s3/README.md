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
| <a name="input_bucket_arn"></a> [bucket\_arn](#input\_bucket\_arn) | Mandatory if the "Collect logs via S3 Bucket" switch is on. It is a required parameter for collecting logs via the AWS S3 Bucket. | `string` | `null` | no |
| <a name="input_bucket_list_prefix"></a> [bucket\_list\_prefix](#input\_bucket\_list\_prefix) | Prefix to apply for the list request to the S3 bucket. | `string` | `null` | no |
| <a name="input_cloudtrail_digest_regex"></a> [cloudtrail\_digest\_regex](#input\_cloudtrail\_digest\_regex) | Regex to match path of CloudTrail Digest S3 Objects.  If<br>blank CloudTrail Digest logs will be skipped. | `string` | `"/CloudTrail-Digest/"` | no |
| <a name="input_cloudtrail_insight_regex"></a> [cloudtrail\_insight\_regex](#input\_cloudtrail\_insight\_regex) | Regex to match path of CloudTrail Insight S3 Objects.  If<br>blank CloudTrail Insight logs will be skipped. | `string` | `"/CloudTrail-Insight/"` | no |
| <a name="input_cloudtrail_regex"></a> [cloudtrail\_regex](#input\_cloudtrail\_regex) | Regex to match path of CloudTrail S3 Objects.  If blank<br>CloudTrail logs will be skipped. | `string` | `"/CloudTrail/"` | no |
| <a name="input_collect_s3_logs"></a> [collect\_s3\_logs](#input\_collect\_s3\_logs) | To Collect logs via S3 bucket enable the toggle switch. By default, it will collect logs via SQS Queue. | `bool` | `false` | no |
| <a name="input_credential_profile_name"></a> [credential\_profile\_name](#input\_credential\_profile\_name) | n/a | `string` | `null` | no |
| <a name="input_default_region"></a> [default\_region](#input\_default\_region) | Default region to use prior to connecting to region specific services/endpoints if no AWS region is set from environment variable, credentials or instance profile. If none of the above are set and no default region is set as well, `us-east-1` is used. A region, either from environment variable, credentials or instance profile or from this default region setting, needs to be set when using regions in non-regular AWS environments such as AWS China or US Government Isolated. | `string` | `""` | no |
| <a name="input_fips_enabled"></a> [fips\_enabled](#input\_fips\_enabled) | Enabling this option changes the service name from `s3` to `s3-fips` for connecting to the correct service endpoint. | `bool` | `false` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the aws package to use. | `string` | `"1.50.0"` | no |
| <a name="input_interval"></a> [interval](#input\_interval) | Time interval for polling listing of the S3 bucket. NOTE: Supported units for this parameter are h/m/s. | `string` | `"1m"` | no |
| <a name="input_max_number_of_messages"></a> [max\_number\_of\_messages](#input\_max\_number\_of\_messages) | The maximum number of SQS messages that can be inflight at any time. | `number` | `5` | no |
| <a name="input_number_of_workers"></a> [number\_of\_workers](#input\_number\_of\_workers) | Number of workers that will process the S3 objects listed. | `number` | `5` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original` | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_proxy_url"></a> [proxy\_url](#input\_proxy\_url) | URL to proxy connections in the form of http\[s\]://<user>:<password>@<server name/ip>:<port> | `string` | `null` | no |
| <a name="input_queue_url"></a> [queue\_url](#input\_queue\_url) | Mandatory if the "Collect logs via S3 Bucket" switch is off. URL of the AWS SQS queue that messages will be received from. | `string` | `null` | no |
| <a name="input_role_arn"></a> [role\_arn](#input\_role\_arn) | n/a | `string` | `null` | no |
| <a name="input_secret_access_key"></a> [secret\_access\_key](#input\_secret\_access\_key) | n/a | `string` | `null` | no |
| <a name="input_session_token"></a> [session\_token](#input\_session\_token) | n/a | `string` | `null` | no |
| <a name="input_shared_credential_file"></a> [shared\_credential\_file](#input\_shared\_credential\_file) | Directory of the shared credentials file | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded",<br>  "aws-cloudtrail"<br>]</pre> | no |
| <a name="input_visibility_timeout"></a> [visibility\_timeout](#input\_visibility\_timeout) | The duration that the received messages are hidden from subsequent retrieve requests after being retrieved by a ReceiveMessage request.  The maximum is 12 hours. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->