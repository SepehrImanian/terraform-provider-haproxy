resource "haproxy_bind" "bind_test" {
  name        = "bind_test"
  port        = 8080
  address     = "0.0.0.0"
  parent_name = "front_test"
  parent_type = "frontend"
  maxconn = 3000
  depends_on = [ haproxy_frontend.front_test ]
}