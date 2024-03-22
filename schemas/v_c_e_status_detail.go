package schemas

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

func VCEStatusDetailModel(d *schema.ResourceData) *models.VCEStatusDetail {
	activationKey, _ := d.Get("activation_key").(string)
	activationKeyExpires, _ := d.Get("activation_key_expires").(strfmt.DateTime)
	activationState, _ := d.Get("activation_state").(string)
	activationTime, _ := d.Get("activation_time").(strfmt.DateTime)
	alertsEnabled, _ := d.Get("alerts_enabled").(bool)
	bastionState, _ := d.Get("bastion_state").(string)
	created, _ := d.Get("created").(strfmt.DateTime)
	edgeState, _ := d.Get("edge_state").(string)
	edgeStateTime, _ := d.Get("edge_state_time").(strfmt.DateTime)
	haLastContact, _ := d.Get("ha_last_contact").(strfmt.DateTime)
	haPreviousState, _ := d.Get("ha_previous_state").(string)
	haSerialNumber, _ := d.Get("ha_serial_number").(string)
	haState, _ := d.Get("ha_state").(string)
	isHub, _ := d.Get("is_hub").(bool)
	isLive, _ := d.Get("is_live").(bool)
	lastContact, _ := d.Get("last_contact").(strfmt.DateTime)
	lteRegion, _ := d.Get("lte_region").(string)
	modified, _ := d.Get("modified").(strfmt.DateTime)
	operatorAlertsEnabled, _ := d.Get("operator_alerts_enabled").(bool)
	serialNumber, _ := d.Get("serial_number").(string)
	serviceState, _ := d.Get("service_state").(string)
	serviceUpSince, _ := d.Get("service_up_since").(strfmt.DateTime)
	statusFetchedAt, _ := d.Get("status_fetched_at").(strfmt.DateTime)
	systemUpSince, _ := d.Get("system_up_since").(strfmt.DateTime)
	vceEdgeIDInt, _ := d.Get("vce_edge_id").(int)
	vceEdgeID := int64(vceEdgeIDInt)
	vceURL, _ := d.Get("vce_url").(string)
	return &models.VCEStatusDetail{
		ActivationKey:         activationKey,
		ActivationKeyExpires:  activationKeyExpires,
		ActivationState:       activationState,
		ActivationTime:        activationTime,
		AlertsEnabled:         alertsEnabled,
		BastionState:          bastionState,
		Created:               created,
		EdgeState:             edgeState,
		EdgeStateTime:         edgeStateTime,
		HaLastContact:         haLastContact,
		HaPreviousState:       haPreviousState,
		HaSerialNumber:        haSerialNumber,
		HaState:               haState,
		IsHub:                 isHub,
		IsLive:                isLive,
		LastContact:           lastContact,
		LteRegion:             lteRegion,
		Modified:              modified,
		OperatorAlertsEnabled: operatorAlertsEnabled,
		SerialNumber:          serialNumber,
		ServiceState:          serviceState,
		ServiceUpSince:        serviceUpSince,
		StatusFetchedAt:       statusFetchedAt,
		SystemUpSince:         systemUpSince,
		VceEdgeID:             vceEdgeID,
		VceURL:                vceURL,
	}
}

func VCEStatusDetailModelFromMap(m map[string]interface{}) *models.VCEStatusDetail {
	activationKey := m["activation_key"].(string)
	activationKeyExpires := m["activation_key_expires"].(strfmt.DateTime)
	activationState := m["activation_state"].(string)
	activationTime := m["activation_time"].(strfmt.DateTime)
	alertsEnabled := m["alerts_enabled"].(bool)
	bastionState := m["bastion_state"].(string)
	created := m["created"].(strfmt.DateTime)
	edgeState := m["edge_state"].(string)
	edgeStateTime := m["edge_state_time"].(strfmt.DateTime)
	haLastContact := m["ha_last_contact"].(strfmt.DateTime)
	haPreviousState := m["ha_previous_state"].(string)
	haSerialNumber := m["ha_serial_number"].(string)
	haState := m["ha_state"].(string)
	isHub := m["is_hub"].(bool)
	isLive := m["is_live"].(bool)
	lastContact := m["last_contact"].(strfmt.DateTime)
	lteRegion := m["lte_region"].(string)
	modified := m["modified"].(strfmt.DateTime)
	operatorAlertsEnabled := m["operator_alerts_enabled"].(bool)
	serialNumber := m["serial_number"].(string)
	serviceState := m["service_state"].(string)
	serviceUpSince := m["service_up_since"].(strfmt.DateTime)
	statusFetchedAt := m["status_fetched_at"].(strfmt.DateTime)
	systemUpSince := m["system_up_since"].(strfmt.DateTime)
	vceEdgeID := int64(m["vce_edge_id"].(int)) // int64
	vceURL := m["vce_url"].(string)
	return &models.VCEStatusDetail{
		ActivationKey:         activationKey,
		ActivationKeyExpires:  activationKeyExpires,
		ActivationState:       activationState,
		ActivationTime:        activationTime,
		AlertsEnabled:         alertsEnabled,
		BastionState:          bastionState,
		Created:               created,
		EdgeState:             edgeState,
		EdgeStateTime:         edgeStateTime,
		HaLastContact:         haLastContact,
		HaPreviousState:       haPreviousState,
		HaSerialNumber:        haSerialNumber,
		HaState:               haState,
		IsHub:                 isHub,
		IsLive:                isLive,
		LastContact:           lastContact,
		LteRegion:             lteRegion,
		Modified:              modified,
		OperatorAlertsEnabled: operatorAlertsEnabled,
		SerialNumber:          serialNumber,
		ServiceState:          serviceState,
		ServiceUpSince:        serviceUpSince,
		StatusFetchedAt:       statusFetchedAt,
		SystemUpSince:         systemUpSince,
		VceEdgeID:             vceEdgeID,
		VceURL:                vceURL,
	}
}

// Update the underlying VCEStatusDetail resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVCEStatusDetailResourceData(d *schema.ResourceData, m *models.VCEStatusDetail) {
	d.Set("activation_key", m.ActivationKey)
	d.Set("activation_key_expires", m.ActivationKeyExpires)
	d.Set("activation_state", m.ActivationState)
	d.Set("activation_time", m.ActivationTime)
	d.Set("alerts_enabled", m.AlertsEnabled)
	d.Set("bastion_state", m.BastionState)
	d.Set("created", m.Created)
	d.Set("edge_state", m.EdgeState)
	d.Set("edge_state_time", m.EdgeStateTime)
	d.Set("ha_last_contact", m.HaLastContact)
	d.Set("ha_previous_state", m.HaPreviousState)
	d.Set("ha_serial_number", m.HaSerialNumber)
	d.Set("ha_state", m.HaState)
	d.Set("is_hub", m.IsHub)
	d.Set("is_live", m.IsLive)
	d.Set("last_contact", m.LastContact)
	d.Set("lte_region", m.LteRegion)
	d.Set("modified", m.Modified)
	d.Set("operator_alerts_enabled", m.OperatorAlertsEnabled)
	d.Set("serial_number", m.SerialNumber)
	d.Set("service_state", m.ServiceState)
	d.Set("service_up_since", m.ServiceUpSince)
	d.Set("status_fetched_at", m.StatusFetchedAt)
	d.Set("system_up_since", m.SystemUpSince)
	d.Set("vce_edge_id", m.VceEdgeID)
	d.Set("vce_url", m.VceURL)
}

// Iterate through and update the VCEStatusDetail resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVCEStatusDetailSubResourceData(m []*models.VCEStatusDetail) (d []*map[string]interface{}) {
	for _, VCEStatusDetailModel := range m {
		if VCEStatusDetailModel != nil {
			properties := make(map[string]interface{})
			properties["activation_key"] = VCEStatusDetailModel.ActivationKey
			properties["activation_key_expires"] = VCEStatusDetailModel.ActivationKeyExpires.String()
			properties["activation_state"] = VCEStatusDetailModel.ActivationState
			properties["activation_time"] = VCEStatusDetailModel.ActivationTime.String()
			properties["alerts_enabled"] = VCEStatusDetailModel.AlertsEnabled
			properties["bastion_state"] = VCEStatusDetailModel.BastionState
			properties["created"] = VCEStatusDetailModel.Created.String()
			properties["edge_state"] = VCEStatusDetailModel.EdgeState
			properties["edge_state_time"] = VCEStatusDetailModel.EdgeStateTime.String()
			properties["ha_last_contact"] = VCEStatusDetailModel.HaLastContact.String()
			properties["ha_previous_state"] = VCEStatusDetailModel.HaPreviousState
			properties["ha_serial_number"] = VCEStatusDetailModel.HaSerialNumber
			properties["ha_state"] = VCEStatusDetailModel.HaState
			properties["is_hub"] = VCEStatusDetailModel.IsHub
			properties["is_live"] = VCEStatusDetailModel.IsLive
			properties["last_contact"] = VCEStatusDetailModel.LastContact.String()
			properties["lte_region"] = VCEStatusDetailModel.LteRegion
			properties["modified"] = VCEStatusDetailModel.Modified.String()
			properties["operator_alerts_enabled"] = VCEStatusDetailModel.OperatorAlertsEnabled
			properties["serial_number"] = VCEStatusDetailModel.SerialNumber
			properties["service_state"] = VCEStatusDetailModel.ServiceState
			properties["service_up_since"] = VCEStatusDetailModel.ServiceUpSince.String()
			properties["status_fetched_at"] = VCEStatusDetailModel.StatusFetchedAt.String()
			properties["system_up_since"] = VCEStatusDetailModel.SystemUpSince.String()
			properties["vce_edge_id"] = VCEStatusDetailModel.VceEdgeID
			properties["vce_url"] = VCEStatusDetailModel.VceURL
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the VCEStatusDetail resource defined in the Terraform configuration
func VCEStatusDetailSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"activation_key": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"activation_key_expires": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"activation_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"activation_time": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"alerts_enabled": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"bastion_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"created": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"edge_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"edge_state_time": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"ha_last_contact": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"ha_previous_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ha_serial_number": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ha_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"is_hub": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"is_live": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"last_contact": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"lte_region": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"modified": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"operator_alerts_enabled": {
			Description: ``,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"serial_number": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"service_state": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"service_up_since": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"status_fetched_at": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"system_up_since": {
			Description:  ``,
			Type:         schema.TypeString,
			ValidateFunc: validation.IsRFC3339Time,
			Optional:     true,
		},

		"vce_edge_id": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"vce_url": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the VCEStatusDetail resource
func GetVCEStatusDetailPropertyFields() (t []string) {
	return []string{
		"activation_key",
		"activation_key_expires",
		"activation_state",
		"activation_time",
		"alerts_enabled",
		"bastion_state",
		"created",
		"edge_state",
		"edge_state_time",
		"ha_last_contact",
		"ha_previous_state",
		"ha_serial_number",
		"ha_state",
		"is_hub",
		"is_live",
		"last_contact",
		"lte_region",
		"modified",
		"operator_alerts_enabled",
		"serial_number",
		"service_state",
		"service_up_since",
		"status_fetched_at",
		"system_up_since",
		"vce_edge_id",
		"vce_url",
	}
}
