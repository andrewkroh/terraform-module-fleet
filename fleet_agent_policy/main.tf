resource "elasticstack_fleet_agent_policy" "agent_policy" {
  name                 = var.name
  description          = var.description
  namespace            = var.namespace
  monitoring_output_id = var.monitoring_output_id
  data_output_id       = var.data_output_id
  fleet_server_host_id = var.fleet_server_host_id
  monitor_logs         = var.monitor_logs
  monitor_metrics      = var.monitor_metrics
  skip_destroy         = var.skip_destroy
}

data "elasticstack_fleet_enrollment_tokens" "enrollment_token" {
  policy_id = elasticstack_fleet_agent_policy.agent_policy.id
}
