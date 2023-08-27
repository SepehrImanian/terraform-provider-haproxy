# data "haproxy_health" "test" {
#   name = "test"
# }

# output "haproxy_health" {
#   value = data.haproxy_health.test.health
# }


resource "haproxy_backend" "backend_test" {
  name = "backend_test"
  mode = "http"
  balance {
    algorithm = "roundrobin"
  }
}

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

  depends_on = [haproxy_backend.backend_test]
}

