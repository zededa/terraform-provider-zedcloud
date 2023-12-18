package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func SimpleUserModel(d *schema.ResourceData) *models.SimpleUser {
	firstName, _ := d.Get("first_name").(string)
	id, _ := d.Get("id").(string)
	realmID, _ := d.Get("realm_id").(string)
	var state *models.UserState // UserState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewUserState(models.UserState(stateModel))
	}
	username, _ := d.Get("username").(string)
	return &models.SimpleUser{
		FirstName: firstName,
		ID:        id,
		RealmID:   realmID,
		State:     state,
		Username:  username,
	}
}

func SimpleUserModelFromMap(m map[string]interface{}) *models.SimpleUser {
	firstName := m["first_name"].(string)
	id := m["id"].(string)
	realmID := m["realm_id"].(string)
	var state *models.UserState // UserState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewUserState(models.UserState(stateModel))
	}
	username := m["username"].(string)
	return &models.SimpleUser{
		FirstName: firstName,
		ID:        id,
		RealmID:   realmID,
		State:     state,
		Username:  username,
	}
}

func SetSimpleUserResourceData(d *schema.ResourceData, m *models.SimpleUser) {
	d.Set("first_name", m.FirstName)
	d.Set("id", m.ID)
	d.Set("realm_id", m.RealmID)
	d.Set("state", m.State)
	d.Set("username", m.Username)
}

func SetSimpleUserSubResourceData(m []*models.SimpleUser) (d []*map[string]interface{}) {
	for _, SimpleUserModel := range m {
		if SimpleUserModel != nil {
			properties := make(map[string]interface{})
			properties["first_name"] = SimpleUserModel.FirstName
			properties["id"] = SimpleUserModel.ID
			properties["realm_id"] = SimpleUserModel.RealmID
			properties["state"] = SimpleUserModel.State
			properties["username"] = SimpleUserModel.Username
			d = append(d, &properties)
		}
	}
	return
}

func SimpleUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"first_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: ``,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"realm_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"username": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetSimpleUserPropertyFields() (t []string) {
	return []string{
		"first_name",
		"id",
		"realm_id",
		"state",
		"username",
	}
}
