package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func InlineOpaqueBase64DataModel(d *schema.ResourceData) *models.InlineOpaqueBase64Data {
	base64Data, _ := d.Get("base64_data").(string)
	base64MetaData, _ := d.Get("base64_meta_data").(string)
	fileNameToUse, _ := d.Get("file_name_to_use").(string)
	return &models.InlineOpaqueBase64Data{
		Base64Data:     base64Data,
		Base64MetaData: base64MetaData,
		FileNameToUse:  fileNameToUse,
	}
}

func InlineOpaqueBase64DataModelFromMap(m map[string]interface{}) *models.InlineOpaqueBase64Data {
	base64Data := m["base64_data"].(string)
	base64MetaData := m["base64_meta_data"].(string)
	fileNameToUse := m["file_name_to_use"].(string)
	return &models.InlineOpaqueBase64Data{
		Base64Data:     base64Data,
		Base64MetaData: base64MetaData,
		FileNameToUse:  fileNameToUse,
	}
}

func SetInlineOpaqueBase64DataResourceData(d *schema.ResourceData, m *models.InlineOpaqueBase64Data) {
	d.Set("base64_data", m.Base64Data)
	d.Set("base64_meta_data", m.Base64MetaData)
	d.Set("file_name_to_use", m.FileNameToUse)
}

func SetInlineOpaqueBase64DataSubResourceData(m []*models.InlineOpaqueBase64Data) (d []*map[string]interface{}) {
	for _, InlineOpaqueBase64DataModel := range m {
		if InlineOpaqueBase64DataModel != nil {
			properties := make(map[string]interface{})
			properties["base64_data"] = InlineOpaqueBase64DataModel.Base64Data
			properties["base64_meta_data"] = InlineOpaqueBase64DataModel.Base64MetaData
			properties["file_name_to_use"] = InlineOpaqueBase64DataModel.FileNameToUse
			d = append(d, &properties)
		}
	}
	return
}

func InlineOpaqueBase64DataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"base64_data": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"base64_meta_data": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"file_name_to_use": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetInlineOpaqueBase64DataPropertyFields() (t []string) {
	return []string{
		"base64_data",
		"base64_meta_data",
		"file_name_to_use",
	}
}
