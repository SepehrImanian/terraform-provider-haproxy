resource "haproxy_frontend" "front_test" {
  name = "front_test"
  backend = "backend_test"
  http_connection_mode = "http-keep-alive"
  max_connection = 3000
  mode = "tcp"
  depends_on = [
    haproxy_backend.backend_test
  ]
}

resource "haproxy_backend" "backend_test" {
  name = "backend_test"
  mode         = "tcp"
  balance_algorithm = "roundrobin"
}

resource "haproxy_server" "server_test" {
  name        = "server_test"
  port        = 8080
  address     = "172.16.13.15"
  parent_name = "backend_test"
  parent_type = "backend"
  send_proxy  = true
  check       = true
  inter       = 3
  rise        = 3
  fall        = 3
  depends_on = [
    haproxy_backend.backend_test
  ]
}

data "haproxy_backend" "backend_test" {
  name = "backend_test"
}

output "haproxy_backend" {
  value = haproxy_backend.backend_test
}