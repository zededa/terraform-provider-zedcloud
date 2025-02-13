package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func EnterpriseModel(d *schema.ResourceData) *models.Enterprise {
	hubspotID, _ := d.Get("hubspot_id").(string)
	sfdcID, _ := d.Get("sfdc_id").(string)
	aPITokenExpiryInSecondsInt, _ := d.Get("api_token_expiry_in_seconds").(int)
	aPITokenExpiryInSeconds := int64(aPITokenExpiryInSecondsInt)
	attributes := map[string]string{}
	attributesInterface, attributesIsSet := d.GetOk("attributes")
	if attributesIsSet {
		attributesMap := attributesInterface.(map[string]interface{})
		for k, v := range attributesMap {
			if v == nil {
				continue
			}
			attributes[k] = v.(string)
		}
	}

	azureSubID, _ := d.Get("azure_sub_id").(string)
	var childEnterprises []*models.EnterpriseSummary // []*EnterpriseSummary
	childEnterprisesInterface, childEnterprisesIsSet := d.GetOk("child_enterprises")
	if childEnterprisesIsSet {
		var items []interface{}
		if listItems, isList := childEnterprisesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = childEnterprisesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := EnterpriseSummaryModelFromMap(v.(map[string]interface{}))
			childEnterprises = append(childEnterprises, m)
		}
	}
	controllerHostURL, _ := d.Get("controller_host_url").(string)
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	inheritAuthFromParent, _ := d.Get("inherit_auth_from_parent").(bool)
	name, _ := d.Get("name").(string)
	parentEntpID, _ := d.Get("parent_entp_id").(string)
	var policyList *models.PolicyVersionList // PolicyVersionList
	policyListInterface, policyListIsSet := d.GetOk("policy_list")
	if policyListIsSet && policyListInterface != nil {
		policyListMap := policyListInterface.([]interface{})
		if len(policyListMap) > 0 {
			policyList = PolicyVersionListModelFromMap(policyListMap[0].(map[string]interface{}))
		}
	}
	var realms []string
	realmsInterface, realmsIsSet := d.GetOk("realms")
	if realmsIsSet {
		var items []interface{}
		if listItems, isList := realmsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = realmsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			realms = append(realms, v.(string))
		}
	}
	var state *models.EnterpriseState // EnterpriseState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewEnterpriseState(models.EnterpriseState(stateModel))
	}
	var streamEvents *models.DataStream // DataStream
	streamEventsInterface, streamEventsIsSet := d.GetOk("stream_events")
	if streamEventsIsSet && streamEventsInterface != nil {
		streamEventsMap := streamEventsInterface.([]interface{})
		if len(streamEventsMap) > 0 {
			streamEvents = DataStreamModelFromMap(streamEventsMap[0].(map[string]interface{}))
		}
	}
	title, _ := d.Get("title").(string)
	var totpSettings *models.TOTPSettings // TOTPSettings
	totpSettingsInterface, totpSettingsIsSet := d.GetOk("totp_settings")
	if totpSettingsIsSet && totpSettingsInterface != nil {
		totpSettingsMap := totpSettingsInterface.([]interface{})
		if len(totpSettingsMap) > 0 {
			totpSettings = TOTPSettingsModelFromMap(totpSettingsMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.EnterpriseType // EnterpriseType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewEnterpriseType(models.EnterpriseType(typeModel))
	}
	return &models.Enterprise{
		HubspotID:               hubspotID,
		SfdcID:                  sfdcID,
		APITokenExpiryInSeconds: aPITokenExpiryInSeconds,
		Attributes:              attributes,
		AzureSubID:              azureSubID,
		ChildEnterprises:        childEnterprises,
		ControllerHostURL:       controllerHostURL,
		Description:             description,
		ID:                      id,
		InheritAuthFromParent:   inheritAuthFromParent,
		Name:                    &name, // string
		ParentEntpID:            parentEntpID,
		PolicyList:              policyList,
		Realms:                  realms,
		State:                   state,
		StreamEvents:            streamEvents,
		Title:                   &title, // string
		TotpSettings:            totpSettings,
		Type:                    typeVar,
	}
}

func EnterpriseModelFromMap(m map[string]interface{}) *models.Enterprise {
	hubspotID := m["hubspot_id"].(string)
	sfdcID := m["sfdc_id"].(string)
	aPITokenExpiryInSeconds := int64(m["api_token_expiry_in_seconds"].(int)) // int64
	attributes := map[string]string{}
	attributesInterface, attributesIsSet := m["attributes"]
	if attributesIsSet {
		attributesMap := attributesInterface.(map[string]interface{})
		for k, v := range attributesMap {
			if v == nil {
				continue
			}
			attributes[k] = v.(string)
		}
	}

	azureSubID := m["azure_sub_id"].(string)
	var childEnterprises []*models.EnterpriseSummary // []*EnterpriseSummary
	childEnterprisesInterface, childEnterprisesIsSet := m["child_enterprises"]
	if childEnterprisesIsSet {
		var items []interface{}
		if listItems, isList := childEnterprisesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = childEnterprisesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := EnterpriseSummaryModelFromMap(v.(map[string]interface{}))
			childEnterprises = append(childEnterprises, m)
		}
	}
	controllerHostURL := m["controller_host_url"].(string)
	description := m["description"].(string)
	id := m["id"].(string)
	inheritAuthFromParent := m["inherit_auth_from_parent"].(bool)
	name := m["name"].(string)
	parentEntpID := m["parent_entp_id"].(string)
	var policyList *models.PolicyVersionList // PolicyVersionList
	policyListInterface, policyListIsSet := m["policy_list"]
	if policyListIsSet && policyListInterface != nil {
		policyListMap := policyListInterface.([]interface{})
		if len(policyListMap) > 0 {
			policyList = PolicyVersionListModelFromMap(policyListMap[0].(map[string]interface{}))
		}
	}
	//
	var realms []string
	realmsInterface, realmsIsSet := m["realms"]
	if realmsIsSet {
		var items []interface{}
		if listItems, isList := realmsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = realmsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			realms = append(realms, v.(string))
		}
	}
	var state *models.EnterpriseState // EnterpriseState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewEnterpriseState(models.EnterpriseState(stateModel))
	}
	var streamEvents *models.DataStream // DataStream
	streamEventsInterface, streamEventsIsSet := m["stream_events"]
	if streamEventsIsSet && streamEventsInterface != nil {
		streamEventsMap := streamEventsInterface.([]interface{})
		if len(streamEventsMap) > 0 {
			streamEvents = DataStreamModelFromMap(streamEventsMap[0].(map[string]interface{}))
		}
	}
	//
	title := m["title"].(string)
	var totpSettings *models.TOTPSettings // TOTPSettings
	totpSettingsInterface, totpSettingsIsSet := m["totp_settings"]
	if totpSettingsIsSet && totpSettingsInterface != nil {
		totpSettingsMap := totpSettingsInterface.([]interface{})
		if len(totpSettingsMap) > 0 {
			totpSettings = TOTPSettingsModelFromMap(totpSettingsMap[0].(map[string]interface{}))
		}
	}
	//
	var typeVar *models.EnterpriseType // EnterpriseType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewEnterpriseType(models.EnterpriseType(typeModel))
	}
	return &models.Enterprise{
		HubspotID:               hubspotID,
		SfdcID:                  sfdcID,
		APITokenExpiryInSeconds: aPITokenExpiryInSeconds,
		Attributes:              attributes,
		AzureSubID:              azureSubID,
		ChildEnterprises:        childEnterprises,
		ControllerHostURL:       controllerHostURL,
		Description:             description,
		ID:                      id,
		InheritAuthFromParent:   inheritAuthFromParent,
		Name:                    &name,
		ParentEntpID:            parentEntpID,
		PolicyList:              policyList,
		Realms:                  realms,
		State:                   state,
		StreamEvents:            streamEvents,
		Title:                   &title,
		TotpSettings:            totpSettings,
		Type:                    typeVar,
	}
}

func SetEnterpriseResourceData(d *schema.ResourceData, m *models.Enterprise) {
	d.Set("hubspot_id", m.HubspotID)
	d.Set("sfdc_id", m.SfdcID)
	d.Set("api_token_expiry_in_seconds", m.APITokenExpiryInSeconds)
	d.Set("attributes", m.Attributes)
	d.Set("azure_sub_id", m.AzureSubID)
	d.Set("child_enterprises", SetEnterpriseSummarySubResourceData(m.ChildEnterprises))
	d.Set("controller_host_url", m.ControllerHostURL)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("inherit_auth_from_parent", m.InheritAuthFromParent)
	d.Set("name", m.Name)
	d.Set("parent_entp_id", m.ParentEntpID)
	d.Set("policy_list", SetPolicyVersionListSubResourceData([]*models.PolicyVersionList{m.PolicyList}))
	d.Set("realms", m.Realms)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("state", m.State)
	d.Set("stream_events", SetDataStreamSubResourceData([]*models.DataStream{m.StreamEvents}))
	d.Set("title", m.Title)
	d.Set("totp_settings", SetTOTPSettingsSubResourceData([]*models.TOTPSettings{m.TotpSettings}))
	d.Set("type", m.Type)
}

func SetEnterpriseSubResourceData(m []*models.Enterprise) (d []*map[string]interface{}) {
	for _, EnterpriseModel := range m {
		if EnterpriseModel != nil {
			properties := make(map[string]interface{})
			properties["hubspot_id"] = EnterpriseModel.HubspotID
			properties["sfdc_id"] = EnterpriseModel.SfdcID
			properties["api_token_expiry_in_seconds"] = EnterpriseModel.APITokenExpiryInSeconds
			properties["attributes"] = EnterpriseModel.Attributes
			properties["azure_sub_id"] = EnterpriseModel.AzureSubID
			properties["child_enterprises"] = SetEnterpriseSummarySubResourceData(EnterpriseModel.ChildEnterprises)
			properties["controller_host_url"] = EnterpriseModel.ControllerHostURL
			properties["description"] = EnterpriseModel.Description
			properties["id"] = EnterpriseModel.ID
			properties["inherit_auth_from_parent"] = EnterpriseModel.InheritAuthFromParent
			properties["name"] = EnterpriseModel.Name
			properties["parent_entp_id"] = EnterpriseModel.ParentEntpID
			properties["policy_list"] = SetPolicyVersionListSubResourceData([]*models.PolicyVersionList{EnterpriseModel.PolicyList})
			properties["realms"] = EnterpriseModel.Realms
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{EnterpriseModel.Revision})
			properties["state"] = EnterpriseModel.State
			properties["stream_events"] = SetDataStreamSubResourceData([]*models.DataStream{EnterpriseModel.StreamEvents})
			properties["title"] = EnterpriseModel.Title
			properties["totp_settings"] = SetTOTPSettingsSubResourceData([]*models.TOTPSettings{EnterpriseModel.TotpSettings})
			properties["type"] = EnterpriseModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func EnterpriseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hubspot_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"sfdc_id": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"api_token_expiry_in_seconds": {
			Description: `Enterprise settings for API Token expiry to be set in seconds`,
			Type:        schema.TypeInt,
			Optional:    true,
			Default: 	 7776000,
		},

		"attributes": {
			Description: `enterprise level key-value pairs`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"azure_sub_id": {
			Description: `Azure subscription ID tied to this enterprise`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"child_enterprises": {
			Description: `List of all child enterprises`,
			Type:        schema.TypeList, //GoType: []*EnterpriseSummary
			Elem: &schema.Resource{
				Schema: EnterpriseSummarySchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"controller_host_url": {
			Description: `zedcontrol host`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"description": {
			Description: `Detailed description of the enterprise`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `Unique system defined enterprise ID`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"inherit_auth_from_parent": {
			Description: `Perform authorization using parent enterprise`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"name": {
			Description: `User defined name of the enterprise. Once enterprise is created, name can’t be changed`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"parent_entp_id": {
			Description: `Parent enterprise ID`,
			Type:        schema.TypeString,
			Optional:    true,
			DiffSuppressFunc: diffSupressStringNonConfigChanges("parent_entp_id"),
		},

		"policy_list": {
			Description: `Policy version list`,
			Type:        schema.TypeList, //GoType: PolicyVersionList
			Elem: &schema.Resource{
				Schema: PolicyVersionListSchema(),
			},
			Optional: true,
			DiffSuppressFunc: diffSupressMapInterfaceNonConfigChanges("policy_list", ""),
		},

		"realms": {
			Description: `List of realms associated with the enterprise`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"revision": {
			Description: `system defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"state": {
			Description: `Enterprise state`,
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "ENTERPRISE_STATE_ACTIVE",
		},

		"stream_events": {
			Description: `Enable / Disable streaming of events to 3rd party end point`,
			Type:        schema.TypeList, //GoType: DataStream
			Elem: &schema.Resource{
				Schema: DataStreamSchema(),
			},
			Optional: true,
		},

		"title": {
			Description: `User defined title for the enterprise. Title can be changed any time`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"totp_settings": {
			Description: `Enterprise settings to enforce TOTP`,
			Type:        schema.TypeList, //GoType: TOTPSettings
			Elem: &schema.Resource{
				Schema: TOTPSettingsSchema(),
			},
			Optional: true,
			DiffSuppressFunc: diffSupressMapInterfaceNonConfigChanges("totp_settings", "enforce"),
		},

		"type": {
			Description: `Enterprise type`,
			Type:        schema.TypeString,
			Optional:    true,
			Default:	 "ENTERPRISE_TYPE_UNSPECIFIED",
		},
	}
}

func GetEnterprisePropertyFields() (t []string) {
	return []string{
		"hubspot_id",
		"sfdc_id",
		"api_token_expiry_in_seconds",
		"attributes",
		"azure_sub_id",
		"child_enterprises",
		"controller_host_url",
		"description",
		"id",
		"inherit_auth_from_parent",
		"name",
		"parent_entp_id",
		"policy_list",
		"realms",
		"state",
		"stream_events",
		"title",
		"totp_settings",
		"type",
	}
}
