package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoResourceStatusMapModel(d *schema.ResourceData) *models.GitRepoResourceStatusMap {
	aPIVersion, _ := d.Get("api_version").(string)
	kind, _ := d.Get("kind").(string)
	name, _ := d.Get("name").(string)
	namespace, _ := d.Get("namespace").(string)
	state, _ := d.Get("state").(string)
	return &models.GitRepoResourceStatusMap{
		APIVersion: aPIVersion,
		Kind:       kind,
		Name:       name,
		Namespace:  namespace,
		State:      state,
	}
}

func GitRepoResourceStatusMapModelFromMap(m map[string]interface{}) *models.GitRepoResourceStatusMap {
	aPIVersion := m["api_version"].(string)
	kind := m["kind"].(string)
	name := m["name"].(string)
	namespace := m["namespace"].(string)
	state := m["state"].(string)
	return &models.GitRepoResourceStatusMap{
		APIVersion: aPIVersion,
		Kind:       kind,
		Name:       name,
		Namespace:  namespace,
		State:      state,
	}
}

func SetGitRepoResourceStatusMapResourceData(d *schema.ResourceData, m *models.GitRepoResourceStatusMap) {
	d.Set("api_version", m.APIVersion)
	d.Set("kind", m.Kind)
	d.Set("name", m.Name)
	d.Set("namespace", m.Namespace)
	d.Set("state", m.State)
}

func SetGitRepoResourceStatusMapSubResourceData(m []*models.GitRepoResourceStatusMap) (d []*map[string]interface{}) {
	for _, GitRepoResourceStatusMapModel := range m {
		if GitRepoResourceStatusMapModel != nil {
			properties := make(map[string]interface{})
			properties["api_version"] = GitRepoResourceStatusMapModel.APIVersion
			properties["kind"] = GitRepoResourceStatusMapModel.Kind
			properties["name"] = GitRepoResourceStatusMapModel.Name
			properties["namespace"] = GitRepoResourceStatusMapModel.Namespace
			properties["state"] = GitRepoResourceStatusMapModel.State
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoResourceStatusMapSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_version": {
			Description: `API version of the resource`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"kind": {
			Description: `Kind of Kubernetes resource`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"name": {
			Description: `Name of the resource`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"namespace": {
			Description: `Namespace of the resource`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"state": {
			Description: `State of the resource`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGitRepoResourceStatusMapPropertyFields() (t []string) {
	return []string{
		"api_version",
		"kind",
		"name",
		"namespace",
		"state",
	}
}
