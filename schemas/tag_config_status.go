package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate TagConfigStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TagConfigStatusModel(d *schema.ResourceData) *models.TagConfigStatus {
	appPolicyID, _ := d.Get("app_policy_id").(string)
	appPolicyName, _ := d.Get("app_policy_name").(string)
	var attestPolicyType *models.AttestPolicyType // AttestPolicyType
	attestPolicyTypeInterface, attestPolicyTypeIsSet := d.GetOk("attest_policy_type")
	if attestPolicyTypeIsSet {
		attestPolicyTypeModel := attestPolicyTypeInterface.(string)
		attestPolicyType = models.NewAttestPolicyType(models.AttestPolicyType(attestPolicyTypeModel))
	}
	var cloudPolicyType *models.PolicyType // PolicyType
	cloudPolicyTypeInterface, cloudPolicyTypeIsSet := d.GetOk("cloud_policy_type")
	if cloudPolicyTypeIsSet {
		cloudPolicyTypeModel := cloudPolicyTypeInterface.(string)
		cloudPolicyType = models.NewPolicyType(models.PolicyType(cloudPolicyTypeModel))
	}
	edgeviewAllow, _ := d.Get("edgeview_allow").(bool)
	edgeviewSessionCountInt, _ := d.Get("edgeview_session_count").(int)
	edgeviewSessionCount := int64(edgeviewSessionCountInt)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	networkDeviceDefault, _ := d.Get("network_device_default").(string)
	networkPolicyID, _ := d.Get("network_policy_id").(string)
	return &models.TagConfigStatus{
		AppPolicyID:          appPolicyID,
		AppPolicyName:        appPolicyName,
		AttestPolicyType:     attestPolicyType,
		CloudPolicyType:      cloudPolicyType,
		EdgeviewAllow:        edgeviewAllow,
		EdgeviewSessionCount: edgeviewSessionCount,
		ID:                   id,
		Name:                 name,
		NetworkDeviceDefault: networkDeviceDefault,
		NetworkPolicyID:      networkPolicyID,
	}
}

func TagConfigStatusModelFromMap(m map[string]interface{}) *models.TagConfigStatus {
	appPolicyID := m["app_policy_id"].(string)
	appPolicyName := m["app_policy_name"].(string)
	attestPolicyType := m["attest_policy_type"].(*models.AttestPolicyType) // AttestPolicyType
	cloudPolicyType := m["cloud_policy_type"].(*models.PolicyType)         // PolicyType
	edgeviewAllow := m["edgeview_allow"].(bool)
	edgeviewSessionCount := int64(m["edgeview_session_count"].(int)) // int64 false false false
	id := m["id"].(string)
	name := m["name"].(string)
	networkDeviceDefault := m["network_device_default"].(string)
	networkPolicyID := m["network_policy_id"].(string)
	return &models.TagConfigStatus{
		AppPolicyID:          appPolicyID,
		AppPolicyName:        appPolicyName,
		AttestPolicyType:     attestPolicyType,
		CloudPolicyType:      cloudPolicyType,
		EdgeviewAllow:        edgeviewAllow,
		EdgeviewSessionCount: edgeviewSessionCount,
		ID:                   id,
		Name:                 name,
		NetworkDeviceDefault: networkDeviceDefault,
		NetworkPolicyID:      networkPolicyID,
	}
}

// Update the underlying TagConfigStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTagConfigStatusResourceData(d *schema.ResourceData, m *models.TagConfigStatus) {
	d.Set("app_policy_id", m.AppPolicyID)
	d.Set("app_policy_name", m.AppPolicyName)
	d.Set("attest_policy_type", m.AttestPolicyType)
	d.Set("cloud_policy_type", m.CloudPolicyType)
	d.Set("edgeview_allow", m.EdgeviewAllow)
	d.Set("edgeview_session_count", m.EdgeviewSessionCount)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("network_device_default", m.NetworkDeviceDefault)
	d.Set("network_policy_id", m.NetworkPolicyID)
	d.Set("status", m.Status)
}

// Iterate through and update the TagConfigStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTagConfigStatusSubResourceData(m []*models.TagConfigStatus) (d []*map[string]interface{}) {
	for _, TagConfigStatusModel := range m {
		if TagConfigStatusModel != nil {
			properties := make(map[string]interface{})
			properties["app_policy_id"] = TagConfigStatusModel.AppPolicyID
			properties["app_policy_name"] = TagConfigStatusModel.AppPolicyName
			properties["attest_policy_type"] = TagConfigStatusModel.AttestPolicyType
			properties["cloud_policy_type"] = TagConfigStatusModel.CloudPolicyType
			properties["edgeview_allow"] = TagConfigStatusModel.EdgeviewAllow
			properties["edgeview_session_count"] = TagConfigStatusModel.EdgeviewSessionCount
			properties["id"] = TagConfigStatusModel.ID
			properties["name"] = TagConfigStatusModel.Name
			properties["network_device_default"] = TagConfigStatusModel.NetworkDeviceDefault
			properties["network_policy_id"] = TagConfigStatusModel.NetworkPolicyID
			properties["status"] = TagConfigStatusModel.Status
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the TagConfigStatus resource defined in the Terraform configuration
func TagConfigStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_policy_id": {
			Description: `app policy Id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"app_policy_name": {
			Description: `app policy name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"attest_policy_type": {
			Description: `type of attestation policy`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"cloud_policy_type": {
			Description: `type of cloud policy`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"edgeview_allow": {
			Description: `edgeview is allowed or not for devices in project`,
			Type:        schema.TypeBool,
			Optional:    true,
		},

		"edgeview_session_count": {
			Description: `total count of devices enabled with edgeview session`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"id": {
			Description: `System defined universally unique Id of the resource group.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the resource group, unique across the enterprise. Once resource group is created, name can’t be changed.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"network_device_default": {
			Description: `flag to indicate if this is the default network instance for the device`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"network_policy_id": {
			Description: `network policy Id`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"status": {
			Description: `Resource group status`,
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

// Retrieve property field names for updating the TagConfigStatus resource
func GetTagConfigStatusPropertyFields() (t []string) {
	return []string{
		"app_policy_id",
		"app_policy_name",
		"attest_policy_type",
		"cloud_policy_type",
		"edgeview_allow",
		"edgeview_session_count",
		"id",
		"name",
		"network_device_default",
		"network_policy_id",
	}
}
