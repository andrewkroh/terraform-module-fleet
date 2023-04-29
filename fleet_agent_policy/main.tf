resource "restapi_object" "agent_policy" {
  path         = "/api/fleet/agent_policies"
  id_attribute = "item/id"

  // BUG (andrewkroh, 2023-04-21): https://github.com/elastic/kibana/issues/155543
  // Cannot delete this policy using this resource type. This policy will be orphaned.
  destroy_path = "/api/fleet/enrollment_api_keys/delete"

  data = jsonencode({
    name        = var.name
    description = var.description
    namespace   = var.namespace
    monitoring_enabled = compact([
      var.monitor_logs ? "logs" : "",
      var.monitor_metrics ? "metrics" : "",
    ])
    inactivity_timeout   = var.inactivity_timeout_sec
    monitoring_output_id = var.monitoring_output_id
    data_output_id       = var.data_output_id
    fleet_server_host_id = var.fleet_server_host_id
  })
}

data "restapi_object" "enrollment_token" {
  path         = "/api/fleet/enrollment_api_keys"
  results_key  = "list"
  search_key   = "policy_id"
  search_value = restapi_object.agent_policy.id
}
