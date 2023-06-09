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
| [restapi_object.fleet_host](https://registry.terraform.io/providers/Mastercard/restapi/latest/docs/resources/object) | resource |

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