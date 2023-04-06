package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func AttestationPolicyModel(d *schema.ResourceData) *models.AttestationPolicy {
	id, _ := d.Get("id").(string)
	var typeVar *models.AttestPolicyType // AttestPolicyType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAttestPolicyType(models.AttestPolicyType(typeModel))
	}
	return &models.AttestationPolicy{
		ID:   id,
		Type: typeVar,
	}
}

func AttestationPolicyModelFromMap(m map[string]interface{}) *models.AttestationPolicy {
	id := m["id"].(string)
	var typeVar *models.AttestPolicyType // AttestPolicyType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewAttestPolicyType(models.AttestPolicyType(typeModel))
	}
	return &models.AttestationPolicy{
		ID:   id,
		Type: typeVar,
	}
}

func SetAttestationPolicyResourceData(d *schema.ResourceData, m *models.AttestationPolicy) {
	d.Set("id", m.ID)
	d.Set("type", m.Type)
}

func SetAttestationPolicySubResourceData(m []*models.AttestationPolicy) (d []*map[string]interface{}) {
	for _, AttestationPolicyModel := range m {
		if AttestationPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = AttestationPolicyModel.ID
			properties["type"] = AttestationPolicyModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func AttestationPolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `unique policy id`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"type": {
			Description: `Attestation policy type`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetAttestationPolicyPropertyFields() (t []string) {
	return []string{
		"id",
		"type",
	}
}
