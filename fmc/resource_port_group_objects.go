package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var port_group_type string = "PortObjectGroup"

func resourcePortGroupObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Port Group Objects in FMC\n" +
			"\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"## Example\n" +
			"resource \"fmc_port_group_objects\" \"port-group\" {\n" +
			"    name = \"TCP-ICMP\"\n" +
			"    description = \"Combo ports\"\n" +
			"    objects {\n" +
			"        id = fmc_port_objects.http.id\n" +
			"        type = fmc_port_objects.http.type\n" +
			"    }\n" +
			"    objects {\n" +
			"        id = fmc_icmpv4_objects.wrong-proto.id\n" +
			"        type = fmc_icmpv4_objects.wrong-proto.type\n" +
			"    }\n" +
			"}\n" +
			"```",
		CreateContext: resourcePortGroupObjectsCreate,
		ReadContext:   resourcePortGroupObjectsRead,
		UpdateContext: resourcePortGroupObjectsUpdate,
		DeleteContext: resourcePortGroupObjectsDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of this resource",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of this resource",
			},
			"objects": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of this resource",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of this resource",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of this resource",
						},
					},
				},
				Description: "The list of port groups to add",
			},
		},
	}
}

func resourcePortGroupObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	var objs []PortGroupObjectObjects

	if inputObjs, ok := d.GetOk("objects"); ok {
		for _, obj := range inputObjs.([]interface{}) {
			obji := obj.(map[string]interface{})
			objs = append(objs, PortGroupObjectObjects{
				ID:   obji["id"].(string),
				Name: obji["name"].(string),
				Type: obji["type"].(string),
			})
		}
	}

	res, err := c.CreateFmcPortGroupObject(ctx, &PortGroupObject{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        port_group_type,
		Objects:     objs,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create port group object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourcePortGroupObjectsRead(ctx, d, m)
}

func resourcePortGroupObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcPortGroupObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port group object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port group object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port group object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port group object",
			Detail:   err.Error(),
		})
		return diags
	}

	objects := make([]interface{}, len(item.Objects))
	for _, obj := range item.Objects {
		obji := make(map[string]interface{})
		obji["id"] = obj.ID
		obji["name"] = obj.Name
		obji["type"] = obj.Type
		objects = append(objects, obji)
	}

	if err := d.Set("objects", objects); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port group object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourcePortGroupObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "objects") {
		var objs []PortGroupObjectObjects

		if inputObjs, ok := d.GetOk("objects"); ok {
			for _, obj := range inputObjs.([]interface{}) {
				obji := obj.(map[string]interface{})
				objs = append(objs, PortGroupObjectObjects{
					ID:   obji["id"].(string),
					Name: obji["name"].(string),
					Type: obji["type"].(string),
				})
			}
		}

		_, err := c.UpdateFmcPortGroupObject(ctx, id, &PortGroupObjectUpdateInput{
			ID:          d.Id(),
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Type:        port_group_type,
			Objects:     objs,
		})

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update port group object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourcePortGroupObjectsRead(ctx, d, m)
}

func resourcePortGroupObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcPortGroupObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete port group object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
