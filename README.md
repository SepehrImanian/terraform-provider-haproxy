<p align="center">
  <a href="https://github.com/SepehrImanian/terraform-provider-haproxy">
    <img src="./assets/haproxy.png" alt="minio-provider-terraform" width="200">
  </a>
  <h1 align="center" style="font-weight: bold">Terraform Provider for HAProxy</h1>
  <p align="center">
    <a href="https://golang.org/doc/devel/release.html">
      <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/SepehrImanian/terraform-provider-haproxy?style=flat-square">
    </a>
    <a href="https://github.com/SepehrImanian/terraform-provider-haproxy/actions?query=workflow%3A%22Terraform+Provider+CI%22">
      <img alt="GitHub Workflow Status" src="https://img.shields.io/github/v/release/SepehrImanian/terraform-provider-haproxy?style=flat-square">
    </a>
    <a href="https://github.com/SepehrImanian/terraform-provider-haproxy/releases">
      <img alt="GitHub release (latest by date including pre-releases)" src="https://img.shields.io/github/license/SepehrImanian/terraform-provider-haproxy?style=flat-square">
    </a>
  </p>
  <p align="center">
    <a href="https://github.com/SepehrImanian/terraform-provider-haproxy/tree/master/docs"><strong>Explore the docs »</strong></a>
  </p>
</p>

> ⚠️ **Disclaimer:** This project is under **Active Development** and is not ready for production use. Please use at your own risk.

## Table of Contents
- [Table of Contents](#table-of-contents)
  - [About this project](#about-this-project)
  - [Usage](#usage)
  - [Examples](#examples)
  - [Building and Installing](#building-and-installing)
  - [License](#license)

### About this project

A [Terraform](https://www.terraform.io) provider to manage [HAProxy](https://www.haproxy.com/).

### Usage

See our [examples](./examples/) folder.

```hcl
terraform {
  required_providers {
    haproxy = {
      source  = "SepehrImanian/haproxy"
      version = "0.0.3"
    }
  }
}

provider "haproxy" {
  url      = "http://haproxy.example.com:8080"
  username = "username"
  password = "password"
}
```

### Examples

Manage all HAProxy with a single provider:

```hcl
resource "haproxy_global" "global" {
  user          = "haproxy"
  group         = "haproxy"
  chroot        = "/var/lib/haproxy"
  daemon        = true
  master_worker = true
  maxconn       = 2000
  pidfile       = "/var/run/haproxy.pid"
  ulimit_n      = 2000
  crt_base      = "/etc/ssl/certs"
  ca_base       = "/etc/ssl/private"
  stats_maxconn = 100
  stats_timeout = 60
}

resource "haproxy_defaults" "default_test" {
  name                    = "default_test"
  mode                    = "http"
  backlog                 = 10000
  httplog                 = true
  httpslog                = true
  tcplog                  = false
  retries                 = 3
  check_timeout           = 10
  client_timeout          = 10
  connect_timeout         = 10
  http_keep_alive_timeout = 10
  http_request_timeout    = 10
  queue_timeout           = 10
  server_timeout          = 9
  server_fin_timeout      = 10
  maxconn                 = 2000
}

resource "haproxy_acl" "acl_test" {
  name        = "acl_test"
  index       = 0
  parent_name = "backend_test"
  parent_type = "backend"
  criterion   = "hdr_dom(host)"
  value       = "example.com"
  depends_on  = [haproxy_backend.backend_test]
}

resource "haproxy_frontend" "front_test" {
  name                        = "front_test"
  backend                     = "backend_test"
  http_connection_mode        = "http-keep-alive"
  accept_invalid_http_request = true
  maxconn                     = 100
  mode                        = "http"
  backlog                     = 1000
  http_keep_alive_timeout     = 10
  http_request_timeout        = 10
  http_use_proxy_header       = true
  httplog                     = true
  httpslog                    = true
  tcplog                      = false

  compression {
    algorithms = ["gzip", "identity"]
    offload    = true
    types      = ["text/html", "text/plain", "text/css", "application/javascript"]
  }

  forwardfor {
    enabled = true
    # except  = ".example2.com"
    header = "X-Forwarded-For"
    ifnone = true
  }

  depends_on = [haproxy_backend.backend_test]
}

resource "haproxy_bind" "bind_test" {
  name        = "bind_test"
  port        = 8080
  address     = "0.0.0.0"
  parent_name = "front_test"
  parent_type = "frontend"
  maxconn     = 3000
  depends_on  = [haproxy_frontend.front_test]
}

resource "haproxy_backend" "backend_test" {
  name                 = "backend_test"
  mode                 = "http"
  http_connection_mode = "http-keep-alive"
  server_timeout       = 9
  check_timeout        = 10
  connect_timeout      = 10
  queue_timeout        = 10
  check_cache          = true

  balance {
    algorithm = "roundrobin"
  }

  httpchk_params {
    uri     = "/health"
    version = "HTTP/1.1"
    method  = "GET"
  }

  forwardfor {
    enabled = true
  }
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
  depends_on  = [haproxy_backend.backend_test]
}
```

### Building and Installing

Prebuilt versions of this provider are available on the [releases page](https://github.com/SepehrImanian/terraform-provider-haproxy/releases/latest).

But if you need to build it yourself, changes vars in **Makefile** then:

```bash
git clone https://github.com/SepehrImanian/terraform-provider-haproxy.git
cd terraform-provider-haproxy
make build
```

### License

Distributed under the Apache License. See [LICENSE](./LICENSE) for more information.

Made with <span style="color: #e25555;">&#9829;</span> using [Go](https://golang.org/).