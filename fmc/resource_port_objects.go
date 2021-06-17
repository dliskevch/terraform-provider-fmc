package fmc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var port_type string = "ProtocolPortObject"

func resourcePortObjects() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for Port Objects in FMC\n" +
			"\n" +
			"## Example\n" +
			"An example is shown below: \n" +
			"```hcl\n" +
			"resource \"fmc_port_objects\" \"http\" {\n" +
			"    name = \"HTTP\"\n" +
			"    port = \"80\"\n" +
			"    protocol = \"TCP\"\n" +
			"}\n" +
			"```",
		CreateContext: resourcePortObjectsCreate,
		ReadContext:   resourcePortObjectsRead,
		UpdateContext: resourcePortObjectsUpdate,
		DeleteContext: resourcePortObjectsDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this resource",
			},
			"port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Port for this resource",
			},
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Protocol for this resource",
			},
			"overridable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sets this resource as overridable",
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
		},
	}
}

func resourcePortObjectsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics
	var diags diag.Diagnostics

	res, err := c.CreateFmcPortObject(ctx, &PortObject{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Port:        d.Get("port").(string),
		Protocol:    d.Get("protocol").(string),
		Overridable: d.Get("overridable").(bool),
		Type:        port_type,
	})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to create port object",
			Detail:   err.Error(),
		})
		return diags
	}
	d.SetId(res.ID)
	return resourcePortObjectsRead(ctx, d, m)
}

func resourcePortObjectsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	item, err := c.GetFmcPortObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}
	if err := d.Set("name", item.Name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", item.Description); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("port", item.Port); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("protocol", item.Protocol); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("type", item.Type); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to read port object",
			Detail:   err.Error(),
		})
		return diags
	}

	return diags
}

func resourcePortObjectsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	id := d.Id()
	if d.HasChanges("name", "description", "value") {
		_, err := c.UpdateFmcPortObject(ctx, id, &PortObjectUpdateInput{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Port:        d.Get("port").(string),
			Protocol:    d.Get("protocol").(string),
			Overridable: d.Get("overridable").(bool),
			Type:        port_type,
			ID:          id,
		})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "unable to update port object",
				Detail:   err.Error(),
			})
			return diags
		}
	}
	return resourcePortObjectsRead(ctx, d, m)
}

func resourcePortObjectsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteFmcPortObject(ctx, id)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "unable to delete port object",
			Detail:   err.Error(),
		})
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
