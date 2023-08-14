terraform {
  required_providers {
    haproxy = {
      version = "~> 1.0.0"
      source  = "terraform-example.com/haproxy-provider/haproxy"
    }
  }
}

provider "haproxy" {
  url         = var.haproxy_url
  username    = var.haproxy_username
  password    = var.haproxy_password
}