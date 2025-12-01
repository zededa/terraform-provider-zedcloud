package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/go-openapi/strfmt"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoGetResponseModel(d *schema.ResourceData) *models.GitRepoGetResponse {
	createdAt, _ := d.Get("created_at").(strfmt.DateTime)
	createdBy, _ := d.Get("created_by").(string)
	var data *models.GitRepoData // GitRepoData
	dataInterface, dataIsSet := d.GetOk("data")
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = GitRepoDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	description, _ := d.Get("description").(string)
	var events []*models.GitRepoEvents // []*GitRepoEvents
	eventsInterface, eventsIsSet := d.GetOk("events")
	if eventsIsSet {
		var items []interface{}
		if listItems, isList := eventsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = eventsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := GitRepoEventsModelFromMap(v.(map[string]interface{}))
			events = append(events, m)
		}
	}
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	var status *models.GitRepoGetStatus // GitRepoGetStatus
	statusInterface, statusIsSet := d.GetOk("status")
	if statusIsSet && statusInterface != nil {
		statusMap := statusInterface.([]interface{})
		if len(statusMap) > 0 {
			status = GitRepoGetStatusModelFromMap(statusMap[0].(map[string]interface{}))
		}
	}
	title, _ := d.Get("title").(string)
	updatedAt, _ := d.Get("updated_at").(strfmt.DateTime)
	updatedBy, _ := d.Get("updated_by").(string)
	workspace, _ := d.Get("workspace").(string)
	return &models.GitRepoGetResponse{
		CreatedAt:   createdAt,
		CreatedBy:   createdBy,
		Data:        data,
		Description: description,
		Events:      events,
		ID:          id,
		Name:        name,
		Status:      status,
		Title:       title,
		UpdatedAt:   updatedAt,
		UpdatedBy:   updatedBy,
		Workspace:   workspace,
	}
}

func GitRepoGetResponseModelFromMap(m map[string]interface{}) *models.GitRepoGetResponse {
	createdAt := m["created_at"].(strfmt.DateTime)
	createdBy := m["created_by"].(string)
	var data *models.GitRepoData // GitRepoData
	dataInterface, dataIsSet := m["data"]
	if dataIsSet && dataInterface != nil {
		dataMap := dataInterface.([]interface{})
		if len(dataMap) > 0 {
			data = GitRepoDataModelFromMap(dataMap[0].(map[string]interface{}))
		}
	}
	//
	description := m["description"].(string)
	var events []*models.GitRepoEvents // []*GitRepoEvents
	eventsInterface, eventsIsSet := m["events"]
	if eventsIsSet {
		var items []interface{}
		if listItems, isList := eventsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = eventsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := GitRepoEventsModelFromMap(v.(map[string]interface{}))
			events = append(events, m)
		}
	}
	id := m["id"].(string)
	name := m["name"].(string)
	var status *models.GitRepoGetStatus // GitRepoGetStatus
	statusInterface, statusIsSet := m["status"]
	if statusIsSet && statusInterface != nil {
		statusMap := statusInterface.([]interface{})
		if len(statusMap) > 0 {
			status = GitRepoGetStatusModelFromMap(statusMap[0].(map[string]interface{}))
		}
	}
	//
	title := m["title"].(string)
	updatedAt := m["updated_at"].(strfmt.DateTime)
	updatedBy := m["updated_by"].(string)
	workspace := m["workspace"].(string)
	return &models.GitRepoGetResponse{
		CreatedAt:   createdAt,
		CreatedBy:   createdBy,
		Data:        data,
		Description: description,
		Events:      events,
		ID:          id,
		Name:        name,
		Status:      status,
		Title:       title,
		UpdatedAt:   updatedAt,
		UpdatedBy:   updatedBy,
		Workspace:   workspace,
	}
}

func SetGitRepoGetResponseResourceData(d *schema.ResourceData, m *models.GitRepoGetResponse) {
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by", m.CreatedBy)
	d.Set("data", SetGitRepoDataSubResourceData([]*models.GitRepoData{m.Data}))
	d.Set("description", m.Description)
	d.Set("events", SetGitRepoEventsSubResourceData(m.Events))
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("status", SetGitRepoGetStatusSubResourceData([]*models.GitRepoGetStatus{m.Status}))
	d.Set("title", m.Title)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by", m.UpdatedBy)
	d.Set("workspace", m.Workspace)
}

func SetGitRepoGetResponseSubResourceData(m []*models.GitRepoGetResponse) (d []*map[string]interface{}) {
	for _, GitRepoGetResponseModel := range m {
		if GitRepoGetResponseModel != nil {
			properties := make(map[string]interface{})
			properties["created_at"] = GitRepoGetResponseModel.CreatedAt.String()
			properties["created_by"] = GitRepoGetResponseModel.CreatedBy
			properties["data"] = SetGitRepoDataSubResourceData([]*models.GitRepoData{GitRepoGetResponseModel.Data})
			properties["description"] = GitRepoGetResponseModel.Description
			properties["events"] = SetGitRepoEventsSubResourceData(GitRepoGetResponseModel.Events)
			properties["id"] = GitRepoGetResponseModel.ID
			properties["name"] = GitRepoGetResponseModel.Name
			properties["status"] = SetGitRepoGetStatusSubResourceData([]*models.GitRepoGetStatus{GitRepoGetResponseModel.Status})
			properties["title"] = GitRepoGetResponseModel.Title
			properties["updated_at"] = GitRepoGetResponseModel.UpdatedAt.String()
			properties["updated_by"] = GitRepoGetResponseModel.UpdatedBy
			properties["workspace"] = GitRepoGetResponseModel.Workspace
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoGetResponseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Description:  `Timestamp when the GitRepo was created`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"created_by": {
			Description: `User who created the GitRepo`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"data": {
			Description: `Specification details`,
			Type:        schema.TypeList, //GoType: GitRepoData
			Elem: &schema.Resource{
				Schema: GitRepoDataSchema(),
			},
			Optional: true,
		},

		"description": {
			Description: `Description of the GitRepo configuration`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"events": {
			Description: `List of events related to the GitRepo`,
			Type:        schema.TypeList, //GoType: []*GitRepoEvents
			Elem: &schema.Resource{
				Schema: GitRepoEventsSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Optional: true,
		},

		"id": {
			Description: `Unique identifier of the GitRepo`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `Name of the GitRepo configuration`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"status": {
			Description: `Current status information`,
			Type:        schema.TypeList, //GoType: GitRepoGetStatus
			Elem: &schema.Resource{
				Schema: GitRepoGetStatusSchema(),
			},
			Optional: true,
		},

		"title": {
			Description: `Display title of the GitRepo configuration`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"updated_at": {
			Description:  `Timestamp when the GitRepo was last updated`,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"updated_by": {
			Description: `User who last updated the GitRepo`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"workspace": {
			Description: `Workspace where the GitRepo is deployed`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetGitRepoGetResponsePropertyFields() (t []string) {
	return []string{
		"created_at",
		"created_by",
		"data",
		"description",
		"events",
		"id",
		"name",
		"status",
		"title",
		"updated_at",
		"updated_by",
		"workspace",
	}
}
