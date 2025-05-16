package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func RPCStatusModel(d *schema.ResourceData) *models.RPCStatus {
	codeInt, _ := d.Get("code").(int)
	code := int32(codeInt)
	var details []*models.ProtobufAny // []*ProtobufAny
	detailsInterface, detailsIsSet := d.GetOk("details")
	if detailsIsSet {
		var items []interface{}
		if listItems, isList := detailsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = detailsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ProtobufAnyModelFromMap(v.(map[string]interface{}))
			details = append(details, m)
		}
	}
	message, _ := d.Get("message").(string)
	return &models.RPCStatus{
		Code:    code,
		Details: details,
		Message: message,
	}
}

func RPCStatusModelFromMap(m map[string]interface{}) *models.RPCStatus {
	code := int32(m["code"].(int))    // int32
	var details []*models.ProtobufAny // []*ProtobufAny
	detailsInterface, detailsIsSet := m["details"]
	if detailsIsSet {
		var items []interface{}
		if listItems, isList := detailsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = detailsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ProtobufAnyModelFromMap(v.(map[string]interface{}))
			details = append(details, m)
		}
	}
	message := m["message"].(string)
	return &models.RPCStatus{
		Code:    code,
		Details: details,
		Message: message,
	}
}

func SetRPCStatusResourceData(d *schema.ResourceData, m *models.RPCStatus) {
	d.Set("code", m.Code)
	d.Set("details", SetProtobufAnySubResourceData(m.Details))
	d.Set("message", m.Message)
}

func SetRPCStatusSubResourceData(m []*models.RPCStatus) (d []*map[string]interface{}) {
	for _, RPCStatusModel := range m {
		if RPCStatusModel != nil {
			properties := make(map[string]interface{})
			properties["code"] = RPCStatusModel.Code
			properties["details"] = SetProtobufAnySubResourceData(RPCStatusModel.Details)
			properties["message"] = RPCStatusModel.Message
			d = append(d, &properties)
		}
	}
	return
}

func RPCStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"code": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"details": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*ProtobufAny
			Elem: &schema.Resource{
				Schema: ProtobufAnySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"message": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetRPCStatusPropertyFields() (t []string) {
	return []string{
		"code",
		"details",
		"message",
	}
}
