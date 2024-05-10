package awx

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/islandrum/go-ansible-awx-sdk/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"awx_host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWX_HOST", "http://127.0.0.1"),
			},
			"awx_username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWX_USERNAME", "admin"),
			},
			"awx_password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AWX_PASSWORD", "password"),
			},
			"validate_certs": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Validate SSL certificates of AWX",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ansible-awx_inventory":          resourceInventory(),
			"ansible-awx_organization":       resourceOrganization(),
			"ansible-awx_inventory_source":   resourceInventorySource(),
			"ansible-awx_inventory_script":   resourceInventoryScript(),
			"ansible-awx_project":            resourceProject(),
			"ansible-awx_job_template":       resourceJobTemplate(),
			"ansible-awx_credential_scm":     resourceCredentialSCM(),
			"ansible-awx_credential_machine": resourceCredentialMachine(),
			"ansible-awx_credential_type":    resourceCredentialType(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	hostname := d.Get("awx_host").(string)
	username := d.Get("awx_username").(string)
	password := d.Get("awx_password").(string)

	client := http.DefaultClient
	if !d.Get("validate_certs").(bool) {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	var diags diag.Diagnostics
	c, err := awx.NewAWX(hostname, username, password, client)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Authentication error",
			Detail:   "Check Host , Username and Password",
		})
		return nil, diags
	}

	return c, diags
}
