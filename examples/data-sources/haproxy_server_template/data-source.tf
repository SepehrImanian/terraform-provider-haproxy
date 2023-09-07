data "haproxy_server_template" "server_template_test" {
  backend    = haproxy_backend.backend_test.name
  prefix     = "srv"
  depends_on = [haproxy_server_template.server_template_test]
}
