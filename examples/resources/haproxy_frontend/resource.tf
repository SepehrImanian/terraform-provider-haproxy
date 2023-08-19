resource "haproxy_frontend" "front_test" {
  name = "front_test"
  backend = "backend_test"
  http_connection_mode = "http-keep-alive"
  max_connection = 3000
  mode = "http"
  depends_on = [ haproxy_backend.backend_test ]
}