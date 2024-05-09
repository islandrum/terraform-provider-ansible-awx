
# Ansible AWX Provider

Ansible AWX Provider for handle AWX Projects with [rest](https://ansible.readthedocs.io/projects/awx/en/latest/rest_api/api_ref.html)

## Example Usage

```hcl
provider "ansible-awx" {
  awx_host = "http://127.0.0.1"
  awx_username = "admin"
  awx_password = "password"
}
```
### Environment variables

You can also provide your credentials via the environment variables, TOWER_HOST, TOWER_USERNAME, MONGO_USR, and TOWER_PASSWORD respectively:

```hcl
provider "ansible-awx" {

}
```

Usage (prefix the export commands with a space to avoid the keys being recorded in OS history):

```shell
$  export TOWER_HOST="xxxx"
$  export TOWER_USERNAME="xxxx"
$  export TOWER_PASSWORD="xxxx"
$ terraform plan
```


## Argument Reference

In addition to [generic `provider`
arguments](https://www.terraform.io/docs/configuration/providers.html) (e.g.
`alias` and `version`), the following arguments are supported in the MongoDB
`provider` block:

* `awx_host` - (Optional) This is the host your ansible awx Server. It must be
  provided, but it can also be sourced from the `TOWER_HOST`
  environment variable.

* `awx_username ` - (Optional) Specifies a username with which to authenticate to the ansible awx Server. It must be
  provided, but it can also be sourced from the `TOWER_USERNAME`
  environment variable.
* `awx_password  ` - (Optional) Specifies a password with which to authenticate to the ansible awx Server. It must be
  provided, but it can also be sourced from the `TOWER_PASSWORD`
  environment variable.

