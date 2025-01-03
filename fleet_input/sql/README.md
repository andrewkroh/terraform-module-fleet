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
| <a name="input_condition"></a> [condition](#input\_condition) | Condition to filter when to apply this datastream. Refer to [Conditions](https://www.elastic.co/guide/en/fleet/current/dynamic-input-configuration.html#conditions) on how to use the available keys in conditions. | `string` | `null` | no |
| <a name="input_data_stream_dataset"></a> [data\_stream\_dataset](#input\_data\_stream\_dataset) | Set the name for your dataset. Once selected a dataset cannot be changed without creating a new integration policy. You can't use - in the name of a dataset and only valid characters for Elasticsearch index names are permitted. | `string` | `"sql"` | no |
| <a name="input_driver"></a> [driver](#input\_driver) | n/a | `string` | `"mysql"` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the sql package to use. | `string` | `"0.5.2"` | no |
| <a name="input_hosts"></a> [hosts](#input\_hosts) | n/a | `list(string)` | <pre>[<br>  "root:test@tcp(127.0.0.1:3306)/"<br>]</pre> | no |
| <a name="input_merge_results"></a> [merge\_results](#input\_merge\_results) | Merge results from multiple queries to a single event (restrictions apply) | `bool` | `false` | no |
| <a name="input_period"></a> [period](#input\_period) | n/a | `string` | `"10s"` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the events are shipped. See [Processors](https://www.elastic.co/guide/en/fleet/current/elastic-agent-processor-configuration.html) for details. | `string` | `null` | no |
| <a name="input_sql_queries_yaml"></a> [sql\_queries\_yaml](#input\_sql\_queries\_yaml) | n/a | `string` | `"- query: SHOW GLOBAL STATUS LIKE 'Innodb_system%'\n  response_format: variables\n        \n"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->