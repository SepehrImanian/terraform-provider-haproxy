resource "haproxy_backend" "backend_test" {
  name = "backend_test"
  mode = "http"

  balance {
    algorithm = "roundrobin"
  }
}

resource "haproxy_server_template" "server_template_test" {
  prefix       = "srv"
  backend      = haproxy_backend.backend_test.name
  fqdn         = "google.com"
  port         = 80
  num_or_range = "1-3"
}

data "haproxy_server_template" "server_template_test" {
  backend    = haproxy_backend.backend_test.name
  prefix     = "srv"
  depends_on = [haproxy_server_template.server_template_test]
}

output "server_template_test" {
  value = data.haproxy_server_template.server_template_test
}
