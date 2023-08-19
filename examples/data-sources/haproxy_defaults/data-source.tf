data "haproxy_defaults" "default_test" {
  name = "default_test"
  depends_on = [ haproxy_defaults.default_test ]
}