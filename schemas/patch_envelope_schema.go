package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func PatchEnvelopeModel(d *schema.ResourceData) *models.PatchEnvelope {
	var action *models.PatchEnvelopeAction // PatchEnvelopeAction
	actionInterface, actionIsSet := d.GetOk("action")
	if actionIsSet {
		actionModel := actionInterface.(string)
		action = models.NewPatchEnvelopeAction(models.PatchEnvelopeAction(actionModel))
	}
	var artifacts []*models.BinaryArtifact // []*BinaryArtifact
	artifactsInterface, artifactsIsSet := d.GetOk("artifacts")
	if artifactsIsSet {
		var items []interface{}
		if listItems, isList := artifactsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = artifactsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := BinaryArtifactModelFromMap(v.(map[string]interface{}))
			artifacts = append(artifacts, m)
		}
	}
	description, _ := d.Get("description").(string)
	deviceCountInt, _ := d.Get("device_count").(int)
	deviceCount := int64(deviceCountInt)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	projectName, _ := d.Get("project_name").(string)
	title, _ := d.Get("title").(string)
	userDefinedVersion, _ := d.Get("user_defined_version").(string)
	return &models.PatchEnvelope{
		Action:             action,
		Artifacts:          artifacts,
		Description:        description,
		DeviceCount:        deviceCount,
		ID:                 id,
		Name:               &name, // string
		ProjectID:          projectID,
		ProjectName:        projectName,
		Title:              &title, // string
		UserDefinedVersion: userDefinedVersion,
	}
}

func PatchEnvelopeModelFromMap(m map[string]interface{}) *models.PatchEnvelope {
	var action *models.PatchEnvelopeAction // PatchEnvelopeAction
	actionInterface, actionIsSet := m["action"]
	if actionIsSet {
		actionModel := actionInterface.(string)
		action = models.NewPatchEnvelopeAction(models.PatchEnvelopeAction(actionModel))
	}
	var artifacts []*models.BinaryArtifact // []*BinaryArtifact
	artifactsInterface, artifactsIsSet := m["artifacts"]
	if artifactsIsSet {
		var items []interface{}
		if listItems, isList := artifactsInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = artifactsInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			m := BinaryArtifactModelFromMap(v.(map[string]interface{}))
			artifacts = append(artifacts, m)
		}
	}
	description := m["description"].(string)
	deviceCount := int64(m["device_count"].(int)) // int64
	id := m["id"].(string)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	projectName := m["project_name"].(string)
	title := m["title"].(string)
	userDefinedVersion := m["user_defined_version"].(string)
	return &models.PatchEnvelope{
		Action:             action,
		Artifacts:          artifacts,
		Description:        description,
		DeviceCount:        deviceCount,
		ID:                 id,
		Name:               &name,
		ProjectID:          projectID,
		ProjectName:        projectName,
		Title:              &title,
		UserDefinedVersion: userDefinedVersion,
	}
}

func SetPatchEnvelopeResourceData(d *schema.ResourceData, m *models.PatchEnvelope) {
	d.Set("action", m.Action)
	d.Set("artifacts", SetBinaryArtifactSubResourceData(m.Artifacts))
	d.Set("description", m.Description)
	d.Set("device_count", m.DeviceCount)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("project_name", m.ProjectName)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("title", m.Title)
	d.Set("user_defined_version", m.UserDefinedVersion)
}

func SetPatchEnvelopeSubResourceData(m []*models.PatchEnvelope) (d []*map[string]interface{}) {
	for _, PatchEnvelopeModel := range m {
		if PatchEnvelopeModel != nil {
			properties := make(map[string]interface{})
			properties["action"] = PatchEnvelopeModel.Action
			properties["artifacts"] = SetBinaryArtifactSubResourceData(PatchEnvelopeModel.Artifacts)
			properties["description"] = PatchEnvelopeModel.Description
			properties["device_count"] = PatchEnvelopeModel.DeviceCount
			properties["id"] = PatchEnvelopeModel.ID
			properties["name"] = PatchEnvelopeModel.Name
			properties["project_id"] = PatchEnvelopeModel.ProjectID
			properties["project_name"] = PatchEnvelopeModel.ProjectName
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{PatchEnvelopeModel.Revision})
			properties["title"] = PatchEnvelopeModel.Title
			properties["user_defined_version"] = PatchEnvelopeModel.UserDefinedVersion
			d = append(d, &properties)
		}
	}
	return
}

func PatchEnvelopeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action": {
			Description: `Flag to represent whether device needs to present it to app instance`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"artifacts": {
			Description: `Patch envelope artifacts`,
			Type:        schema.TypeList, //GoType: []*BinaryArtifact
			Elem: &schema.Resource{
				Schema: BinaryArtifactSchema(),
			},
			// ConfigMode: schema.SchemaConfigModeAttr,
			Required: true,
		},

		"description": {
			Description: `Detailed description of the patch envelope.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"device_count": {
			Description: `number of devices referencing this patch envelope`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the patch envelope.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the patch envelope, unique across the enterprise. Once patch envelope is created, name can’t be changed.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project_id": {
			Description: `project id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"project_name": {
			Description: `project name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"revision": {
			Description: `system defined info`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Computed: true,
		},

		"title": {
			Description: `User defined title of the patch envelope. Title can be changed at any time.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"user_defined_version": {
			Description: `User defined version for the given patch envelope`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetPatchEnvelopePropertyFields() (t []string) {
	return []string{
		"action",
		"artifacts",
		"description",
		"device_count",
		"id",
		"name",
		"project_id",
		"project_name",
		"title",
		"user_defined_version",
	}
}
