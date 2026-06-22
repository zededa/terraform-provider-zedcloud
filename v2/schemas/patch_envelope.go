package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
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
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
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
		Revision:           revision,
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
	// The backend does not preserve artifact ordering, so reorder the response
	// to match the order already in state before writing it. artifacts is a
	// positional TypeList; without this a reordered response shows up as a
	// perpetual (and flaky) plan diff.
	d.Set("artifacts", SetBinaryArtifactSubResourceData(reorderBinaryArtifactsToMatchState(d, m.Artifacts)))
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

// reorderBinaryArtifactsToMatchState returns apiArtifacts reordered to match the
// order of the "artifacts" list currently held in d, matching elements by a
// content signature (artifact type + file name) that is present in both the
// configured/state value and the API response. If the lists cannot be matched
// 1:1 (e.g. different length, or an identifying field changed), the API order
// is returned unchanged so no data is dropped or duplicated.
func reorderBinaryArtifactsToMatchState(d *schema.ResourceData, apiArtifacts []*models.BinaryArtifact) []*models.BinaryArtifact {
	raw, ok := d.GetOk("artifacts")
	if !ok {
		return apiArtifacts
	}
	var stateList []interface{}
	switch v := raw.(type) {
	case []interface{}:
		stateList = v
	case *schema.Set:
		stateList = v.List()
	default:
		return apiArtifacts
	}
	if len(stateList) != len(apiArtifacts) {
		return apiArtifacts
	}

	used := make([]bool, len(apiArtifacts))
	ordered := make([]*models.BinaryArtifact, 0, len(apiArtifacts))
	for _, s := range stateList {
		sm, ok := s.(map[string]interface{})
		if !ok {
			return apiArtifacts
		}
		key := binaryArtifactStateKey(sm)
		matched := -1
		for i, a := range apiArtifacts {
			if !used[i] && binaryArtifactModelKey(a) == key {
				matched = i
				break
			}
		}
		if matched < 0 {
			return apiArtifacts
		}
		ordered = append(ordered, apiArtifacts[matched])
		used[matched] = true
	}
	return ordered
}

// binaryArtifactModelKey builds the match signature for an API artifact.
func binaryArtifactModelKey(a *models.BinaryArtifact) string {
	switch {
	case a == nil:
		return ""
	case a.Base64Artifact != nil:
		return "inline|" + a.Base64Artifact.FileNameToUse
	case a.BinaryArtifact != nil:
		return "binary|" + a.BinaryArtifact.FileNameToUse
	}
	return "unknown"
}

// binaryArtifactStateKey builds the match signature for an artifact held in state.
func binaryArtifactStateKey(m map[string]interface{}) string {
	if name, ok := nestedArtifactFileName(m, "base64_artifact"); ok {
		return "inline|" + name
	}
	if name, ok := nestedArtifactFileName(m, "binary_artifact"); ok {
		return "binary|" + name
	}
	return "unknown"
}

// nestedArtifactFileName extracts file_name_to_use from a nested base64_artifact
// or binary_artifact block, reporting whether that block is present.
func nestedArtifactFileName(m map[string]interface{}, key string) (string, bool) {
	var list []interface{}
	switch v := m[key].(type) {
	case []interface{}:
		list = v
	case *schema.Set:
		list = v.List()
	default:
		return "", false
	}
	if len(list) == 0 {
		return "", false
	}
	inner, ok := list[0].(map[string]interface{})
	if !ok {
		return "", false
	}
	name, _ := inner["file_name_to_use"].(string)
	return name, true
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
			Required:    true,
		},

		"project_name": {
			Description: `project name`,
			Type:        schema.TypeString,
			Required:    true,
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
