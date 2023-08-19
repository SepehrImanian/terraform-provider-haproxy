resource "haproxy_backend" "backend_test" {
  name = "backend_test"
  mode         = "http"
  balance_algorithm = "roundrobin"
}