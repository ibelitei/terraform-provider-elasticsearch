package terraform_provider_elasticsearch

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

// schema and CRUD functions for the Elasticsearch index template
func resourceIndexTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceIndexTemplateCreate,
		Read:   resourceIndexTemplateRead,
		Update: resourceIndexTemplateUpdate,
		Delete: resourceIndexTemplateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the elasticsearch template",
			},
			"template": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "JSON body for the template you want to use for an elasticsearch index",
			},
		},
	}
}

func resourceIndexTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	template := d.Get("template").(string)

	client, err := getClientES(meta.(*providerOpts))
	if err != nil {
		return err
	}

	// valid JSON string
	req, err := client.Indices.PutTemplate(name, strings.NewReader(template))
	if req.IsError() {
		return fmt.Errorf("Error creating index template: %s", req.String())
	}

	d.SetId(name)
	return resourceIndexTemplateRead(d, meta)
}

func resourceIndexTemplateRead(d *schema.ResourceData, meta interface{}) error {
	client, err := getClientES(meta.(*providerOpts))
	if err != nil {
		return err
	}

	req, err := client.Indices.GetTemplate(client.Indices.GetTemplate.WithName(d.Id()))
	if err != nil || req.IsError() {
		return fmt.Errorf("Error reading index template: %s", req.String())
	}

	// response and update the schema
	return nil
}

func resourceIndexTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIndexTemplateCreate(d, meta)
}

func resourceIndexTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	client, err := getClientES(meta.(*providerOpts))
	if err != nil {
		return err
	}

	req, err := client.Indices.DeleteTemplate(d.Id())
	if req.IsError() {
		return fmt.Errorf("Error deleting index template: %s", req.String())
	}

	return nil
}
