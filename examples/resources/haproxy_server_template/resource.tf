resource "haproxy_server_template" "server_template_test" {
  prefix       = "srv"
  backend      = haproxy_backend.backend_test.name
  fqdn         = "google.com"
  port         = 80
  num_or_range = "1-3"
}