package awx

import (
	"context"
	"fmt"
	awx "github.com/islandrum/go-ansible-awx-sdk/client"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
)

func resourceInventory() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInventoryCreate,
		ReadContext:   resourceInventoryRead,
		DeleteContext: resourceInventoryDelete,
		UpdateContext: resourceInventoryUpdate,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"organization_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"kind": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"host_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"inv_var": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceInventoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventoriesService
	var vars []Variable
	list := d.Get("inv_var").(*schema.Set).List()
	err := mapstructure.Decode(list, &vars)
	if err != nil {
		return DiagsError(InventoryResourceName, err)
	}
	inventoryVars := CreateInventoryVariables(vars)
	result, err := awxService.CreateInventory(map[string]interface{}{
		"name":         d.Get("name").(string),
		"organization": d.Get("organization_id").(string),
		"description":  d.Get("description").(string),
		"kind":         d.Get("kind").(string),
		"host_filter":  d.Get("host_filter").(string),
		"variables":    inventoryVars,
	}, map[string]string{})
	if err != nil {
		return DiagsError(InventoryResourceName, err)
	}
	d.SetId(getStateID(result.ID))
	return resourceInventoryRead(ctx, d, m)

}

func resourceInventoryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventoriesService
	stateID := d.State().ID
	id, err := decodeStateId(stateID)

	if err != nil {
		return DiagNotFoundFail(InventoryResourceName, id, err)
	}
	var vars []Variable
	list := d.Get("inv_var").(*schema.Set).List()
	err = mapstructure.Decode(list, &vars)
	if err != nil {
		return DiagsError(InventoryResourceName, err)
	}
	inventoryVars := CreateInventoryVariables(vars)
	_, err = awxService.UpdateInventory(id, map[string]interface{}{
		"name":         d.Get("name").(string),
		"organization": d.Get("organization_id").(string),
		"description":  d.Get("description").(string),
		"kind":         d.Get("kind").(string),
		"host_filter":  d.Get("host_filter").(string),
		"variables":    inventoryVars,
	}, nil)

	if err != nil {
		return DiagUpdateFail(InventoryResourceName, id, err)
	}

	return resourceInventoryRead(ctx, d, m)

}

func resourceInventoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventoriesService
	stateID := d.State().ID
	id, err := decodeStateId(stateID)
	if err != nil {
		return DiagsError(InventoryResourceName, err)
	}
	r, err := awxService.GetInventory(id, map[string]string{})

	if err != nil {
		return DiagNotFoundFail(InventoryResourceName, id, err)
	}
	d = setInventoryResourceData(d, r)
	return nil
}

func resourceInventoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventoriesService
	stateID := d.State().ID
	id, err := decodeStateId(stateID)
	if err != nil {
		return DiagsError(InventoryResourceName, err)
	}
	if _, err := awxService.DeleteInventory(id); err != nil {
		return DiagDeleteFail(
			InventoryResourceName,
			fmt.Sprintf(
				"%s %v, got %s ",
				InventoryResourceName, id, err.Error(),
			),
		)
	}
	d.SetId("")
	return nil
}

//nolint:errcheck
func setInventoryResourceData(d *schema.ResourceData, r *awx.Inventory) *schema.ResourceData {
	d.Set("name", r.Name)
	d.Set("organization_id", strconv.Itoa(r.Organization))
	d.Set("description", r.Description)
	d.Set("kind", r.Kind)
	d.Set("host_filter", r.HostFilter)
	d.SetId(getStateID(r.ID))
	return d
}
