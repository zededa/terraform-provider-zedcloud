package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func VolInstResourceModel(d *schema.ResourceData) *models.VolInstResource {
	curSizeBytes, _ := d.Get("cur_size_bytes").(string)
	maxSizeBytes, _ := d.Get("max_size_bytes").(string)
	return &models.VolInstResource{
		CurSizeBytes: curSizeBytes,
		MaxSizeBytes: maxSizeBytes,
	}
}

func VolInstResourceModelFromMap(m map[string]interface{}) *models.VolInstResource {
	curSizeBytes := m["cur_size_bytes"].(string)
	maxSizeBytes := m["max_size_bytes"].(string)
	return &models.VolInstResource{
		CurSizeBytes: curSizeBytes,
		MaxSizeBytes: maxSizeBytes,
	}
}

// Update the underlying VolInstResource resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVolInstResourceResourceData(d *schema.ResourceData, m *models.VolInstResource) {
	d.Set("cur_size_bytes", m.CurSizeBytes)
	d.Set("max_size_bytes", m.MaxSizeBytes)
}

// Iterate through and update the VolInstResource resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVolInstResourceSubResourceData(m []*models.VolInstResource) (d []*map[string]interface{}) {
	for _, VolInstResourceModel := range m {
		if VolInstResourceModel != nil {
			properties := make(map[string]interface{})
			properties["cur_size_bytes"] = VolInstResourceModel.CurSizeBytes
			properties["max_size_bytes"] = VolInstResourceModel.MaxSizeBytes
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VolInstResource resource defined in the Terraform configuration
func VolInstResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cur_size_bytes": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"max_size_bytes": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VolInstResource resource
func GetVolInstResourcePropertyFields() (t []string) {
	return []string{
		"cur_size_bytes",
		"max_size_bytes",
	}
}
