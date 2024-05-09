

**# ansible-awx_inventory_script

`ansible-awx_inventory_script` custom inventory scripts available in AWX


## Example Usage

```hcl
resource "ansible-awx_inventory_script" "script" {
  name = "test script"
  description = "description"
  organization_id = ansible-awx_oragnization.oragnization.id
  script = <<EOT
#!/usr/bin/env python
echo "hey"
EOT

}
```



## Argument Reference

The following arguments are supported:

* `name` - Name of this custom inventory script. (string, required)
* `description` - Optional description of this custom inventory script. (string, default="")
* `organization_id` - Organization owning this inventory script (id, required)
* `script` - (string, required)

## Import

Ansible AWX Inventory script can be imported using the id, e.g. for an Inventory script with id : 120

```sh  
$ terraform import ansible-awx_inventory_script 120  
```**
