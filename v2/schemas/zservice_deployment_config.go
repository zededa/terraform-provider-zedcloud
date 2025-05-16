package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceDeploymentConfigModel(d *schema.ResourceData) *models.ZserviceDeploymentConfig {
	var chart *models.ZserviceChart // ZserviceChart
	chartInterface, chartIsSet := d.GetOk("chart")
	if chartIsSet && chartInterface != nil {
		chartMap := chartInterface.([]interface{})
		if len(chartMap) > 0 {
			chart = ZserviceChartModelFromMap(chartMap[0].(map[string]interface{}))
		}
	}
	deploymentID, _ := d.Get("deployment_id").(string)
	deploymentName, _ := d.Get("deployment_name").(string)
	var overrideAssetValues *models.ZserviceAssetValues // ZserviceAssetValues
	overrideAssetValuesInterface, overrideAssetValuesIsSet := d.GetOk("override_asset_values")
	if overrideAssetValuesIsSet && overrideAssetValuesInterface != nil {
		overrideAssetValuesMap := overrideAssetValuesInterface.([]interface{})
		if len(overrideAssetValuesMap) > 0 {
			overrideAssetValues = ZserviceAssetValuesModelFromMap(overrideAssetValuesMap[0].(map[string]interface{}))
		}
	}
	var targetAsset *models.ZserviceAsset // ZserviceAsset
	targetAssetInterface, targetAssetIsSet := d.GetOk("target_asset")
	if targetAssetIsSet && targetAssetInterface != nil {
		targetAssetMap := targetAssetInterface.([]interface{})
		if len(targetAssetMap) > 0 {
			targetAsset = ZserviceAssetModelFromMap(targetAssetMap[0].(map[string]interface{}))
		}
	}
	return &models.ZserviceDeploymentConfig{
		Chart:               chart,
		DeploymentID:        deploymentID,
		DeploymentName:      deploymentName,
		OverrideAssetValues: overrideAssetValues,
		TargetAsset:         targetAsset,
	}
}

func ZserviceDeploymentConfigModelFromMap(m map[string]interface{}) *models.ZserviceDeploymentConfig {
	var chart *models.ZserviceChart // ZserviceChart
	chartInterface, chartIsSet := m["chart"]
	if chartIsSet && chartInterface != nil {
		chartMap := chartInterface.([]interface{})
		if len(chartMap) > 0 {
			chart = ZserviceChartModelFromMap(chartMap[0].(map[string]interface{}))
		}
	}
	//
	deploymentID := m["deployment_id"].(string)
	deploymentName := m["deployment_name"].(string)
	var overrideAssetValues *models.ZserviceAssetValues // ZserviceAssetValues
	overrideAssetValuesInterface, overrideAssetValuesIsSet := m["override_asset_values"]
	if overrideAssetValuesIsSet && overrideAssetValuesInterface != nil {
		overrideAssetValuesMap := overrideAssetValuesInterface.([]interface{})
		if len(overrideAssetValuesMap) > 0 {
			overrideAssetValues = ZserviceAssetValuesModelFromMap(overrideAssetValuesMap[0].(map[string]interface{}))
		}
	}
	//
	var targetAsset *models.ZserviceAsset // ZserviceAsset
	targetAssetInterface, targetAssetIsSet := m["target_asset"]
	if targetAssetIsSet && targetAssetInterface != nil {
		targetAssetMap := targetAssetInterface.([]interface{})
		if len(targetAssetMap) > 0 {
			targetAsset = ZserviceAssetModelFromMap(targetAssetMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ZserviceDeploymentConfig{
		Chart:               chart,
		DeploymentID:        deploymentID,
		DeploymentName:      deploymentName,
		OverrideAssetValues: overrideAssetValues,
		TargetAsset:         targetAsset,
	}
}

func SetZserviceDeploymentConfigResourceData(d *schema.ResourceData, m *models.ZserviceDeploymentConfig) {
	d.Set("chart", SetZserviceChartSubResourceData([]*models.ZserviceChart{m.Chart}))
	d.Set("deployment_id", m.DeploymentID)
	d.Set("deployment_name", m.DeploymentName)
	d.Set("override_asset_values", SetZserviceAssetValuesSubResourceData([]*models.ZserviceAssetValues{m.OverrideAssetValues}))
	d.Set("target_asset", SetZserviceAssetSubResourceData([]*models.ZserviceAsset{m.TargetAsset}))
}

func SetZserviceDeploymentConfigSubResourceData(m []*models.ZserviceDeploymentConfig) (d []*map[string]interface{}) {
	for _, ZserviceDeploymentConfigModel := range m {
		if ZserviceDeploymentConfigModel != nil {
			properties := make(map[string]interface{})
			properties["chart"] = SetZserviceChartSubResourceData([]*models.ZserviceChart{ZserviceDeploymentConfigModel.Chart})
			properties["deployment_id"] = ZserviceDeploymentConfigModel.DeploymentID
			properties["deployment_name"] = ZserviceDeploymentConfigModel.DeploymentName
			properties["override_asset_values"] = SetZserviceAssetValuesSubResourceData([]*models.ZserviceAssetValues{ZserviceDeploymentConfigModel.OverrideAssetValues})
			properties["target_asset"] = SetZserviceAssetSubResourceData([]*models.ZserviceAsset{ZserviceDeploymentConfigModel.TargetAsset})
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceDeploymentConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"chart": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZserviceChart
			Elem: &schema.Resource{
				Schema: ZserviceChartSchema(),
			},
			Optional: true,
		},

		"deployment_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"deployment_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"override_asset_values": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZserviceAssetValues
			Elem: &schema.Resource{
				Schema: ZserviceAssetValuesSchema(),
			},
			Optional: true,
		},

		"target_asset": {
			Description: ``,
			Type:        schema.TypeList, //GoType: ZserviceAsset
			Elem: &schema.Resource{
				Schema: ZserviceAssetSchema(),
			},
			Optional: true,
		},
	}
}

func GetZserviceDeploymentConfigPropertyFields() (t []string) {
	return []string{
		"chart",
		"deployment_id",
		"deployment_name",
		"override_asset_values",
		"target_asset",
	}
}
