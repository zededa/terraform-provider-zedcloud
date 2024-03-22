package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func ObjectTagFilterModel(d *schema.ResourceData) *models.ObjectTagFilter {
	objID, _ := d.Get("obj_id").(string)
	objName, _ := d.Get("obj_name").(string)
	var objType *models.ObjectType // ObjectType
	objTypeInterface, objTypeIsSet := d.GetOk("obj_type")
	if objTypeIsSet {
		objTypeModel := objTypeInterface.(string)
		objType = models.NewObjectType(models.ObjectType(objTypeModel))
	}
	return &models.ObjectTagFilter{
		ObjID:   objID,
		ObjName: objName,
		ObjType: objType,
	}
}

func ObjectTagFilterModelFromMap(m map[string]interface{}) *models.ObjectTagFilter {
	objID := m["obj_id"].(string)
	objName := m["obj_name"].(string)
	var objType *models.ObjectType // ObjectType
	objTypeInterface, objTypeIsSet := m["obj_type"]
	if objTypeIsSet {
		objTypeModel := objTypeInterface.(string)
		objType = models.NewObjectType(models.ObjectType(objTypeModel))
	}
	return &models.ObjectTagFilter{
		ObjID:   objID,
		ObjName: objName,
		ObjType: objType,
	}
}

// Update the underlying ObjectTagFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetObjectTagFilterResourceData(d *schema.ResourceData, m *models.ObjectTagFilter) {
	d.Set("obj_id", m.ObjID)
	d.Set("obj_name", m.ObjName)
	d.Set("obj_type", m.ObjType)
}

// Iterate through and update the ObjectTagFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetObjectTagFilterSubResourceData(m []*models.ObjectTagFilter) (d []*map[string]interface{}) {
	for _, ObjectTagFilterModel := range m {
		if ObjectTagFilterModel != nil {
			properties := make(map[string]interface{})
			properties["obj_id"] = ObjectTagFilterModel.ObjID
			properties["obj_name"] = ObjectTagFilterModel.ObjName
			properties["obj_type"] = ObjectTagFilterModel.ObjType
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the ObjectTagFilter resource defined in the Terraform configuration
func ObjectTagFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"obj_id": {
			Description: `Object Id which tags are associated.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"obj_name": {
			Description: `Object name which tags are associated.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"obj_type": {
			Description: `Object type`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the ObjectTagFilter resource
func GetObjectTagFilterPropertyFields() (t []string) {
	return []string{
		"obj_id",
		"obj_name",
		"obj_type",
	}
}
