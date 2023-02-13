package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ContainerDetail resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ContainerDetailModel(d *schema.ResourceData) *models.ContainerDetail {
	containerCreateOption, _ := d.Get("container_create_option").(string)
	return &models.ContainerDetail{
		ContainerCreateOption: containerCreateOption,
	}
}

func ContainerDetailModelFromMap(m map[string]interface{}) *models.ContainerDetail {
	containerCreateOption := m["container_create_option"].(string)
	return &models.ContainerDetail{
		ContainerCreateOption: containerCreateOption,
	}
}

// Update the underlying ContainerDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetContainerDetailResourceData(d *schema.ResourceData, m *models.ContainerDetail) {
	d.Set("container_create_option", m.ContainerCreateOption)
}

// Iterate through and update the ContainerDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetContainerDetailSubResourceData(m []*models.ContainerDetail) (d []*map[string]interface{}) {
	for _, ContainerDetailModel := range m {
		if ContainerDetailModel != nil {
			properties := make(map[string]interface{})
			properties["container_create_option"] = ContainerDetailModel.ContainerCreateOption
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ContainerDetail resource defined in the Terraform configuration
func ContainerDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"container_create_option": {
			Description: `Create options direct the creation of the Docker container`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ContainerDetail resource
func GetContainerDetailPropertyFields() (t []string) {
	return []string{
		"container_create_option",
	}
}
