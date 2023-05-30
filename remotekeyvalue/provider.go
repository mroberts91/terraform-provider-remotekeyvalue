package remotekeyvalue

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"remotekeyvalue_pair": dataSourceKeyValuePair(),
		},
		Schema: map[string]*schema.Schema{
			"uri": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URI of the API endpoint to retrive the key value. This serves as the base of all requests.",
			},
			"api_key_header_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the header used to send an API Key with the request.",
			},
			"api_key_header_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Value of the header used to send an API Key with the request.",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     100,
				Description: "HTTP Request Timeout",
			},
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {

	opt := &apiClientOpt{
		uri:                  d.Get("uri").(string),
		timeout:              d.Get("timeout").(int),
		api_key_header_name:  d.Get("api_key_header_name").(string),
		api_key_header_value: d.Get("api_key_header_value").(string),
	}

	client, err := NewAPIClient(opt)
	return client, err
}
