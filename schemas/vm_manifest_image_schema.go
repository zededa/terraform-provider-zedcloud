package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func VMManifestImageModel(d *schema.ResourceData) *models.VMManifestImage {
	cleartext, _ := d.Get("cleartext").(bool)
	drvtype, _ := d.Get("drvtype").(string)
	ignorepurge, _ := d.Get("ignorepurge").(bool)
	var imageformat *models.ConfigFormat // ConfigFormat
	imageformatInterface, imageformatIsSet := d.GetOk("imageformat")
	if imageformatIsSet {
		imageformatModel := imageformatInterface.(string)
		imageformat = models.NewConfigFormat(models.ConfigFormat(imageformatModel))
	}
	imageid, _ := d.Get("imageid").(string)
	imagename, _ := d.Get("imagename").(string)
	maxsize, _ := d.Get("maxsize").(string)
	mountpath, _ := d.Get("mountpath").(string)
	var params []*models.Param // []*Param
	paramsInterface, paramsIsSet := d.GetOk("params")
	if paramsIsSet {
		var items []interface{}
		if listItems, isList := paramsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = paramsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ParamModelFromMap(v.(map[string]interface{}))
			params = append(params, m)
		}
	}
	preserve, _ := d.Get("preserve").(bool)
	readonly, _ := d.Get("readonly").(bool)
	target, _ := d.Get("target").(string)
	volumelabel, _ := d.Get("volumelabel").(string)
	return &models.VMManifestImage{
		Cleartext:   cleartext,
		Drvtype:     drvtype,
		Ignorepurge: ignorepurge,
		Imageformat: imageformat,
		Imageid:     imageid,
		Imagename:   imagename,
		Maxsize:     maxsize,
		Mountpath:   mountpath,
		Params:      params,
		Preserve:    preserve,
		Readonly:    readonly,
		Target:      target,
		Volumelabel: volumelabel,
	}
}

func VMManifestImageModelFromMap(m map[string]interface{}) *models.VMManifestImage {
	cleartext := m["cleartext"].(bool)
	drvtype := m["drvtype"].(string)
	ignorepurge := m["ignorepurge"].(bool)
	var imageformat *models.ConfigFormat // ConfigFormat
	imageformatInterface, imageformatIsSet := m["imageformat"]
	if imageformatIsSet {
		imageformatModel := imageformatInterface.(string)
		imageformat = models.NewConfigFormat(models.ConfigFormat(imageformatModel))
	}
	imageid := m["imageid"].(string)
	imagename := m["imagename"].(string)
	maxsize := m["maxsize"].(string)
	mountpath := m["mountpath"].(string)
	var params []*models.Param // []*Param
	paramsInterface, paramsIsSet := m["params"]
	if paramsIsSet {
		var items []interface{}
		if listItems, isList := paramsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = paramsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ParamModelFromMap(v.(map[string]interface{}))
			params = append(params, m)
		}
	}
	preserve := m["preserve"].(bool)
	readonly := m["readonly"].(bool)
	target := m["target"].(string)
	volumelabel := m["volumelabel"].(string)
	return &models.VMManifestImage{
		Cleartext:   cleartext,
		Drvtype:     drvtype,
		Ignorepurge: ignorepurge,
		Imageformat: imageformat,
		Imageid:     imageid,
		Imagename:   imagename,
		Maxsize:     maxsize,
		Mountpath:   mountpath,
		Params:      params,
		Preserve:    preserve,
		Readonly:    readonly,
		Target:      target,
		Volumelabel: volumelabel,
	}
}

// Update the underlying VMManifestImage resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVMManifestImageResourceData(d *schema.ResourceData, m *models.VMManifestImage) {
	d.Set("cleartext", m.Cleartext)
	d.Set("drvtype", m.Drvtype)
	d.Set("ignorepurge", m.Ignorepurge)
	d.Set("imageformat", m.Imageformat)
	d.Set("imageid", m.Imageid)
	d.Set("imagename", m.Imagename)
	d.Set("maxsize", m.Maxsize)
	d.Set("mountpath", m.Mountpath)
	d.Set("params", SetParamSubResourceData(m.Params))
	d.Set("preserve", m.Preserve)
	d.Set("readonly", m.Readonly)
	d.Set("target", m.Target)
	d.Set("volumelabel", m.Volumelabel)
}

// Iterate through and update the VMManifestImage resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVMManifestImageSubResourceData(m []*models.VMManifestImage) (d []*map[string]interface{}) {
	for _, VMManifestImageModel := range m {
		if VMManifestImageModel != nil {
			properties := make(map[string]interface{})
			properties["cleartext"] = VMManifestImageModel.Cleartext
			properties["drvtype"] = VMManifestImageModel.Drvtype
			properties["ignorepurge"] = VMManifestImageModel.Ignorepurge
			properties["imageformat"] = VMManifestImageModel.Imageformat
			properties["imageid"] = VMManifestImageModel.Imageid
			properties["imagename"] = VMManifestImageModel.Imagename
			properties["maxsize"] = VMManifestImageModel.Maxsize
			properties["mountpath"] = VMManifestImageModel.Mountpath
			properties["params"] = SetParamSubResourceData(VMManifestImageModel.Params)
			properties["preserve"] = VMManifestImageModel.Preserve
			properties["readonly"] = VMManifestImageModel.Readonly
			properties["target"] = VMManifestImageModel.Target
			properties["volumelabel"] = VMManifestImageModel.Volumelabel
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VMManifestImage resource defined in the Terraform configuration
func VMManifestImageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cleartext": {
			Description: `UI map: AppEditPage:DrivesPane:Cleartext, AppDetailsPage:DrivesPane:ClearText_Field`,
			Type:        schema.TypeBool,

			Optional: true,
		},

		"drvtype": {
			Description: `UI map: AppEditPage:DrivesPane:Drive_Type_Field, AppDetailsPage:DrivesPane:Drive_Type_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ignorepurge": {
			Description: `UI map: AppEditPage:DrivesPane:Ignorepurge, AppDetailsPage:DrivesPane:Ignorepurgee_Field`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"imageformat": {
			Description: `UI map: AppEditPage:DrivesPane:Image_Format_Field, AppDetailsPage:DrivesPane:Image_Format_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"imageid": {
			Description: `UI map: AppEditPage:DrivesPane:Image_ID_Field, AppDetailsPage:DrivesPane:Image_ID_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"imagename": {
			Description: `UI map: AppEditPage:DrivesPane:Image_Name_Field, AppDetailsPage:DrivesPane:Image_Name_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"maxsize": {
			Description: `UI map: AppEditPage:DrivesPane:Max_Size_Field, AppDetailsPage:DrivesPane:Max_Size_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"mountpath": {
			Description: `UI map: AppEditPage:DrivesPane:Mountpath, AppDetailsPage:DrivesPane:Mountpath_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"params": {
			Description: ``,
			Type:        schema.TypeList, //GoType: []*Param
			Elem: &schema.Resource{
				Schema: ParamSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"preserve": {
			Description: `UI map: AppEditPage:DrivesPane:Preserve_Field, AppDetailsPage:DrivesPane:Preserve_Field`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"readonly": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"target": {
			Description: `UI map: AppEditPage:DrivesPane:Target_Field, AppDetailsPage:DrivesPane:Target_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"volumelabel": {
			Description: `UI map: AppEditPage:DrivesPane:Volume_Label, AppDetailsPage:DrivesPane:Volume_Label`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VMManifestImage resource
func GetVMManifestImagePropertyFields() (t []string) {
	return []string{
		"cleartext",
		"drvtype",
		"ignorepurge",
		"imageformat",
		"imageid",
		"imagename",
		"maxsize",
		"mountpath",
		"params",
		"preserve",
		"readonly",
		"target",
		"volumelabel",
	}
}
