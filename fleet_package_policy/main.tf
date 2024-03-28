locals {
  policy_name = var.package_policy_name != null ? var.package_policy_name : "${var.package_name}-${var.namespace}"

  unused_data_streams = [for data_stream in var.all_data_streams :
    data_stream if data_stream != var.data_stream
  ]
  unused_policy_template_inputs = [for policy_template_input in var.all_policy_template_inputs :
    policy_template_input if policy_template_input != "${var.policy_template}-${var.input_type}"
  ]

  # NOTE: This is a really painful part of Fleet. If package contains an input or stream that is enabled by default
  # and it contains required variables, then you must explicitly disable the streams in order to pass
  # validation. Beats had the same issue originally that made doing IaC hard (see
  # https://github.com/elastic/beats/issues/17256). Any time a package adds a new data stream that is
  # enabled by default AND contains required variables then upgrades become breaking changes.
  disabled_stream_config = { for data_stream in local.unused_data_streams :
    "${var.package_name}.${data_stream}" => {
      enabled = false
    }
  }
  disabled_inputs_config = { for policy_template_input in local.unused_policy_template_inputs :
    policy_template_input => {
      enabled = false
    }
  }
}

// Control the version of the installed package. Fleet will not allow
// downgrades while there are package policies using the package. Only one
// version of any package may be installed at one time.
resource "elasticstack_fleet_integration" "assets" {
  name    = var.package_name
  version = var.package_version
  force   = true
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
    inputs = merge({
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
    }, local.disabled_inputs_config)
  })

  depends_on = [
    elasticstack_fleet_integration.assets,
  ]
}
