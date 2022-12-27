package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate DeviceStatusConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeviceStatusConfigModel(d *schema.ResourceData) *models.DeviceStatusConfig {
	adminState := d.Get("admin_state").(*models.AdminState) // AdminState
	appInstCount := d.Get("app_inst_count").(string)
	var dinfo *models.DeviceInfo // DeviceInfo
	dinfoInterface, dinfoIsSet := d.GetOk("dinfo")
	if dinfoIsSet {
		dinfoMap := dinfoInterface.([]interface{})[0].(map[string]interface{})
		dinfo = DeviceInfoModelFromMap(dinfoMap)
	}
	eveImageName := d.Get("eve_image_name").(string)
	id := d.Get("id").(string)
	isEveLatest := d.Get("is_eve_latest").(string)
	location := d.Get("location").(string)
	name := d.Get("name").(string)
	projectID := d.Get("project_id").(string)
	projectName := d.Get("project_name").(string)
	runState := d.Get("run_state").(*models.RunState) // RunState
	serialNo := d.Get("serial_no").(string)
	swInfo := d.Get("sw_info").([]*models.DeviceSWInfo) // []*DeviceSWInfo
	title := d.Get("title").(string)
	return &models.DeviceStatusConfig{
		AdminState:   adminState,
		AppInstCount: appInstCount,
		Dinfo:        dinfo,
		EveImageName: eveImageName,
		ID:           id,
		IsEveLatest:  isEveLatest,
		Location:     location,
		Name:         &name, // string true false false
		ProjectID:    projectID,
		ProjectName:  projectName,
		RunState:     runState,
		SerialNo:     &serialNo, // string true false false
		SwInfo:       swInfo,
		Title:        title,
	}
}

func DeviceStatusConfigModelFromMap(m map[string]interface{}) *models.DeviceStatusConfig {
	adminState := m["admin_state"].(*models.AdminState) // AdminState
	appInstCount := m["app_inst_count"].(string)
	var dinfo *models.DeviceInfo // DeviceInfo
	dinfoInterface, dinfoIsSet := m["dinfo"]
	if dinfoIsSet {
		dinfoMap := dinfoInterface.([]interface{})[0].(map[string]interface{})
		dinfo = DeviceInfoModelFromMap(dinfoMap)
	}
	//
	eveImageName := m["eve_image_name"].(string)
	id := m["id"].(string)
	isEveLatest := m["is_eve_latest"].(string)
	location := m["location"].(string)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	projectName := m["project_name"].(string)
	runState := m["run_state"].(*models.RunState) // RunState
	serialNo := m["serial_no"].(string)
	swInfo := m["sw_info"].([]*models.DeviceSWInfo) // []*DeviceSWInfo
	title := m["title"].(string)
	return &models.DeviceStatusConfig{
		AdminState:   adminState,
		AppInstCount: appInstCount,
		Dinfo:        dinfo,
		EveImageName: eveImageName,
		ID:           id,
		IsEveLatest:  isEveLatest,
		Location:     location,
		Name:         &name,
		ProjectID:    projectID,
		ProjectName:  projectName,
		RunState:     runState,
		SerialNo:     &serialNo,
		SwInfo:       swInfo,
		Title:        title,
	}
}

// Update the underlying DeviceStatusConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeviceStatusConfigResourceData(d *schema.ResourceData, m *models.DeviceStatusConfig) {
	d.Set("admin_state", m.AdminState)
	d.Set("app_inst_count", m.AppInstCount)
	d.Set("dinfo", SetDeviceInfoSubResourceData([]*models.DeviceInfo{m.Dinfo}))
	d.Set("eve_image_name", m.EveImageName)
	d.Set("id", m.ID)
	d.Set("is_eve_latest", m.IsEveLatest)
	d.Set("location", m.Location)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("run_state", m.RunState)
	d.Set("serial_no", m.SerialNo)
	d.Set("sw_info", SetDeviceSWInfoSubResourceData(m.SwInfo))
	d.Set("title", m.Title)
}

// Iterate throught and update the DeviceStatusConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeviceStatusConfigSubResourceData(m []*models.DeviceStatusConfig) (d []*map[string]interface{}) {
	for _, DeviceStatusConfigModel := range m {
		if DeviceStatusConfigModel != nil {
			properties := make(map[string]interface{})
			properties["admin_state"] = DeviceStatusConfigModel.AdminState
			properties["app_inst_count"] = DeviceStatusConfigModel.AppInstCount
			properties["dinfo"] = SetDeviceInfoSubResourceData([]*models.DeviceInfo{DeviceStatusConfigModel.Dinfo})
			properties["eve_image_name"] = DeviceStatusConfigModel.EveImageName
			properties["id"] = DeviceStatusConfigModel.ID
			properties["is_eve_latest"] = DeviceStatusConfigModel.IsEveLatest
			properties["location"] = DeviceStatusConfigModel.Location
			properties["name"] = DeviceStatusConfigModel.Name
			properties["project_id"] = DeviceStatusConfigModel.ProjectID
			properties["project_name"] = DeviceStatusConfigModel.ProjectName
			properties["run_state"] = DeviceStatusConfigModel.RunState
			properties["serial_no"] = DeviceStatusConfigModel.SerialNo
			properties["sw_info"] = SetDeviceSWInfoSubResourceData(DeviceStatusConfigModel.SwInfo)
			properties["title"] = DeviceStatusConfigModel.Title
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the DeviceStatusConfig resource defined in the Terraform configuration
func DeviceStatusConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_state": {
			Description: `Admin state of the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_inst_count": {
			Description: `Number of app instance that is running on the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"dinfo": {
			Description: `Device info like storage, arch, memory etc`,
			Type:        schema.TypeList, //GoType: DeviceInfo
			Elem: &schema.Resource{
				Schema: DeviceInfoSchema(),
			},
			Optional: true,
		},

		"eve_image_name": {
			Description: `Eve image name that was running on the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `system generated unique id for a device`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"is_eve_latest": {
			Description: `Boolean that tells whether the active eve image is latest or not`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"location": {
			Description: `Device location`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `user specified device name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project_id": {
			Description: `project Id of the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name": {
			Description: `project name of the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"run_state": {
			Description: `Run state of the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"serial_no": {
			Description: `Device serial number`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"sw_info": {
			Description: `software info reported by the device`,
			Type:        schema.TypeList, //GoType: []*DeviceSWInfo
			Elem: &schema.Resource{
				Schema: DeviceSWInfoSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"title": {
			Description: `user specified title`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DeviceStatusConfig resource
func GetDeviceStatusConfigPropertyFields() (t []string) {
	return []string{
		"admin_state",
		"app_inst_count",
		"dinfo",
		"eve_image_name",
		"id",
		"is_eve_latest",
		"location",
		"name",
		"project_id",
		"project_name",
		"run_state",
		"serial_no",
		"sw_info",
		"title",
	}
}
