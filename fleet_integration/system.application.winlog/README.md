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
| <a name="input_custom_yaml"></a> [custom\_yaml](#input\_custom\_yaml) | YAML configuration options for winlog input. Be careful, this may break the integration. | `string` | `"# Winlog configuration example\n#batch_read_size: 100"` | no |
| <a name="input_event_id"></a> [event\_id](#input\_event\_id) | A list of included and excluded (blocked) event IDs. The value is a comma-separated list. The accepted values are single event IDs to include (e.g. 4624), a range of event IDs to include (e.g. 4700-4800), and single event IDs to exclude (e.g. -4735).  Limit 22 clauses, lower in some situations. See integration documentation for more details. | `string` | `null` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_force"></a> [fleet\_package\_policy\_force](#input\_fleet\_package\_policy\_force) | Force reinstallation of the package even if already installed. When true, bypasses "already installed" checks and triggers complete re-installation. This deletes and recreates Kibana assets (dashboards, visualizations), removes transforms and their destination indices, and overwrites ingest pipelines and templates. | `bool` | `true` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the system package to use. | `string` | `"2.12.0"` | no |
| <a name="input_ignore_older"></a> [ignore\_older](#input\_ignore\_older) | If this option is specified, events that are older than the specified amount of time are ignored. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". | `string` | `"72h"` | no |
| <a name="input_language"></a> [language](#input\_language) | The language ID the events will be rendered in. The language will be forced regardless of the system language. A complete list of language IDs can be found https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-lcid/a9eac961-e77d-41a6-90a5-ce1a8b0cdb9c[here]. It defaults to `0`, which indicates to use the system language. E.g.: 0x0409 for en-US | `string` | `0` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original XML event, added to the field `event.original` | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->