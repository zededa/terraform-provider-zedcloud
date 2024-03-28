package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func LispConfigModel(d *schema.ResourceData) *models.LispConfig {
	allocate, _ := d.Get("allocate").(bool)
	allocationprefix, _ := d.Get("allocationprefix").(strfmt.Base64)
	allocationprefixlenInt, _ := d.Get("allocationprefixlen").(int)
	allocationprefixlen := int64(allocationprefixlenInt)
	exportprivate, _ := d.Get("exportprivate").(bool)
	lispiidInt, _ := d.Get("lispiid").(int)
	lispiid := int64(lispiidInt)
	var sp []*models.ServicePoint // []*ServicePoint
	spInterface, spIsSet := d.GetOk("sp")
	if spIsSet {
		var items []interface{}
		if listItems, isList := spInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = spInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ServicePointModelFromMap(v.(map[string]interface{}))
			sp = append(sp, m)
		}
	}
	return &models.LispConfig{
		Allocate:            allocate,
		Allocationprefix:    string(allocationprefix),
		Allocationprefixlen: allocationprefixlen,
		Exportprivate:       exportprivate,
		Lispiid:             lispiid,
		Sp:                  sp,
	}
}

func LispConfigModelFromMap(m map[string]interface{}) *models.LispConfig {
	allocate := m["allocate"].(bool)
	allocationprefix := m["allocationprefix"].(strfmt.Base64)
	allocationprefixlen := int64(m["allocationprefixlen"].(int)) // int64
	exportprivate := m["exportprivate"].(bool)
	lispiid := int64(m["lispiid"].(int)) // int64
	var sp []*models.ServicePoint        // []*ServicePoint
	spInterface, spIsSet := m["sp"]
	if spIsSet {
		var items []interface{}
		if listItems, isList := spInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = spInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ServicePointModelFromMap(v.(map[string]interface{}))
			sp = append(sp, m)
		}
	}
	return &models.LispConfig{
		Allocate:            allocate,
		Allocationprefix:    string(allocationprefix),
		Allocationprefixlen: allocationprefixlen,
		Exportprivate:       exportprivate,
		Lispiid:             lispiid,
		Sp:                  sp,
	}
}

func SetLispConfigResourceData(d *schema.ResourceData, m *models.LispConfig) {
	d.Set("allocate", m.Allocate)
	d.Set("allocationprefix", m.Allocationprefix)
	d.Set("allocationprefixlen", m.Allocationprefixlen)
	d.Set("exportprivate", m.Exportprivate)
	d.Set("lispiid", m.Lispiid)
	d.Set("sp", SetServicePointSubResourceData(m.Sp))
}

func SetLispConfigSubResourceData(m []*models.LispConfig) (d []*map[string]interface{}) {
	for _, LispConfigModel := range m {
		if LispConfigModel != nil {
			properties := make(map[string]interface{})
			properties["allocate"] = LispConfigModel.Allocate
			properties["allocationprefix"] = LispConfigModel.Allocationprefix
			properties["allocationprefixlen"] = LispConfigModel.Allocationprefixlen
			properties["exportprivate"] = LispConfigModel.Exportprivate
			properties["lispiid"] = LispConfigModel.Lispiid
			properties["sp"] = SetServicePointSubResourceData(LispConfigModel.Sp)
			d = append(d, &properties)
		}
	}
	return
}

func LispConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocate": {
			Description: `Allocate flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"allocationprefix": {
			Description: `Allocation Prefix`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"allocationprefixlen": {
			Description: `Allocation Prefix Length`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"exportprivate": {
			Description: `Export Private flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"lispiid": {
			Description: `lisp id`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"sp": {
			Description: `Service Point List`,
			Type:        schema.TypeList, //GoType: []*ServicePoint
			Elem: &schema.Resource{
				Schema: ServicePoint(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetLispConfigPropertyFields() (t []string) {
	return []string{
		"allocate",
		"allocationprefix",
		"allocationprefixlen",
		"exportprivate",
		"lispiid",
		"sp",
	}
}
