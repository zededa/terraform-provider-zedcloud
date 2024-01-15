package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func CredentialsModel(d *schema.ResourceData) *models.Credentials {
	var list []*models.Credential // []*Credential
	listInterface, listIsSet := d.GetOk("list")
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := CredentialModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := d.GetOk("summary_by_state")
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	return &models.Credentials{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func CredentialsModelFromMap(m map[string]interface{}) *models.Credentials {
	var list []*models.Credential // []*Credential
	listInterface, listIsSet := m["list"]
	if listIsSet {
		var items []interface{}
		if listItems, isList := listInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = listInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := CredentialModelFromMap(v.(map[string]interface{}))
			list = append(list, m)
		}
	}
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	var summaryByState *models.Summary // Summary
	summaryByStateInterface, summaryByStateIsSet := m["summary_by_state"]
	if summaryByStateIsSet && summaryByStateInterface != nil {
		summaryByStateMap := summaryByStateInterface.([]interface{})
		if len(summaryByStateMap) > 0 {
			summaryByState = SummaryModelFromMap(summaryByStateMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.Credentials{
		List:           list,
		Next:           next,
		SummaryByState: summaryByState,
	}
}

func SetCredentialsResourceData(d *schema.ResourceData, m *models.Credentials) {
	d.Set("list", SetCredentialSubResourceData(m.List))
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("summary_by_state", SetSummarySubResourceData([]*models.Summary{m.SummaryByState}))
}

func SetCredentialsSubResourceData(m []*models.Credentials) (d []*map[string]interface{}) {
	for _, CredentialsModel := range m {
		if CredentialsModel != nil {
			properties := make(map[string]interface{})
			properties["list"] = SetCredentialSubResourceData(CredentialsModel.List)
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{CredentialsModel.Next})
			properties["summary_by_state"] = SetSummarySubResourceData([]*models.Summary{CredentialsModel.SummaryByState})
			d = append(d, &properties)
		}
	}
	return
}

func CredentialsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"list": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*Credential
			Elem: &schema.Resource{
				Schema: CredentialSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"next": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"summary_by_state": {
			Description: ``,
			Type:        schema.TypeList, //GoType: Summary
			Elem: &schema.Resource{
				Schema: SummarySchema(),
			},
			Optional: true,
		},
	}
}

func GetCredentialsPropertyFields() (t []string) {
	return []string{
		"list",
		"next",
		"summary_by_state",
	}
}
