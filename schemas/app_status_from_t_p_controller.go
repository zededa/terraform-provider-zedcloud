package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func AppStatusFromTPControllerModel(d *schema.ResourceData) *models.AppStatusFromTPController {
	appID, _ := d.Get("app_id").(string)
	appName, _ := d.Get("app_name").(string)
	var azureStatus *models.AzureStatus // AzureStatus
	azureStatusInterface, azureStatusIsSet := d.GetOk("azure_status")
	if azureStatusIsSet && azureStatusInterface != nil {
		azureStatusMap := azureStatusInterface.([]interface{})
		if len(azureStatusMap) > 0 {
			azureStatus = AzureStatusModelFromMap(azureStatusMap[0].(map[string]interface{}))
		}
	}
	enterpriseID, _ := d.Get("enterprise_id").(string)
	eveDeviceID, _ := d.Get("eve_device_id").(string)
	var typeVar *models.ControllerType // ControllerType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewControllerType(models.ControllerType(typeModel))
	}
	var vceStatus *models.VCEStatus // VCEStatus
	vceStatusInterface, vceStatusIsSet := d.GetOk("vce_status")
	if vceStatusIsSet && vceStatusInterface != nil {
		vceStatusMap := vceStatusInterface.([]interface{})
		if len(vceStatusMap) > 0 {
			vceStatus = VCEStatusModelFromMap(vceStatusMap[0].(map[string]interface{}))
		}
	}
	return &models.AppStatusFromTPController{
		AppID:        appID,
		AppName:      appName,
		AzureStatus:  azureStatus,
		EnterpriseID: enterpriseID,
		EveDeviceID:  eveDeviceID,
		Type:         typeVar,
		VceStatus:    vceStatus,
	}
}

func AppStatusFromTPControllerModelFromMap(m map[string]interface{}) *models.AppStatusFromTPController {
	appID := m["app_id"].(string)
	appName := m["app_name"].(string)
	var azureStatus *models.AzureStatus // AzureStatus
	azureStatusInterface, azureStatusIsSet := m["azure_status"]
	if azureStatusIsSet && azureStatusInterface != nil {
		azureStatusMap := azureStatusInterface.([]interface{})
		if len(azureStatusMap) > 0 {
			azureStatus = AzureStatusModelFromMap(azureStatusMap[0].(map[string]interface{}))
		}
	}
	//
	enterpriseID := m["enterprise_id"].(string)
	eveDeviceID := m["eve_device_id"].(string)
	var typeVar *models.ControllerType // ControllerType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewControllerType(models.ControllerType(typeModel))
	}
	var vceStatus *models.VCEStatus // VCEStatus
	vceStatusInterface, vceStatusIsSet := m["vce_status"]
	if vceStatusIsSet && vceStatusInterface != nil {
		vceStatusMap := vceStatusInterface.([]interface{})
		if len(vceStatusMap) > 0 {
			vceStatus = VCEStatusModelFromMap(vceStatusMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.AppStatusFromTPController{
		AppID:        appID,
		AppName:      appName,
		AzureStatus:  azureStatus,
		EnterpriseID: enterpriseID,
		EveDeviceID:  eveDeviceID,
		Type:         typeVar,
		VceStatus:    vceStatus,
	}
}

// Update the underlying AppStatusFromTPController resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppStatusFromTPControllerResourceData(d *schema.ResourceData, m *models.AppStatusFromTPController) {
	d.Set("app_id", m.AppID)
	d.Set("app_name", m.AppName)
	d.Set("azure_status", SetAzureStatusSubResourceData([]*models.AzureStatus{m.AzureStatus}))
	d.Set("enterprise_id", m.EnterpriseID)
	d.Set("eve_device_id", m.EveDeviceID)
	d.Set("type", m.Type)
	d.Set("vce_status", SetVCEStatusSubResourceData([]*models.VCEStatus{m.VceStatus}))
}

// Iterate through and update the AppStatusFromTPController resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppStatusFromTPControllerSubResourceData(m []*models.AppStatusFromTPController) (d []*map[string]interface{}) {
	for _, AppStatusFromTPControllerModel := range m {
		if AppStatusFromTPControllerModel != nil {
			properties := make(map[string]interface{})
			properties["app_id"] = AppStatusFromTPControllerModel.AppID
			properties["app_name"] = AppStatusFromTPControllerModel.AppName
			properties["azure_status"] = SetAzureStatusSubResourceData([]*models.AzureStatus{AppStatusFromTPControllerModel.AzureStatus})
			properties["enterprise_id"] = AppStatusFromTPControllerModel.EnterpriseID
			properties["eve_device_id"] = AppStatusFromTPControllerModel.EveDeviceID
			properties["type"] = AppStatusFromTPControllerModel.Type
			properties["vce_status"] = SetVCEStatusSubResourceData([]*models.VCEStatus{AppStatusFromTPControllerModel.VceStatus})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppStatusFromTPController resource defined in the Terraform configuration
func AppStatusFromTPControllerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"azure_status": {
			Description: ``,
			Type:        schema.TypeList, //GoType: AzureStatus
			Elem: &schema.Resource{
				Schema: AzureStatusSchema(),
			},
			Optional: true,
		},

		"enterprise_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"eve_device_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"type": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"vce_status": {
			Description: ``,
			Type:        schema.TypeList, //GoType: VCEStatus
			Elem: &schema.Resource{
				Schema: VCEStatusSchema(),
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the AppStatusFromTPController resource
func GetAppStatusFromTPControllerPropertyFields() (t []string) {
	return []string{
		"app_id",
		"app_name",
		"azure_status",
		"enterprise_id",
		"eve_device_id",
		"type",
		"vce_status",
	}
}
