terraform {
  required_providers {
    haproxy = {
      version = "~> 1.0.0"
      source  = "terraform-example.com/haproxy-provider/haproxy"
    }
  }
}

provider "haproxy" {
  url         = "http://haproxy.example.com:8080"
  username    = "username"
  password    = "password"
}