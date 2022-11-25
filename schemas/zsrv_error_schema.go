package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ZsrvError resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ZsrvErrorModel(d *schema.ResourceData) *models.ZsrvError {
	details := d.Get("details").(string)
	ec := d.Get("ec").(*models.ZsrvErrorCode) // ZsrvErrorCode
	location := d.Get("location").(string)
	return &models.ZsrvError{
		Details:  details,
		Ec:       ec,
		Location: location,
	}
}

func ZsrvErrorModelFromMap(m map[string]interface{}) *models.ZsrvError {
	details := m["details"].(string)
	ec := m["ec"].(*models.ZsrvErrorCode) // ZsrvErrorCode
	location := m["location"].(string)
	return &models.ZsrvError{
		Details:  details,
		Ec:       ec,
		Location: location,
	}
}

// Update the underlying ZsrvError resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetZsrvErrorResourceData(d *schema.ResourceData, m *models.ZsrvError) {
	d.Set("details", m.Details)
	d.Set("ec", m.Ec)
	d.Set("location", m.Location)
}

// Iterate throught and update the ZsrvError resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetZsrvErrorSubResourceData(m []*models.ZsrvError) (d []*map[string]interface{}) {
	for _, ZsrvErrorModel := range m {
		if ZsrvErrorModel != nil {
			properties := make(map[string]interface{})
			properties["details"] = ZsrvErrorModel.Details
			properties["ec"] = ZsrvErrorModel.Ec
			properties["location"] = ZsrvErrorModel.Location
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ZsrvError resource defined in the Terraform configuration
func ZsrvErrorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"details": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ec": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZsrvErrorCode
			Elem: &schema.Resource{
				Schema: ZsrvErrorCodeSchema(),
			},
			Optional: true,
		},

		"location": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ZsrvError resource
func GetZsrvErrorPropertyFields() (t []string) {
	return []string{
		"details",
		"ec",
		"location",
	}
}
