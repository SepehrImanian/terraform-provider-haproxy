provider "haproxy" {
  url         = "http://103.75.196.148:5555"
  username    = "admin"
  password    = "adminpwd"
}

## Resources

resource "haproxy_backend" "test" {
  backend_name = "test"
  mode         = "http"
  balance_algorithm = "roundrobin"
}

resource "haproxy_backend" "test2" {
  backend_name = "test2"
  mode         = "http"
  balance_algorithm = "leastconn"
}
