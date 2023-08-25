data "haproxy_frontend" "front_test" {
  name       = "front_test"
  depends_on = [haproxy_frontend.front_test]
}
