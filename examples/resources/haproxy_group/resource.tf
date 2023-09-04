resource "haproxy_group" "something" {
  name       = "something"
  userlist   = "userslist"
  depends_on = [haproxy_userlist.userslist]
}
