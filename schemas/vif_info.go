package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func VifInfoModel(d *schema.ResourceData) *models.VifInfo {
	appName, _ := d.Get("app_name").(string)
	macAddress, _ := d.Get("mac_address").(string)
	vifName, _ := d.Get("vif_name").(string)
	return &models.VifInfo{
		AppName:    appName,
		MacAddress: macAddress,
		VifName:    vifName,
	}
}

func VifInfoModelFromMap(m map[string]interface{}) *models.VifInfo {
	appName := m["app_name"].(string)
	macAddress := m["mac_address"].(string)
	vifName := m["vif_name"].(string)
	return &models.VifInfo{
		AppName:    appName,
		MacAddress: macAddress,
		VifName:    vifName,
	}
}

// Update the underlying VifInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVifInfoResourceData(d *schema.ResourceData, m *models.VifInfo) {
	d.Set("app_name", m.AppName)
	d.Set("mac_address", m.MacAddress)
	d.Set("vif_name", m.VifName)
}

// Iterate through and update the VifInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVifInfoSubResourceData(m []*models.VifInfo) (d []*map[string]interface{}) {
	for _, VifInfoModel := range m {
		if VifInfoModel != nil {
			properties := make(map[string]interface{})
			properties["app_name"] = VifInfoModel.AppName
			properties["mac_address"] = VifInfoModel.MacAddress
			properties["vif_name"] = VifInfoModel.VifName
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VifInfo resource defined in the Terraform configuration
func VifInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"mac_address": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"vif_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VifInfo resource
func GetVifInfoPropertyFields() (t []string) {
	return []string{
		"app_name",
		"mac_address",
		"vif_name",
	}
}
