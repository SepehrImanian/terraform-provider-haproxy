resource "haproxy_backend" "backend_test" {
  name                 = "backend_test"
  mode                 = "http"
  http_connection_mode = "http-keep-alive"
  server_timeout       = 9
  check_timeout        = 20
  connect_timeout      = 20
  queue_timeout        = 20
  tarpit_timeout       = 20
  tunnel_timeout       = 20
  check_cache          = true

  balance {
    algorithm = "roundrobin"
  }

  httpchk_params {
    uri     = "/health"
    version = "HTTP/1.1"
    method  = "GET"
  }

  forwardfor {
    enabled = true
  }
}
