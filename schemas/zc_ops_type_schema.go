package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ZcOpsType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ZcOpsTypeModel(d *schema.ResourceData) *models.ZcOpsType {
	zcOpsType := d.Get("zc_ops_type").(models.ZcOpsType)
	return &zcOpsType
}

func ZcOpsTypeModelFromMap(m map[string]interface{}) *models.ZcOpsType {
	zcOpsType := m["zc_ops_type"].(models.ZcOpsType)
	return &zcOpsType
}

// Update the underlying ZcOpsType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetZcOpsTypeResourceData(d *schema.ResourceData, m *models.ZcOpsType) {
}

// Iterate throught and update the ZcOpsType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetZcOpsTypeSubResourceData(m []*models.ZcOpsType) (d []*map[string]interface{}) {
	for _, ZcOpsTypeModel := range m {
		if ZcOpsTypeModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ZcOpsType resource defined in the Terraform configuration
func ZcOpsTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the ZcOpsType resource
func GetZcOpsTypePropertyFields() (t []string) {
	return []string{}
}
