resource "haproxy_backend" "backend_test" {
  name = "backend_test"
  mode = "http"

  balance {
    algorithm = "roundrobin"
  }
}

resource "haproxy_filter" "filter_test" {
  parent_name       = haproxy_backend.backend_test.name
  parent_type       = "backend"
  index             = 0
  name              = "something"
  type              = "trace"
  trace_rnd_parsing = false
}

data "haproxy_filter" "filter_test_data" {
  parent_name = haproxy_backend.backend_test.name
  parent_type = "backend"
  index       = 0
  name        = "something"
  depends_on  = [haproxy_filter.filter_test]
}

output "filter_test_data" {
  value = data.haproxy_filter.filter_test_data
}
