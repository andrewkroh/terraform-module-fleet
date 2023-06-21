resource "elasticstack_fleet_output" "elasticsearch_output" {
  name = var.name
  type = "elasticsearch"
  config_yaml = yamlencode({
    ssl = {
      verification_mode = var.insecure ? "none" : "certificate"
    }
  })
  default_integrations = false
  default_monitoring   = false
  hosts = [
    var.elasticsearch_url
  ]
}
