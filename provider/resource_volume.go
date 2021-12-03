package provider

import (
	"fmt"
	"regexp"
	"strings"

	// "github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/richmondang/custom-tf/api/client"
	"github.com/richmondang/custom-tf/api/server"
)


func resourceVolume() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"volume_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: false,
				Description:  "Unique volume ID",
				ValidateFunc: validateVolumeId,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Description:  "Volume Name",
				Optional: true,
			},
			"appliance_id": &schema.Schema{
				Type:     schema.TypeString,
				Description:  "ADSS appliance ID",
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
		Create: resourceCreateVolume,
		Read:   resourceReadVolume,
		Update: resourceUpdateVolume,
		Delete: resourceDeleteVolume,
		Exists: resourceExistsVolume,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateVolume(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)


	volume := server.Volume{
		ID:        d.Get("volume_id").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		ApplianceID: d.Get("appliance_id").(string),
		Size: d.Get("size").(int),
	}

	err := apiClient.NewVolume(&volume)

	if err != nil {
		return err
	}
	d.SetId(volume.ID)
	return nil
}

func resourceReadVolume(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	volumeId := d.Id()
	volume, err := apiClient.GetVolume(volumeId)
	if err != nil {
		if strings.Contains(err.Error(), "volume not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding volume with ID %s", volumeId)
		}
	}

	d.SetId(volume.ID)
	d.Set("name", volume.Name)
	d.Set("description", volume.Description)
	d.Set("appliance_id", volume.ApplianceID)
	d.Set("size", volume.Size)
	
	return nil
}

func resourceUpdateVolume(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	volume := server.Volume{
		ID:        d.Get("volume_id").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		ApplianceID: d.Get("appliance_id").(string),
		Size: d.Get("size").(int),
	}

	err := apiClient.UpdateVolume(&volume)
	if err != nil {
		return err
	}
	return nil
}

func resourceDeleteVolume(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	volumeId := d.Id()

	err := apiClient.DeleteVolume(volumeId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsVolume(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	volumeId := d.Id()
	_, err := apiClient.GetVolume(volumeId)
	if err != nil {
		if strings.Contains(err.Error(), "volume not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func validateVolumeId(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs
}