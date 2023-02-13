package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ZedCloudOpsCmd resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ZedCloudOpsCmdModel(d *schema.ResourceData) *models.ZedCloudOpsCmd {
	counterInt, _ := d.Get("counter").(int)
	counter := int64(counterInt)
	opsTime, _ := d.Get("ops_time").(interface{}) // interface{}
	return &models.ZedCloudOpsCmd{
		Counter: counter,
		OpsTime: opsTime,
	}
}

func ZedCloudOpsCmdModelFromMap(m map[string]interface{}) *models.ZedCloudOpsCmd {
	counter := int64(m["counter"].(int)) // int64 false false false
	opsTime := m["ops_time"].(interface{})
	return &models.ZedCloudOpsCmd{
		Counter: counter,
		OpsTime: opsTime,
	}
}

// Update the underlying ZedCloudOpsCmd resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetZedCloudOpsCmdResourceData(d *schema.ResourceData, m *models.ZedCloudOpsCmd) {
	d.Set("counter", m.Counter)
	d.Set("ops_time", m.OpsTime)
}

// Iterate through and update the ZedCloudOpsCmd resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetZedCloudOpsCmdSubResourceData(m []*models.ZedCloudOpsCmd) (d []*map[string]interface{}) {
	for _, ZedCloudOpsCmdModel := range m {
		if ZedCloudOpsCmdModel != nil {
			properties := make(map[string]interface{})
			properties["counter"] = ZedCloudOpsCmdModel.Counter
			properties["ops_time"] = ZedCloudOpsCmdModel.OpsTime
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ZedCloudOpsCmd resource defined in the Terraform configuration
func ZedCloudOpsCmdSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"counter": {
			Description: `counter`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"ops_time": {
			Description: `Timestamp: Operational time`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the ZedCloudOpsCmd resource
func GetZedCloudOpsCmdPropertyFields() (t []string) {
	return []string{
		"counter",
		"ops_time",
	}
}
