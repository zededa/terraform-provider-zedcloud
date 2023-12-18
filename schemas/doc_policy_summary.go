package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DocPolicySummaryModel(d *schema.ResourceData) *models.DocPolicySummary {
	fileURL, _ := d.Get("file_url").(string)
	policy, _ := d.Get("policy").(string)
	version, _ := d.Get("version").(string)
	return &models.DocPolicySummary{
		FileURL: fileURL,
		Policy:  policy,
		Version: version,
	}
}

func DocPolicySummaryModelFromMap(m map[string]interface{}) *models.DocPolicySummary {
	fileURL := m["file_url"].(string)
	policy := m["policy"].(string)
	version := m["version"].(string)
	return &models.DocPolicySummary{
		FileURL: fileURL,
		Policy:  policy,
		Version: version,
	}
}

func SetDocPolicySummaryResourceData(d *schema.ResourceData, m *models.DocPolicySummary) {
	d.Set("file_url", m.FileURL)
	d.Set("policy", m.Policy)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("version", m.Version)
}

func SetDocPolicySummarySubResourceData(m []*models.DocPolicySummary) (d []*map[string]interface{}) {
	for _, DocPolicySummaryModel := range m {
		if DocPolicySummaryModel != nil {
			properties := make(map[string]interface{})
			properties["file_url"] = DocPolicySummaryModel.FileURL
			properties["policy"] = DocPolicySummaryModel.Policy
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{DocPolicySummaryModel.Revision})
			properties["version"] = DocPolicySummaryModel.Version
			d = append(d, &properties)
		}
	}
	return
}

func DocPolicySummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"file_url": {
			Description: `Policy doc fileURL`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"policy": {
			Description: `User defined name of the docpolicy. Name cannot be changed once created`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"revision": {
			Description: `System defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"version": {
			Description: `Policy doc version`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetDocPolicySummaryPropertyFields() (t []string) {
	return []string{
		"file_url",
		"policy",
		"version",
	}
}
