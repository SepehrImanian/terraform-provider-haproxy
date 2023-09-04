resource "haproxy_userlist" "userslist2" {
  name = "userslist2"
}

data "haproxy_userlist" "user1" {
  name = haproxy_userlist.userslist2.name
}

output "user1" {
  value = data.haproxy_userlist.user1
}
