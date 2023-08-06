terraform {
  required_providers {
    haproxy = {
      version = "~> 1.0.0"
      source  = "terraform-example.com/haproxy-provider/haproxy"
    }
  }
}

provider "haproxy" {
  haproxy_server = "103.75.196.148:5555"
  haproxy_user    = "admin"
  haproxy_password    = "adminpwd"
  haproxy_insecure    = true
}

## Resources
resource "haproxy_backend" "test" {
  backend_name = "test"
  mode         = "http"
  balance_algorithm = "roundrobin"
}
