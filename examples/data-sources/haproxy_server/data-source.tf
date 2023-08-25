data "haproxy_server" "server_test" {
  name        = "server_test"
  parent_name = "backend_test"
  parent_type = "backend"
  depends_on  = [haproxy_server.server_test]
}
