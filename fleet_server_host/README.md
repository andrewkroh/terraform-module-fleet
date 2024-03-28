<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_elasticstack"></a> [elasticstack](#requirement\_elasticstack) | >= 0.11.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_elasticstack"></a> [elasticstack](#provider\_elasticstack) | >= 0.11.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [elasticstack_fleet_server_host.fleet_host](https://registry.terraform.io/providers/elastic/elasticstack/latest/docs/resources/fleet_server_host) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_fleet_url"></a> [fleet\_url](#input\_fleet\_url) | Fleet URL | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Name of the Fleet host. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | n/a |
<!-- END_TF_DOCS -->