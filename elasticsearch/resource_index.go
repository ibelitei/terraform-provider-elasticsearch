package elasticsearch

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

// schema and CRUD functions for the elasticsearch index
func resourceIndex() *schema.Resource {
	return &schema.Resource{
		Create: resourceIndexCreate,
		Read:   resourceIndexRead,
		Update: resourceIndexUpdate,
		Delete: resourceIndexDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the elasticsearch index",
			},
			"settings": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "settings for elasticsearch index",
			},
		},
	}
}

func resourceIndexCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	settings := d.Get("settings").(string)

	client, err := getClientES(meta.(*providerOpts))
	if err != nil {
		return fmt.Errorf("error elasticsearch client: %v", err)
	}

	// create index
	req, err := client.Indices.Create(name, client.Indices.Create.WithBody(strings.NewReader(settings)))
	if err != nil {
		return fmt.Errorf("error creating index: %v", err)
	}
	if req.IsError() {
		return fmt.Errorf("error creating index: %s", req.String())
	}

	d.SetId(name)
	return nil
}

func resourceIndexRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIndexUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIndexCreate(d, meta)
}

func resourceIndexDelete(d *schema.ResourceData, meta interface{}) error {
	client, err := getClientES(meta.(*providerOpts))
	if err != nil {
		return err
	}

	req, err := client.Indices.Delete([]string{d.Id()}) // Corrected to delete an index, not a template
	if err != nil {
		return fmt.Errorf("error deleting index: %v", err)
	}
	if req.IsError() {
		return fmt.Errorf("error deleting index: %s", req.String())
	}

	return nil
}
