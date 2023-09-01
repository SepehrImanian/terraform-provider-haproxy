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

resource "haproxy_server" "server_test_1" {
  name        = "server_test_1"
  port        = 8080
  address     = "172.16.13.13"
  parent_name = "backend_test_acl"
  parent_type = "backend"
  depends_on  = [haproxy_backend.backend_test_acl]
}

resource "haproxy_server" "server_test_2" {
  name        = "server_test_2"
  port        = 8080
  address     = "172.16.13.14"
  parent_name = "backend_test_acl"
  parent_type = "backend"
  depends_on  = [haproxy_backend.backend_test_acl]
}

resource "haproxy_server" "server_test_3" {
  name        = "server_test_3"
  port        = 8080
  address     = "172.16.13.15"
  parent_name = "backend_test_acl"
  parent_type = "backend"
  depends_on  = [haproxy_backend.backend_test_acl]
}
