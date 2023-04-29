resource "restapi_object" "fleet_host" {
  path         = "/api/fleet/fleet_server_hosts"
  id_attribute = "item/id"

  data = jsonencode({
    name       = var.name
    is_default = false
    host_urls = [
      var.fleet_url,
    ]
  })
}
