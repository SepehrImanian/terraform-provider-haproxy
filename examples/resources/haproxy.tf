data "haproxy_health" "test" {
  name = "test"
}

output "haproxy_health" {
  value = data.haproxy_health.test.health
}
