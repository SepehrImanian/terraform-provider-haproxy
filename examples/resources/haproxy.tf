resource "haproxy_userlist" "userslist" {
  name = "userslist"
}

resource "haproxy_user" "sepehr" {
  username        = "sepehr"
  userlist        = haproxy_userlist.userslist.name
  password        = "123456999"
  secure_password = true
  groups          = haproxy_group.something.name
  depends_on      = [haproxy_userlist.userslist, haproxy_group.something]
}

resource "haproxy_group" "something" {
  name       = "something"
  userlist   = "userslist"
  depends_on = [haproxy_userlist.userslist]
}

data "haproxy_group" "something_data" {
  name     = haproxy_group.something.name
  userlist = "userslist"
}

output "something" {
  value = data.haproxy_group.something_data
}
