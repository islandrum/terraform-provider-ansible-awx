terraform {
  required_version = ">= 0.13"

  required_providers {
    ansible-awx = {
      source = "registry.terraform.io/islandrum/ansible-awx"
    }
  }
}
variable "username" {
  description = "ansible awx username"
  default = "admin"
}
variable "password" {
  description = "ansible awx password"
  default = "password"
}

variable "host" {
  description = "ansible awx host"
  default = "http://127.0.0.1"
}
provider "ansible-awx" {
  awx_host = var.host
  awx_username = var.username
  awx_password = var.password
}

resource "ansible-awx_organisation" "organisation" {
  name = "test organisation"
  description = "desc"
}

resource "ansible-awx_inventory" "inventory" {
  name = "test inventory"
  description = "test dsd"
  organisation_id = ansible-awx_organisation.organisation.id
  kind = ""
  host_filter = ""
  inv_var {
    key = "sas"
    value = "sasaa"
  }
  inv_var {
    key = "monta"
    value = "[ a , b ]"
  }
}

resource "ansible-awx_inventory_script" "script" {
  name = "tf scriptssdsdddsds"
  description = "dsdsd"
  organization_id = ansible-awx_organisation.organisation.id
  script = <<EOT
#!/usr/bin/env python
echo "hey"
EOT

}
resource "ansible-awx_inventory_source" "source_custom_script" {
  name = "cxcdsfdsffffx"
  inventory_id = ansible-awx_inventory.inventory.id
  source = "custom"
  source_script = ansible-awx_inventory_script.script.id
}
resource "ansible-awx_credential_scm" "credential" {
  organisation_id = ansible-awx_organisation.organisation.id
  name            = "acc-scm-credential"
  username        = "test"
  ssh_key_data    = file("${path.module}/files/id_rsa")
}


resource "ansible-awx_credential_machine" "credential" {
  organisation_id     = ansible-awx_organisation.organisation.id
  name                = "acc-machine-credential"
  username            = "test"
  ssh_key_data        = file("${path.module}/files/id_rsa")
  ssh_public_key_data = file("${path.module}/files/id_rsa.pub")

}

resource "ansible-awx_project" "vault" {
  name                 = "test playbook"
  scm_type             = "git"
  scm_url              = "https://github.com/islandrum/ansible-playbook-awx-test"
  scm_branch           = "main"
  scm_update_on_launch = true
  organisation_id      = ansible-awx_organisation.organisation.id
//  scm_credential_id    = ansible-awx_credential_scm.credential.id
}
resource "ansible-awx_inventory_source" "source" {
  name = "cfdfdxcx"
  inventory_id = ansible-awx_inventory.inventory.id
  source_project_id = ansible-awx_project.vault.id
  source_path= ""
  source = "scm"

}
resource "ansible-awx_job_template" "template" {
  name           = "test-job-template"
  inventory_id   = ansible-awx_inventory.inventory.id
  project_id     = ansible-awx_project.vault.id
  playbook       = "main.yml"
  job_type       = "run"
  become_enabled = true
}

resource "ansible-awx_credential_type" "type2" {
  name           = "credential_type"
  input {
    id = "username"
    type = "string"
    label = "USERNAME"
  }
  input {
    id = "password"
    type = "string"
    label = "PASSWORD"
    secret = true
  }

  input {
    id = "url"
    type = "string"
    label = "URI"
    format = "url"
    multiline = false
  }
}
