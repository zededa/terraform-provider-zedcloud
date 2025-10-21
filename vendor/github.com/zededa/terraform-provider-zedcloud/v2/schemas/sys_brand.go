package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SysBrandModel(d *schema.ResourceData) *models.SysBrand {
	attr := map[string]string{}
	attrInterface, attrIsSet := d.GetOk("attr")
	if attrIsSet {
		attrMap := attrInterface.(map[string]interface{})
		for k, v := range attrMap {
			if v == nil {
				continue
			}
			attr[k] = v.(string)
		}
	}

	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	logo := map[string]string{}
	logoInterface, logoIsSet := d.GetOk("logo")
	if logoIsSet {
		logoMap := logoInterface.(map[string]interface{})
		for k, v := range logoMap {
			if v == nil {
				continue
			}
			logo[k] = v.(string)
		}
	}

	name, _ := d.Get("name").(string)
	var originType *models.Origin // Origin
	originTypeInterface, originTypeIsSet := d.GetOk("origin_type")
	if originTypeIsSet {
		originTypeModel := originTypeInterface.(string)
		originType = models.NewOrigin(models.Origin(originTypeModel))
	}
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	var state *models.SysModelState // SysModelState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewSysModelState(models.SysModelState(stateModel))
	}
	svg, _ := d.Get("svg").(string)
	systemMfgName, _ := d.Get("system_mfg_name").(string)
	title, _ := d.Get("title").(string)
	return &models.SysBrand{
		Attr:          attr,
		Description:   description,
		ID:            id,
		Logo:          logo,
		Name:          &name, // string
		OriginType:    originType,
		Revision:      revision,
		State:         state,
		Svg:           svg,
		SystemMfgName: systemMfgName,
		Title:         &title, // string
	}
}

func SysBrandModelFromMap(m map[string]interface{}) *models.SysBrand {
	attr := map[string]string{}
	attrInterface, attrIsSet := m["attr"]
	if attrIsSet {
		attrMap := attrInterface.(map[string]interface{})
		for k, v := range attrMap {
			if v == nil {
				continue
			}
			attr[k] = v.(string)
		}
	}

	description := m["description"].(string)
	id := m["id"].(string)
	logo := map[string]string{}
	logoInterface, logoIsSet := m["logo"]
	if logoIsSet {
		logoMap := logoInterface.(map[string]interface{})
		for k, v := range logoMap {
			if v == nil {
				continue
			}
			logo[k] = v.(string)
		}
	}

	name := m["name"].(string)
	var originType *models.Origin // Origin
	originTypeInterface, originTypeIsSet := m["origin_type"]
	if originTypeIsSet {
		originTypeModel := originTypeInterface.(string)
		originType = models.NewOrigin(models.Origin(originTypeModel))
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
	var state *models.SysModelState // SysModelState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewSysModelState(models.SysModelState(stateModel))
	}
	svg := m["svg"].(string)
	systemMfgName := m["system_mfg_name"].(string)
	title := m["title"].(string)
	return &models.SysBrand{
		Attr:          attr,
		Description:   description,
		ID:            id,
		Logo:          logo,
		Name:          &name,
		OriginType:    originType,
		Revision:      revision,
		State:         state,
		Svg:           svg,
		SystemMfgName: systemMfgName,
		Title:         &title,
	}
}

func SetSysBrandResourceData(d *schema.ResourceData, m *models.SysBrand) {
	d.Set("attr", m.Attr)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("logo", m.Logo)
	d.Set("name", m.Name)
	d.Set("origin_type", m.OriginType)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("state", m.State)
	d.Set("svg", m.Svg)
	d.Set("system_mfg_name", m.SystemMfgName)
	d.Set("title", m.Title)
}

func SetSysBrandSubResourceData(m []*models.SysBrand) (d []*map[string]interface{}) {
	for _, SysBrandModel := range m {
		if SysBrandModel != nil {
			properties := make(map[string]interface{})
			properties["attr"] = SysBrandModel.Attr
			properties["description"] = SysBrandModel.Description
			properties["id"] = SysBrandModel.ID
			properties["logo"] = SysBrandModel.Logo
			properties["name"] = SysBrandModel.Name
			properties["origin_type"] = SysBrandModel.OriginType
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{SysBrandModel.Revision})
			properties["state"] = SysBrandModel.State
			properties["svg"] = SysBrandModel.Svg
			properties["system_mfg_name"] = SysBrandModel.SystemMfgName
			properties["title"] = SysBrandModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func SysBrandSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attr": {
			Description: `Map of <string, string>`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"description": {
			Description: `Detailed description of the image.`,
			Type:        schema.TypeString,
			Optional:    true,
			DiffSuppressFunc: supress(),
		},

		"id": {
			Description: `System defined universally unique Id of the brand.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"logo": {
			Description: `Map of <string, string> which holds the key:url for the logo artifact of the the brand`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"name": {
			Description: `user defined sys brand name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"origin_type": {
			Description: `origin of object`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"revision": {
			Description: `Object Revision  of the sys brand`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Optional: true,
			DiffSuppressFunc: supress(),
		},

		"state": {
			Description: `Sys Model Status`,
			Type:        schema.TypeString,
			Optional:    true,
			DiffSuppressFunc: supress(),
		},

		"svg": {
			Description: `Deprecated: base64 encoded string of svg file`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"system_mfg_name": {
			Description: `System Manufacturer name`,
			Type:        schema.TypeString,
			Optional:    true,
			DiffSuppressFunc: supress(),
		},

		"title": {
			Description: `user defined title for sys brand`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

func GetSysBrandPropertyFields() (t []string) {
	return []string{
		"attr",
		"description",
		"id",
		"logo",
		"name",
		"origin_type",
		"revision",
		"state",
		"svg",
		"system_mfg_name",
		"title",
	}
}
