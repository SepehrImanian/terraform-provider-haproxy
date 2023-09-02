data "haproxy_nameserver" "nameserver_1" {
  name     = haproxy_nameserver.nameserver_1.name
  resolver = haproxy_resolvers.resolvers_test.name
}
