package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func VfsModel(d *schema.ResourceData) *models.Vfs {
	countInt, _ := d.Get("count").(int)
	count := int64(countInt)
	return &models.Vfs{
		Count: count,
	}
}

func VfsModelFromMap(m map[string]interface{}) *models.Vfs {
	count := int64(m["count"].(int)) // int64
	return &models.Vfs{
		Count: count,
	}
}

func SetVfsResourceData(d *schema.ResourceData, m *models.Vfs) {
	d.Set("count", m.Count)
}

func SetVfsSubResourceData(m []*models.Vfs) (d []*map[string]interface{}) {
	for _, VfsModel := range m {
		if VfsModel != nil {
			properties := make(map[string]interface{})
			properties["count"] = VfsModel.Count
			d = append(d, &properties)
		}
	}
	return
}

func VfsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"count": {
			Description: `The number of Virtual Functions for given Physical Function. Only applies for IO_TYPE_ETH_PF.`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetVfsPropertyFields() (t []string) {
	return []string{
		"count",
	}
}
