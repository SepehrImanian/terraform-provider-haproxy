resource "haproxy_backend" "backend_test" {
  name = "backend_test"
  mode = "http"

  balance {
    algorithm = "roundrobin"
  }
}

resource "haproxy_httpcheck" "httpcheck_test" {
  parent_name = haproxy_backend.backend_test.name
  parent_type = "backend"
  index       = 0
  type        = "send"
  method      = "GET"
  uri         = "/health"
  port        = 80

  headers {
    name = "Host"
    fmt  = "example.com"
  }

  headers {
    name = "User-Agent"
    fmt  = "Mozilla/5.0"
  }
}

data "haproxy_httpcheck" "httpcheck_test" {
  parent_name = haproxy_backend.backend_test.name
  parent_type = "backend"
  index       = 0
  type        = "send"
  depends_on  = [haproxy_backend.backend_test]
}

output "name" {
  value = data.haproxy_httpcheck.httpcheck_test
}
