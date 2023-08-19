---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "haproxy_frontend Data Source - terraform-provider-haproxy"
subcategory: ""
description: |-
  
---

# haproxy_frontend (Data Source)

Retrieve information about an existing frontend.

## Example Usage

```terraform
data "haproxy_frontend" "front_test" {
  name = "front_test"
  depends_on = [ haproxy_frontend.front_test ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the frontend. It must be unique and cannot be changed.

### Read-Only

- `id` (String) The ID of this resource.