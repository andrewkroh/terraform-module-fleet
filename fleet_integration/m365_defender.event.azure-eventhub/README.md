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
| <a name="input_connection_string"></a> [connection\_string](#input\_connection\_string) | The connection string required to communicate with Event Hubs. See [Get an Event Hubs connection string](https://docs.microsoft.com/en-us/azure/event-hubs/event-hubs-get-connection-string) to learn more. | `string` | n/a | yes |
| <a name="input_consumer_group"></a> [consumer\_group](#input\_consumer\_group) | We recommend using a dedicated consumer group for the azure input. Reusing consumer groups among non-related consumers can cause unexpected behavior and possibly lost events. | `string` | `"$Default"` | no |
| <a name="input_eventhub"></a> [eventhub](#input\_eventhub) | Elastic recommends using one event hub for each integration. Visit [Create an event hub](https://docs.elastic.co/integrations/azure#create-an-event-hub) to learn more. Use event hub names up to 30 characters long to avoid compatibility issues. | `string` | n/a | yes |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the m365\_defender package to use. | `string` | `"5.5.0"` | no |
| <a name="input_preserve_duplicate_custom_fields"></a> [preserve\_duplicate\_custom\_fields](#input\_preserve\_duplicate\_custom\_fields) | Preserve m365\_defender.event fields that were copied to Elastic Common Schema (ECS) fields. | `bool` | `false` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original`. | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_resource_manager_endpoint"></a> [resource\_manager\_endpoint](#input\_resource\_manager\_endpoint) | By default we are using the azure public environment, to override, users can provide a specific resource manager endpoint in order to use a different azure environment. | `string` | `null` | no |
| <a name="input_storage_account"></a> [storage\_account](#input\_storage\_account) | The name of the storage account where the consumer group's state/offsets will be stored and updated. | `string` | n/a | yes |
| <a name="input_storage_account_container"></a> [storage\_account\_container](#input\_storage\_account\_container) | The storage account container where the integration stores the checkpoint data for the consumer group. It is an advanced option to use with extreme care. You MUST use a dedicated storage account container for each Azure log type. DO NOT REUSE the same container name for more than one Azure log type. See [Container Names](https://docs.microsoft.com/en-us/rest/api/storageservices/naming-and-referencing-containers--blobs--and-metadata#container-names) for details on naming rules from Microsoft. The integration generates a default container name if not specified. | `string` | `null` | no |
| <a name="input_storage_account_key"></a> [storage\_account\_key](#input\_storage\_account\_key) | The storage account key, this key will be used to authorize access to data in your storage account. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded",<br>  "m365_defender-event"<br>]</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->