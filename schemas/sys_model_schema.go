package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate SysModel resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func SysModelModel(d *schema.ResourceData) *models.SysModel {
	pCRTemplates := d.Get("p_c_r_templates").([]*models.PCRTemplate) // []*PCRTemplate
	attr := d.Get("attr").(map[string]string)                        // map[string]string
	brandID := d.Get("brand_id").(string)
	description := d.Get("description").(string)
	id := d.Get("id").(string)
	ioMemberList := d.Get("io_member_list").([]*models.IoMember) // []*IoMember
	isImported := d.Get("is_imported").(bool)
	logo := d.Get("logo").(map[string]string) // map[string]string
	name := d.Get("name").(string)
	originType := d.Get("origin_type").(*models.Origin) // Origin
	var parentDetail *models.ObjectParentDetail         // ObjectParentDetail
	parentDetailInterface, parentDetailIsSet := d.GetOk("parent_detail")
	if parentDetailIsSet {
		parentDetailMap := parentDetailInterface.([]interface{})[0].(map[string]interface{})
		parentDetail = ObjectParentDetailModelFromMap(parentDetailMap)
	}
	productStatus := d.Get("product_status").(string)
	productURL := d.Get("product_url").(string)
	state := d.Get("state").(*models.SysModelState) // SysModelState
	title := d.Get("title").(string)
	typeVar := d.Get("type").(*models.ModelArchType) // ModelArchType
	return &models.SysModel{
		PCRTemplates:  pCRTemplates,
		Attr:          attr,
		BrandID:       &brandID, // string true false false
		Description:   description,
		ID:            id,
		IoMemberList:  ioMemberList,
		IsImported:    isImported,
		Logo:          logo,
		Name:          &name, // string true false false
		OriginType:    originType,
		ParentDetail:  parentDetail,
		ProductStatus: productStatus,
		ProductURL:    productURL,
		State:         state,
		Title:         &title, // string true false false
		Type:          typeVar,
	}
}

func SysModelModelFromMap(m map[string]interface{}) *models.SysModel {
	pCRTemplates := m["p_c_r_templates"].([]*models.PCRTemplate) // []*PCRTemplate
	attr := m["attr"].(map[string]string)
	brandID := m["brand_id"].(string)
	description := m["description"].(string)
	id := m["id"].(string)
	ioMemberList := m["io_member_list"].([]*models.IoMember) // []*IoMember
	isImported := m["is_imported"].(bool)
	logo := m["logo"].(map[string]string)
	name := m["name"].(string)
	originType := m["origin_type"].(*models.Origin) // Origin
	var parentDetail *models.ObjectParentDetail     // ObjectParentDetail
	parentDetailInterface, parentDetailIsSet := m["parent_detail"]
	if parentDetailIsSet {
		parentDetailMap := parentDetailInterface.([]interface{})[0].(map[string]interface{})
		parentDetail = ObjectParentDetailModelFromMap(parentDetailMap)
	}
	//
	productStatus := m["product_status"].(string)
	productURL := m["product_url"].(string)
	state := m["state"].(*models.SysModelState) // SysModelState
	title := m["title"].(string)
	typeVar := m["type"].(*models.ModelArchType) // ModelArchType
	return &models.SysModel{
		PCRTemplates:  pCRTemplates,
		Attr:          attr,
		BrandID:       &brandID,
		Description:   description,
		ID:            id,
		IoMemberList:  ioMemberList,
		IsImported:    isImported,
		Logo:          logo,
		Name:          &name,
		OriginType:    originType,
		ParentDetail:  parentDetail,
		ProductStatus: productStatus,
		ProductURL:    productURL,
		State:         state,
		Title:         &title,
		Type:          typeVar,
	}
}

// Update the underlying SysModel resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetSysModelResourceData(d *schema.ResourceData, m *models.SysModel) {
	d.Set("p_c_r_templates", SetPCRTemplateSubResourceData(m.PCRTemplates))
	d.Set("attr", m.Attr)
	d.Set("brand_id", m.BrandID)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("io_member_list", SetIoMemberSubResourceData(m.IoMemberList))
	d.Set("is_imported", m.IsImported)
	d.Set("logo", m.Logo)
	d.Set("name", m.Name)
	d.Set("origin_type", m.OriginType)
	d.Set("parent_detail", SetObjectParentDetailSubResourceData([]*models.ObjectParentDetail{m.ParentDetail}))
	d.Set("product_status", m.ProductStatus)
	d.Set("product_url", m.ProductURL)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("state", m.State)
	d.Set("title", m.Title)
	d.Set("type", m.Type)
}

// Iterate throught and update the SysModel resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetSysModelSubResourceData(m []*models.SysModel) (d []*map[string]interface{}) {
	for _, SysModelModel := range m {
		if SysModelModel != nil {
			properties := make(map[string]interface{})
			properties["p_c_r_templates"] = SetPCRTemplateSubResourceData(SysModelModel.PCRTemplates)
			properties["attr"] = SysModelModel.Attr
			properties["brand_id"] = SysModelModel.BrandID
			properties["description"] = SysModelModel.Description
			properties["id"] = SysModelModel.ID
			properties["io_member_list"] = SetIoMemberSubResourceData(SysModelModel.IoMemberList)
			properties["is_imported"] = SysModelModel.IsImported
			properties["logo"] = SysModelModel.Logo
			properties["name"] = SysModelModel.Name
			properties["origin_type"] = SysModelModel.OriginType
			properties["parent_detail"] = SetObjectParentDetailSubResourceData([]*models.ObjectParentDetail{SysModelModel.ParentDetail})
			properties["product_status"] = SysModelModel.ProductStatus
			properties["product_url"] = SysModelModel.ProductURL
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{SysModelModel.Revision})
			properties["state"] = SysModelModel.State
			properties["title"] = SysModelModel.Title
			properties["type"] = SysModelModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the SysModel resource defined in the Terraform configuration
func SysModelSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"p_c_r_templates": {
			Description: `PCR templates keyed by EVE version`,
			Type:        schema.TypeList, //GoType: []*PCRTemplate
			Elem: &schema.Resource{
				Schema: PCRTemplateSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"attr": {
			Description: `Map of <string, string> which defines attr`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},

		"brand_id": {
			Description: `System defined universally unique Id of the brand.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"description": {
			Description: `Detailed description of the model.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the model.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"io_member_list": {
			Description: `List of IoMembers`,
			Type:        schema.TypeList, //GoType: []*IoMember
			Elem: &schema.Resource{
				Schema: IoMemberSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"is_imported": {
			Description: `Flag to represent whether sysModel is imported`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"logo": {
			Description: `Map of <string, string> which holds the key:url for the logo artifact of the model`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"name": {
			Description: `user defined model name`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"origin_type": {
			Description: `origin of object`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},

		"parent_detail": {
			Description: `origin and parent related details`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Optional: true,
		},

		"product_status": {
			Description: `Product status`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"product_url": {
			Description: `Product URL`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"revision": {
			Description: `Object Revision  of the model`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Computed: true,
		},

		"state": {
			Description: `SysModel State which denotes the status of the model`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},

		"title": {
			Description: `User defined title of the model. Title can be changed at any time.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"type": {
			Description: `Defines the Architecture type of the model`,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Retrieve property field names for updating the SysModel resource
func GetSysModelPropertyFields() (t []string) {
	return []string{
		"p_c_r_templates",
		"attr",
		"brand_id",
		"description",
		"id",
		"io_member_list",
		"is_imported",
		"logo",
		"name",
		"origin_type",
		"parent_detail",
		"product_status",
		"product_url",
		"state",
		"title",
		"type",
	}
}
