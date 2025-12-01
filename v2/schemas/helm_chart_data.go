package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func HelmChartDataModel(d *schema.ResourceData) *models.HelmChartData {
	var chart *models.Chart // Chart
	chartInterface, chartIsSet := d.GetOk("chart")
	if chartIsSet && chartInterface != nil {
		chartMap := chartInterface.([]interface{})
		if len(chartMap) > 0 {
			chart = ChartModelFromMap(chartMap[0].(map[string]interface{}))
		}
	}
	var targetClusters *models.TargetClusters // TargetClusters
	targetClustersInterface, targetClustersIsSet := d.GetOk("target_clusters")
	if targetClustersIsSet && targetClustersInterface != nil {
		targetClustersMap := targetClustersInterface.([]interface{})
		if len(targetClustersMap) > 0 {
			targetClusters = TargetClustersModelFromMap(targetClustersMap[0].(map[string]interface{}))
		}
	}
	return &models.HelmChartData{
		Chart:          chart,
		TargetClusters: targetClusters,
	}
}

func HelmChartDataModelFromMap(m map[string]interface{}) *models.HelmChartData {
	var chart *models.Chart // Chart
	chartInterface, chartIsSet := m["chart"]
	if chartIsSet && chartInterface != nil {
		chartMap := chartInterface.([]interface{})
		if len(chartMap) > 0 {
			chart = ChartModelFromMap(chartMap[0].(map[string]interface{}))
		}
	}
	//
	var targetClusters *models.TargetClusters // TargetClusters
	targetClustersInterface, targetClustersIsSet := m["target_clusters"]
	if targetClustersIsSet && targetClustersInterface != nil {
		targetClustersMap := targetClustersInterface.([]interface{})
		if len(targetClustersMap) > 0 {
			targetClusters = TargetClustersModelFromMap(targetClustersMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.HelmChartData{
		Chart:          chart,
		TargetClusters: targetClusters,
	}
}

func SetHelmChartDataResourceData(d *schema.ResourceData, m *models.HelmChartData) {
	d.Set("chart", SetChartSubResourceData([]*models.Chart{m.Chart}))
	d.Set("target_clusters", SetTargetClustersSubResourceData([]*models.TargetClusters{m.TargetClusters}))
}

func SetHelmChartDataSubResourceData(m []*models.HelmChartData) (d []*map[string]interface{}) {
	for _, HelmChartDataModel := range m {
		if HelmChartDataModel != nil {
			properties := make(map[string]interface{})
			properties["chart"] = SetChartSubResourceData([]*models.Chart{HelmChartDataModel.Chart})
			properties["target_clusters"] = SetTargetClustersSubResourceData([]*models.TargetClusters{HelmChartDataModel.TargetClusters})
			d = append(d, &properties)
		}
	}
	return
}

func HelmChartDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"chart": {
			Description: `Helm chart details for the deployment`,
			Type:        schema.TypeList, //GoType: Chart
			Elem: &schema.Resource{
				Schema: ChartSchema(),
			},
			Required: true,
		},

		"target_clusters": {
			Description: `Target clusters for the deployment`,
			Type:        schema.TypeList, //GoType: TargetClusters
			Elem: &schema.Resource{
				Schema: TargetClustersSchema(),
			},
			Required: true,
		},
	}
}

func GetHelmChartDataPropertyFields() (t []string) {
	return []string{
		"chart",
		"target_clusters",
	}
}
