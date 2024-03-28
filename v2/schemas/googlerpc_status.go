package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GooglerpcStatusModel(d *schema.ResourceData) *models.GooglerpcStatus {
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
	return &models.GooglerpcStatus{
		Code:    code,
		Details: details,
		Message: message,
	}
}

func GooglerpcStatusModelFromMap(m map[string]interface{}) *models.GooglerpcStatus {
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
	return &models.GooglerpcStatus{
		Code:    code,
		Details: details,
		Message: message,
	}
}

// Update the underlying GooglerpcStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetGooglerpcStatusResourceData(d *schema.ResourceData, m *models.GooglerpcStatus) {
	d.Set("code", m.Code)
	d.Set("details", SetProtobufAnySubResourceData(m.Details))
	d.Set("message", m.Message)
}

// Iterate through and update the GooglerpcStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetGooglerpcStatusSubResourceData(m []*models.GooglerpcStatus) (d []*map[string]interface{}) {
	for _, GooglerpcStatusModel := range m {
		if GooglerpcStatusModel != nil {
			properties := make(map[string]interface{})
			properties["code"] = GooglerpcStatusModel.Code
			properties["details"] = SetProtobufAnySubResourceData(GooglerpcStatusModel.Details)
			properties["message"] = GooglerpcStatusModel.Message
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the GooglerpcStatus resource defined in the Terraform configuration
func GooglerpcStatusSchema() map[string]*schema.Schema {
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

// Retrieve property field names for updating the GooglerpcStatus resource
func GetGooglerpcStatusPropertyFields() (t []string) {
	return []string{
		"code",
		"details",
		"message",
	}
}
