package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func AzureDevStatusDetailModel(d *schema.ResourceData) *models.AzureDevStatusDetail {
	aState, _ := d.Get("a_state").(string)
	connectionStateUpdatedTime, _ := d.Get("connection_state_updated_time").(string)
	lastActivityTime, _ := d.Get("last_activity_time").(string)
	oState, _ := d.Get("o_state").(string)
	statusUpdateTime, _ := d.Get("status_update_time").(string)
	var twin *models.TwinDetail // TwinDetail
	twinInterface, twinIsSet := d.GetOk("twin")
	if twinIsSet && twinInterface != nil {
		twinMap := twinInterface.([]interface{})
		if len(twinMap) > 0 {
			twin = TwinDetailModelFromMap(twinMap[0].(map[string]interface{}))
		}
	}
	return &models.AzureDevStatusDetail{
		AState:                     aState,
		ConnectionStateUpdatedTime: connectionStateUpdatedTime,
		LastActivityTime:           lastActivityTime,
		OState:                     oState,
		StatusUpdateTime:           statusUpdateTime,
		Twin:                       twin,
	}
}

func AzureDevStatusDetailModelFromMap(m map[string]interface{}) *models.AzureDevStatusDetail {
	aState := m["a_state"].(string)
	connectionStateUpdatedTime := m["connection_state_updated_time"].(string)
	lastActivityTime := m["last_activity_time"].(string)
	oState := m["o_state"].(string)
	statusUpdateTime := m["status_update_time"].(string)
	var twin *models.TwinDetail // TwinDetail
	twinInterface, twinIsSet := m["twin"]
	if twinIsSet && twinInterface != nil {
		twinMap := twinInterface.([]interface{})
		if len(twinMap) > 0 {
			twin = TwinDetailModelFromMap(twinMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AzureDevStatusDetail{
		AState:                     aState,
		ConnectionStateUpdatedTime: connectionStateUpdatedTime,
		LastActivityTime:           lastActivityTime,
		OState:                     oState,
		StatusUpdateTime:           statusUpdateTime,
		Twin:                       twin,
	}
}

// Update the underlying AzureDevStatusDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAzureDevStatusDetailResourceData(d *schema.ResourceData, m *models.AzureDevStatusDetail) {
	d.Set("a_state", m.AState)
	d.Set("connection_state_updated_time", m.ConnectionStateUpdatedTime)
	d.Set("last_activity_time", m.LastActivityTime)
	d.Set("o_state", m.OState)
	d.Set("status_update_time", m.StatusUpdateTime)
	d.Set("twin", SetTwinDetailSubResourceData([]*models.TwinDetail{m.Twin}))
}

// Iterate through and update the AzureDevStatusDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAzureDevStatusDetailSubResourceData(m []*models.AzureDevStatusDetail) (d []*map[string]interface{}) {
	for _, AzureDevStatusDetailModel := range m {
		if AzureDevStatusDetailModel != nil {
			properties := make(map[string]interface{})
			properties["a_state"] = AzureDevStatusDetailModel.AState
			properties["connection_state_updated_time"] = AzureDevStatusDetailModel.ConnectionStateUpdatedTime
			properties["last_activity_time"] = AzureDevStatusDetailModel.LastActivityTime
			properties["o_state"] = AzureDevStatusDetailModel.OState
			properties["status_update_time"] = AzureDevStatusDetailModel.StatusUpdateTime
			properties["twin"] = SetTwinDetailSubResourceData([]*models.TwinDetail{AzureDevStatusDetailModel.Twin})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AzureDevStatusDetail resource defined in the Terraform configuration
func AzureDevStatusDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"a_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"connection_state_updated_time": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"last_activity_time": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"o_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"status_update_time": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"twin": {
			Description: ``,
			Type:        schema.TypeList, //GoType: TwinDetail
			Elem: &schema.Resource{
				Schema: TwinDetailSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the AzureDevStatusDetail resource
func GetAzureDevStatusDetailPropertyFields() (t []string) {
	return []string{
		"a_state",
		"connection_state_updated_time",
		"last_activity_time",
		"o_state",
		"status_update_time",
		"twin",
	}
}
