package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate LispConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func LispConfigModel(d *schema.ResourceData) *models.LispConfig {
	allocate := d.Get("allocate").(bool)
	allocationprefix := d.Get("allocationprefix").(strfmt.Base64)
	allocationprefixlen := int64(d.Get("allocationprefixlen").(int))
	exportprivate := d.Get("exportprivate").(bool)
	lispiid := int64(d.Get("lispiid").(int))
	sp := d.Get("sp").([]*models.ServicePoint) // []*ServicePoint
	return &models.LispConfig{
		Allocate:            allocate,
		Allocationprefix:    allocationprefix,
		Allocationprefixlen: allocationprefixlen,
		Exportprivate:       exportprivate,
		Lispiid:             lispiid,
		Sp:                  sp,
	}
}

func LispConfigModelFromMap(m map[string]interface{}) *models.LispConfig {
	allocate := m["allocate"].(bool)
	allocationprefix := m["allocationprefix"].(strfmt.Base64)
	allocationprefixlen := int64(m["allocationprefixlen"].(int)) // int64 false false false
	exportprivate := m["exportprivate"].(bool)
	lispiid := int64(m["lispiid"].(int))   // int64 false false false
	sp := m["sp"].([]*models.ServicePoint) // []*ServicePoint
	return &models.LispConfig{
		Allocate:            allocate,
		Allocationprefix:    allocationprefix,
		Allocationprefixlen: allocationprefixlen,
		Exportprivate:       exportprivate,
		Lispiid:             lispiid,
		Sp:                  sp,
	}
}

// Update the underlying LispConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetLispConfigResourceData(d *schema.ResourceData, m *models.LispConfig) {
	d.Set("allocate", m.Allocate)
	d.Set("allocationprefix", m.Allocationprefix)
	d.Set("allocationprefixlen", m.Allocationprefixlen)
	d.Set("exportprivate", m.Exportprivate)
	d.Set("lispiid", m.Lispiid)
	d.Set("sp", SetServicePointSubResourceData(m.Sp))
}

// Iterate throught and update the LispConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the LispConfig resource defined in the Terraform configuration
func LispConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocate": {
			Description: `Allocate flag`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"allocationprefix": {
			Description:  `Allocation Prefix`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
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
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},
	}
}

// Retrieve property field names for updating the LispConfig resource
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
