---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "haproxy_userlist Resource - terraform-provider-haproxy"
subcategory: ""
description: |-
  
---

# haproxy_userlist (Resource)



## Example Usage

```terraform
resource "haproxy_userlist" "userslist" {
  name = "userslist"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the Userlist. It must be unique

### Read-Only

- `id` (String) The ID of this resource.
