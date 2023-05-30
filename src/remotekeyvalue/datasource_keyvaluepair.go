package remotekeyvalue

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKeyValuePair() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKeyValuePairRead,
		Schema: map[string]*schema.Schema{
			"path": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The path used to retrieve a key-value pair.",
			},
			"method": {
				Optional:    true,
				Default:     "GET",
				Type:        schema.TypeString,
				Description: "HTTP method to use when retrieving a key-value pair. Defaults to GET",
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sensitive": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceKeyValuePairRead(d *schema.ResourceData, meta interface{}) error {
	path := d.Get("path").(string)
	method := d.Get("method").(string)

	client := meta.(*api_client)
	client.path = path
	client.method = method

	jsonBytes, err := client.send_request()
	if err != nil {
		return err
	}

	rawValue := string(jsonBytes)
	log.Printf("json response:\n%s", rawValue)
	item, err := UnmarshalApiResponseItem(jsonBytes)

	if err == nil {
		/* Setting terraform ID tells terraform the object was created or it exists */
		log.Printf("datasource_.go: Data resource. Returned id is '%d'\n", item.ID)
		d.SetId(fmt.Sprint(item.ID))
		d.Set("id", item.ID)
		d.Set("key", item.Key)
		d.Set("value", item.Value)
		d.Set("sensitive", item.IsSensitive)
	}

	return err
}
