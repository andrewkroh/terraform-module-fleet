output "agent_policy_id" {
  value = module.agent_policy_github.id
}

output "fleet_enrollment_token" {
  sensitive = true
  value     = module.agent_policy_github.enrollment_token
}