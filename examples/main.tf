provider "haproxy" {
  url         = "http://103.75.196.148:5555"
  username    = "admin"
  password    = "adminpwd"
}

## Resources
resource "haproxy_backend" "test_2" {
  backend_name = "test_2"
  mode         = "http"
  balance_algorithm = "roundrobin"
}

resource "haproxy_backend" "test_1" {
  backend_name = "test_1"
  mode         = "http"
  balance_algorithm = "roundrobin"
}

resource "haproxy_backend" "test_3" {
  backend_name = "test_3"
  mode         = "http"
  balance_algorithm = "roundrobin"
}