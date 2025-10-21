package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PhyAdapterModel(d *schema.ResourceData) *models.PhyAdapter {
	name, _ := d.Get("name").(string)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := d.GetOk("tags")
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	var typeVar *models.IoType // IoType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewIoType(models.IoType(typeModel))
	}
	return &models.PhyAdapter{
		Name: name,
		Tags: tags,
		Type: typeVar,
	}
}

func PhyAdapterModelFromMap(m map[string]interface{}) *models.PhyAdapter {
	name := m["name"].(string)
	tags := map[string]string{}
	tagsInterface, tagsIsSet := m["tags"]
	if tagsIsSet {
		tagsMap := tagsInterface.(map[string]interface{})
		for k, v := range tagsMap {
			if v == nil {
				continue
			}
			tags[k] = v.(string)
		}
	}

	var typeVar *models.IoType // IoType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewIoType(models.IoType(typeModel))
	}
	return &models.PhyAdapter{
		Name: name,
		Tags: tags,
		Type: typeVar,
	}
}

func SetPhyAdapterResourceData(d *schema.ResourceData, m *models.PhyAdapter) {
	d.Set("name", m.Name)
	d.Set("tags", m.Tags)
	d.Set("type", m.Type)
}

func SetPhyAdapterSubResourceData(m []*models.PhyAdapter) (d []*map[string]interface{}) {
	for _, PhyAdapterModel := range m {
		if PhyAdapterModel != nil {
			properties := make(map[string]interface{})
			properties["name"] = PhyAdapterModel.Name
			properties["tags"] = PhyAdapterModel.Tags
			properties["type"] = PhyAdapterModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func PhyAdapter() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Description: `Physical Adapter name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"tags": {
			Description: `Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"type": {
			Description: `IoType specifies the type of the Input output of the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPhyAdapterPropertyFields() (t []string) {
	return []string{
		"name",
		"tags",
		"type",
	}
}
