---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "prefect_variable Data Source - prefect"
subcategory: ""
description: |-
  Data Source representing a Prefect variable
---

# prefect_variable (Data Source)

Data Source representing a Prefect variable

## Example Usage

```terraform
data "prefect_variable" "existing_by_id" {
  id = "variable-UUID"
}

data "prefect_variable" "existing_by_name" {
  name = "my-variable-name"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) Account UUID, defaults to the account set in the provider
- `id` (String) Variable UUID
- `name` (String) Name of the variable
- `workspace_id` (String) Workspace UUID, defaults to the workspace set in the provider

### Read-Only

- `created` (String) Date and time of the variable creation in RFC 3339 format
- `tags` (List of String) Tags associated with the variable
- `updated` (String) Date and time that the variable was last updated in RFC 3339 format
- `value` (String) Value of the variable