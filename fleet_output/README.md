<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_elasticstack"></a> [elasticstack](#requirement\_elasticstack) | ~> 0.6.2 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_elasticstack"></a> [elasticstack](#provider\_elasticstack) | ~> 0.6.2 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [elasticstack_fleet_output.elasticsearch_output](https://registry.terraform.io/providers/elastic/elasticstack/latest/docs/resources/fleet_output) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_elasticsearch_url"></a> [elasticsearch\_url](#input\_elasticsearch\_url) | Elasticsearch URL | `string` | n/a | yes |
| <a name="input_insecure"></a> [insecure](#input\_insecure) | Disable TLS verification. | `bool` | `false` | no |
| <a name="input_name"></a> [name](#input\_name) | Name of output. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | n/a |
<!-- END_TF_DOCS -->