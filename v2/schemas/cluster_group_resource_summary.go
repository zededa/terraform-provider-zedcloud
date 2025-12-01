package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ClusterGroupResourceSummaryModel(d *schema.ResourceData) *models.ClusterGroupResourceSummary {
	missingInt, _ := d.Get("missing").(int)
	missing := int64(missingInt)
	modifiedInt, _ := d.Get("modified").(int)
	modified := int64(modifiedInt)
	notReadyInt, _ := d.Get("not_ready").(int)
	notReady := int64(notReadyInt)
	orphanedInt, _ := d.Get("orphaned").(int)
	orphaned := int64(orphanedInt)
	readyInt, _ := d.Get("ready").(int)
	ready := int64(readyInt)
	unknownInt, _ := d.Get("unknown").(int)
	unknown := int64(unknownInt)
	waitAppliedInt, _ := d.Get("wait_applied").(int)
	waitApplied := int64(waitAppliedInt)
	return &models.ClusterGroupResourceSummary{
		Missing:     missing,
		Modified:    modified,
		NotReady:    notReady,
		Orphaned:    orphaned,
		Ready:       ready,
		Unknown:     unknown,
		WaitApplied: waitApplied,
	}
}

func ClusterGroupResourceSummaryModelFromMap(m map[string]interface{}) *models.ClusterGroupResourceSummary {
	missing := int64(m["missing"].(int))          // int64
	modified := int64(m["modified"].(int))        // int64
	notReady := int64(m["not_ready"].(int))       // int64
	orphaned := int64(m["orphaned"].(int))        // int64
	ready := int64(m["ready"].(int))              // int64
	unknown := int64(m["unknown"].(int))          // int64
	waitApplied := int64(m["wait_applied"].(int)) // int64
	return &models.ClusterGroupResourceSummary{
		Missing:     missing,
		Modified:    modified,
		NotReady:    notReady,
		Orphaned:    orphaned,
		Ready:       ready,
		Unknown:     unknown,
		WaitApplied: waitApplied,
	}
}

func SetClusterGroupResourceSummaryResourceData(d *schema.ResourceData, m *models.ClusterGroupResourceSummary) {
	d.Set("missing", m.Missing)
	d.Set("modified", m.Modified)
	d.Set("not_ready", m.NotReady)
	d.Set("orphaned", m.Orphaned)
	d.Set("ready", m.Ready)
	d.Set("unknown", m.Unknown)
	d.Set("wait_applied", m.WaitApplied)
}

func SetClusterGroupResourceSummarySubResourceData(m []*models.ClusterGroupResourceSummary) (d []*map[string]interface{}) {
	for _, ClusterGroupResourceSummaryModel := range m {
		if ClusterGroupResourceSummaryModel != nil {
			properties := make(map[string]interface{})
			properties["missing"] = ClusterGroupResourceSummaryModel.Missing
			properties["modified"] = ClusterGroupResourceSummaryModel.Modified
			properties["not_ready"] = ClusterGroupResourceSummaryModel.NotReady
			properties["orphaned"] = ClusterGroupResourceSummaryModel.Orphaned
			properties["ready"] = ClusterGroupResourceSummaryModel.Ready
			properties["unknown"] = ClusterGroupResourceSummaryModel.Unknown
			properties["wait_applied"] = ClusterGroupResourceSummaryModel.WaitApplied
			d = append(d, &properties)
		}
	}
	return
}

func ClusterGroupResourceSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"missing": {
			Description: `Number of resources in the cluster group that are missing`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"modified": {
			Description: `Number of resources in the cluster group that are modified`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"not_ready": {
			Description: `Number of resources in the cluster group that are not ready`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"orphaned": {
			Description: `Number of resources in the cluster group that are orphaned`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"ready": {
			Description: `Number of resources in the cluster group that are ready`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"unknown": {
			Description: `Number of resources in the cluster group that are in unknown state`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"wait_applied": {
			Description: `Number of resources in the cluster group that are waiting for applied state`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetClusterGroupResourceSummaryPropertyFields() (t []string) {
	return []string{
		"missing",
		"modified",
		"not_ready",
		"orphaned",
		"ready",
		"unknown",
		"wait_applied",
	}
}
