# Terraform provider Ansible-AWX

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/islandrum/terraform-provider-ansible-awx?logo=go&style=flat-square)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/islandrum/terraform-provider-ansible-awx?logo=git&style=flat-square)
![GitHub](https://img.shields.io/github/license/islandrum/terraform-provider-ansible-awx?color=yellow&style=flat-square)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/islandrum/terraform-provider-ansible-awx/golangci?logo=github&style=flat-square)
![GitHub issues](https://img.shields.io/github/issues/islandrum/terraform-provider-ansible-awx?logo=github&style=flat-square)


This repository is a [Terraform](https://www.terraform.io) Provider for Ansible AWX  
 
### Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13
- [Go](https://golang.org/doc/install) >= 1.17

### Installation

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the `make install` command:

````bash
git clone https://github.com/islandrum/terraform-provider-mongodb
cd terraform-provider-ansible-awx
make install
````

### To test locally

**1.1: launch awx**


````bash
cd examples
docker-compose up -d
````

*follow the instruction in this link*

https://debugthis.dev/posts/2020/04/setting-up-ansible-awx-using-a-docker-environment-part-2-the-docker-compose-approach/


**1.4 :  user in awx**

* default user : admin
* default password : password

**2: Build the provider**

follow the [Installation](#Installation)

**3: Use the provider**

````bash
cd examples
make apply
````
