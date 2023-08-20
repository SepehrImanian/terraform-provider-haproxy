data "haproxy_resolvers" "resolvers_test" {
  name = "resolvers_test"
  depends_on = [ haproxy_resolvers.resolvers_test ]
}