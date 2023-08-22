resource "haproxy_cache" "cash_test" {
  name = "cash_test"
  max_age = 3600
  max_object_size = 100000
  max_secondary_entries = 10000
  process_vary  = true
  total_max_size = 112
}