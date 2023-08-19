resource "haproxy_defaults" "default_test" {
  name = "default_test"
  mode = "http"
  backlog = 10000
  httplog = true
  httpslog = true
  tcplog = false
  retries = 3
  check_timeout = 10
  client_timeout = 10
  connect_timeout = 10
  http_keep_alive_timeout = 10
  http_request_timeout = 10
  queue_timeout = 10
  server_timeout = 9
  server_fin_timeout = 10
  maxconn = 2000
}