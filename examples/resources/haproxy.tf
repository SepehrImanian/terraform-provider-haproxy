resource "haproxy_userlist" "users" {
  name = "userslist"
}

data "haproxy_userlist" "user1" {
  name = haproxy_userlist.users.name
}

output "user1" {
  value = data.haproxy_userlist.user1
}
