
# ansible-awx_inventory

`ansible-awx_inventory` An Inventory is a collection of hosts against which jobs may be launched, the same as an Ansible inventory file. Inventories are divided into groups and these groups contain the actual hosts. Groups may be sourced manually, by entering host names into AWX, or from one of Ansible AWX's supported cloud providers.

## Example Usage

```hcl

variable "db_pwd" {
  description = "database pwd passed as var to inventory"
}

resource "ansible-awx_inventory" "default" {
  name            = "acc-test"
  organization_id = "1"
  inv_var {
    key = "database_user"
    value = "root"
  }
  inv_var {
    key = "database_pwd"
    value = var.db_pwd
  }
}
```

## Example Usage with  organization

```hcl
data "ansible-awx_organization" "example" {
  name = "example"
}

resource "ansible-awx_inventory" "example" {
  name            = "acc-test"
  organization_id = data.ansible-awx_organization.example.id
  inv_var {
    key = "example_key"
    value = "value_value"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of this inventory. (string, required) 
* `organization_id` - Organization containing this inventory. (id, required)
* `description` - Optional description of this inventory. (string, default="")
* `host_filter` - (Optional)  Filter that will be applied to the hosts of this inventory. (string, default="")
* `kind` - (Optional)  Kind of inventory being represented. (choice) 
  * **""** :  Hosts have a direct link to this inventory. (default)
  * **smart** : Hosts for inventory generated using the host_filter property.
* `inv_var` - (Optional) Inventory variables accept up to 10 


## Import

Ansible AWX Inventory can be imported using the id, e.g. for an Inventory with id : 125

```sh
$ terraform import ansible-awx_inventory.example 125
```
