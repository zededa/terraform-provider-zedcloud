package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate ModulePolicy resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ModulePolicyModel(d *schema.ResourceData) *models.ModulePolicy {
	etag := d.Get("etag").(string)
	apps := d.Get("apps").([]*models.AppConfig) // []*AppConfig
	id := d.Get("id").(string)
	labels := d.Get("labels").(map[string]string) // map[string]string
	var metrics *models.MetricsDetail             // MetricsDetail
	metricsInterface, metricsIsSet := d.GetOk("metrics")
	if metricsIsSet {
		metricsMap := metricsInterface.([]interface{})[0].(map[string]interface{})
		metrics = MetricsDetailModelFromMap(metricsMap)
	}
	priority := int64(d.Get("priority").(int))
	routes := d.Get("routes").(map[string]string) // map[string]string
	targetCondition := d.Get("target_condition").(string)
	targetConditionNew := d.Get("target_condition_new").(map[string]string) // map[string]string
	return &models.ModulePolicy{
		Etag:               etag,
		Apps:               apps,
		ID:                 id,
		Labels:             labels,
		Metrics:            metrics,
		Priority:           priority,
		Routes:             routes,
		TargetCondition:    targetCondition,
		TargetConditionNew: targetConditionNew,
	}
}

func ModulePolicyModelFromMap(m map[string]interface{}) *models.ModulePolicy {
	etag := m["etag"].(string)
	apps := m["apps"].([]*models.AppConfig) // []*AppConfig
	id := m["id"].(string)
	labels := m["labels"].(map[string]string)
	var metrics *models.MetricsDetail // MetricsDetail
	metricsInterface, metricsIsSet := m["metrics"]
	if metricsIsSet {
		metricsMap := metricsInterface.([]interface{})[0].(map[string]interface{})
		metrics = MetricsDetailModelFromMap(metricsMap)
	}
	//
	priority := int64(m["priority"].(int)) // int64 false false false
	routes := m["routes"].(map[string]string)
	targetCondition := m["target_condition"].(string)
	targetConditionNew := m["target_condition_new"].(map[string]string)
	return &models.ModulePolicy{
		Etag:               etag,
		Apps:               apps,
		ID:                 id,
		Labels:             labels,
		Metrics:            metrics,
		Priority:           priority,
		Routes:             routes,
		TargetCondition:    targetCondition,
		TargetConditionNew: targetConditionNew,
	}
}

// Update the underlying ModulePolicy resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetModulePolicyResourceData(d *schema.ResourceData, m *models.ModulePolicy) {
	d.Set("etag", m.Etag)
	d.Set("apps", SetAppConfigSubResourceData(m.Apps))
	d.Set("id", m.ID)
	d.Set("labels", m.Labels)
	d.Set("metrics", SetMetricsDetailSubResourceData([]*models.MetricsDetail{m.Metrics}))
	d.Set("priority", m.Priority)
	d.Set("routes", m.Routes)
	d.Set("target_condition", m.TargetCondition)
	d.Set("target_condition_new", m.TargetConditionNew)
}

// Iterate throught and update the ModulePolicy resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetModulePolicySubResourceData(m []*models.ModulePolicy) (d []*map[string]interface{}) {
	for _, ModulePolicyModel := range m {
		if ModulePolicyModel != nil {
			properties := make(map[string]interface{})
			properties["etag"] = ModulePolicyModel.Etag
			properties["apps"] = SetAppConfigSubResourceData(ModulePolicyModel.Apps)
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

// Schema mapping representing the ModulePolicy resource defined in the Terraform configuration
func ModulePolicySchema() map[string]*schema.Schema {
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
				Schema: AppConfigSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
		},

		"id": {
			Description: `unique id for deployment`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"labels": {
			Description: ``,
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
				Schema: MetricsDetailSchema(),
			},
			Optional: true,
		},

		"priority": {
			Description: `deployment priority of module manifest`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"routes": {
			Description: ``,
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
			Description: ``,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

// Retrieve property field names for updating the ModulePolicy resource
func GetModulePolicyPropertyFields() (t []string) {
	return []string{
		"etag",
		"apps",
		"id",
		"labels",
		"metrics",
		"priority",
		"routes",
		"target_condition",
		"target_condition_new",
	}
}
