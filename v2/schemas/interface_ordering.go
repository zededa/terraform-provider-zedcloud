package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func InterfaceOrderingModel(d *schema.ResourceData) *models.InterfaceOrdering {
	interfaceOrdering, _ := d.Get("interface_ordering").(models.InterfaceOrdering)
	return &interfaceOrdering
}

func InterfaceOrderingModelFromMap(m map[string]interface{}) *models.InterfaceOrdering {
	interfaceOrdering := m["interface_ordering"].(models.InterfaceOrdering)
	return &interfaceOrdering
}

func SetInterfaceOrderingResourceData(d *schema.ResourceData, m *models.InterfaceOrdering) {
}

func SetInterfaceOrderingSubResourceData(m []*models.InterfaceOrdering) (d []*map[string]interface{}) {
	for _, InterfaceOrderingModel := range m {
		if InterfaceOrderingModel != nil {
			properties := make(map[string]interface{})
			d = append(d, &properties)
		}
	}
	return
}

func InterfaceOrderingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func GetInterfaceOrderingPropertyFields() (t []string) {
	return []string{}
}
