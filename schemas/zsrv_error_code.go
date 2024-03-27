package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ZsrvErrorCodeModel(d *schema.ResourceData) *models.ZsrvErrorCode {
	zsrvErrorCode, _ := d.Get("zsrv_error_code").(models.ZsrvErrorCode)
	return &zsrvErrorCode
}

func ZsrvErrorCodeModelFromMap(m map[string]interface{}) *models.ZsrvErrorCode {
	zsrvErrorCode := m["zsrv_error_code"].(models.ZsrvErrorCode)
	return &zsrvErrorCode
}

// Update the underlying ZsrvErrorCode resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetZsrvErrorCodeResourceData(d *schema.ResourceData, m *models.ZsrvErrorCode) {
}

// Iterate through and update the ZsrvErrorCode resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetZsrvErrorCodeSubResourceData(m []*models.ZsrvErrorCode) (d []*map[string]interface{}) {
	for _, ZsrvErrorCodeModel := range m {
		if ZsrvErrorCodeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ZsrvErrorCode resource defined in the Terraform configuration
func ZsrvErrorCodeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the ZsrvErrorCode resource
func GetZsrvErrorCodePropertyFields() (t []string) {
	return []string{}
}
