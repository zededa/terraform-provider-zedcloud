package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ZcOpsStatusModel(d *schema.ResourceData) *models.ZcOpsStatus {
	zcOpsStatus, _ := d.Get("zc_ops_status").(models.ZcOpsStatus)
	return &zcOpsStatus
}

func ZcOpsStatusModelFromMap(m map[string]interface{}) *models.ZcOpsStatus {
	zcOpsStatus := m["zc_ops_status"].(models.ZcOpsStatus)
	return &zcOpsStatus
}

// Update the underlying ZcOpsStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetZcOpsStatusResourceData(d *schema.ResourceData, m *models.ZcOpsStatus) {
}

// Iterate through and update the ZcOpsStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetZcOpsStatusSubResourceData(m []*models.ZcOpsStatus) (d []*map[string]interface{}) {
	for _, ZcOpsStatusModel := range m {
		if ZcOpsStatusModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ZcOpsStatus resource defined in the Terraform configuration
func ZcOpsStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the ZcOpsStatus resource
func GetZcOpsStatusPropertyFields() (t []string) {
	return []string{}
}
