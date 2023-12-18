package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PatchEnvelopeStateModel(d *schema.ResourceData) *models.PatchEnvelopeState {
	patchEnvelopeState, _ := d.Get("patch_envelope_state").(models.PatchEnvelopeState)
	return &patchEnvelopeState
}

func PatchEnvelopeStateModelFromMap(m map[string]interface{}) *models.PatchEnvelopeState {
	patchEnvelopeState := m["patch_envelope_state"].(models.PatchEnvelopeState)
	return &patchEnvelopeState
}

func SetPatchEnvelopeStateResourceData(d *schema.ResourceData, m *models.PatchEnvelopeState) {
}

func SetPatchEnvelopeStateSubResourceData(m []*models.PatchEnvelopeState) (d []*map[string]interface{}) {
	for _, PatchEnvelopeStateModel := range m {
		if PatchEnvelopeStateModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func PatchEnvelopeStateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetPatchEnvelopeStatePropertyFields() (t []string) {
	return []string{}
}
