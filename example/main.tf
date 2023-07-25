provider "haproxy" {
  server_addr = "10.100.0.130:5555"
  username    = "CHANGE_ME"
  password    = "CHANGE_ME"
  insecure    = true
}

resource "haproxy_frontend" "front-name" {
   name    = "backend-name"
   backend = "backend-name"
}

## Resources
resource "haproxy_backend" "backend-name" {}
resource "haproxy_server" "server-name" {}
resource "haproxy_global" "name" {}
resource "haproxy_defaults" "name" {}
resource "haproxy_acl" "acl-name" {}

## Datas
data "haproxy_backend" "backend-name" {}
data "haproxy_frontend" "front-name" {}
data "haproxy_server" "server-name" {}
data "haproxy_acl" "acl-name" {}