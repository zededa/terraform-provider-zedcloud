package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ModulePolicyModel(d *schema.ResourceData) *models.ModulePolicy {
	etag, _ := d.Get("etag").(string)
	var apps []*models.AppConfig // []*AppConfig
	appsInterface, appsIsSet := d.GetOk("apps")
	if appsIsSet {
		var items []interface{}
		if listItems, isList := appsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = appsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppConfigModelFromMap(v.(map[string]interface{}))
			apps = append(apps, m)
		}
	}
	var azureEdgeAgent *models.AppConfig // AppConfig
	azureEdgeAgentInterface, azureEdgeAgentIsSet := d.GetOk("azure_edge_agent")
	if azureEdgeAgentIsSet && azureEdgeAgentInterface != nil {
		azureEdgeAgentMap := azureEdgeAgentInterface.([]interface{})
		if len(azureEdgeAgentMap) > 0 {
			azureEdgeAgent = AppConfigModelFromMap(azureEdgeAgentMap[0].(map[string]interface{}))
		}
	}
	var azureEdgeHub *models.AppConfig // AppConfig
	azureEdgeHubInterface, azureEdgeHubIsSet := d.GetOk("azure_edge_hub")
	if azureEdgeHubIsSet && azureEdgeHubInterface != nil {
		azureEdgeHubMap := azureEdgeHubInterface.([]interface{})
		if len(azureEdgeHubMap) > 0 {
			azureEdgeHub = AppConfigModelFromMap(azureEdgeHubMap[0].(map[string]interface{}))
		}
	}
	id, _ := d.Get("id").(string)
	labels := map[string]string{}
	labelsInterface, labelsIsSet := d.GetOk("labels")
	if labelsIsSet {
		labelsMap := labelsInterface.(map[string]interface{})
		for k, v := range labelsMap {
			if v == nil {
				continue
			}
			labels[k] = v.(string)
		}
	}

	var metrics *models.MetricsDetail // MetricsDetail
	metricsInterface, metricsIsSet := d.GetOk("metrics")
	if metricsIsSet && metricsInterface != nil {
		metricsMap := metricsInterface.([]interface{})
		if len(metricsMap) > 0 {
			metrics = MetricsDetailModelFromMap(metricsMap[0].(map[string]interface{}))
		}
	}
	priorityInt, _ := d.Get("priority").(int)
	priority := int64(priorityInt)
	routes := map[string]string{}
	routesInterface, routesIsSet := d.GetOk("routes")
	if routesIsSet {
		routesMap := routesInterface.(map[string]interface{})
		for k, v := range routesMap {
			if v == nil {
				continue
			}
			routes[k] = v.(string)
		}
	}

	targetCondition, _ := d.Get("target_condition").(string)
	targetConditionNew := map[string]string{}
	targetConditionNewInterface, targetConditionNewIsSet := d.GetOk("targetConditionNew")
	if targetConditionNewIsSet {
		targetConditionNewMap := targetConditionNewInterface.(map[string]interface{})
		for k, v := range targetConditionNewMap {
			if v == nil {
				continue
			}
			targetConditionNew[k] = v.(string)
		}
	}

	return &models.ModulePolicy{
		Etag:               etag,
		Apps:               apps,
		AzureEdgeAgent:     azureEdgeAgent,
		AzureEdgeHub:       azureEdgeHub,
		ID:                 id,
		Labels:             labels,
		Metrics:            metrics,
		Priority:           &priority, // int64 true false false
		Routes:             routes,
		TargetCondition:    targetCondition,
		TargetConditionNew: targetConditionNew,
	}
}

func ModulePolicyModelFromMap(m map[string]interface{}) *models.ModulePolicy {
	etag := m["etag"].(string)
	var apps []*models.AppConfig // []*AppConfig
	appsInterface, appsIsSet := m["apps"]
	if appsIsSet {
		var items []interface{}
		if listItems, isList := appsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = appsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := AppConfigModelFromMap(v.(map[string]interface{}))
			apps = append(apps, m)
		}
	}
	var azureEdgeAgent *models.AppConfig // AppConfig
	azureEdgeAgentInterface, azureEdgeAgentIsSet := m["azure_edge_agent"]
	if azureEdgeAgentIsSet && azureEdgeAgentInterface != nil {
		azureEdgeAgentMap := azureEdgeAgentInterface.([]interface{})
		if len(azureEdgeAgentMap) > 0 {
			azureEdgeAgent = AppConfigModelFromMap(azureEdgeAgentMap[0].(map[string]interface{}))
		}
	}
	//
	var azureEdgeHub *models.AppConfig // AppConfig
	azureEdgeHubInterface, azureEdgeHubIsSet := m["azure_edge_hub"]
	if azureEdgeHubIsSet && azureEdgeHubInterface != nil {
		azureEdgeHubMap := azureEdgeHubInterface.([]interface{})
		if len(azureEdgeHubMap) > 0 {
			azureEdgeHub = AppConfigModelFromMap(azureEdgeHubMap[0].(map[string]interface{}))
		}
	}
	//
	id := m["id"].(string)
	labels := map[string]string{}
	labelsInterface, labelsIsSet := m["labels"]
	if labelsIsSet {
		labelsMap := labelsInterface.(map[string]interface{})
		for k, v := range labelsMap {
			if v == nil {
				continue
			}
			labels[k] = v.(string)
		}
	}

	var metrics *models.MetricsDetail // MetricsDetail
	metricsInterface, metricsIsSet := m["metrics"]
	if metricsIsSet && metricsInterface != nil {
		metricsMap := metricsInterface.([]interface{})
		if len(metricsMap) > 0 {
			metrics = MetricsDetailModelFromMap(metricsMap[0].(map[string]interface{}))
		}
	}
	//
	priority := int64(m["priority"].(int)) // int64
	routes := map[string]string{}
	routesInterface, routesIsSet := m["routes"]
	if routesIsSet {
		routesMap := routesInterface.(map[string]interface{})
		for k, v := range routesMap {
			if v == nil {
				continue
			}
			routes[k] = v.(string)
		}
	}

	targetCondition := m["target_condition"].(string)
	targetConditionNew := map[string]string{}
	targetConditionNewInterface, targetConditionNewIsSet := m["target_condition_new"]
	if targetConditionNewIsSet {
		targetConditionNewMap := targetConditionNewInterface.(map[string]interface{})
		for k, v := range targetConditionNewMap {
			if v == nil {
				continue
			}
			targetConditionNew[k] = v.(string)
		}
	}

	return &models.ModulePolicy{
		Etag:               etag,
		Apps:               apps,
		AzureEdgeAgent:     azureEdgeAgent,
		AzureEdgeHub:       azureEdgeHub,
		ID:                 id,
		Labels:             labels,
		Metrics:            metrics,
		Priority:           &priority,
		Routes:             routes,
		TargetCondition:    targetCondition,
		TargetConditionNew: targetConditionNew,
	}
}

func SetModulePolicyResourceData(d *schema.ResourceData, m *models.ModulePolicy) {
	d.Set("etag", m.Etag)
	d.Set("apps", SetAppConfigSubResourceData(m.Apps))
	d.Set("azure_edge_agent", SetAppConfigSubResourceData([]*models.AppConfig{m.AzureEdgeAgent}))
	d.Set("azure_edge_hub", SetAppConfigSubResourceData([]*models.AppConfig{m.AzureEdgeHub}))
	d.Set("id", m.ID)
	d.Set("labels", m.Labels)
	d.Set("metrics", SetMetricsDetailSubResourceData([]*models.MetricsDetail{m.Metrics}))
	d.Set("priority", m.Priority)
	d.Set("routes", m.Routes)
	d.Set("target_condition", m.TargetCondition)
	d.Set("target_condition_new", m.TargetConditionNew)
}

func SetModulePolicySubResourceData(m []*models.ModulePolicy) (d []*map[string]interface{}) {
	for _, ModulePolicyModel := range m {
		if ModulePolicyModel != nil {
			properties := make(map[string]interface{})
			properties["etag"] = ModulePolicyModel.Etag
			properties["apps"] = SetAppConfigSubResourceData(ModulePolicyModel.Apps)
			properties["azure_edge_agent"] = SetAppConfigSubResourceData([]*models.AppConfig{ModulePolicyModel.AzureEdgeAgent})
			properties["azure_edge_hub"] = SetAppConfigSubResourceData([]*models.AppConfig{ModulePolicyModel.AzureEdgeHub})
			properties["id"] = ModulePolicyModel.ID
			properties["labels"] = ModulePolicyModel.Labels
			properties["metrics"] = SetMetricsDetailSubResourceData([]*models.MetricsDetail{ModulePolicyModel.Metrics})
			properties["priority"] = ModulePolicyModel.Priority
			properties["routes"] = ModulePolicyModel.Routes
			properties["target_condition"] = ModulePolicyModel.TargetCondition
			properties["target_condition_new"] = ModulePolicyModel.TargetConditionNew
			d = append(d, &properties)
		}
	}
	return
}

func ModulePolicy() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"etag": {
			Description: `etag for deployment`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"apps": {
			Description: `list of app details that will be provisioned on all the devices of the project to which this policy is attached`,
			Type:        schema.TypeList, //GoType: []*AppConfig
			Elem: &schema.Resource{
				Schema: AppConfig(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"azure_edge_agent": {
			Description: `app that describes the azure edge agent to be deployed on the Azure runtime`,
			Type:        schema.TypeList, //GoType: AppConfig
			Elem: &schema.Resource{
				Schema: AppConfig(),
			},
			Optional: true,
		},

		"azure_edge_hub": {
			Description: `app that describes the azure edge hub to be deployed on the Azure runtime`,
			Type:        schema.TypeList, //GoType: AppConfig
			Elem: &schema.Resource{
				Schema: AppConfig(),
			},
			Optional: true,
		},

		"id": {
			Description: `unique id for deployment`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"labels": {
			Description: `Mapping of label variable keys and value`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"metrics": {
			Description: `custom metrics for deployment`,
			Type:        schema.TypeList, //GoType: MetricsDetail
			Elem: &schema.Resource{
				Schema: MetricsDetail(),
			},
			Optional: true,
		},

		"priority": {
			Description: `deployment priority of module manifest`,
			Type:        schema.TypeInt,
			Required:    true,
		},

		"routes": {
			Description: `Mapping of routes variable keys and value`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"target_condition": {
			Description: `target condition for deployment that matches single device or group of devices`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"target_condition_new": {
			Description: `target condition for deployment that matches single device or group of devices`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func GetModulePolicyPropertyFields() (t []string) {
	return []string{
		"etag",
		"apps",
		"azure_edge_agent",
		"azure_edge_hub",
		"id",
		"labels",
		"metrics",
		"priority",
		"routes",
		"target_condition",
		"target_condition_new",
	}
}
