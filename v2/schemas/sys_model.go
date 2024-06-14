package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SysModelModel(d *schema.ResourceData) *models.SysModel {
	var pCRTemplates []*models.PCRTemplate // []*PCRTemplate
	PCRTemplatesInterface, PCRTemplatesIsSet := d.GetOk("p_c_r_templates")
	if PCRTemplatesIsSet {
		var items []interface{}
		if listItems, isList := PCRTemplatesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = PCRTemplatesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PCRTemplateModelFromMap(v.(map[string]interface{}))
			pCRTemplates = append(pCRTemplates, m)
		}
	}
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

	brandID, _ := d.Get("brand_id").(string)
	description, _ := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	var ioMemberList []*models.IoMember // []*IoMember
	ioMemberListInterface, ioMemberListIsSet := d.GetOk("io_member_list")
	if ioMemberListIsSet {
		var items []interface{}
		if listItems, isList := ioMemberListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = ioMemberListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IoMemberModelFromMap(v.(map[string]interface{}))
			ioMemberList = append(ioMemberList, m)
		}
	}
	isImported, _ := d.Get("is_imported").(bool)
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
	var parentDetail *models.ObjectParentDetail // ObjectParentDetail
	parentDetailInterface, parentDetailIsSet := d.GetOk("parent_detail")
	if parentDetailIsSet && parentDetailInterface != nil {
		parentDetailMap := parentDetailInterface.([]interface{})
		if len(parentDetailMap) > 0 {
			parentDetail = ObjectParentDetailModelFromMap(parentDetailMap[0].(map[string]interface{}))
		}
	}
	productStatus, _ := d.Get("product_status").(string)
	productURL, _ := d.Get("product_url").(string)
	var state *models.SysModelState // SysModelState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewSysModelState(models.SysModelState(stateModel))
	}
	title, _ := d.Get("title").(string)
	var typeVar *models.ModelArchType // ModelArchType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewModelArchType(models.ModelArchType(typeModel))
	}
	return &models.SysModel{
		PCRTemplates:  pCRTemplates,
		Attr:          attr,
		BrandID:       &brandID, // string
		Description:   description,
		ID:            id,
		IoMemberList:  ioMemberList,
		IsImported:    isImported,
		Logo:          logo,
		Name:          &name, // string
		OriginType:    originType,
		ParentDetail:  parentDetail,
		ProductStatus: productStatus,
		ProductURL:    productURL,
		State:         state,
		Title:         &title, // string
		Type:          typeVar,
	}
}

func SysModelModelFromMap(m map[string]interface{}) *models.SysModel {
	var pCRTemplates []*models.PCRTemplate // []*PCRTemplate
	PCRTemplatesInterface, PCRTemplatesIsSet := m["p_c_r_templates"]
	if PCRTemplatesIsSet {
		var items []interface{}
		if listItems, isList := PCRTemplatesInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = PCRTemplatesInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := PCRTemplateModelFromMap(v.(map[string]interface{}))
			pCRTemplates = append(pCRTemplates, m)
		}
	}
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

	brandID := m["brand_id"].(string)
	description := m["description"].(string)
	id := m["id"].(string)
	var ioMemberList []*models.IoMember // []*IoMember
	ioMemberListInterface, ioMemberListIsSet := m["io_member_list"]
	if ioMemberListIsSet {
		var items []interface{}
		if listItems, isList := ioMemberListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = ioMemberListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := IoMemberModelFromMap(v.(map[string]interface{}))
			ioMemberList = append(ioMemberList, m)
		}
	}
	isImported := m["is_imported"].(bool)
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
	var parentDetail *models.ObjectParentDetail // ObjectParentDetail
	parentDetailInterface, parentDetailIsSet := m["parent_detail"]
	if parentDetailIsSet && parentDetailInterface != nil {
		parentDetailMap := parentDetailInterface.([]interface{})
		if len(parentDetailMap) > 0 {
			parentDetail = ObjectParentDetailModelFromMap(parentDetailMap[0].(map[string]interface{}))
		}
	}
	//
	productStatus := m["product_status"].(string)
	productURL := m["product_url"].(string)
	var state *models.SysModelState // SysModelState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet {
		stateModel := stateInterface.(string)
		state = models.NewSysModelState(models.SysModelState(stateModel))
	}
	title := m["title"].(string)
	var typeVar *models.ModelArchType // ModelArchType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewModelArchType(models.ModelArchType(typeModel))
	}
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
			DiffSuppressFunc: diffSuppressIoMemberListOrder("io_member_list"),
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
			Type:        schema.TypeString,
			Required:    true,
		},

		"parent_detail": {
			Description: `origin and parent related details`,
			Type:        schema.TypeList, //GoType: ObjectParentDetail
			Elem: &schema.Resource{
				Schema: ObjectParentDetail(),
			},
			Optional: true,
			DiffSuppressFunc: supress(),
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
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"state": {
			Description: `SysModel State which denotes the status of the model`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"title": {
			Description: `User defined title of the model. Title can be changed at any time.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"type": {
			Description: `Defines the Architecture type of the model`,
			Type:        schema.TypeString,
			Required:    true,
		},
	}
}

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
