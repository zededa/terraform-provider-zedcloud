package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ManifestInfoModel(d *schema.ResourceData) *models.ManifestInfo {
	bundleVersion, _ := d.Get("bundle_version").(string)
	var details *models.TransDetails // TransDetails
	detailsInterface, detailsIsSet := d.GetOk("details")
	if detailsIsSet && detailsInterface != nil {
		detailsMap := detailsInterface.([]interface{})
		if len(detailsMap) > 0 {
			details = TransDetailsModelFromMap(detailsMap[0].(map[string]interface{}))
		}
	}
	nextBundleVersion, _ := d.Get("next_bundle_version").(string)
	params := map[string]string{}
	paramsInterface, paramsIsSet := d.GetOk("params")
	if paramsIsSet {
		paramsMap := paramsInterface.(map[string]interface{})
		for k, v := range paramsMap {
			if v == nil {
				continue
			}
			params[k] = v.(string)
		}
	}

	var transitionAction *models.InstanceTransitionAction // InstanceTransitionAction
	transitionActionInterface, transitionActionIsSet := d.GetOk("transition_action")
	if transitionActionIsSet {
		transitionActionModel := transitionActionInterface.(string)
		transitionAction = models.NewInstanceTransitionAction(models.InstanceTransitionAction(transitionActionModel))
	}
	return &models.ManifestInfo{
		BundleVersion:     bundleVersion,
		Details:           details,
		NextBundleVersion: nextBundleVersion,
		Params:            params,
		TransitionAction:  transitionAction,
	}
}

func ManifestInfoModelFromMap(m map[string]interface{}) *models.ManifestInfo {
	bundleVersion := m["bundle_version"].(string)
	var details *models.TransDetails // TransDetails
	detailsInterface, detailsIsSet := m["details"]
	if detailsIsSet && detailsInterface != nil {
		detailsMap := detailsInterface.([]interface{})
		if len(detailsMap) > 0 {
			details = TransDetailsModelFromMap(detailsMap[0].(map[string]interface{}))
		}
	}
	//
	nextBundleVersion := m["next_bundle_version"].(string)
	params := map[string]string{}
	paramsInterface, paramsIsSet := m["params"]
	if paramsIsSet {
		paramsMap := paramsInterface.(map[string]interface{})
		for k, v := range paramsMap {
			if v == nil {
				continue
			}
			params[k] = v.(string)
		}
	}

	var transitionAction *models.InstanceTransitionAction // InstanceTransitionAction
	transitionActionInterface, transitionActionIsSet := m["transition_action"]
	if transitionActionIsSet {
		transitionActionModel := transitionActionInterface.(string)
		transitionAction = models.NewInstanceTransitionAction(models.InstanceTransitionAction(transitionActionModel))
	}
	return &models.ManifestInfo{
		BundleVersion:     bundleVersion,
		Details:           details,
		NextBundleVersion: nextBundleVersion,
		Params:            params,
		TransitionAction:  transitionAction,
	}
}

// Update the underlying ManifestInfo resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetManifestInfoResourceData(d *schema.ResourceData, m *models.ManifestInfo) {
	d.Set("bundle_version", m.BundleVersion)
	d.Set("details", SetTransDetailsSubResourceData([]*models.TransDetails{m.Details}))
	d.Set("next_bundle_version", m.NextBundleVersion)
	d.Set("params", m.Params)
	d.Set("transition_action", m.TransitionAction)
}

// Iterate through and update the ManifestInfo resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetManifestInfoSubResourceData(m []*models.ManifestInfo) (d []*map[string]interface{}) {
	for _, ManifestInfoModel := range m {
		if ManifestInfoModel != nil {
			properties := make(map[string]interface{})
			properties["bundle_version"] = ManifestInfoModel.BundleVersion
			properties["details"] = SetTransDetailsSubResourceData([]*models.TransDetails{ManifestInfoModel.Details})
			properties["next_bundle_version"] = ManifestInfoModel.NextBundleVersion
			properties["params"] = ManifestInfoModel.Params
			properties["transition_action"] = ManifestInfoModel.TransitionAction
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ManifestInfo resource defined in the Terraform configuration
func ManifestInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bundle_version": {
			Description: `Current version of edge application being used`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"details": {
			Description: `Details for recommended transition action`,
			Type:        schema.TypeList, //GoType: TransDetails
			Elem: &schema.Resource{
				Schema: TransDetailsSchema(),
			},
			Optional: true,
		},

		"next_bundle_version": {
			Description: `Next version of edge application available`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"params": {
			Description: `Common and Custom parameters for deciding on transition actions`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"transition_action": {
			Description: `Recommended transition action`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ManifestInfo resource
func GetManifestInfoPropertyFields() (t []string) {
	return []string{
		"bundle_version",
		"details",
		"next_bundle_version",
		"params",
		"transition_action",
	}
}
