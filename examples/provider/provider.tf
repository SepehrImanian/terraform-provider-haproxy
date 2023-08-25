terraform {
  required_providers {
    haproxy = {
      source  = "SepehrImanian/haproxy"
      version = "0.0.4"
    }
  }
}

provider "haproxy" {
  url      = "http://haproxy.example.com:8080"
  username = "username"
  password = "password"
}
