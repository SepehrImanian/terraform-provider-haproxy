terraform {
  required_providers {
    haproxy = {
      version = "~> 1.0.0"
      source  = "terraform-example.com/haproxy-provider/haproxy"
    }
  }
}