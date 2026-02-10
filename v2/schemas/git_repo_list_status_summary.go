package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoListStatusSummaryModel(d *schema.ResourceData) *models.GitRepoListStatusSummary {
	description, _ := d.Get("description").(string)
	totalInt, _ := d.Get("total").(int)
	total := int32(totalInt)
	values := map[string]int32{}
	valuesInterface, valuesIsSet := d.GetOk("values")
	if valuesIsSet {
		valuesMap := valuesInterface.(map[string]interface{})
		for k, v := range valuesMap {
			if v == nil {
				continue
			}
			values[k] = int32(v.(int))
		}
	}
	return &models.GitRepoListStatusSummary{
		Description: description,
		Total:       total,
		Values:      values,
	}
}

func GitRepoListStatusSummaryModelFromMap(m map[string]interface{}) *models.GitRepoListStatusSummary {
	description := m["description"].(string)
	total := int32(m["total"].(int)) // int32
	values := map[string]int32{}
	valuesInterface, valuesIsSet := m["values"]
	if valuesIsSet {
		valuesMap := valuesInterface.(map[string]interface{})
		for k, v := range valuesMap {
			if v == nil {
				continue
			}
			values[k] = int32(v.(int))
		}
	}
	return &models.GitRepoListStatusSummary{
		Description: description,
		Total:       total,
		Values:      values,
	}
}

func SetGitRepoListStatusSummaryResourceData(d *schema.ResourceData, m *models.GitRepoListStatusSummary) {
	d.Set("description", m.Description)
	d.Set("total", m.Total)
	d.Set("values", m.Values)
}

func SetGitRepoListStatusSummarySubResourceData(m []*models.GitRepoListStatusSummary) (d []*map[string]interface{}) {
	for _, GitRepoListStatusSummaryModel := range m {
		if GitRepoListStatusSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["description"] = GitRepoListStatusSummaryModel.Description
			properties["total"] = GitRepoListStatusSummaryModel.Total
			properties["values"] = GitRepoListStatusSummaryModel.Values
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoListStatusSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: `Description of the status summary`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"total": {
			Description: `Total number of GitRepos`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"values": {
			Description: `Map of status values and their counts`,
			Type:        schema.TypeMap, //GoType: map[string]int32
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetGitRepoListStatusSummaryPropertyFields() (t []string) {
	return []string{
		"description",
		"total",
		"values",
	}
}
