package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PatchReferenceUpdateConfigModel(d *schema.ResourceData) *models.PatchReferenceUpdateConfig {
	var appInstIDList []string
	appInstIDListInterface, appInstIDListIsSet := d.GetOk("appInstIDList")
	if appInstIDListIsSet {
		var items []interface{}
		if listItems, isList := appInstIDListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = appInstIDListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			appInstIDList = append(appInstIDList, v.(string))
		}
	}
	patchenvelopeID, _ := d.Get("patchenvelope_id").(string)
	projectID, _ := d.Get("project_id").(string)
	return &models.PatchReferenceUpdateConfig{
		AppInstIDList:   appInstIDList,
		PatchenvelopeID: patchenvelopeID,
		ProjectID:       projectID,
	}
}

func PatchReferenceUpdateConfigModelFromMap(m map[string]interface{}) *models.PatchReferenceUpdateConfig {
	var appInstIDList []string
	appInstIDListInterface, appInstIDListIsSet := m["appInstIDList"]
	if appInstIDListIsSet {
		var items []interface{}
		if listItems, isList := appInstIDListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = appInstIDListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			appInstIDList = append(appInstIDList, v.(string))
		}
	}
	patchenvelopeID := m["patchenvelope_id"].(string)
	projectID := m["project_id"].(string)
	return &models.PatchReferenceUpdateConfig{
		AppInstIDList:   appInstIDList,
		PatchenvelopeID: patchenvelopeID,
		ProjectID:       projectID,
	}
}

func SetPatchReferenceUpdateConfigResourceData(d *schema.ResourceData, m *models.PatchReferenceUpdateConfig) {
	d.Set("app_inst_id_list", m.AppInstIDList)
	d.Set("patchenvelope_id", m.PatchenvelopeID)
	d.Set("project_id", m.ProjectID)
}

func SetPatchReferenceUpdateConfigSubResourceData(m []*models.PatchReferenceUpdateConfig) (d []*map[string]interface{}) {
	for _, PatchReferenceUpdateConfigModel := range m {
		if PatchReferenceUpdateConfigModel != nil {
			properties := make(map[string]interface{})
			properties["app_inst_id_list"] = PatchReferenceUpdateConfigModel.AppInstIDList
			properties["patchenvelope_id"] = PatchReferenceUpdateConfigModel.PatchenvelopeID
			properties["project_id"] = PatchReferenceUpdateConfigModel.ProjectID
			d = append(d, &properties)
		}
	}
	return
}

func PatchReferenceUpdateConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_inst_id_list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"patchenvelope_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPatchReferenceUpdateConfigPropertyFields() (t []string) {
	return []string{
		"app_inst_id_list",
		"patchenvelope_id",
		"project_id",
	}
}
