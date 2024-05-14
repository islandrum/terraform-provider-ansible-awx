package awx

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/islandrum/go-ansible-awx-sdk/client"
	"github.com/mitchellh/mapstructure"
)

func resourceCredentialMachine() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCredentialMachineCreate,
		ReadContext:   resourceCredentialMachineRead,
		UpdateContext: resourceCredentialMachineUpdate,
		DeleteContext: resourceCredentialMachineDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"organization_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"ssh_key_data": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"ssh_public_key_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_key_unlock": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"become_method": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"become_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"become_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"team_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCredentialMachineDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	stateID := d.State().ID
	id, err := decodeStateId(stateID)
	if err != nil {
		return DiagsError(CredentialMachineResourceName, err)
	}
	client := m.(*awx.AWX)
	err = client.CredentialsService.DeleteCredentialsByID(id, map[string]string{})
	if err != nil {
		return DiagDeleteFail(CredentialMachineResourceName, fmt.Sprintf(
			"%s %v, got %s ",
			CredentialMachineResourceName, id, err.Error(),
		))
	}

	return diags
}

func resourceCredentialMachineCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var err error
	newCredential := map[string]interface{}{
		"name":            d.Get("name").(string),
		"description":     d.Get("description").(string),
		"organization":    d.Get("organization_id").(int),
		"credential_type": 1, // SSH
		"inputs": map[string]interface{}{
			"username":            d.Get("username").(string),
			"password":            d.Get("password").(string),
			"ssh_key_data":        d.Get("ssh_key_data").(string),
			"ssh_public_key_data": d.Get("ssh_public_key_data").(string),
			"ssh_key_unlock":      d.Get("ssh_key_unlock").(string),
			"become_method":       d.Get("become_method").(string),
			"become_username":     d.Get("become_username").(string),
			"become_password":     d.Get("become_password").(string),
		},
		"team": d.Get("team_id").(int),
		"user": d.Get("user_id").(int),
	}
	if newCredential["organization"] == 0 {
		newCredential["organization"] = nil
	}
	if newCredential["team"] == 0 {
		newCredential["team"] = nil
	}
	if newCredential["user"] == 0 {
		newCredential["user"] = nil
	}

	client := m.(*awx.AWX)
	cred, err := client.CredentialsService.CreateCredentials(newCredential, map[string]string{})
	if err != nil {
		return DiagsError(CredentialMachineResourceName, err)
	}

	d.SetId(getStateID(cred.ID))
	resourceCredentialMachineRead(ctx, d, m)

	return diags
}

func resourceCredentialMachineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*awx.AWX)
	stateID := d.State().ID
	id, err := decodeStateId(stateID)
	if err != nil {
		return DiagsError(CredentialMachineResourceName, err)
	}
	cred, err := client.CredentialsService.GetCredentialsByID(id, map[string]string{})
	if err != nil {
		return DiagNotFoundFail(CredentialMachineResourceName, id, err)
	}

	setErr := d.Set("name", cred.Name)
	if setErr != nil {
		return DiagsError(CredentialMachineResourceName, setErr)
	}
	setErr = d.Set("description", cred.Description)
	if setErr != nil {
		return DiagsError(CredentialMachineResourceName, setErr)
	}
	setErr = d.Set("username", cred.Inputs["username"])
	if setErr != nil {
		return DiagsError(CredentialMachineResourceName, setErr)
	}
	setErr = d.Set("password", cred.Inputs["password"])

	if setErr != nil {
		return DiagsError(CredentialMachineResourceName, setErr)
	}
	setErr = d.Set("ssh_key_data", cred.Inputs["ssh_key_data"])
	if setErr != nil {
		return DiagsError(CredentialMachineResourceName, setErr)
	}
	setErr = d.Set("ssh_public_key_data", cred.Inputs["ssh_public_key_data"])
	if setErr != nil {
		return DiagsError(CredentialMachineResourceName, setErr)
	}
	setErr = d.Set("ssh_key_unlock", cred.Inputs["ssh_key_unlock"])
	if setErr != nil {
		return DiagsError(CredentialMachineResourceName, setErr)
	}
	setErr = d.Set("become_method", cred.Inputs["become_method"])
	if setErr != nil {
		return DiagsError(CredentialMachineResourceName, setErr)
	}
	setErr = d.Set("become_username", cred.Inputs["become_username"])
	if setErr != nil {
		return DiagsError(CredentialMachineResourceName, setErr)
	}
	setErr = d.Set("become_password", cred.Inputs["become_password"])
	if setErr != nil {
		return DiagsError(CredentialMachineResourceName, setErr)
	}
	owners, err := GetCredentialOwnersFromSummaryFields(cred.SummaryFields)
	if err != nil {
		return DiagsError(CredentialMachineResourceName, err)
	}
	if cred.OrganizationID != nil {
		setErr = d.Set("organization_id", cred.OrganizationID)
		if setErr != nil {
			return DiagsError(CredentialMachineResourceName, setErr)
		}
	} else if owners.OrganizationID != 0 {
		setErr = d.Set("organization_id", owners.OrganizationID)
		if setErr != nil {
			return DiagsError(CredentialMachineResourceName, setErr)
		}
	} else {
		setErr = d.Set("organization_id", nil)
		if setErr != nil {
			return DiagsError(CredentialMachineResourceName, setErr)
		}
	}
	if owners.TeamID != 0 {
		setErr = d.Set("team_id", owners.TeamID)
		if setErr != nil {
			return DiagsError(CredentialMachineResourceName, setErr)
		}
	}
	if owners.UserID != 0 {
		setErr = d.Set("user_id", owners.UserID)
		if setErr != nil {
			return DiagsError(CredentialMachineResourceName, setErr)
		}
	}

	return diags
}

func resourceCredentialMachineUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	keys := []string{
		"name",
		"description",
		"username",
		"password",
		"ssh_key_data",
		"ssh_public_key_data",
		"ssh_key_unlock",
		"become_method",
		"become_username",
		"become_password",
		"organization_id",
		"teams",
		"owners",
	}

	if d.HasChanges(keys...) {
		var err error

		stateID := d.State().ID
		id, err := decodeStateId(stateID)
		if err != nil {
			return DiagsError(CredentialMachineResourceName, err)
		}
		updatedCredential := map[string]interface{}{
			"name":            d.Get("name").(string),
			"description":     d.Get("description").(string),
			"organization":    d.Get("organization_id").(int),
			"credential_type": 1, // SSH
			"inputs": map[string]interface{}{
				"username":            d.Get("username").(string),
				"password":            d.Get("password").(string),
				"ssh_key_data":        d.Get("ssh_key_data").(string),
				"ssh_public_key_data": d.Get("ssh_public_key_data").(string),
				"ssh_key_unlock":      d.Get("ssh_key_unlock").(string),
				"become_method":       d.Get("become_method").(string),
				"become_username":     d.Get("become_username").(string),
				"become_password":     d.Get("become_password").(string),
			},
			"team": d.Get("team_id").(int),
			"user": d.Get("user_id").(int),
		}
		if updatedCredential["organization"] == 0 {
			updatedCredential["organization"] = nil
		}
		if updatedCredential["team"] == 0 {
			updatedCredential["team"] = nil
		}
		if updatedCredential["user"] == 0 {
			updatedCredential["user"] = nil
		}

		client := m.(*awx.AWX)
		_, err = client.CredentialsService.UpdateCredentialsByID(id, updatedCredential, map[string]string{})
		if err != nil {
			return DiagUpdateFail(CredentialMachineResourceName, id, err)
		}
	}

	return resourceCredentialMachineRead(ctx, d, m)
}

type OwnerIDs struct {
	UserID         int
	TeamID         int
	OrganizationID int
}

type SummaryFieldOwner struct {
	ID          int    `mapstructure:"id"`
	Type        string `mapstructure:"type"`
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Url         string `mapstructure:"url"`
}

type SummaryFieldOwners struct {
	Owners []SummaryFieldOwner `mapstructure:"owners"`
}

func GetCredentialOwnersFromSummaryFields(SummaryFields map[string]interface{}) (OwnerIDs, error) {
	var summaryFieldOwners SummaryFieldOwners
	var owners OwnerIDs
	err := mapstructure.Decode(SummaryFields, &summaryFieldOwners)
	if err != nil {
		return owners, err
	}
	for _, owner := range summaryFieldOwners.Owners {
		switch strings.ToLower(owner.Type) {
		case "organization":
			owners.OrganizationID = owner.ID
		case "team":
			owners.TeamID = owner.ID
		case "user":
			owners.UserID = owner.ID
		}
	}
	return owners, nil
}
