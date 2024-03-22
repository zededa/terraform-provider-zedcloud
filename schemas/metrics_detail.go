package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func MetricsDetailModel(d *schema.ResourceData) *models.MetricsDetail {
	queries := map[string]string{}
	queriesInterface, queriesIsSet := d.GetOk("queries")
	if queriesIsSet {
		queriesMap := queriesInterface.(map[string]interface{})
		for k, v := range queriesMap {
			if v == nil {
				continue
			}
			queries[k] = v.(string)
		}
	}

	results := map[string]string{}
	resultsInterface, resultsIsSet := d.GetOk("results")
	if resultsIsSet {
		resultsMap := resultsInterface.(map[string]interface{})
		for k, v := range resultsMap {
			if v == nil {
				continue
			}
			results[k] = v.(string)
		}
	}

	return &models.MetricsDetail{
		Queries: queries,
		Results: results,
	}
}

func MetricsDetailModelFromMap(m map[string]interface{}) *models.MetricsDetail {
	queries := map[string]string{}
	queriesInterface, queriesIsSet := m["queries"]
	if queriesIsSet {
		queriesMap := queriesInterface.(map[string]interface{})
		for k, v := range queriesMap {
			if v == nil {
				continue
			}
			queries[k] = v.(string)
		}
	}

	results := map[string]string{}
	resultsInterface, resultsIsSet := m["results"]
	if resultsIsSet {
		resultsMap := resultsInterface.(map[string]interface{})
		for k, v := range resultsMap {
			if v == nil {
				continue
			}
			results[k] = v.(string)
		}
	}

	return &models.MetricsDetail{
		Queries: queries,
		Results: results,
	}
}

func SetMetricsDetailResourceData(d *schema.ResourceData, m *models.MetricsDetail) {
	d.Set("queries", m.Queries)
	d.Set("results", m.Results)
}

func SetMetricsDetailSubResourceData(m []*models.MetricsDetail) (d []*map[string]interface{}) {
	for _, MetricsDetailModel := range m {
		if MetricsDetailModel != nil {
			properties := make(map[string]interface{})
			properties["queries"] = MetricsDetailModel.Queries
			properties["results"] = MetricsDetailModel.Results
			d = append(d, &properties)
		}
	}
	return
}

func MetricsDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"queries": {
			Description: `Mapping of queries variable keys and value`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"results": {
			Description: `Mapping of results variable keys and value`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetMetricsDetailPropertyFields() (t []string) {
	return []string{
		"queries",
		"results",
	}
}
