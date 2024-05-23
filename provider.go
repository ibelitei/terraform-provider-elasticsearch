package terraform_provider_elasticsearch

import (
	"net/url"

	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerOpts the options for the provider
type providerOpts struct {
	url *url.URL
}

// Provider _
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"elasticsearch_index": resourceIndexTemplate(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ELASTICSEARCH_URL", nil),
				Description: "The URL of the elasticsearch",
			},
		},
		ConfigureFunc: configureProvider,
	}
}

// provider configuration
func configureProvider(d *schema.ResourceData) (interface{}, error) {
	urles := d.Get("url").(string)
	parsedURL, err := url.Parse(urles)
	if err != nil {
		return nil, err
	}

	return &providerOpts{
		url: parsedURL,
	}, nil
}

// creates Elasticsearch client from the provider
func getClientES(opts *providerOpts) (*es.Client, error) {
	cfg := es.Config{
		Addresses: []string{opts.url.String()},
	}
	return es.NewClient(cfg)
}
