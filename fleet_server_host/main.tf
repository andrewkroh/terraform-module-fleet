resource "elasticstack_fleet_server_host" "fleet_host" {
  name    = var.name
  default = false
  hosts = [
    var.fleet_url
  ]
}
