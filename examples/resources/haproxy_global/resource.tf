resource "haproxy_global" "global" {
    user = "haproxy"
    group = "haproxy"
    chroot = "/var/lib/haproxy"
    daemon = true
    master_worker = true
    maxcompcpuusage = 0
    maxpipes = 0
    maxsslconn = 0
    maxconn = 2000
    nbproc = 1 // just before version haproxy 2.5
    nbthread = 1
    pidfile = "/var/run/haproxy.pid"
    ulimit_n = 2000
    crt_base = "/etc/ssl/certs"
    ca_base = "/etc/ssl/private"
    stats_maxconn = 100
    stats_timeout = 60
    ssl_default_bind_ciphers = "ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256"
    ssl_default_bind_options = "no-sslv3 no-tlsv10 no-tlsv11 no-tls-tickets"
}