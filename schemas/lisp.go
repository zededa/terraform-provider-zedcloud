package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func LispModel(d *schema.ResourceData) *models.LispConfig {
	allocate, _ := d.Get("allocate").(bool)
	allocationprefix, _ := d.Get("allocationprefix").(string)
	allocationprefixlenInt, _ := d.Get("allocationprefixlen").(int)
	allocationprefixlen := int64(allocationprefixlenInt)
	exportprivate, _ := d.Get("exportprivate").(bool)
	lispiidInt, _ := d.Get("lispiid").(int)
	lispiid := int64(lispiidInt)
	var sp []*models.ServicePoint // []*ServicePoint
	spInterface, spIsSet := d.GetOk("sp")
	if spIsSet {
		for _, v := range spInterface.([]interface{}) {
			if v == nil {
				continue
			}
			m := ServicePointModelFromMap(v.(map[string]interface{}))
			sp = append(sp, m)
		}
	}
	return &models.LispConfig{
		Allocate:            allocate,
		Allocationprefix:    allocationprefix,
		Allocationprefixlen: allocationprefixlen,
		Exportprivate:       exportprivate,
		Lispiid:             lispiid,
		Sp:                  sp,
	}
}

func LispModelFromMap(m map[string]interface{}) *models.LispConfig {
	allocate := m["allocate"].(bool)
	allocationprefix := m["allocationprefix"].(string)
	allocationprefixlen := int64(m["allocationprefixlen"].(int)) // int64
	exportprivate := m["exportprivate"].(bool)
	lispiid := int64(m["lispiid"].(int)) // int64
	var sp []*models.ServicePoint        // []*ServicePoint
	spInterface, spIsSet := m["sp"]
	if spIsSet {
		for _, v := range spInterface.([]interface{}) {
			if v == nil {
				continue
			}
			m := ServicePointModelFromMap(v.(map[string]interface{}))
			sp = append(sp, m)
		}
	}
	return &models.LispConfig{
		Allocate:            allocate,
		Allocationprefix:    allocationprefix,
		Allocationprefixlen: allocationprefixlen,
		Exportprivate:       exportprivate,
		Lispiid:             lispiid,
		Sp:                  sp,
	}
}

func SetLispResourceData(d *schema.ResourceData, m *models.LispConfig) {
	d.Set("allocate", m.Allocate)
	d.Set("allocationprefix", m.Allocationprefix)
	d.Set("allocationprefixlen", m.Allocationprefixlen)
	d.Set("exportprivate", m.Exportprivate)
	d.Set("lispiid", m.Lispiid)
	d.Set("sp", SetServicePointSubResourceData(m.Sp))
}

func SetLispSubResourceData(m []*models.LispConfig) (d []*map[string]interface{}) {
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

func Lisp() map[string]*schema.Schema {
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
				Schema: ServicePointSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},
	}
}

func GetLispPropertyFields() (t []string) {
	return []string{
		"allocate",
		"allocationprefix",
		"allocationprefixlen",
		"exportprivate",
		"lispiid",
		"sp",
	}
}
