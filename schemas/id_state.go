package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func IDStateModel(d *schema.ResourceData) *models.IDState {
	iDState, _ := d.Get("id_state").(models.IDState)
	return &iDState
}

func IDStateModelFromMap(m map[string]interface{}) *models.IDState {
	iDState := m["id_state"].(models.IDState)
	return &iDState
}

func SetIDStateResourceData(d *schema.ResourceData, m *models.IDState) {
}

func SetIDStateSubResourceData(m []*models.IDState) (d []*map[string]interface{}) {
	for _, IDStateModel := range m {
		if IDStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func IDStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetIDStatePropertyFields() (t []string) {
	return []string{}
}
