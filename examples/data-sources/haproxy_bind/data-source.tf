data "haproxy_bind" "bind_test" {
  name = "bind_test"
  parent_name = "front_test"
  parent_type = "frontend"
  depends_on = [ haproxy_bind.bind_test ]
}