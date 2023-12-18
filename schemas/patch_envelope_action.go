package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PatchEnvelopeActionModel(d *schema.ResourceData) *models.PatchEnvelopeAction {
	patchEnvelopeAction, _ := d.Get("patch_envelope_action").(models.PatchEnvelopeAction)
	return &patchEnvelopeAction
}

func PatchEnvelopeActionModelFromMap(m map[string]interface{}) *models.PatchEnvelopeAction {
	patchEnvelopeAction := m["patch_envelope_action"].(models.PatchEnvelopeAction)
	return &patchEnvelopeAction
}

func SetPatchEnvelopeActionResourceData(d *schema.ResourceData, m *models.PatchEnvelopeAction) {
}

func SetPatchEnvelopeActionSubResourceData(m []*models.PatchEnvelopeAction) (d []*map[string]interface{}) {
	for _, PatchEnvelopeActionModel := range m {
		if PatchEnvelopeActionModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func PatchEnvelopeActionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetPatchEnvelopeActionPropertyFields() (t []string) {
	return []string{}
}
