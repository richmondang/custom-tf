package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	// "log"
	// "net/http"
)

// Volume - volume properties
// type Volume struct {
// 	ID                       string               `json:"id,omitempty"`
// 	Name                     string               `json:"name,omitempty"`
// 	Description              string               `json:"description,omitempty"`
// 	Size                     int                  `json:"size,omitempty"`
// }

// var Volumes = []Volume{
//     {
//         ID:     "123456",
//         Name:  "Test_Volume1",
//         Description: "APEX Data Storage Services Volume 1",
//         Size:   156150,
//     },
// }

func resourceVolume() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Create: resourceVolumeCreate,
		Read:   resourceVolumeRead,
		Update: resourceVolumeUpdate,
		Delete: resourceVolumeDelete,
		Schema: map[string]*schema.Schema{
			"resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceVolumeRead(d *schema.ResourceData, meta interface{}) error {

	return nil

}

//placeholder for resource volume create
func resourceVolumeCreate(d *schema.ResourceData, meta interface{}) error {

	resource_id := d.Get("resource_id").(string)
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	size := d.Get("size").(int)

	d.Set("resource_id", resource_id)
	d.Set("name", name)
	d.Set("description", description)
	d.Set("size", size)

	d.SetId(resource_id)

	return resourceVolumeRead(d, meta)
}

//placeholder for resource volume update
func resourceVolumeUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceVolumeRead(d, meta)
}

//placeholder for resource volume delete
func resourceVolumeDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")

	return nil
}
