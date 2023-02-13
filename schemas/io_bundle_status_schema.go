package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func IoBundleStatusModel(d *schema.ResourceData) *models.IoBundleStatus {
	appName, _ := d.Get("app_name").(string)
	var err *models.DeviceError // DeviceError
	errInterface, errIsSet := d.GetOk("err")
	if errIsSet && errInterface != nil {
		errMap := errInterface.([]interface{})
		if len(errMap) > 0 {
			err = DeviceErrorModelFromMap(errMap[0].(map[string]interface{}))
		}
	}
	var lteInfo *models.LTEAdapter // LTEAdapter
	lteInfoInterface, lteInfoIsSet := d.GetOk("lte_info")
	if lteInfoIsSet && lteInfoInterface != nil {
		lteInfoMap := lteInfoInterface.([]interface{})
		if len(lteInfoMap) > 0 {
			lteInfo = LTEAdapterModelFromMap(lteInfoMap[0].(map[string]interface{}))
		}
	}
	var memberList []*models.IoMemberStatus // []*IoMemberStatus
	memberListInterface, memberListIsSet := d.GetOk("member_list")
	if memberListIsSet {
		var items []interface{}
		if listItems, isList := memberListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = memberListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IoMemberStatusModelFromMap(v.(map[string]interface{}))
			memberList = append(memberList, m)
		}
	}
	var members []string
	membersInterface, membersIsSet := d.GetOk("members")
	if membersIsSet {
		membersSlice := membersInterface.([]interface{})
		for _, i := range membersSlice {
			membersSlice = append(membersSlice, i.(string))
		}
	}
	name, _ := d.Get("name").(string)
	var typeVar *models.IoType // IoType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewIoType(models.IoType(typeModel))
	}
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
	if errIsSet && errInterface != nil {
		errMap := errInterface.([]interface{})
		if len(errMap) > 0 {
			err = DeviceErrorModelFromMap(errMap[0].(map[string]interface{}))
		}
	}
	//
	var lteInfo *models.LTEAdapter // LTEAdapter
	lteInfoInterface, lteInfoIsSet := m["lte_info"]
	if lteInfoIsSet && lteInfoInterface != nil {
		lteInfoMap := lteInfoInterface.([]interface{})
		if len(lteInfoMap) > 0 {
			lteInfo = LTEAdapterModelFromMap(lteInfoMap[0].(map[string]interface{}))
		}
	}
	//
	var memberList []*models.IoMemberStatus // []*IoMemberStatus
	memberListInterface, memberListIsSet := m["member_list"]
	if memberListIsSet {
		var items []interface{}
		if listItems, isList := memberListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = memberListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IoMemberStatusModelFromMap(v.(map[string]interface{}))
			memberList = append(memberList, m)
		}
	}
	var members []string
	membersInterface, membersIsSet := m["members"]
	if membersIsSet {
		membersSlice := membersInterface.([]interface{})
		for _, i := range membersSlice {
			membersSlice = append(membersSlice, i.(string))
		}
	}
	name := m["name"].(string)
	var typeVar *models.IoType // IoType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewIoType(models.IoType(typeModel))
	}
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

// Iterate through and update the IoBundleStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
			Type:        schema.TypeList, //GoType: DeviceError
			Elem: &schema.Resource{
				Schema: DeviceErrorSchema(),
			},
			Optional: true,
		},

		"lte_info": {
			Description: `LTE information`,
			Type:        schema.TypeList, //GoType: LTEAdapter
			Elem: &schema.Resource{
				Schema: LTEAdapterSchema(),
			},
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
			Type:        schema.TypeString,
			Required:    true,
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
