package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterReadinessModel(d *schema.ResourceData) *models.ClusterReadiness {
	readyInt, _ := d.Get("ready").(int)
	ready := int64(readyInt)
	totalInt, _ := d.Get("total").(int)
	total := int64(totalInt)
	return &models.ClusterReadiness{
		Ready: ready,
		Total: total,
	}
}

func ClusterReadinessModelFromMap(m map[string]interface{}) *models.ClusterReadiness {
	ready := int64(m["ready"].(int)) // int64
	total := int64(m["total"].(int)) // int64
	return &models.ClusterReadiness{
		Ready: ready,
		Total: total,
	}
}

func SetClusterReadinessResourceData(d *schema.ResourceData, m *models.ClusterReadiness) {
	d.Set("ready", m.Ready)
	d.Set("total", m.Total)
}

func SetClusterReadinessSubResourceData(m []*models.ClusterReadiness) (d []*map[string]interface{}) {
	for _, ClusterReadinessModel := range m {
		if ClusterReadinessModel != nil {
			properties := make(map[string]interface{})
			properties["ready"] = ClusterReadinessModel.Ready
			properties["total"] = ClusterReadinessModel.Total
			d = append(d, &properties)
		}
	}
	return
}

func ClusterReadinessSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ready": {
			Description: `Number of clusters in the cluster group that are ready`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"total": {
			Description: `Total number of clusters in the cluster group`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetClusterReadinessPropertyFields() (t []string) {
	return []string{
		"ready",
		"total",
	}
}
