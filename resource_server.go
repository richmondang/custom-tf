package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"net/http"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"server_count": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {

	server_count := d.Get("server_count").(string)

	d.SetId(server_count)

	// https://www.uuidtools.com/api/generate/v1/count/uuid_count
	resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/" + server_count)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

// func resourceRestAPIDelete(d *schema.ResourceData, meta interface{}) error {
// 	obj, err := makeAPIObject(d, meta)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("resource_api_object.go: Delete routine called. Object built:\n%s\n", obj.toString())

// 	err = obj.deleteObject()
// 	if err != nil {
// 		if strings.Contains(err.Error(), "404") {
// 			/* 404 means it doesn't exist. Call that good enough */
// 			err = nil
// 		}
// 	}
// 	return err
// }

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
