data "haproxy_filter" "filter_test_data" {
  parent_name = haproxy_backend.backend_test.name
  parent_type = "backend"
  index       = 0
  name        = "something"
  depends_on  = [haproxy_filter.filter_test]
}
