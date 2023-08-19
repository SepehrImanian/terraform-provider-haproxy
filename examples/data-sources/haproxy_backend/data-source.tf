data "haproxy_backend" "backend_test" {
  name = "backend_test"
  depends_on = [ haproxy_backend.backend_test ]
}