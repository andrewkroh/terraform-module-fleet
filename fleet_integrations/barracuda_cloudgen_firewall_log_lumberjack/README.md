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
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_description"></a> [fleet\_data\_stream\_description](#input\_fleet\_data\_stream\_description) | Description to use for the data stream. | `string` | `""` | no |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the barracuda\_cloudgen\_firewall package to use. | `string` | `"1.3.0"` | no |
| <a name="input_fleet_policy_name_suffix"></a> [fleet\_policy\_name\_suffix](#input\_fleet\_policy\_name\_suffix) | Suffix to use at the end of the generated Policy Name. | `string` | `""` | no |
| <a name="input_listen_address"></a> [listen\_address](#input\_listen\_address) | The bind address to listen for TCP connections. Set to `0.0.0.0` to bind to all available interfaces. | `string` | `"localhost"` | no |
| <a name="input_listen_port"></a> [listen\_port](#input\_listen\_port) | The TCP port number to listen on. | `number` | `5044` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original` | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the logs are parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_ssl_yaml"></a> [ssl\_yaml](#input\_ssl\_yaml) | Options for enabling TLS mode. See the [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/configuration-ssl.html) for a list of all options. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "barracuda_cloudgen_firewall-log",<br>  "forwarded"<br>]</pre> | no |
| <a name="input_versions"></a> [versions](#input\_versions) | n/a | `list(string)` | <pre>[<br>  "v2"<br>]</pre> | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->