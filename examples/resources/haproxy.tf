# resource "haproxy_defaults" "default_test" {
#   name = "default_test"
#   mode = "http"
#   backlog = 1000
#   httplog = true
#   httpslog = true
#   tcplog = false
#   retries = 3
#   check_timeout = 10
#   client_timeout = 10
#   connect_timeout = 10
#   http_keep_alive_timeout = 10
#   http_request_timeout = 10
#   queue_timeout = 10
#   server_timeout = 9
#   server_fin_timeout = 10
#   maxconn = 1000
# }

# resource "haproxy_acl" "acl_test" {
#   name        = "acl_test"
#   index       = 0
#   parent_name = "backend_test"
#   parent_type = "backend"
#   criterion   = "hdr_dom(host)"
#   value       = "example.com"
#   depends_on = [ haproxy_backend.backend_test ]
# }

# resource "haproxy_frontend" "front_test" {
#   name = "front_test"
#   backend = "backend_test"
#   http_connection_mode = "http-keep-alive"
#   max_connection = 100
#   mode = "http"
#   depends_on = [ haproxy_backend.backend_test ]
# }

# resource "haproxy_bind" "bind_test" {
#   name        = "bind_test"
#   port        = 8080
#   address     = "0.0.0.0"
#   parent_name = "front_test"
#   parent_type = "frontend"
#   maxconn = 100
#   depends_on = [ haproxy_frontend.front_test ]
# }

# resource "haproxy_backend" "backend_test" {
#   name = "backend_test"
#   mode         = "http"
#   balance_algorithm = "roundrobin"
# }

# resource "haproxy_server" "server_test" {
#   name        = "server_test"
#   port        = 8080
#   address     = "172.16.13.15"
#   parent_name = "backend_test"
#   parent_type = "backend"
#   send_proxy  = true
#   check       = true
#   inter       = 3
#   rise        = 3
#   fall        = 3
#   depends_on = [ haproxy_backend.backend_test ]
# }

# data "haproxy_backend" "backend_test" {
#   name = "backend_test"
#   depends_on = [ haproxy_backend.backend_test ]
# }

# data "haproxy_frontend" "front_test" {
#   name = "front_test"
#   depends_on = [ haproxy_frontend.front_test ]
# }

# data "haproxy_server" "server_test" {
#   name = "server_test"
#   parent_name = "backend_test"
#   parent_type = "backend"
#   depends_on = [ haproxy_server.server_test ]
# }

# data "haproxy_bind" "bind_test" {
#   name = "bind_test"
#   parent_name = "front_test"
#   parent_type = "frontend"
#   depends_on = [ haproxy_bind.bind_test ]
# }

# data "haproxy_defaults" "default_test" {
#   name = "default_test"
#   depends_on = [ haproxy_defaults.default_test ]
# }

# data "haproxy_acl" "acl_test" {
#   name = "acl_test"
#   index = 0
#   parent_name = "backend_test"
#   parent_type = "backend"
#   depends_on = [ haproxy_acl.acl_test ]
# }

# output "haproxy_backend" {
#   value = haproxy_backend.backend_test
# }

# output "haproxy_frontend" {
#   value = haproxy_frontend.front_test
# }

# output "haproxy_server" {
#   value = haproxy_server.server_test
# }

# output "haproxy_bind" {
#   value = haproxy_bind.bind_test
# }

# output "haproxy_defaults" {
#   value = haproxy_defaults.default_test
# }

# output "haproxy_acl" {
#   value = haproxy_acl.acl_test
# }

# resource "haproxy_resolvers" "resolvers_test" {
#   name = "resolvers_test"
#   accepted_payload_size = 8192
#   hold_nx = 30
#   hold_other = 30
#   hold_refused = 30
#   hold_timeout = 30
#   hold_valid = 10
#   parse_resolv_conf = true
#   resolve_retries = 3
#   timeout_resolve = 1
#   timeout_retry = 1
# }

# data "haproxy_resolvers" "resolvers_test" {
#   name = "resolvers_test"
#   depends_on = [ haproxy_resolvers.resolvers_test ]
# }


# resource "haproxy_cache" "cash_test" {
#   name = "cash_test"
#   max_age = 3600
#   max_object_size = 100000
#   max_secondary_entries = 10000
#   process_vary  = true
#   total_max_size = 112
# }

# data "haproxy_cache" "cash_test" {
#   name = "cash_test"
#   depends_on = [ haproxy_cache.cash_test ]
# }

# resource "haproxy_global" "global" {
#     user = "haproxy"
#     group = "haproxy"
#     chroot = "/var/lib/haproxy"
#     daemon = true
#     master_worker = true
#     maxcompcpuusage = 0
#     maxpipes = 0
#     maxsslconn = 0
#     maxconn = 2000
#     # nbproc = 1
#     nbthread = 1
#     pidfile = "/var/run/haproxy.pid"
#     ulimit_n = 2000
#     crt_base = "/etc/ssl/certs"
#     ca_base = "/etc/ssl/private"
#     stats_maxconn = 100
#     stats_timeout = 60
#     ssl_default_bind_ciphers = "ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256"
#     ssl_default_bind_options = "no-sslv3 no-tlsv10 no-tlsv11 no-tls-tickets"
# }

# data "haproxy_global" "global" {
#     name = "global"
#     depends_on = [ haproxy_global.global ]
# }

resource "haproxy_backend" "backend_test" {
  name                 = "backend_test"
  mode                 = "http"
  http_connection_mode = "http-keep-alive"
  server_timeout       = 9
  check_timeout        = 20
  connect_timeout      = 20
  queue_timeout        = 20
  tarpit_timeout       = 20
  tunnel_timeout       = 20
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

data "haproxy_backend" "backend_test" {
  name       = "backend_test"
  depends_on = [haproxy_backend.backend_test]
}

output "backend_test" {
  value = data.haproxy_backend.backend_test
}
