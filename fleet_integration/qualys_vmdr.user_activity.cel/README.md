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
| <a name="input_batch_size"></a> [batch\_size](#input\_batch\_size) | Batch size for the response of the Qualys Server API. This parameter specifies the truncation limit for the response. Specify 0 for no truncation limit. | `number` | `1000` | no |
| <a name="input_enable_request_tracer"></a> [enable\_request\_tracer](#input\_enable\_request\_tracer) | The request tracer logs HTTP requests and responses to the agent's local file-system for debugging configurations. Enabling this request tracing compromises security and should only be used for debugging. See [documentation](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-input-cel.html#_resource_tracer_filename) for details. | `bool` | `null` | no |
| <a name="input_fleet_agent_policy_id"></a> [fleet\_agent\_policy\_id](#input\_fleet\_agent\_policy\_id) | Agent policy ID to add the package policy to. | `string` | n/a | yes |
| <a name="input_fleet_data_stream_namespace"></a> [fleet\_data\_stream\_namespace](#input\_fleet\_data\_stream\_namespace) | Namespace to use for the data stream. | `string` | `"default"` | no |
| <a name="input_fleet_package_policy_description"></a> [fleet\_package\_policy\_description](#input\_fleet\_package\_policy\_description) | Description to use for the package policy. | `string` | `""` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the qualys\_vmdr package to use. | `string` | `"5.5.0"` | no |
| <a name="input_http_client_timeout"></a> [http\_client\_timeout](#input\_http\_client\_timeout) | Duration before declaring that the HTTP client connection has timed out. Give a timeout of more than 1 minute when retrieving data which is more than 15 days old. Supported time units are ns, us, ms, s, m, h. Requests may take significant time, so short timeouts are not recommended. | `string` | `"10m"` | no |
| <a name="input_initial_interval"></a> [initial\_interval](#input\_initial\_interval) | How far back to pull the User Activity log data from Qualys VMDR. Supported units for this parameter are s, m, h. | `string` | `"24h"` | no |
| <a name="input_interval"></a> [interval](#input\_interval) | Interval between two REST API calls. User can choose interval as per their plan mentioned in [Qualys API Limits](https://www.qualys.com/docs/qualys-api-limits.pdf). Supported units for this parameter are h/m/s. | `string` | `"4h"` | no |
| <a name="input_password"></a> [password](#input\_password) | Password for the Qualys VMDR. | `string` | n/a | yes |
| <a name="input_preserve_duplicate_custom_fields"></a> [preserve\_duplicate\_custom\_fields](#input\_preserve\_duplicate\_custom\_fields) | Preserve qualys\_vmdr.user\_activity fields that were copied to Elastic Common Schema (ECS) fields. | `bool` | `false` | no |
| <a name="input_preserve_original_event"></a> [preserve\_original\_event](#input\_preserve\_original\_event) | Preserves a raw copy of the original event, added to the field `event.original`. | `bool` | `false` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the data is parsed. See [Processors](https://www.elastic.co/guide/en/beats/filebeat/current/filtering-and-enhancing-data.html) for details. | `string` | `null` | no |
| <a name="input_proxy_url"></a> [proxy\_url](#input\_proxy\_url) | URL to proxy connections in the form of http[s]://<user>:<password>@<server name/ip>:<port>. Please ensure your username and password are in URL encoded format. | `string` | `null` | no |
| <a name="input_ssl_yaml"></a> [ssl\_yaml](#input\_ssl\_yaml) | i.e. certificate\_authorities, supported\_protocols, verification\_mode etc. | `string` | `"#certificate_authorities:\n#  - |\n#    -----BEGIN CERTIFICATE-----\n#    MIIDCjCCAfKgAwIBAgITJ706Mu2wJlKckpIvkWxEHvEyijANBgkqhkiG9w0BAQsF\n#    ADAUMRIwEAYDVQQDDAlsb2NhbGhvc3QwIBcNMTkwNzIyMTkyOTA0WhgPMjExOTA2\n#    MjgxOTI5MDRaMBQxEjAQBgNVBAMMCWxvY2FsaG9zdDCCASIwDQYJKoZIhvcNAQEB\n#    BQADggEPADCCAQoCggEBANce58Y/JykI58iyOXpxGfw0/gMvF0hUQAcUrSMxEO6n\n#    fZRA49b4OV4SwWmA3395uL2eB2NB8y8qdQ9muXUdPBWE4l9rMZ6gmfu90N5B5uEl\n#    94NcfBfYOKi1fJQ9i7WKhTjlRkMCgBkWPkUokvBZFRt8RtF7zI77BSEorHGQCk9t\n#    /D7BS0GJyfVEhftbWcFEAG3VRcoMhF7kUzYwp+qESoriFRYLeDWv68ZOvG7eoWnP\n#    PsvZStEVEimjvK5NSESEQa9xWyJOmlOKXhkdymtcUd/nXnx6UTCFgnkgzSdTWV41\n#    CI6B6aJ9svCTI2QuoIq2HxX/ix7OvW1huVmcyHVxyUECAwEAAaNTMFEwHQYDVR0O\n#    BBYEFPwN1OceFGm9v6ux8G+DZ3TUDYxqMB8GA1UdIwQYMBaAFPwN1OceFGm9v6ux\n#    8G+DZ3TUDYxqMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAG5D\n#    874A4YI7YUwOVsVAdbWtgp1d0zKcPRR+r2OdSbTAV5/gcS3jgBJ3i1BN34JuDVFw\n#    3DeJSYT3nxy2Y56lLnxDeF8CUTUtVQx3CuGkRg1ouGAHpO/6OqOhwLLorEmxi7tA\n#    H2O8mtT0poX5AnOAhzVy7QW0D/k4WaoLyckM5hUa6RtvgvLxOwA0U+VGurCDoctu\n#    8F4QOgTAWyh8EZIwaKCliFRSynDpv3JTUwtfZkxo6K6nce1RhCWFAsMvDZL8Dgc0\n#    yvgJ38BRsFOtkRuAGSf6ZUwTO8JJRRIFnpUzXflAnGivK9M13D5GEQMmIl6U9Pvk\n#    sxSmbIUfc2SGJGCJD4I=\n#    -----END CERTIFICATE-----\n"` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | n/a | `list(string)` | <pre>[<br>  "forwarded",<br>  "qualys_vmdr-user_activity"<br>]</pre> | no |
| <a name="input_url"></a> [url](#input\_url) | Base URL of the Qualys Server API. | `string` | n/a | yes |
| <a name="input_username"></a> [username](#input\_username) | Username for the Qualys VMDR. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->