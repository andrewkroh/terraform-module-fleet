// This is used for making API calls to Kibana to control Fleet.
provider "restapi" {
  uri                  = var.kibana_url
  username             = var.kibana_username
  password             = var.kibana_password
  insecure             = var.kibana_insecure
  write_returns_object = true
  headers = {
    kbn-xsrf = true
  }
}

// Create a new empty Fleet agent policy that will be used for running the GitHub integration.
module "agent_policy_github" {
  source      = "../../fleet_agent_policy"
  name        = "GitHub Integration Example"
  description = "Ingest issues from GitHub."
}

// Add GitHub Issues data stream to the agent policy.
module "integration_github_issues" {
  source                = "../../fleet_integration/github.issues.httpjson"
  fleet_agent_policy_id = module.agent_policy_github.id
  access_token          = var.github_token
  owner                 = "elastic"
}