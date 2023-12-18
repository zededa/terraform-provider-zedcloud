package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func UserStateModel(d *schema.ResourceData) *models.UserState {
	userState, _ := d.Get("user_state").(models.UserState)
	return &userState
}

func UserStateModelFromMap(m map[string]interface{}) *models.UserState {
	userState := m["user_state"].(models.UserState)
	return &userState
}

func SetUserStateResourceData(d *schema.ResourceData, m *models.UserState) {
}

func SetUserStateSubResourceData(m []*models.UserState) (d []*map[string]interface{}) {
	for _, UserStateModel := range m {
		if UserStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func UserStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetUserStatePropertyFields() (t []string) {
	return []string{}
}
