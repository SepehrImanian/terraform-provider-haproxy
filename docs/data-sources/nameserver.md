---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "haproxy_nameserver Data Source - terraform-provider-haproxy"
subcategory: ""
description: |-
  
---

# haproxy_nameserver (Data Source)



## Example Usage

```terraform
data "haproxy_nameserver" "nameserver_1" {
  name     = haproxy_nameserver.nameserver_1.name
  resolver = haproxy_resolvers.resolvers_test.name
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the Nameserver. It must be unique
- `resolver` (String) The name of the parent object

### Read-Only

- `id` (String) The ID of this resource.
