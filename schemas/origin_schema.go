package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Origin resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func OriginModel(d *schema.ResourceData) *models.Origin {
	origin, _ := d.Get("origin").(models.Origin)
	return &origin
}

func OriginModelFromMap(m map[string]interface{}) *models.Origin {
	origin := m["origin"].(models.Origin)
	return &origin
}

// Update the underlying Origin resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetOriginResourceData(d *schema.ResourceData, m *models.Origin) {
}

// Iterate through and update the Origin resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetOriginSubResourceData(m []*models.Origin) (d []*map[string]interface{}) {
	for _, OriginModel := range m {
		if OriginModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Origin resource defined in the Terraform configuration
func OriginSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the Origin resource
func GetOriginPropertyFields() (t []string) {
	return []string{}
}
