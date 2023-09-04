data "haproxy_user" "sepehr" {
  username = "sepehr"
  userlist = haproxy_userlist.userslist.name
}