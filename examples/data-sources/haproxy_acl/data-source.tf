data "haproxy_acl" "acl_test" {
  name        = "acl_test"
  index       = 0
  parent_name = "backend_test"
  parent_type = "backend"
  depends_on  = [haproxy_acl.acl_test]
}
