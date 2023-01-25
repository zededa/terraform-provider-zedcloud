package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func NetworkFilterModel(d *schema.ResourceData) *models.NetworkFilter {
	var dist *models.NetworkWirelessType // NetworkWirelessType
	distInterface, distIsSet := d.GetOk("dist")
	if distIsSet {
		distModel := distInterface.(string)
		dist = models.NewNetworkWirelessType(models.NetworkWirelessType(distModel))
	}
	var kind *models.NetworkKind // NetworkKind
	kindInterface, kindIsSet := d.GetOk("kind")
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkKind(models.NetworkKind(kindModel))
	}
	namePattern, _ := d.Get("name_pattern").(string)
	projectName, _ := d.Get("project_name").(string)
	projectNamePattern, _ := d.Get("project_name_pattern").(string)
	return &models.NetworkFilter{
		Dist:               dist,
		Kind:               kind,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
	}
}

func NetworkFilterModelFromMap(m map[string]interface{}) *models.NetworkFilter {
	var dist *models.NetworkWirelessType // NetworkWirelessType
	distInterface, distIsSet := m["dist"]
	if distIsSet {
		distModel := distInterface.(string)
		dist = models.NewNetworkWirelessType(models.NetworkWirelessType(distModel))
	}
	var kind *models.NetworkKind // NetworkKind
	kindInterface, kindIsSet := m["kind"]
	if kindIsSet {
		kindModel := kindInterface.(string)
		kind = models.NewNetworkKind(models.NetworkKind(kindModel))
	}
	namePattern := m["name_pattern"].(string)
	projectName := m["project_name"].(string)
	projectNamePattern := m["project_name_pattern"].(string)
	return &models.NetworkFilter{
		Dist:               dist,
		Kind:               kind,
		NamePattern:        namePattern,
		ProjectName:        projectName,
		ProjectNamePattern: projectNamePattern,
	}
}

// Update the underlying NetworkFilter resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNetworkFilterResourceData(d *schema.ResourceData, m *models.NetworkFilter) {
	d.Set("dist", m.Dist)
	d.Set("kind", m.Kind)
	d.Set("name_pattern", m.NamePattern)
	d.Set("project_name", m.ProjectName)
	d.Set("project_name_pattern", m.ProjectNamePattern)
}

// Iterate through and update the NetworkFilter resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNetworkFilterSubResourceData(m []*models.NetworkFilter) (d []*map[string]interface{}) {
	for _, NetworkFilterModel := range m {
		if NetworkFilterModel != nil {
			properties := make(map[string]interface{})
			properties["dist"] = NetworkFilterModel.Dist
			properties["kind"] = NetworkFilterModel.Kind
			properties["name_pattern"] = NetworkFilterModel.NamePattern
			properties["project_name"] = NetworkFilterModel.ProjectName
			properties["project_name_pattern"] = NetworkFilterModel.ProjectNamePattern
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the NetworkFilter resource defined in the Terraform configuration
func NetworkFilterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dist": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"kind": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name_pattern": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name_pattern": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the NetworkFilter resource
func GetNetworkFilterPropertyFields() (t []string) {
	return []string{
		"dist",
		"kind",
		"name_pattern",
		"project_name",
		"project_name_pattern",
	}
}
