resource "haproxy_frontend" "front_test" {
  name                        = "front_test"
  backend                     = "backend_test"
  http_connection_mode        = "http-keep-alive"
  accept_invalid_http_request = true
  maxconn                     = 100
  mode                        = "http"
  backlog                     = 1000
  http_keep_alive_timeout     = 10
  http_request_timeout        = 10
  http_use_proxy_header       = true
  httplog                     = true
  httpslog                    = true
  tcplog                      = false

  compression {
    algorithms = ["gzip", "identity"]
    offload    = true
    types      = ["text/html", "text/plain", "text/css", "application/javascript"]
  }

  forwardfor {
    enabled = true
    header  = "X-Forwarded-For"
    ifnone  = true
  }

  depends_on = [haproxy_backend.backend_test]
}
