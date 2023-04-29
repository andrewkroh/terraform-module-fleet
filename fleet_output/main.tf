resource "restapi_object" "elasticsearch_output" {
  path         = "/api/fleet/outputs"
  id_attribute = "item/id"

  data = jsonencode({
    name                  = var.name
    type                  = "elasticsearch"
    is_default            = false
    is_default_monitoring = false
    hosts = [
      var.elasticsearch_url
    ]
    config_yaml = yamlencode({
      ssl = {
        verification_mode = var.insecure ? "none" : "certificate"
      }
    })
  })
}
