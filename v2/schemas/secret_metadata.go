package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SecretMetadataModel(d *schema.ResourceData) *models.SecretMetadata {
	creationTimestamp, _ := d.Get("creation_timestamp").(string)
	name, _ := d.Get("name").(string)
	projectID, _ := d.Get("project_id").(string)
	var state *models.SecretState // SecretState
	stateInterface, stateIsSet := d.GetOk("state")
	if stateIsSet && stateInterface != nil {
		stateMap := stateInterface.([]interface{})
		if len(stateMap) > 0 {
			state = SecretStateModelFromMap(stateMap[0].(map[string]interface{}))
		}
	}
	var typeVar *models.SecretType // SecretType
	typeInterface, typeIsSet := d.GetOk("type")
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewSecretType(models.SecretType(typeModel))
	}
	return &models.SecretMetadata{
		CreationTimestamp: creationTimestamp,
		Name:              &name, // string
		ProjectID:         projectID,
		State:             state,
		Type:              typeVar,
	}
}

func SecretMetadataModelFromMap(m map[string]interface{}) *models.SecretMetadata {
	creationTimestamp := m["creation_timestamp"].(string)
	name := m["name"].(string)
	projectID := m["project_id"].(string)
	var state *models.SecretState // SecretState
	stateInterface, stateIsSet := m["state"]
	if stateIsSet && stateInterface != nil {
		stateMap := stateInterface.([]interface{})
		if len(stateMap) > 0 {
			state = SecretStateModelFromMap(stateMap[0].(map[string]interface{}))
		}
	}
	//
	var typeVar *models.SecretType // SecretType
	typeInterface, typeIsSet := m["type"]
	if typeIsSet {
		typeModel := typeInterface.(string)
		typeVar = models.NewSecretType(models.SecretType(typeModel))
	}
	return &models.SecretMetadata{
		CreationTimestamp: creationTimestamp,
		Name:              &name,
		ProjectID:         projectID,
		State:             state,
		Type:              typeVar,
	}
}

func SetSecretMetadataResourceData(d *schema.ResourceData, m *models.SecretMetadata) {
	d.Set("creation_timestamp", m.CreationTimestamp)
	d.Set("name", m.Name)
	d.Set("project_id", m.ProjectID)
	d.Set("state", SetSecretStateSubResourceData([]*models.SecretState{m.State}))
	d.Set("type", m.Type)
}

func SetSecretMetadataSubResourceData(m []*models.SecretMetadata) (d []*map[string]interface{}) {
	for _, SecretMetadataModel := range m {
		if SecretMetadataModel != nil {
			properties := make(map[string]interface{})
			properties["creation_timestamp"] = SecretMetadataModel.CreationTimestamp
			properties["name"] = SecretMetadataModel.Name
			properties["project_id"] = SecretMetadataModel.ProjectID
			properties["state"] = SetSecretStateSubResourceData([]*models.SecretState{SecretMetadataModel.State})
			properties["type"] = SecretMetadataModel.Type
			d = append(d, &properties)
		}
	}
	return
}

func SecretMetadataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"creation_timestamp": {
			Description: `Timestamp when the secret was created`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `Name of the Kubernetes secret. This field is required and must be a valid Kubernetes secret name.`,
			Type:        schema.TypeString,
			Required:    true,
		},

		"project_id": {
			Description: `Project identifier that owns this secret`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"state": {
			Description: `Current state of the secret including error status and transition information`,
			Type:        schema.TypeList, //GoType: SecretState
			Elem: &schema.Resource{
				Schema: SecretStateSchema(),
			},
			Computed: true,
		},

		"type": {
			Description: `Type of the secret (SSH, Basic Auth, etc.)`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetSecretMetadataPropertyFields() (t []string) {
	return []string{
		"creation_timestamp",
		"name",
		"project_id",
		"state",
		"type",
	}
}
