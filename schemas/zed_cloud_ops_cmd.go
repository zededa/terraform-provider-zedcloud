package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ZedCloudOpsCmdModel(d *schema.ResourceData) *models.ZedCloudOpsCmd {
	counterInt, _ := d.Get("counter").(int)
	counter := int64(counterInt)

	// opsTime := map[string]string{}
	// opsTimeInterface, opsTimeIsSet := d.GetOk("ops_time")
	// if opsTimeIsSet {
	// 	tagsMap := opsTimeInterface.(map[string]interface{})
	// 	for k, v := range tagsMap {
	// 		if v == nil {
	// 			continue
	// 		}
	// 		opsTime[k] = v.(string)
	// 	}
	// }
	return &models.ZedCloudOpsCmd{
		Counter: counter,
		// OpsTime: opsTime,
	}
}

func ZedCloudOpsCmdModelFromMap(m map[string]interface{}) *models.ZedCloudOpsCmd {
	counter := int64(m["counter"].(int)) // int64 false false false
	// opsTime := map[string]string{}
	// opsTimeInterface, opsTimeIsSet := m["ops_time"]
	// if opsTimeIsSet {
	// 	tagsMap := opsTimeInterface.(map[string]interface{})
	// 	for k, v := range tagsMap {
	// 		if v == nil {
	// 			continue
	// 		}
	// 		opsTime[k] = v.(string)
	// 	}
	// }
	return &models.ZedCloudOpsCmd{
		Counter: counter,
		// OpsTime: opsTime,
	}
}

func SetZedCloudOpsCmdResourceData(d *schema.ResourceData, m *models.ZedCloudOpsCmd) {
	d.Set("counter", m.Counter)
	// d.Set("ops_time", m.OpsTime)
}

func SetZedCloudOpsCmdSubResourceData(m []*models.ZedCloudOpsCmd) (d []*map[string]interface{}) {
	for _, ZedCloudOpsCmdModel := range m {
		if ZedCloudOpsCmdModel != nil {
			properties := make(map[string]interface{})
			properties["counter"] = ZedCloudOpsCmdModel.Counter
			// properties["ops_time"] = ZedCloudOpsCmdModel.OpsTime
			d = append(d, &properties)
		}
	}
	return
}

func ZedCloudOpsCmdSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"counter": {
			Description: `counter`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		// "ops_time": {
		// 	Description: `Timestamp: Operational time`,
		// 	Type:        schema.TypeString, //GoType: interface{}
		// 	Optional:    true,
		// },
	}
}

// Retrieve property field names for updating the ZedCloudOpsCmd resource
func GetZedCloudOpsCmdPropertyFields() (t []string) {
	return []string{
		"counter",
		"ops_time",
	}
}
