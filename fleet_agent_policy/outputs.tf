output "id" {
  value = restapi_object.agent_policy.id
}

output "enrollment_token" {
  description = "Token that can be used to enroll Agents into the policy."
  value       = jsondecode(data.restapi_object.enrollment_token.api_response).item.api_key
  sensitive   = true
}