resource "haproxy_defaults" "default_test" {
  name = "default_test"
  mode = "http"
  backlog = 10000
  httplog = true
  httpslog = true
  tcplog = false
  retries = 3
  check_timeout = 10
  client_timeout = 10
  connect_timeout = 10
  http_keep_alive_timeout = 10
  http_request_timeout = 10
  queue_timeout = 10
  server_timeout = 9
  server_fin_timeout = 10
  maxconn = 2000
}

resource "haproxy_frontend" "front_test" {
  name = "front_test"
  backend = "backend_test"
  http_connection_mode = "http-keep-alive"
  max_connection = 3000
  mode = "http"
  depends_on = [ haproxy_backend.backend_test ]
}

resource "haproxy_bind" "bind_test" {
  name        = "bind_test"
  port        = 8080
  address     = "0.0.0.0"
  parent_name = "front_test"
  parent_type = "frontend"
  maxconn = 3000
  depends_on = [ haproxy_frontend.front_test ]
}

resource "haproxy_backend" "backend_test" {
  name = "backend_test"
  mode         = "http"
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
  depends_on = [ haproxy_backend.backend_test ]
}

data "haproxy_backend" "backend_test" {
  name = "backend_test"
  depends_on = [ haproxy_backend.backend_test ]
}

data "haproxy_frontend" "front_test" {
  name = "front_test"
  depends_on = [ haproxy_frontend.front_test ]
}

data "haproxy_server" "server_test" {
  name = "server_test"
  parent_name = "backend_test"
  parent_type = "backend"
  depends_on = [ haproxy_server.server_test ]
}

data "haproxy_bind" "bind_test" {
  name = "bind_test"
  parent_name = "front_test"
  parent_type = "frontend"
  depends_on = [ haproxy_bind.bind_test ]
}

output "haproxy_backend" {
  value = haproxy_backend.backend_test
}

output "haproxy_frontend" {
  value = haproxy_frontend.front_test
}

output "haproxy_server" {
  value = haproxy_server.server_test
}

output "haproxy_bind" {
  value = haproxy_bind.bind_test
}