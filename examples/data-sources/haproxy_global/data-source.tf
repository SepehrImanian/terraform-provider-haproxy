data "haproxy_global" "global" {
    name = "global"
    depends_on = [ haproxy_global.global ]
}