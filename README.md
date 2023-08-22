# Terraform HAProxy Provider

> ⚠️ **Disclaimer:** This project is under **Active Development** and is not ready for production use. Please use at your own risk.

![GitHub release (latest by date)](https://img.shields.io/github/v/release/SepehrImanian/terraform-provider-haproxy?style=flat-square)


## About this project

A [Terraform](https://www.terraform.io) provider to manage [HAProxy](https://www.haproxy.com/).

## Usage

See our [examples](./examples/) folder.

```hcl
terraform {
  required_providers {
    haproxy = {
      source = "SepehrImanian/haproxy"
      version = "0.0.3"
    }
  }
}

provider "haproxy" {
  url         = "http://haproxy.example.com:8080"
  username    = "username"
  password    = "password"
}
```

## Examples

Manage all HAProxy with a single provider:

```hcl
resource "haproxy_global" "global" {
    user = "haproxy"
    group = "haproxy"
    chroot = "/var/lib/haproxy"
    daemon = true
    master_worker = true
    maxconn = 2000
    pidfile = "/var/run/haproxy.pid"
    ulimit_n = 2000
    crt_base = "/etc/ssl/certs"
    ca_base = "/etc/ssl/private"
    stats_maxconn = 100
    stats_timeout = 60
}

resource "haproxy_defaults" "default_test" {
  name = "default_test"
  mode = "http"
  backlog = 10000
  httplog = true
  httpslog = true
  tcplog = false
  retries = 3
  check_timeout = 10
  client_timeout = 10
  connect_timeout = 10
  http_keep_alive_timeout = 10
  http_request_timeout = 10
  queue_timeout = 10
  server_timeout = 9
  server_fin_timeout = 10
  maxconn = 2000
}

resource "haproxy_acl" "acl_test" {
  name        = "acl_test"
  index       = 0
  parent_name = "backend_test"
  parent_type = "backend"
  criterion   = "hdr_dom(host)"
  value       = "example.com"
  depends_on = [ haproxy_backend.backend_test ]
}

resource "haproxy_frontend" "front_test" {
  name = "front_test"
  backend = "backend_test"
  http_connection_mode = "http-keep-alive"
  max_connection = 3000
  mode = "http"
  depends_on = [ haproxy_backend.backend_test ]
}

resource "haproxy_bind" "bind_test" {
  name        = "bind_test"
  port        = 8080
  address     = "0.0.0.0"
  parent_name = "front_test"
  parent_type = "frontend"
  maxconn = 3000
  depends_on = [ haproxy_frontend.front_test ]
}

resource "haproxy_backend" "backend_test" {
  name = "backend_test"
  mode         = "http"
  balance_algorithm = "roundrobin"
}

resource "haproxy_server" "server_test" {
  name        = "server_test"
  port        = 8080
  address     = "172.16.13.15"
  parent_name = "backend_test"
  parent_type = "backend"
  send_proxy  = true
  check       = true
  inter       = 3
  rise        = 3
  fall        = 3
  depends_on = [ haproxy_backend.backend_test ]
}
```

## Building and Installing

Prebuilt versions of this provider are available on the [releases page](https://github.com/SepehrImanian/terraform-provider-haproxy/releases/latest).

But if you need to build it yourself, changes vars in **Makefile** then:

```bash
git clone https://github.com/SepehrImanian/terraform-provider-haproxy.git
cd terraform-provider-haproxy
make build
```

## License

Distributed under the Apache License. See [LICENSE](./LICENSE) for more information.

Made with <span style="color: #e25555;">&#9829;</span> using [Go](https://golang.org/).