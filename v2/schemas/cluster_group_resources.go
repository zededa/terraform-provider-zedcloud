package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterGroupResourcesModel(d *schema.ResourceData) *models.ClusterGroupResources {
	amountInt, _ := d.Get("amount").(int)
	amount := int64(amountInt)
	var summary *models.ClusterGroupResourceSummary // ClusterGroupResourceSummary
	summaryInterface, summaryIsSet := d.GetOk("summary")
	if summaryIsSet && summaryInterface != nil {
		summaryMap := summaryInterface.([]interface{})
		if len(summaryMap) > 0 {
			summary = ClusterGroupResourceSummaryModelFromMap(summaryMap[0].(map[string]interface{}))
		}
	}
	return &models.ClusterGroupResources{
		Amount:  amount,
		Summary: summary,
	}
}

func ClusterGroupResourcesModelFromMap(m map[string]interface{}) *models.ClusterGroupResources {
	amount := int64(m["amount"].(int))              // int64
	var summary *models.ClusterGroupResourceSummary // ClusterGroupResourceSummary
	summaryInterface, summaryIsSet := m["summary"]
	if summaryIsSet && summaryInterface != nil {
		summaryMap := summaryInterface.([]interface{})
		if len(summaryMap) > 0 {
			summary = ClusterGroupResourceSummaryModelFromMap(summaryMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ClusterGroupResources{
		Amount:  amount,
		Summary: summary,
	}
}

func SetClusterGroupResourcesResourceData(d *schema.ResourceData, m *models.ClusterGroupResources) {
	d.Set("amount", m.Amount)
	d.Set("summary", SetClusterGroupResourceSummarySubResourceData([]*models.ClusterGroupResourceSummary{m.Summary}))
}

func SetClusterGroupResourcesSubResourceData(m []*models.ClusterGroupResources) (d []*map[string]interface{}) {
	for _, ClusterGroupResourcesModel := range m {
		if ClusterGroupResourcesModel != nil {
			properties := make(map[string]interface{})
			properties["amount"] = ClusterGroupResourcesModel.Amount
			properties["summary"] = SetClusterGroupResourceSummarySubResourceData([]*models.ClusterGroupResourceSummary{ClusterGroupResourcesModel.Summary})
			d = append(d, &properties)
		}
	}
	return
}

func ClusterGroupResourcesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"amount": {
			Description: `Total resources of clusters in the cluster group`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"summary": {
			Description: `Summary of the resources in the cluster group`,
			Type:        schema.TypeList, //GoType: ClusterGroupResourceSummary
			Elem: &schema.Resource{
				Schema: ClusterGroupResourceSummarySchema(),
			},
			Optional: true,
		},
	}
}

func GetClusterGroupResourcesPropertyFields() (t []string) {
	return []string{
		"amount",
		"summary",
	}
}
