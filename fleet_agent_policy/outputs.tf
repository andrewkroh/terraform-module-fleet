output "id" {
  value = elasticstack_fleet_agent_policy.agent_policy.id
}

output "enrollment_token" {
  description = "Token that can be used to enroll Agents into the policy."
  value       = data.elasticstack_fleet_enrollment_tokens.enrollment_token.tokens.0.api_key
  sensitive   = true
}