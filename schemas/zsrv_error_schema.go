package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ZsrvErrorModel(d *schema.ResourceData) *models.ZsrvError {
	details, _ := d.Get("details").(string)
	var ec *models.ZsrvErrorCode // ZsrvErrorCode
	ecInterface, ecIsSet := d.GetOk("ec")
	if ecIsSet {
		ecModel := ecInterface.(string)
		ec = models.NewZsrvErrorCode(models.ZsrvErrorCode(ecModel))
	}
	location, _ := d.Get("location").(string)
	return &models.ZsrvError{
		Details:   details,
		ErrorCode: ec,
		Location:  location,
	}
}

func ZsrvErrorModelFromMap(m map[string]interface{}) *models.ZsrvError {
	details := m["details"].(string)
	var ec *models.ZsrvErrorCode // ZsrvErrorCode
	ecInterface, ecIsSet := m["ec"]
	if ecIsSet {
		ecModel := ecInterface.(string)
		ec = models.NewZsrvErrorCode(models.ZsrvErrorCode(ecModel))
	}
	location := m["location"].(string)
	return &models.ZsrvError{
		Details:   details,
		ErrorCode: ec,
		Location:  location,
	}
}

// Update the underlying ZsrvError resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetZsrvErrorResourceData(d *schema.ResourceData, m *models.ZsrvError) {
	d.Set("details", m.Details)
	d.Set("ec", m.ErrorCode)
	d.Set("location", m.Location)
}

// Iterate through and update the ZsrvError resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetZsrvErrorSubResourceData(m []*models.ZsrvError) (d []*map[string]interface{}) {
	for _, ZsrvErrorModel := range m {
		if ZsrvErrorModel != nil {
			properties := make(map[string]interface{})
			properties["details"] = ZsrvErrorModel.Details
			properties["ec"] = ZsrvErrorModel.ErrorCode
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
			Type:        schema.TypeString,
			Optional:    true,
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
