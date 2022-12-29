package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SysBrand resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SysBrandModel(d *schema.ResourceData) *models.SysBrand {
	attr, _ := d.Get("attr").(map[string]string) // map[string]string
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	logo, _ := d.Get("logo").(map[string]string) // map[string]string
	name, _ := d.Get("name").(string)
	originTypeModel, ok := d.Get("origin_type").(models.Origin) // Origin
	originType := &originTypeModel
	if !ok {
		originType = nil
	}
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	stateModel, ok := d.Get("state").(models.SysModelState) // SysModelState
	state := &stateModel
	if !ok {
		state = nil
	}
	svg, _ := d.Get("svg").(string)
	systemMfgName, _ := d.Get("system_mfg_name").(string)
	title, _ := d.Get("title").(string)
	return &models.SysBrand{
		Attr:          attr,
		Description:   description,
		ID:            id,
		Logo:          logo,
		Name:          &name, // string true false false
		OriginType:    originType,
		Revision:      revision,
		State:         state,
		Svg:           svg,
		SystemMfgName: systemMfgName,
		Title:         &title, // string true false false
	}
}

func SysBrandModelFromMap(m map[string]interface{}) *models.SysBrand {
	attr := m["attr"].(map[string]string)
	description := m["description"].(string)
	id := m["id"].(string)
	logo := m["logo"].(map[string]string)
	name := m["name"].(string)
	originType := m["origin_type"].(*models.Origin) // Origin
	var revision *models.ObjectRevision             // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	//
	state := m["state"].(*models.SysModelState) // SysModelState
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

// Update the underlying SysBrand resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
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

// Iterate through and update the SysBrand resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Schema mapping representing the SysBrand resource defined in the Terraform configuration
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
				Schema: ObjectRevisionSchema(),
			},
			Optional: true,
		},

		"state": {
			Description: `Sys Model Status`,
			Type:        schema.TypeString,
			Optional:    true,
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
		},

		"title": {
			Description: `user defined title for sys brand`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

// Retrieve property field names for updating the SysBrand resource
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
