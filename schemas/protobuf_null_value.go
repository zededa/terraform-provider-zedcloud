package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate ProtobufNullValue resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ProtobufNullValueModel(d *schema.ResourceData) *models.ProtobufNullValue {
	protobufNullValue, _ := d.Get("protobuf_null_value").(models.ProtobufNullValue)
	return &protobufNullValue
}

func ProtobufNullValueModelFromMap(m map[string]interface{}) *models.ProtobufNullValue {
	protobufNullValue := m["protobuf_null_value"].(models.ProtobufNullValue)
	return &protobufNullValue
}

// Update the underlying ProtobufNullValue resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetProtobufNullValueResourceData(d *schema.ResourceData, m *models.ProtobufNullValue) {
}

// Iterate through and update the ProtobufNullValue resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetProtobufNullValueSubResourceData(m []*models.ProtobufNullValue) (d []*map[string]interface{}) {
	for _, ProtobufNullValueModel := range m {
		if ProtobufNullValueModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ProtobufNullValue resource defined in the Terraform configuration
func ProtobufNullValueSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

// Retrieve property field names for updating the ProtobufNullValue resource
func GetProtobufNullValuePropertyFields() (t []string) {
	return []string{}
}
