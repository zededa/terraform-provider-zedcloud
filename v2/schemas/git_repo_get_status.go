package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoGetStatusModel(d *schema.ResourceData) *models.GitRepoGetStatus {
	var conditions []*models.GitRepoCondition // []*GitRepoCondition
	conditionsInterface, conditionsIsSet := d.GetOk("conditions")
	if conditionsIsSet {
		var items []interface{}
		if listItems, isList := conditionsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = conditionsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := GitRepoConditionModelFromMap(v.(map[string]interface{}))
			conditions = append(conditions, m)
		}
	}
	var display *models.GitRepoStatusDisplay // GitRepoStatusDisplay
	displayInterface, displayIsSet := d.GetOk("display")
	if displayIsSet && displayInterface != nil {
		displayMap := displayInterface.([]interface{})
		if len(displayMap) > 0 {
			display = GitRepoStatusDisplayModelFromMap(displayMap[0].(map[string]interface{}))
		}
	}
	var resourceCounts *models.GitRepoResourceCounts // GitRepoResourceCounts
	resourceCountsInterface, resourceCountsIsSet := d.GetOk("resource_counts")
	if resourceCountsIsSet && resourceCountsInterface != nil {
		resourceCountsMap := resourceCountsInterface.([]interface{})
		if len(resourceCountsMap) > 0 {
			resourceCounts = GitRepoResourceCountsModelFromMap(resourceCountsMap[0].(map[string]interface{}))
		}
	}
	var resourceStatusMap []*models.GitRepoResourceStatusMap // []*GitRepoResourceStatusMap
	resourceStatusMapInterface, resourceStatusMapIsSet := d.GetOk("resource_status_map")
	if resourceStatusMapIsSet {
		var items []interface{}
		if listItems, isList := resourceStatusMapInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = resourceStatusMapInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := GitRepoResourceStatusMapModelFromMap(v.(map[string]interface{}))
			resourceStatusMap = append(resourceStatusMap, m)
		}
	}
	var state *models.GitRepoState // GitRepoState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet && stateInterface != nil {
		stateMap := stateInterface.([]interface{})
		if len(stateMap) > 0 {
			state = GitRepoStateModelFromMap(stateMap[0].(map[string]interface{}))
		}
	}
	summary, _ := d.Get("summary").(any) // any
	return &models.GitRepoGetStatus{
		Conditions:        conditions,
		Display:           display,
		ResourceCounts:    resourceCounts,
		ResourceStatusMap: resourceStatusMap,
		State:             state,
		Summary:           summary,
	}
}

func GitRepoGetStatusModelFromMap(m map[string]interface{}) *models.GitRepoGetStatus {
	var conditions []*models.GitRepoCondition // []*GitRepoCondition
	conditionsInterface, conditionsIsSet := m["conditions"]
	if conditionsIsSet {
		var items []interface{}
		if listItems, isList := conditionsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = conditionsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := GitRepoConditionModelFromMap(v.(map[string]interface{}))
			conditions = append(conditions, m)
		}
	}
	var display *models.GitRepoStatusDisplay // GitRepoStatusDisplay
	displayInterface, displayIsSet := m["display"]
	if displayIsSet && displayInterface != nil {
		displayMap := displayInterface.([]interface{})
		if len(displayMap) > 0 {
			display = GitRepoStatusDisplayModelFromMap(displayMap[0].(map[string]interface{}))
		}
	}
	//
	var resourceCounts *models.GitRepoResourceCounts // GitRepoResourceCounts
	resourceCountsInterface, resourceCountsIsSet := m["resource_counts"]
	if resourceCountsIsSet && resourceCountsInterface != nil {
		resourceCountsMap := resourceCountsInterface.([]interface{})
		if len(resourceCountsMap) > 0 {
			resourceCounts = GitRepoResourceCountsModelFromMap(resourceCountsMap[0].(map[string]interface{}))
		}
	}
	//
	var resourceStatusMap []*models.GitRepoResourceStatusMap // []*GitRepoResourceStatusMap
	resourceStatusMapInterface, resourceStatusMapIsSet := m["resource_status_map"]
	if resourceStatusMapIsSet {
		var items []interface{}
		if listItems, isList := resourceStatusMapInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = resourceStatusMapInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := GitRepoResourceStatusMapModelFromMap(v.(map[string]interface{}))
			resourceStatusMap = append(resourceStatusMap, m)
		}
	}
	var state *models.GitRepoState // GitRepoState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet && stateInterface != nil {
		stateMap := stateInterface.([]interface{})
		if len(stateMap) > 0 {
			state = GitRepoStateModelFromMap(stateMap[0].(map[string]interface{}))
		}
	}
	//
	summary := m["summary"].(any)
	return &models.GitRepoGetStatus{
		Conditions:        conditions,
		Display:           display,
		ResourceCounts:    resourceCounts,
		ResourceStatusMap: resourceStatusMap,
		State:             state,
		Summary:           summary,
	}
}

func SetGitRepoGetStatusResourceData(d *schema.ResourceData, m *models.GitRepoGetStatus) {
	d.Set("conditions", SetGitRepoConditionSubResourceData(m.Conditions))
	d.Set("display", SetGitRepoStatusDisplaySubResourceData([]*models.GitRepoStatusDisplay{m.Display}))
	d.Set("resource_counts", SetGitRepoResourceCountsSubResourceData([]*models.GitRepoResourceCounts{m.ResourceCounts}))
	d.Set("resource_status_map", SetGitRepoResourceStatusMapSubResourceData(m.ResourceStatusMap))
	d.Set("state", SetGitRepoStateSubResourceData([]*models.GitRepoState{m.State}))
	d.Set("summary", m.Summary)
}

func SetGitRepoGetStatusSubResourceData(m []*models.GitRepoGetStatus) (d []*map[string]interface{}) {
	for _, GitRepoGetStatusModel := range m {
		if GitRepoGetStatusModel != nil {
			properties := make(map[string]interface{})
			properties["conditions"] = SetGitRepoConditionSubResourceData(GitRepoGetStatusModel.Conditions)
			properties["display"] = SetGitRepoStatusDisplaySubResourceData([]*models.GitRepoStatusDisplay{GitRepoGetStatusModel.Display})
			properties["resource_counts"] = SetGitRepoResourceCountsSubResourceData([]*models.GitRepoResourceCounts{GitRepoGetStatusModel.ResourceCounts})
			properties["resource_status_map"] = SetGitRepoResourceStatusMapSubResourceData(GitRepoGetStatusModel.ResourceStatusMap)
			properties["state"] = SetGitRepoStateSubResourceData([]*models.GitRepoState{GitRepoGetStatusModel.State})
			properties["summary"] = GitRepoGetStatusModel.Summary
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoGetStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"conditions": {
			Description: `List of conditions describing the current state`,
			Type:        schema.TypeList, //GoType: []*GitRepoCondition
			Elem: &schema.Resource{
				Schema: GitRepoConditionSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"display": {
			Description: `Display status information`,
			Type:        schema.TypeList, //GoType: GitRepoStatusDisplay
			Elem: &schema.Resource{
				Schema: GitRepoStatusDisplaySchema(),
			},
			Optional: true,
		},

		"resource_counts": {
			Description: `Resource count information`,
			Type:        schema.TypeList, //GoType: GitRepoResourceCounts
			Elem: &schema.Resource{
				Schema: GitRepoResourceCountsSchema(),
			},
			Optional: true,
		},

		"resource_status_map": {
			Description: `Detailed status map of individual resources managed by this GitRepo`,
			Type:        schema.TypeList, //GoType: []*GitRepoResourceStatusMap
			Elem: &schema.Resource{
				Schema: GitRepoResourceStatusMapSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"state": {
			Description: `Current state information for the GitRepo`,
			Type:        schema.TypeList, //GoType: GitRepoState
			Elem: &schema.Resource{
				Schema: GitRepoStateSchema(),
			},
			Optional: true,
		},

		"summary": {
			Description: `Summary information for the GitRepo`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGitRepoGetStatusPropertyFields() (t []string) {
	return []string{
		"conditions",
		"display",
		"resource_counts",
		"resource_status_map",
		"state",
		"summary",
	}
}
