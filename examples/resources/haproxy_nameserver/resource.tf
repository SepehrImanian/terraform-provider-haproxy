resource "haproxy_nameserver" "nameserver_1" {
  name     = "nameserver_1"
  resolver = haproxy_resolvers.resolvers_test.name
  address  = "192.168.1.3"
  port     = 53
}
