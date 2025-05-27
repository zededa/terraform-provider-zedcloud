package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZserviceDeploymentReadROModel(d *schema.ResourceData) *models.ZserviceDeploymentReadRO {
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
	latestRevisionInt, _ := d.Get("latest_revision").(int)
	latestRevision := int64(latestRevisionInt)
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
	return &models.ZserviceDeploymentReadRO{
		Chart:               chart,
		DeploymentID:        deploymentID,
		DeploymentName:      deploymentName,
		LatestRevision:      latestRevision,
		OverrideAssetValues: overrideAssetValues,
		TargetAsset:         targetAsset,
	}
}

func ZserviceDeploymentReadROModelFromMap(m map[string]interface{}) *models.ZserviceDeploymentReadRO {
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
	latestRevision := int64(m["latest_revision"].(int)) // int64
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
	return &models.ZserviceDeploymentReadRO{
		Chart:               chart,
		DeploymentID:        deploymentID,
		DeploymentName:      deploymentName,
		LatestRevision:      latestRevision,
		OverrideAssetValues: overrideAssetValues,
		TargetAsset:         targetAsset,
	}
}

func SetZserviceDeploymentReadROResourceData(d *schema.ResourceData, m *models.ZserviceDeploymentReadRO) {
	d.Set("chart", SetZserviceChartSubResourceData([]*models.ZserviceChart{m.Chart}))
	d.Set("deployment_id", m.DeploymentID)
	d.Set("deployment_name", m.DeploymentName)
	d.Set("latest_revision", m.LatestRevision)
	d.Set("override_asset_values", SetZserviceAssetValuesSubResourceData([]*models.ZserviceAssetValues{m.OverrideAssetValues}))
	d.Set("target_asset", SetZserviceAssetSubResourceData([]*models.ZserviceAsset{m.TargetAsset}))
}

func SetZserviceDeploymentReadROSubResourceData(m []*models.ZserviceDeploymentReadRO) (d []*map[string]interface{}) {
	for _, ZserviceDeploymentReadROModel := range m {
		if ZserviceDeploymentReadROModel != nil {
			properties := make(map[string]interface{})
			properties["chart"] = SetZserviceChartSubResourceData([]*models.ZserviceChart{ZserviceDeploymentReadROModel.Chart})
			properties["deployment_id"] = ZserviceDeploymentReadROModel.DeploymentID
			properties["deployment_name"] = ZserviceDeploymentReadROModel.DeploymentName
			properties["latest_revision"] = ZserviceDeploymentReadROModel.LatestRevision
			properties["override_asset_values"] = SetZserviceAssetValuesSubResourceData([]*models.ZserviceAssetValues{ZserviceDeploymentReadROModel.OverrideAssetValues})
			properties["target_asset"] = SetZserviceAssetSubResourceData([]*models.ZserviceAsset{ZserviceDeploymentReadROModel.TargetAsset})
			d = append(d, &properties)
		}
	}
	return
}

func ZserviceDeploymentReadROSchema() map[string]*schema.Schema {
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

		"latest_revision": {
			Description: ``,
			Type:        schema.TypeInt,
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

func GetZserviceDeploymentReadROPropertyFields() (t []string) {
	return []string{
		"chart",
		"deployment_id",
		"deployment_name",
		"latest_revision",
		"override_asset_values",
		"target_asset",
	}
}
