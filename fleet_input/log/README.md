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
| <a name="input_custom_yaml"></a> [custom\_yaml](#input\_custom\_yaml) | Additional settings to be added to the configuration. Be careful using this as it might break the input as those settings are not validated and can override the settings specified above. See [`log` input settings docs](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-log.html) for details. | `string` | `""` | no |
| <a name="input_data_stream_dataset"></a> [data\_stream\_dataset](#input\_data\_stream\_dataset) | Set the name for your dataset. Changing the dataset will send the data to a different index. You can't use `-` in the name of a dataset and only valid characters for [Elasticsearch index names](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html). | `string` | n/a | yes |
| <a name="input_exclude_files"></a> [exclude\_files](#input\_exclude\_files) | Patterns to be ignored | `list(string)` | `null` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_force"></a> [fleet\_package\_policy\_force](#input\_fleet\_package\_policy\_force) | Force reinstallation of the package even if already installed. When true, bypasses "already installed" checks and triggers complete re-installation. This deletes and recreates Kibana assets (dashboards, visualizations), removes transforms and their destination indices, and overwrites ingest pipelines and templates. | `bool` | `true` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the log package to use. | `string` | `"2.4.4"` | no |
| <a name="input_ignore_older"></a> [ignore\_older](#input\_ignore\_older) | If this option is specified, events that are older than the specified amount of time are ignored. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". | `string` | `"72h"` | no |
| <a name="input_paths"></a> [paths](#input\_paths) | Path to log files to be collected | `list(string)` | n/a | yes |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Tags to include in the published event | `list(string)` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->