package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PolicyCommonModel(d *schema.ResourceData) *models.PolicyCommon {
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	policyTargetCondition := map[string]string{}
	policyTargetConditionInterface, policyTargetConditionIsSet := d.GetOk("policyTargetCondition")
	if policyTargetConditionIsSet {
		policyTargetConditionMap := policyTargetConditionInterface.(map[string]interface{})
		for k, v := range policyTargetConditionMap {
			if v == nil {
				continue
			}
			policyTargetCondition[k] = v.(string)
		}
	}

	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	title, _ := d.Get("title").(string)
	return &models.PolicyCommon{
		ID:                    id,
		Name:                  name,
		PolicyTargetCondition: policyTargetCondition,
		Revision:              revision,
		Tags:                  tags,
		Title:                 title,
	}
}

func PolicyCommonModelFromMap(m map[string]interface{}) *models.PolicyCommon {
	id := m["id"].(string)
	name := m["name"].(string)
	policyTargetCondition := map[string]string{}
	policyTargetConditionInterface, policyTargetConditionIsSet := m["policy_target_condition"]
	if policyTargetConditionIsSet {
		policyTargetConditionMap := policyTargetConditionInterface.(map[string]interface{})
		for k, v := range policyTargetConditionMap {
			if v == nil {
				continue
			}
			policyTargetCondition[k] = v.(string)
		}
	}

	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	//
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	title := m["title"].(string)
	return &models.PolicyCommon{
		ID:                    id,
		Name:                  name,
		PolicyTargetCondition: policyTargetCondition,
		Revision:              revision,
		Tags:                  tags,
		Title:                 title,
	}
}

func SetPolicyCommonResourceData(d *schema.ResourceData, m *models.PolicyCommon) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("policy_target_condition", m.PolicyTargetCondition)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("tags", m.Tags)
	d.Set("title", m.Title)
}

func SetPolicyCommonSubResourceData(m []*models.PolicyCommon) (d []*map[string]interface{}) {
	for _, PolicyCommonModel := range m {
		if PolicyCommonModel != nil {
			properties := make(map[string]interface{})
			properties["id"] = PolicyCommonModel.ID
			properties["name"] = PolicyCommonModel.Name
			properties["policy_target_condition"] = PolicyCommonModel.PolicyTargetCondition
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{PolicyCommonModel.Revision})
			properties["tags"] = PolicyCommonModel.Tags
			properties["title"] = PolicyCommonModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func PolicyCommon() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Description: `system generated unique id for an policy`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `user defined policy name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"policy_target_condition": {
			Description: `user defined key-value pairs of a policy that will be used for targeting by devices`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"revision": {
			Description: `object revision details`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Optional: true,
		},

		"tags": {
			Description: `user defined key-value pairs of a policy`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"title": {
			Description: `user defined policy title`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPolicyCommonPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
		"policy_target_condition",
		"revision",
		"tags",
		"title",
	}
}
