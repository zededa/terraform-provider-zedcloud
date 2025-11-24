package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ZksInstanceStatusSummaryModel(d *schema.ResourceData) *models.ZksInstanceStatusSummary {
	errorInt, _ := d.Get("error").(int)
	error := int32(errorInt)
	initInt, _ := d.Get("init").(int)
	init := int32(initInt)
	offlineInt, _ := d.Get("offline").(int)
	offline := int32(offlineInt)
	onlineInt, _ := d.Get("online").(int)
	online := int32(onlineInt)
	pendingInt, _ := d.Get("pending").(int)
	pending := int32(pendingInt)
	suspectInt, _ := d.Get("suspect").(int)
	suspect := int32(suspectInt)
	unknownInt, _ := d.Get("unknown").(int)
	unknown := int32(unknownInt)
	unspecifiedInt, _ := d.Get("unspecified").(int)
	unspecified := int32(unspecifiedInt)
	return &models.ZksInstanceStatusSummary{
		Error:       error,
		Init:        init,
		Offline:     offline,
		Online:      online,
		Pending:     pending,
		Suspect:     suspect,
		Unknown:     unknown,
		Unspecified: unspecified,
	}
}

func ZksInstanceStatusSummaryModelFromMap(m map[string]interface{}) *models.ZksInstanceStatusSummary {
	error := int32(m["error"].(int))             // int32
	init := int32(m["init"].(int))               // int32
	offline := int32(m["offline"].(int))         // int32
	online := int32(m["online"].(int))           // int32
	pending := int32(m["pending"].(int))         // int32
	suspect := int32(m["suspect"].(int))         // int32
	unknown := int32(m["unknown"].(int))         // int32
	unspecified := int32(m["unspecified"].(int)) // int32
	return &models.ZksInstanceStatusSummary{
		Error:       error,
		Init:        init,
		Offline:     offline,
		Online:      online,
		Pending:     pending,
		Suspect:     suspect,
		Unknown:     unknown,
		Unspecified: unspecified,
	}
}

func SetZksInstanceStatusSummaryResourceData(d *schema.ResourceData, m *models.ZksInstanceStatusSummary) {
	d.Set("error", m.Error)
	d.Set("init", m.Init)
	d.Set("offline", m.Offline)
	d.Set("online", m.Online)
	d.Set("pending", m.Pending)
	d.Set("suspect", m.Suspect)
	d.Set("unknown", m.Unknown)
	d.Set("unspecified", m.Unspecified)
}

func SetZksInstanceStatusSummarySubResourceData(m []*models.ZksInstanceStatusSummary) (d []*map[string]interface{}) {
	for _, ZksInstanceStatusSummaryModel := range m {
		if ZksInstanceStatusSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["error"] = ZksInstanceStatusSummaryModel.Error
			properties["init"] = ZksInstanceStatusSummaryModel.Init
			properties["offline"] = ZksInstanceStatusSummaryModel.Offline
			properties["online"] = ZksInstanceStatusSummaryModel.Online
			properties["pending"] = ZksInstanceStatusSummaryModel.Pending
			properties["suspect"] = ZksInstanceStatusSummaryModel.Suspect
			properties["unknown"] = ZksInstanceStatusSummaryModel.Unknown
			properties["unspecified"] = ZksInstanceStatusSummaryModel.Unspecified
			d = append(d, &properties)
		}
	}
	return
}

func ZksInstanceStatusSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"error": {
			Description: `Number of error nodes in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"init": {
			Description: `Number of nodes in init state in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"offline": {
			Description: `Number of offline nodes in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"online": {
			Description: `Number of online nodes in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"pending": {
			Description: `Number of pending nodes in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"suspect": {
			Description: `Number of suspect nodes in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"unknown": {
			Description: `Number of nodes in unknown state in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"unspecified": {
			Description: `Number of nodes in unspecified state in the zks instance`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetZksInstanceStatusSummaryPropertyFields() (t []string) {
	return []string{
		"error",
		"init",
		"offline",
		"online",
		"pending",
		"suspect",
		"unknown",
		"unspecified",
	}
}
