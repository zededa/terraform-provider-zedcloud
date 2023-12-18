package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DocPolicyModel(d *schema.ResourceData) *models.DocPolicy {
	fileURL, _ := d.Get("file_url").(string)
	id, _ := d.Get("id").(string)
	latest, _ := d.Get("latest").(bool)
	policyName, _ := d.Get("policy_name").(string)
	serverHost, _ := d.Get("server_host").(string)
	version, _ := d.Get("version").(string)
	return &models.DocPolicy{
		FileURL:    fileURL,
		ID:         id,
		Latest:     latest,
		PolicyName: &policyName, // string
		ServerHost: serverHost,
		Version:    version,
	}
}

func DocPolicyModelFromMap(m map[string]interface{}) *models.DocPolicy {
	fileURL := m["file_url"].(string)
	id := m["id"].(string)
	latest := m["latest"].(bool)
	policyName := m["policy_name"].(string)
	serverHost := m["server_host"].(string)
	version := m["version"].(string)
	return &models.DocPolicy{
		FileURL:    fileURL,
		ID:         id,
		Latest:     latest,
		PolicyName: &policyName,
		ServerHost: serverHost,
		Version:    version,
	}
}

func SetDocPolicyResourceData(d *schema.ResourceData, m *models.DocPolicy) {
	d.Set("file_url", m.FileURL)
	d.Set("id", m.ID)
	d.Set("latest", m.Latest)
	d.Set("policy_name", m.PolicyName)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("server_host", m.ServerHost)
	d.Set("version", m.Version)
}

func SetDocPolicySubResourceData(m []*models.DocPolicy) (d []*map[string]interface{}) {
	for _, DocPolicyModel := range m {
		if DocPolicyModel != nil {
			properties := make(map[string]interface{})
			properties["file_url"] = DocPolicyModel.FileURL
			properties["id"] = DocPolicyModel.ID
			properties["latest"] = DocPolicyModel.Latest
			properties["policy_name"] = DocPolicyModel.PolicyName
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{DocPolicyModel.Revision})
			properties["server_host"] = DocPolicyModel.ServerHost
			properties["version"] = DocPolicyModel.Version
			d = append(d, &properties)
		}
	}
	return
}

func DocPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"file_url": {
			Description: `Policy doc fileURL`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"id": {
			Description: `Unique system defined docpolicy ID`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"latest": {
			Description: `Mark latest docpolicy check`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"policy_name": {
			Description: `User defined name of the docpolicy. Name cannot be changed once created`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"revision": {
			Description: `System defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"server_host": {
			Description: `Server Host`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"version": {
			Description: `Policy doc version`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetDocPolicyPropertyFields() (t []string) {
	return []string{
		"file_url",
		"id",
		"latest",
		"policy_name",
		"server_host",
		"version",
	}
}
