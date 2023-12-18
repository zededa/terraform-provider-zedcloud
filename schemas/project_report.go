package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func ProjectReportModel(d *schema.ResourceData) *models.ProjectReport {
	entpID, _ := d.Get("entp_id").(string)
	error, _ := d.Get("error").(string)
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := d.GetOk("next")
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	var projectSummaryReport *models.ProjectSummaryReport // ProjectSummaryReport
	projectSummaryReportInterface, projectSummaryReportIsSet := d.GetOk("project_summary_report")
	if projectSummaryReportIsSet && projectSummaryReportInterface != nil {
		projectSummaryReportMap := projectSummaryReportInterface.([]interface{})
		if len(projectSummaryReportMap) > 0 {
			projectSummaryReport = ProjectSummaryReportModelFromMap(projectSummaryReportMap[0].(map[string]interface{}))
		}
	}
	return &models.ProjectReport{
		EntpID:               entpID,
		Error:                error,
		Next:                 next,
		ProjectSummaryReport: projectSummaryReport,
	}
}

func ProjectReportModelFromMap(m map[string]interface{}) *models.ProjectReport {
	entpID := m["entp_id"].(string)
	error := m["error"].(string)
	var next *models.Cursor // Cursor
	nextInterface, nextIsSet := m["next"]
	if nextIsSet && nextInterface != nil {
		nextMap := nextInterface.([]interface{})
		if len(nextMap) > 0 {
			next = CursorModelFromMap(nextMap[0].(map[string]interface{}))
		}
	}
	//
	var projectSummaryReport *models.ProjectSummaryReport // ProjectSummaryReport
	projectSummaryReportInterface, projectSummaryReportIsSet := m["project_summary_report"]
	if projectSummaryReportIsSet && projectSummaryReportInterface != nil {
		projectSummaryReportMap := projectSummaryReportInterface.([]interface{})
		if len(projectSummaryReportMap) > 0 {
			projectSummaryReport = ProjectSummaryReportModelFromMap(projectSummaryReportMap[0].(map[string]interface{}))
		}
	}
	//
	return &models.ProjectReport{
		EntpID:               entpID,
		Error:                error,
		Next:                 next,
		ProjectSummaryReport: projectSummaryReport,
	}
}

func SetProjectReportResourceData(d *schema.ResourceData, m *models.ProjectReport) {
	d.Set("entp_id", m.EntpID)
	d.Set("error", m.Error)
	d.Set("next", SetCursorSubResourceData([]*models.Cursor{m.Next}))
	d.Set("project_summary_report", SetProjectSummaryReportSubResourceData([]*models.ProjectSummaryReport{m.ProjectSummaryReport}))
}

func SetProjectReportSubResourceData(m []*models.ProjectReport) (d []*map[string]interface{}) {
	for _, ProjectReportModel := range m {
		if ProjectReportModel != nil {
			properties := make(map[string]interface{})
			properties["entp_id"] = ProjectReportModel.EntpID
			properties["error"] = ProjectReportModel.Error
			properties["next"] = SetCursorSubResourceData([]*models.Cursor{ProjectReportModel.Next})
			properties["project_summary_report"] = SetProjectSummaryReportSubResourceData([]*models.ProjectSummaryReport{ProjectReportModel.ProjectSummaryReport})
			d = append(d, &properties)
		}
	}
	return
}

func ProjectReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"entp_id": {
			Description: `Enterprise id for which we want to get summary report for all objects`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"error": {
			Description: `Error while fetching report for enterprise, if any`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"next": {
			Description: `cursor next`,
			Type:        schema.TypeList, //GoType: Cursor
			Elem: &schema.Resource{
				Schema: CursorSchema(),
			},
			Optional: true,
		},

		"project_summary_report": {
			Description: `Enterprise project report`,
			Type:        schema.TypeList, //GoType: ProjectSummaryReport
			Elem: &schema.Resource{
				Schema: ProjectSummaryReportSchema(),
			},
			Optional: true,
		},
	}
}

func GetProjectReportPropertyFields() (t []string) {
	return []string{
		"entp_id",
		"error",
		"next",
		"project_summary_report",
	}
}
