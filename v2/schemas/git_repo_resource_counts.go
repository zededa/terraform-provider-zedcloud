package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoResourceCountsModel(d *schema.ResourceData) *models.GitRepoResourceCounts {
	desiredReadyInt, _ := d.Get("desired_ready").(int)
	desiredReady := int32(desiredReadyInt)
	missingInt, _ := d.Get("missing").(int)
	missing := int32(missingInt)
	modifiedInt, _ := d.Get("modified").(int)
	modified := int32(modifiedInt)
	notReadyInt, _ := d.Get("not_ready").(int)
	notReady := int32(notReadyInt)
	orphanedInt, _ := d.Get("orphaned").(int)
	orphaned := int32(orphanedInt)
	readyInt, _ := d.Get("ready").(int)
	ready := int32(readyInt)
	unknownInt, _ := d.Get("unknown").(int)
	unknown := int32(unknownInt)
	waitAppliedInt, _ := d.Get("wait_applied").(int)
	waitApplied := int32(waitAppliedInt)
	return &models.GitRepoResourceCounts{
		DesiredReady: desiredReady,
		Missing:      missing,
		Modified:     modified,
		NotReady:     notReady,
		Orphaned:     orphaned,
		Ready:        ready,
		Unknown:      unknown,
		WaitApplied:  waitApplied,
	}
}

func GitRepoResourceCountsModelFromMap(m map[string]interface{}) *models.GitRepoResourceCounts {
	desiredReady := int32(m["desired_ready"].(int)) // int32
	missing := int32(m["missing"].(int))            // int32
	modified := int32(m["modified"].(int))          // int32
	notReady := int32(m["not_ready"].(int))         // int32
	orphaned := int32(m["orphaned"].(int))          // int32
	ready := int32(m["ready"].(int))                // int32
	unknown := int32(m["unknown"].(int))            // int32
	waitApplied := int32(m["wait_applied"].(int))   // int32
	return &models.GitRepoResourceCounts{
		DesiredReady: desiredReady,
		Missing:      missing,
		Modified:     modified,
		NotReady:     notReady,
		Orphaned:     orphaned,
		Ready:        ready,
		Unknown:      unknown,
		WaitApplied:  waitApplied,
	}
}

func SetGitRepoResourceCountsResourceData(d *schema.ResourceData, m *models.GitRepoResourceCounts) {
	d.Set("desired_ready", m.DesiredReady)
	d.Set("missing", m.Missing)
	d.Set("modified", m.Modified)
	d.Set("not_ready", m.NotReady)
	d.Set("orphaned", m.Orphaned)
	d.Set("ready", m.Ready)
	d.Set("unknown", m.Unknown)
	d.Set("wait_applied", m.WaitApplied)
}

func SetGitRepoResourceCountsSubResourceData(m []*models.GitRepoResourceCounts) (d []*map[string]interface{}) {
	for _, GitRepoResourceCountsModel := range m {
		if GitRepoResourceCountsModel != nil {
			properties := make(map[string]interface{})
			properties["desired_ready"] = GitRepoResourceCountsModel.DesiredReady
			properties["missing"] = GitRepoResourceCountsModel.Missing
			properties["modified"] = GitRepoResourceCountsModel.Modified
			properties["not_ready"] = GitRepoResourceCountsModel.NotReady
			properties["orphaned"] = GitRepoResourceCountsModel.Orphaned
			properties["ready"] = GitRepoResourceCountsModel.Ready
			properties["unknown"] = GitRepoResourceCountsModel.Unknown
			properties["wait_applied"] = GitRepoResourceCountsModel.WaitApplied
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoResourceCountsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"desired_ready": {
			Description: `Desired ready count`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"missing": {
			Description: `Missing resources count`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"modified": {
			Description: `Modified resources count`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"not_ready": {
			Description: `Not ready resources count`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"orphaned": {
			Description: `Orphaned resources count`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"ready": {
			Description: `Ready resources count`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"unknown": {
			Description: `Unknown resources count`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"wait_applied": {
			Description: `Wait applied resources count`,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetGitRepoResourceCountsPropertyFields() (t []string) {
	return []string{
		"desired_ready",
		"missing",
		"modified",
		"not_ready",
		"orphaned",
		"ready",
		"unknown",
		"wait_applied",
	}
}
