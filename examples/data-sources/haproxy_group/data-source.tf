data "haproxy_group" "something_data" {
  name     = haproxy_group.something.name
  userlist = "userslist"
}
