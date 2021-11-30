package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Resources and data sources names must follow <provider>_<resource_name> convention
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"apex_example_server": resourceServer(),
			"apex_example_volume": resourceVolume(),
		},
	}
}
