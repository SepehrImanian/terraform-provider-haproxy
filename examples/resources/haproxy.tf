resource "haproxy_backend" "backend_test_acl" {
  name        = "backend_test_acl"
  mode        = "http"
  check_cache = true

  balance {
    algorithm = "roundrobin"
  }
}


resource "haproxy_frontend" "front_test" {
  name    = "front_test"
  backend = "backend_test_acl"
  mode    = "http"

  forwardfor {
    enabled = true
    header  = "X-Forwarded-For"
    ifnone  = true
  }

  depends_on = [haproxy_backend.backend_test_acl]
}
