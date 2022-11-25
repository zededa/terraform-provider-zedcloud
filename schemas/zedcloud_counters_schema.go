package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ZedcloudCounters resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ZedcloudCountersModel(d *schema.ResourceData) *models.ZedcloudCounters {
	failures := d.Get("failures").(uint64)
	ifName := d.Get("if_name").(string)
	lastFailure := d.Get("last_failure").(interface{}) // interface{}
	lastSuccess := d.Get("last_success").(interface{}) // interface{}
	success := d.Get("success").(uint64)
	return &models.ZedcloudCounters{
		Failures:    failures,
		IfName:      ifName,
		LastFailure: lastFailure,
		LastSuccess: lastSuccess,
		Success:     success,
	}
}

func ZedcloudCountersModelFromMap(m map[string]interface{}) *models.ZedcloudCounters {
	failures := m["failures"].(uint64)
	ifName := m["if_name"].(string)
	lastFailure := m["last_failure"].(interface{})
	lastSuccess := m["last_success"].(interface{})
	success := m["success"].(uint64)
	return &models.ZedcloudCounters{
		Failures:    failures,
		IfName:      ifName,
		LastFailure: lastFailure,
		LastSuccess: lastSuccess,
		Success:     success,
	}
}

// Update the underlying ZedcloudCounters resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetZedcloudCountersResourceData(d *schema.ResourceData, m *models.ZedcloudCounters) {
	d.Set("failures", m.Failures)
	d.Set("if_name", m.IfName)
	d.Set("last_failure", m.LastFailure)
	d.Set("last_success", m.LastSuccess)
	d.Set("success", m.Success)
}

// Iterate throught and update the ZedcloudCounters resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetZedcloudCountersSubResourceData(m []*models.ZedcloudCounters) (d []*map[string]interface{}) {
	for _, ZedcloudCountersModel := range m {
		if ZedcloudCountersModel != nil {
			properties := make(map[string]interface{})
			properties["failures"] = ZedcloudCountersModel.Failures
			properties["if_name"] = ZedcloudCountersModel.IfName
			properties["last_failure"] = ZedcloudCountersModel.LastFailure
			properties["last_success"] = ZedcloudCountersModel.LastSuccess
			properties["success"] = ZedcloudCountersModel.Success
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ZedcloudCounters resource defined in the Terraform configuration
func ZedcloudCountersSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"failures": {
			Description: `Failures`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"if_name": {
			Description: `ifName`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"last_failure": {
			Description: `Timestamp of last failure`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"last_success": {
			Description: `Timestamp of last success`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"success": {
			Description: `Success`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ZedcloudCounters resource
func GetZedcloudCountersPropertyFields() (t []string) {
	return []string{
		"failures",
		"if_name",
		"last_failure",
		"last_success",
		"success",
	}
}
