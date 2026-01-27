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
| <a name="input_fleet_package_policy_force"></a> [fleet\_package\_policy\_force](#input\_fleet\_package\_policy\_force) | Force reinstallation of the package even if already installed. When true, bypasses "already installed" checks and triggers complete re-installation. This deletes and recreates Kibana assets (dashboards, visualizations), removes transforms and their destination indices, and overwrites ingest pipelines and templates. | `bool` | `true` | no |
| <a name="input_fleet_package_policy_name_suffix"></a> [fleet\_package\_policy\_name\_suffix](#input\_fleet\_package\_policy\_name\_suffix) | Suffix to append to the end of the package policy name. | `string` | `""` | no |
| <a name="input_fleet_package_version"></a> [fleet\_package\_version](#input\_fleet\_package\_version) | Version of the sql package to use. | `string` | `"1.1.0"` | no |
| <a name="input_hosts"></a> [hosts](#input\_hosts) | n/a | `list(string)` | <pre>[<br>  "root:test@tcp(127.0.0.1:3306)/"<br>]</pre> | no |
| <a name="input_merge_results"></a> [merge\_results](#input\_merge\_results) | Merge results from multiple queries to a single event (restrictions apply) | `bool` | `false` | no |
| <a name="input_period"></a> [period](#input\_period) | n/a | `string` | `"10s"` | no |
| <a name="input_processors_yaml"></a> [processors\_yaml](#input\_processors\_yaml) | Processors are used to reduce the number of fields in the exported event or to enhance the event with metadata. This executes in the agent before the events are shipped. See [Processors](https://www.elastic.co/guide/en/fleet/current/elastic-agent-processor-configuration.html) for details. | `string` | `null` | no |
| <a name="input_sql_queries_yaml"></a> [sql\_queries\_yaml](#input\_sql\_queries\_yaml) | n/a | `string` | `"- query: SHOW GLOBAL STATUS LIKE 'Innodb_system%'\n  response_format: variables\n        \n"` | no |
| <a name="input_ssl_yaml"></a> [ssl\_yaml](#input\_ssl\_yaml) | SSL configuration options. See [documentation](https://www.elastic.co/docs/reference/integrations/sql) for details. | `string` | `"# certificate_authorities:\n#  - |\n#    -----BEGIN CERTIFICATE-----\n#    MIIDCjCCAfKgAwIBAgITJ706Mu2wJlKckpIvkWxEHvEyijANBgkqhkiG9w0BAQsF\n#    ADAUMRIwEAYDVQQDDAlsb2NhbGhvc3QwIBcNMTkwNzIyMTkyOTA0WhgPMjExOTA2\n#    MjgxOTI5MDRaMBQxEjAQBgNVBAMMCWxvY2FsaG9zdDCCASIwDQYJKoZIhvcNAQEB\n#    BQADggEPADCCAQoCggEBANce58Y/JykI58iyOXpxGfw0/gMvF0hUQAcUrSMxEO6n\n#    fZRA49b4OV4SwWmA3395uL2eB2NB8y8qdQ9muXUdPBWE4l9rMZ6gmfu90N5B5uEl\n#    94NcfBfYOKi1fJQ9i7WKhTjlRkMCgBkWPkUokvBZFRt8RtF7zI77BSEorHGQCk9t\n#    /D7BS0GJyfVEhftbWcFEAG3VRcoMhF7kUzYwp+qESoriFRYLeDWv68ZOvG7eoWnP\n#    PsvZStEVEimjvK5NSESEQa9xWyJOmlOKXhkdymtcUd/nXnx6UTCFgnkgzSdTWV41\n#    CI6B6aJ9svCTI2QuoIq2HxX/ix7OvW1huVmcyHVxyUECAwEAAaNTMFEwHQYDVR0O\n#    BBYEFPwN1OceFGm9v6ux8G+DZ3TUDYxqMB8GA1UdIwQYMBaAFPwN1OceFGm9v6ux\n#    8G+DZ3TUDYxqMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAG5D\n#    874A4YI7YUwOVsVAdbWtgp1d0zKcPRR+r2OdSbTAV5/gcS3jgBJ3i1BN34JuDVFw\n#    3DeJSYT3nxy2Y56lLnxDeF8CUTUtVQx3CuGkRg1ouGAHpO/6OqOhwLLorEmxi7tA\n#    H2O8mtT0poX5AnOAhzVy7QW0D/k4WaoLyckM5hUa6RtvgvLxOwA0U+VGurCDoctu\n#    8F4QOgTAWyh8EZIwaKCliFRSynDpv3JTUwtfZkxo6K6nce1RhCWFAsMvDZL8Dgc0\n#    yvgJ38BRsFOtkRuAGSf6ZUwTO8JJRRIFnpUzXflAnGivK9M13D5GEQMmIl6U9Pvk\n#    sxSmbIUfc2SGJGCJD4I=\n#    -----END CERTIFICATE-----\n"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | Package policy ID |
<!-- END_TF_DOCS -->