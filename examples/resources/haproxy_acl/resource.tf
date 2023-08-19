resource "haproxy_acl" "acl_test" {
  name        = "acl_test"
  index       = 0
  parent_name = "backend_test"
  parent_type = "backend"
  criterion   = "hdr_dom(host)"
  value       = "example.com"
  depends_on = [ haproxy_backend.backend_test ]
}