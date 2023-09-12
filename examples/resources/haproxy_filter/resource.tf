resource "haproxy_filter" "filter_test" {
  parent_name       = haproxy_backend.backend_test.name
  parent_type       = "backend"
  index             = 0
  name              = "something"
  type              = "trace"
  trace_rnd_parsing = true
}
