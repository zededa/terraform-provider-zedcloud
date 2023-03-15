package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func InterfaceModel(d *schema.ResourceData) *models.Interface {
	var acls []*models.ACL // []*ACL
	aclsInterface, aclsIsSet := d.GetOk("acls")
	if aclsIsSet {
		var items []interface{}
		if listItems, isList := aclsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = aclsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ACLModelFromMap(v.(map[string]interface{}))
			acls = append(acls, m)
		}
	}
	directattach, _ := d.Get("directattach").(bool)
	name, _ := d.Get("name").(string)
	optional, _ := d.Get("optional").(bool)
	privateip, _ := d.Get("privateip").(bool)
	typeVar, _ := d.Get("type").(string)
	return &models.Interface{
		Acls:         acls,
		Directattach: directattach,
		Name:         name,
		Optional:     optional,
		Privateip:    privateip,
		Type:         typeVar,
	}
}

func InterfaceModelFromMap(m map[string]interface{}) *models.Interface {
	var acls []*models.ACL // []*ACL
	aclsInterface, aclsIsSet := m["acls"]
	if aclsIsSet {
		var items []interface{}
		if listItems, isList := aclsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = aclsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := ACLModelFromMap(v.(map[string]interface{}))
			acls = append(acls, m)
		}
	}
	directattach := m["directattach"].(bool)
	name, _ := m["name"].(string)
	optional, _ := m["optional"].(bool)
	privateip, _ := m["privateip"].(bool)
	typeVar, _ := m["type"].(string)
	return &models.Interface{
		Acls:         acls,
		Directattach: directattach,
		Name:         name,
		Optional:     optional,
		Privateip:    privateip,
		Type:         typeVar,
	}
}

// Update the underlying Interface resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetInterfaceResourceData(d *schema.ResourceData, m *models.Interface) {
	d.Set("acls", SetACLSubResourceData(m.Acls))
	d.Set("directattach", m.Directattach)
	d.Set("name", m.Name)
	d.Set("optional", m.Optional)
	d.Set("privateip", m.Privateip)
	d.Set("type", m.Type)
}

// Iterate through and update the Interface resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetInterfaceSubResourceData(m []*models.Interface) (d []*map[string]interface{}) {
	for _, InterfaceModel := range m {
		if InterfaceModel != nil {
			properties := make(map[string]interface{})
			properties["acls"] = SetACLSubResourceData(InterfaceModel.Acls)
			properties["directattach"] = InterfaceModel.Directattach
			properties["name"] = InterfaceModel.Name
			properties["optional"] = InterfaceModel.Optional
			properties["privateip"] = InterfaceModel.Privateip
			properties["type"] = InterfaceModel.Type
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Interface resource defined in the Terraform configuration
func Interface() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"acls": {
			Description: `Traffic access control rules for this interface. Applicable only when "direct attach" flag is false.`,
			Type:        schema.TypeList, //GoType: []*ACL
			Elem: &schema.Resource{
				Schema: ACL(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"directattach": {
			Description: `If true, a physical adapter is assigned to the edge application directly. If false, a network instance is assigned to the edge application.`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"name": {
			Description: `Interface name used by the edge application`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"optional": {
			Description: `Indicates if the interface is optional for edge application.`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"privateip": {
			Description: `If true, DHCP network can't be assigned and user needs to provide a static IP address.`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"type": {
			Description: `Physical Adapter type for this interface. Applicable only when "direct attach" flag is true.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Interface resource
func GetInterfacePropertyFields() (t []string) {
	return []string{
		"acls",
		"directattach",
		"name",
		"optional",
		"privateip",
		"type",
	}
}

func CompareInterfaceLists(a, b []*models.Interface) bool {
	if len(a) != len(b) {
		return false
	}
	// is each element of the new list in the old list?
	for _, newList := range b {
		if newList == nil {
			continue
		}

		found := false
		for _, oldList := range a {
			if oldList == nil {
				continue
			}
			if oldList.Directattach != newList.Directattach {
				continue
			}
			if oldList.Name != newList.Name {
				continue
			}
			if oldList.Optional != newList.Optional {
				continue
			}
			if oldList.Privateip != newList.Privateip {
				continue
			}
			if oldList.Type != newList.Type {
				continue
			}

			if !CompareACLList(oldList.Acls, newList.Acls) {
				continue
			}
			found = true
			break
		}
		if !found {
			return false
		}
	}

	// is each element of the old list also in the new list?
	for _, oldList := range a {
		if oldList == nil {
			continue
		}

		found := false
		for _, newList := range b {
			if newList == nil {
				continue
			}
			if oldList.Directattach != newList.Directattach {
				continue
			}
			if oldList.Name != newList.Name {
				continue
			}
			if oldList.Optional != newList.Optional {
				continue
			}
			if oldList.Privateip != newList.Privateip {
				continue
			}
			if oldList.Type != newList.Type {
				continue
			}

			if !CompareACLList(oldList.Acls, newList.Acls) {
				continue
			}
			found = true
			break
		}
		if !found {
			return false
		}
	}
	return true
}
