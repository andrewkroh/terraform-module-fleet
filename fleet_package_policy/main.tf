locals {
  policy_name = var.package_policy_name != null ? var.package_policy_name : "${var.package_name}-${var.namespace}"

  unused_data_streams = [for data_stream in var.all_data_streams :
    data_stream if data_stream != var.data_stream
  ]

  # NOTE: This is a really painful part of Fleet. If package contains a stream that is enabled by default
  # and it contains required variables, then you must explicitly disable the streams in order to pass
  # validation. Beats had the same issue originally that made doing IaC hard (see
  # https://github.com/elastic/beats/issues/17256). Any time a package adds a new data stream that is
  # enabled by default AND contains required variables then upgrades become breaking changes.
  disabled_stream_config = { for data_stream in local.unused_data_streams :
    "${var.package_name}.${data_stream}" => {
      enabled = false
    }
  }
}

resource "restapi_object" "package_policy" {
  path         = "/api/fleet/package_policies"
  id_attribute = "item/id"

  data = jsonencode({
    policy_id = var.agent_policy_id
    package = {
      name    = var.package_name
      version = var.package_version
    }
    name        = local.policy_name
    namespace   = var.namespace
    description = var.description
    vars        = var.package_variables_json
    inputs = {
      "${var.policy_template}-${var.input_type}" = {
        enabled = true
        vars    = var.input_variables_json == null ? null : jsondecode(var.input_variables_json)
        streams = merge({
          // Input packages don't have a data_stream value and use package_name.policy_template.
          "${var.package_name}.%{if var.data_stream != ""}${var.data_stream}%{else}${var.policy_template}%{endif}" = {
            enabled = true
            vars    = var.data_stream_variables_json == null ? null : jsondecode(var.data_stream_variables_json)
          }
        }, local.disabled_stream_config)
      }
    }
  })
}
