data "haproxy_httpcheck" "httpcheck_test" {
  parent_name = haproxy_backend.backend_test.name
  parent_type = "backend"
  index       = 0
  type        = "send"
  depends_on  = [haproxy_backend.backend_test]
}
