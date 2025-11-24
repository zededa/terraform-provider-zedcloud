package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterGroupManifestOutputModel(d *schema.ResourceData) *models.ClusterGroupManifestOutput {
	var manifest *models.ClusterGroupManifestJaml // ClusterGroupManifestJaml
	manifestInterface, manifestIsSet := d.GetOk("manifest")
	if manifestIsSet && manifestInterface != nil {
		manifestMap := manifestInterface.([]interface{})
		if len(manifestMap) > 0 {
			manifest = ClusterGroupManifestJamlModelFromMap(manifestMap[0].(map[string]interface{}))
		}
	}
	return &models.ClusterGroupManifestOutput{
		Manifest: manifest,
	}
}

func ClusterGroupManifestOutputModelFromMap(m map[string]interface{}) *models.ClusterGroupManifestOutput {
	var manifest *models.ClusterGroupManifestJaml // ClusterGroupManifestJaml
	manifestInterface, manifestIsSet := m["manifest"]
	if manifestIsSet && manifestInterface != nil {
		manifestMap := manifestInterface.([]interface{})
		if len(manifestMap) > 0 {
			manifest = ClusterGroupManifestJamlModelFromMap(manifestMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ClusterGroupManifestOutput{
		Manifest: manifest,
	}
}

func SetClusterGroupManifestOutputResourceData(d *schema.ResourceData, m *models.ClusterGroupManifestOutput) {
	d.Set("manifest", SetClusterGroupManifestJamlSubResourceData([]*models.ClusterGroupManifestJaml{m.Manifest}))
}

func SetClusterGroupManifestOutputSubResourceData(m []*models.ClusterGroupManifestOutput) (d []*map[string]interface{}) {
	for _, ClusterGroupManifestOutputModel := range m {
		if ClusterGroupManifestOutputModel != nil {
			properties := make(map[string]interface{})
			properties["manifest"] = SetClusterGroupManifestJamlSubResourceData([]*models.ClusterGroupManifestJaml{ClusterGroupManifestOutputModel.Manifest})
			d = append(d, &properties)
		}
	}
	return
}

func ClusterGroupManifestOutputSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"manifest": {
			Description: `The cluster group manifest in jaml format`,
			Type:        schema.TypeList, //GoType: ClusterGroupManifestJaml
			Elem: &schema.Resource{
				Schema: ClusterGroupManifestJamlSchema(),
			},
			Optional: true,
		},
	}
}

func GetClusterGroupManifestOutputPropertyFields() (t []string) {
	return []string{
		"manifest",
	}
}
