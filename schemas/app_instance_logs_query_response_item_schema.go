package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider/models"
)

func AppInstanceLogsQueryResponseItemModel(d *schema.ResourceData) *models.AppInstanceLogsQueryResponseItem {
	content, _ := d.Get("content").(string)
	msgid, _ := d.Get("msgid").(string)
	timestamp, _ := d.Get("timestamp").(strfmt.DateTime)
	return &models.AppInstanceLogsQueryResponseItem{
		Content:   content,
		Msgid:     msgid,
		Timestamp: timestamp,
	}
}

func AppInstanceLogsQueryResponseItemModelFromMap(m map[string]interface{}) *models.AppInstanceLogsQueryResponseItem {
	content := m["content"].(string)
	msgid := m["msgid"].(string)
	timestamp := m["timestamp"].(strfmt.DateTime)
	return &models.AppInstanceLogsQueryResponseItem{
		Content:   content,
		Msgid:     msgid,
		Timestamp: timestamp,
	}
}

// Update the underlying AppInstanceLogsQueryResponseItem resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppInstanceLogsQueryResponseItemResourceData(d *schema.ResourceData, m *models.AppInstanceLogsQueryResponseItem) {
	d.Set("content", m.Content)
	d.Set("msgid", m.Msgid)
	d.Set("timestamp", m.Timestamp)
}

// Iterate through and update the AppInstanceLogsQueryResponseItem resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppInstanceLogsQueryResponseItemSubResourceData(m []*models.AppInstanceLogsQueryResponseItem) (d []*map[string]interface{}) {
	for _, AppInstanceLogsQueryResponseItemModel := range m {
		if AppInstanceLogsQueryResponseItemModel != nil {
			properties := make(map[string]interface{})
			properties["content"] = AppInstanceLogsQueryResponseItemModel.Content
			properties["msgid"] = AppInstanceLogsQueryResponseItemModel.Msgid
			properties["timestamp"] = AppInstanceLogsQueryResponseItemModel.Timestamp.String()
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the AppInstanceLogsQueryResponseItem resource defined in the Terraform configuration
func AppInstanceLogsQueryResponseItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"content": {
			Description: `app instance logs`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"msgid": {
			Description: `message Id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"timestamp": {
			Description:  `timestamp of query time`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},
	}
}

// Retrieve property field names for updating the AppInstanceLogsQueryResponseItem resource
func GetAppInstanceLogsQueryResponseItemPropertyFields() (t []string) {
	return []string{
		"content",
		"msgid",
		"timestamp",
	}
}
