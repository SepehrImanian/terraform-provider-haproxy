data "haproxy_cache" "cash_test" {
  name       = "cash_test"
  depends_on = [haproxy_cache.cash_test]
}
