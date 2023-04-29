<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_restapi"></a> [restapi](#requirement\_restapi) | >= 1.18.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_restapi"></a> [restapi](#provider\_restapi) | >= 1.18.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [restapi_object.elasticsearch_output](https://registry.terraform.io/providers/Mastercard/restapi/latest/docs/resources/object) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_elasticsearch_url"></a> [elasticsearch\_url](#input\_elasticsearch\_url) | Elasticsearch URL | `string` | n/a | yes |
| <a name="input_insecure"></a> [insecure](#input\_insecure) | Disable TLS verification. | `bool` | `false` | no |
| <a name="input_name"></a> [name](#input\_name) | Name of output. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_output_id"></a> [output\_id](#output\_output\_id) | n/a |
<!-- END_TF_DOCS -->