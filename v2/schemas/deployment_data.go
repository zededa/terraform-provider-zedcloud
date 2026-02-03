package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DeploymentDataModel(d *schema.ResourceData) *models.DeploymentData {
	var helmChartData *models.HelmChartData // HelmChartData
	helmChartDataInterface, helmChartDataIsSet := d.GetOk("helm_chart_data")
	if helmChartDataIsSet && helmChartDataInterface != nil {
		helmChartDataMap := helmChartDataInterface.([]interface{})
		if len(helmChartDataMap) > 0 {
			helmChartData = HelmChartDataModelFromMap(helmChartDataMap[0].(map[string]interface{}))
		}
	}
	return &models.DeploymentData{
		HelmChartData: helmChartData,
	}
}

func DeploymentDataModelFromMap(m map[string]interface{}) *models.DeploymentData {
	var helmChartData *models.HelmChartData // HelmChartData
	helmChartDataInterface, helmChartDataIsSet := m["helm_chart_data"]
	if helmChartDataIsSet && helmChartDataInterface != nil {
		helmChartDataMap := helmChartDataInterface.([]interface{})
		if len(helmChartDataMap) > 0 {
			helmChartData = HelmChartDataModelFromMap(helmChartDataMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.DeploymentData{
		HelmChartData: helmChartData,
	}
}

func SetDeploymentDataResourceData(d *schema.ResourceData, m *models.DeploymentData) {
	d.Set("helm_chart_data", SetHelmChartDataSubResourceData([]*models.HelmChartData{m.HelmChartData}))
}

func SetDeploymentDataSubResourceData(m []*models.DeploymentData) (d []*map[string]interface{}) {
	for _, DeploymentDataModel := range m {
		if DeploymentDataModel != nil {
			properties := make(map[string]interface{})
			properties["helm_chart_data"] = SetHelmChartDataSubResourceData([]*models.HelmChartData{DeploymentDataModel.HelmChartData})
			d = append(d, &properties)
		}
	}
	return
}

func DeploymentDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"helm_chart_data": {
			Description: `Helm chart data for the deployment`,
			Type:        schema.TypeList, //GoType: HelmChartData
			Elem: &schema.Resource{
				Schema: HelmChartDataSchema(),
			},
			Required: true,
		},
	}
}

func GetDeploymentDataPropertyFields() (t []string) {
	return []string{
		"helm_chart_data",
	}
}
