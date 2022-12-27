package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate IoBundleStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func IoBundleStatusModel(d *schema.ResourceData) *models.IoBundleStatus {
	appName := d.Get("app_name").(string)
	var err *models.DeviceError // DeviceError
	errInterface, errIsSet := d.GetOk("err")
	if errIsSet {
		errMap := errInterface.([]interface{})[0].(map[string]interface{})
		err = DeviceErrorModelFromMap(errMap)
	}
	var lteInfo *models.LTEAdapter // LTEAdapter
	lteInfoInterface, lteInfoIsSet := d.GetOk("lte_info")
	if lteInfoIsSet {
		lteInfoMap := lteInfoInterface.([]interface{})[0].(map[string]interface{})
		lteInfo = LTEAdapterModelFromMap(lteInfoMap)
	}
	memberList := d.Get("member_list").([]*models.IoMemberStatus) // []*IoMemberStatus
	members := d.Get("members").([]string)
	name := d.Get("name").(string)
	typeVar := d.Get("type").(*models.IoType) // IoType
	return &models.IoBundleStatus{
		AppName:    &appName, // string true false false
		Err:        err,
		LteInfo:    lteInfo,
		MemberList: memberList,
		Members:    members,
		Name:       &name, // string true false false
		Type:       typeVar,
	}
}

func IoBundleStatusModelFromMap(m map[string]interface{}) *models.IoBundleStatus {
	appName := m["app_name"].(string)
	var err *models.DeviceError // DeviceError
	errInterface, errIsSet := m["err"]
	if errIsSet {
		errMap := errInterface.([]interface{})[0].(map[string]interface{})
		err = DeviceErrorModelFromMap(errMap)
	}
	//
	var lteInfo *models.LTEAdapter // LTEAdapter
	lteInfoInterface, lteInfoIsSet := m["lte_info"]
	if lteInfoIsSet {
		lteInfoMap := lteInfoInterface.([]interface{})[0].(map[string]interface{})
		lteInfo = LTEAdapterModelFromMap(lteInfoMap)
	}
	//
	memberList := m["member_list"].([]*models.IoMemberStatus) // []*IoMemberStatus
	members := m["members"].([]string)
	name := m["name"].(string)
	typeVar := m["type"].(*models.IoType) // IoType
	return &models.IoBundleStatus{
		AppName:    &appName,
		Err:        err,
		LteInfo:    lteInfo,
		MemberList: memberList,
		Members:    members,
		Name:       &name,
		Type:       typeVar,
	}
}

// Update the underlying IoBundleStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetIoBundleStatusResourceData(d *schema.ResourceData, m *models.IoBundleStatus) {
	d.Set("app_name", m.AppName)
	d.Set("err", SetDeviceErrorSubResourceData([]*models.DeviceError{m.Err}))
	d.Set("lte_info", SetLTEAdapterSubResourceData([]*models.LTEAdapter{m.LteInfo}))
	d.Set("member_list", SetIoMemberStatusSubResourceData(m.MemberList))
	d.Set("members", m.Members)
	d.Set("name", m.Name)
	d.Set("type", m.Type)
}

// Iterate throught and update the IoBundleStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetIoBundleStatusSubResourceData(m []*models.IoBundleStatus) (d []*map[string]interface{}) {
	for _, IoBundleStatusModel := range m {
		if IoBundleStatusModel != nil {
			properties := make(map[string]interface{})
			properties["app_name"] = IoBundleStatusModel.AppName
			properties["err"] = SetDeviceErrorSubResourceData([]*models.DeviceError{IoBundleStatusModel.Err})
			properties["lte_info"] = SetLTEAdapterSubResourceData([]*models.LTEAdapter{IoBundleStatusModel.LteInfo})
			properties["member_list"] = SetIoMemberStatusSubResourceData(IoBundleStatusModel.MemberList)
			properties["members"] = IoBundleStatusModel.Members
			properties["name"] = IoBundleStatusModel.Name
			properties["type"] = IoBundleStatusModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the IoBundleStatus resource defined in the Terraform configuration
func IoBundleStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_name": {
			Description: `Application name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"err": {
			Description: `Device error details`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"lte_info": {
			Description: `LTE information`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"member_list": {
			Description: `List of IO members`,
			Type:        schema.TypeList, //GoType: []*IoMemberStatus
			Elem: &schema.Resource{
				Schema: IoMemberStatusSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"members": {
			Description: `Member Array`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},

		"name": {
			Description: `Io Bundle status name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"type": {
			Description: `IoType specifies the type of the Input output of the device`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Retrieve property field names for updating the IoBundleStatus resource
func GetIoBundleStatusPropertyFields() (t []string) {
	return []string{
		"app_name",
		"err",
		"lte_info",
		"member_list",
		"members",
		"name",
		"type",
	}
}
