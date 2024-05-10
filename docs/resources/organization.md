# ansible-awx_organization
An Organization is a logical collection of Users, Teams, Projects, and Inventories, and is the highest level in the AWX object hierarchy.

## Example Usage

```hcl
resource "ansible-awx_organization" "example" {
  name            = "test"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required)  Name of this organization.
* `custom_virtualenv` - (Optional) Local absolute file path containing a custom Python virtualenv to use
* `description` - (Optional) description of this organization. (string, default="")
* `max_hosts` - (Optional) Maximum number of hosts allowed to be managed by this organization

## Import

Ansible AWX Organisation can be imported using the id, e.g. for an organization with id : 110 

```sh
$ terraform import ansible-awx_organization.example 110
```
