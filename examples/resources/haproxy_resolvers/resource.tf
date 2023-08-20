resource "haproxy_resolvers" "resolvers_test" {
  name = "resolvers_test"
  accepted_payload_size = 8192
  hold_nx = 30
  hold_other = 30
  hold_refused = 30
  hold_timeout = 30
  hold_valid = 10
  parse_resolv_conf = true
  resolve_retries = 3
  timeout_resolve = 1
  timeout_retry = 1
}