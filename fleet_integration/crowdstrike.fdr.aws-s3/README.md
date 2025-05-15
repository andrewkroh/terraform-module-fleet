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
| <a name="input_api_timeout"></a> [api\_timeout](#input\_api\_timeout) | The maximum duration of AWS API can take. The maximum is half of the visibility timeout value. Valid time units are h, m, s. | `string` | `null` | no |
| <a name="input_credential_profile_name"></a> [credential\_profile\_name](#input\_credential\_profile\_name) | n/a | `string` | `null` | no |
| <a name="input_default_region"></a> [default\_region](#input\_default\_region) | Default region to use prior to connecting to region specific services/endpoints if no AWS region is set from environment variable, credentials or instance profile. If none of the above are set and no default region is set as well, `us-east-1` is used. A region, either from environment variable, credentials or instance profile or from this default region setting, needs to be set when using regions in non-regular AWS environments such as AWS China or US Government Isolated. | `string` | `""` | no |
| <a name="input_enable_deduplication"></a> [enable\_deduplication](#input\_enable\_deduplication) | If data deduplication is enabled, it ensures no duplicate events are received. A [fingerprint processor](https://www.elastic.co/guide/en/elasticsearch/reference/current/fingerprint-processor.html) is used to calculate hash from a set of crowdstrike fields that uniquely defines the event. | `bool` | `false` | no |
| <a name="input_endpoint"></a> [endpoint](#input\_endpoint) | URL of the entry point for an AWS web service | `string` | `""` | no |
| <a name="input_enrich_metadata"></a> [enrich\_metadata](#input\_enrich\_metadata) | Uses data in aidmaster and userinfo to add host and user information to events. The aidmaster blob must contain the string "aidmaster" in its path and the userinfo blob path must contain "userinfo", and the FDR Notification Parsing Script must sort events so that aidmaster and userinfo events appear first in the stream. | `bool` | `true` | no |
| <a name="input_fdr_parsing_script"></a> [fdr\_parsing\_script](#input\_fdr\_parsing\_script) | The JS script used to parse the custom format of SQS FDR notifications. | `string` | `"function parse(n) {\n  var m = JSON.parse(n);\n  var evts = [];\n  var files = m.files;\n  var bucket = m.bucket;\n  if (!Array.isArray(files) || (files.length == 0) || bucket == null || bucket == \"\") {\n    return evts;\n  }\n  files.sort(function(a, b) {\n    var isMetadata = function(a) {\n      return a.path && ((a.path.indexOf(\"aidmaster\") !== -1) || (a.path.indexOf(\"userinfo\") !== -1));\n    };\n    var cmp = function(a, b) {\n      if (a < b) {\n        return -1;\n      }\n      if (a > b) {\n        return 1;\n      }\n      return 0;\n    };\n    if (isMetadata(a) === isMetadata(b)) {\n      return cmp(a.path, b.path);\n    }\n    if (isMetadata(a)) {\n      return -1;\n    }\n    return 1;\n  });\n  files.forEach(function(f){\n    var evt = new S3EventV2();\n    evt.SetS3BucketName(bucket);\n    evt.SetS3ObjectKey(f.path);\n    evts.push(evt);\n  });\n  return evts;\n}\n"` | no |
| <a name="input_fips_enabled"></a> [fips\_enabled](#input\_fips\_enabled) | Enabling this option changes the service name from `s3` to `s3-fips` for connecting to the correct service endpoint. | `bool` | `false` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the crowdstrike package to use. | `string` | `"1.67.0"` | no |
| <a name="input_is_fdr_queue"></a> [is\_fdr\_queue](#input\_is\_fdr\_queue) | By default the FDR queue is expected. This option must be set to `false` if you are using your own queue. | `bool` | `true` | no |
| <a name="input_keep_metadata"></a> [keep\_metadata](#input\_keep\_metadata) | Keep the aidmaster and userinfo documents after they have been used for event enrichment. | `bool` | `false` | no |
| <a name="input_long_fields"></a> [long\_fields](#input\_long\_fields) | Choose to `Index` or `Delete` long fields. Fields longer than 1024 bytes (except `event.original`) will be kept (indexed) or deleted based on this option. | `string` | `"index_long_fields"` | no |
| <a name="input_long_fields_max_length"></a> [long\_fields\_max\_length](#input\_long\_fields\_max\_length) | The maximum length of fields (in bytes) to consider them as too long. By default, fields larger than `1024` bytes are considered too long. This option in addition to `Long Fields` option helps users configure how integration should handle long fields. | `number` | `1024` | no |
| <a name="input_max_number_of_messages"></a> [max\_number\_of\_messages](#input\_max\_number\_of\_messages) | Deprecated in agent version 8.16.0, this parameter is ignored if present, use `Number of Workers` instead. The maximum number of SQS messages that can be in flight at any time. | `number` | `5` | no |
| <a name="input_metadata_cache_capacity"></a> [metadata\_cache\_capacity](#input\_metadata\_cache\_capacity) | The maximum amount of metadata objects to cache. Operations that would cause the capacity to be exceeded will result in evictions of the oldest elements. The capacity should not be lower than the number of elements that are expected to be referenced when processing the input as evicted elements are lost. Values at or below zero indicate no limit. <br>WARNING: This setting needs to be set only if the amount of metadata elements is known beforehand, otherwise it might lead to enrichment data loss. If you are not sure, leave it untouched. | `string` | `0` | no |
| <a name="input_metadata_cache_write_interval"></a> [metadata\_cache\_write\_interval](#input\_metadata\_cache\_write\_interval) | The interval between periodic cache writes to the backing file. Valid time units are h, m, s, ms, us/µs and ns. The contents are always written out to the backing file when the processor is closed. Default is zero, no periodic writes. | `string` | `0` | no |
| <a name="input_metadata_ttl"></a> [metadata\_ttl](#input\_metadata\_ttl) | The period of time that metadata is considered valid for. Valid time units are h, m, s, ms, us/µs and ns. | `string` | `"168h"` | no |
| <a name="input_number_of_workers"></a> [number\_of\_workers](#input\_number\_of\_workers) | Number of workers that will process the S3 or SQS objects listed. | `number` | `5` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original` | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_proxy_url"></a> [proxy\_url](#input\_proxy\_url) | URL to proxy connections in the form of http\[s\]://<user>:<password>@<server name/ip>:<port> | `string` | `null` | no |
| <a name="input_prune_fields"></a> [prune\_fields](#input\_prune\_fields) | Prune fields deletes fields that are less likely to be useful. This includes `agent.ephemeral_id`, `ecs.version`, `event.timezone` and `log.offset`. | `bool` | `true` | no |
| <a name="input_queue_url"></a> [queue\_url](#input\_queue\_url) | URL of the AWS SQS queue that messages will be received from. | `string` | n/a | yes |
| <a name="input_role_arn"></a> [role\_arn](#input\_role\_arn) | n/a | `string` | `null` | no |
| <a name="input_secret_access_key"></a> [secret\_access\_key](#input\_secret\_access\_key) | n/a | `string` | `null` | no |
| <a name="input_session_token"></a> [session\_token](#input\_session\_token) | n/a | `string` | `null` | no |
| <a name="input_shared_credential_file"></a> [shared\_credential\_file](#input\_shared\_credential\_file) | Directory of the shared credentials file | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded",<br>  "crowdstrike-fdr"<br>]</pre> | no |
| <a name="input_visibility_timeout"></a> [visibility\_timeout](#input\_visibility\_timeout) | The duration that the received messages are hidden from subsequent retrieve requests after being retrieved by a ReceiveMessage request. The maximum is 12 hours. Valid time units are h, m, s. | `string` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->