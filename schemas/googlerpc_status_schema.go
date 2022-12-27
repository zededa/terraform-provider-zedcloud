package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate GooglerpcStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func GooglerpcStatusModel(d *schema.ResourceData) *models.GooglerpcStatus {
	code := int32(d.Get("code").(int))
	details := d.Get("details").([]*models.ProtobufAny) // []*ProtobufAny
	message := d.Get("message").(string)
	return &models.GooglerpcStatus{
		Code:    code,
		Details: details,
		Message: message,
	}
}

func GooglerpcStatusModelFromMap(m map[string]interface{}) *models.GooglerpcStatus {
	code := int32(m["code"].(int))                  // int32 false false false
	details := m["details"].([]*models.ProtobufAny) // []*ProtobufAny
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

// Iterate throught and update the GooglerpcStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
