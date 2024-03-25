package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func ScopeModel(d *schema.ResourceData) *models.Scope {
	accessApp, _ := d.Get("access_app").(string)
	accessAppInstance, _ := d.Get("access_app_instance").(string)
	accessDevice, _ := d.Get("access_device").(string)
	accessEdgeApp, _ := d.Get("access_edge_app").(string)
	accessEnterprise, _ := d.Get("access_enterprise").(string)
	accessStorage, _ := d.Get("access_storage").(string)
	accessUser, _ := d.Get("access_user").(string)
	var enterpriseFilter []string
	enterpriseFilterInterface, enterpriseFilterIsSet := d.GetOk("enterpriseFilter")
	if enterpriseFilterIsSet {
		var items []interface{}
		if listItems, isList := enterpriseFilterInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = enterpriseFilterInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			enterpriseFilter = append(enterpriseFilter, v.(string))
		}
	}
	var projectFilter []string
	projectFilterInterface, projectFilterIsSet := d.GetOk("projectFilter")
	if projectFilterIsSet {
		var items []interface{}
		if listItems, isList := projectFilterInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = projectFilterInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			projectFilter = append(projectFilter, v.(string))
		}
	}
	return &models.Scope{
		AccessApp:         accessApp,
		AccessAppInstance: accessAppInstance,
		AccessDevice:      accessDevice,
		AccessEdgeApp:     accessEdgeApp,
		AccessEnterprise:  accessEnterprise,
		AccessStorage:     accessStorage,
		AccessUser:        accessUser,
		EnterpriseFilter:  enterpriseFilter,
		ProjectFilter:     projectFilter,
	}
}

func ScopeModelFromMap(m map[string]interface{}) *models.Scope {
	accessApp := m["access_app"].(string)
	accessAppInstance := m["access_app_instance"].(string)
	accessDevice := m["access_device"].(string)
	accessEdgeApp := m["access_edge_app"].(string)
	accessEnterprise := m["access_enterprise"].(string)
	accessStorage := m["access_storage"].(string)
	accessUser := m["access_user"].(string)
	var enterpriseFilter []string
	enterpriseFilterInterface, enterpriseFilterIsSet := m["enterpriseFilter"]
	if enterpriseFilterIsSet {
		var items []interface{}
		if listItems, isList := enterpriseFilterInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = enterpriseFilterInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			enterpriseFilter = append(enterpriseFilter, v.(string))
		}
	}
	var projectFilter []string
	projectFilterInterface, projectFilterIsSet := m["projectFilter"]
	if projectFilterIsSet {
		var items []interface{}
		if listItems, isList := projectFilterInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = projectFilterInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			projectFilter = append(projectFilter, v.(string))
		}
	}
	return &models.Scope{
		AccessApp:         accessApp,
		AccessAppInstance: accessAppInstance,
		AccessDevice:      accessDevice,
		AccessEdgeApp:     accessEdgeApp,
		AccessEnterprise:  accessEnterprise,
		AccessStorage:     accessStorage,
		AccessUser:        accessUser,
		EnterpriseFilter:  enterpriseFilter,
		ProjectFilter:     projectFilter,
	}
}

func SetScopeResourceData(d *schema.ResourceData, m *models.Scope) {
	d.Set("access_app", m.AccessApp)
	d.Set("access_app_instance", m.AccessAppInstance)
	d.Set("access_device", m.AccessDevice)
	d.Set("access_edge_app", m.AccessEdgeApp)
	d.Set("access_enterprise", m.AccessEnterprise)
	d.Set("access_storage", m.AccessStorage)
	d.Set("access_user", m.AccessUser)
	d.Set("enterprise_filter", m.EnterpriseFilter)
	d.Set("project_filter", m.ProjectFilter)
	d.Set("uuid", m.UUID)
}

func SetScopeSubResourceData(m []*models.Scope) (d []*map[string]interface{}) {
	for _, ScopeModel := range m {
		if ScopeModel != nil {
			properties := make(map[string]interface{})
			properties["access_app"] = ScopeModel.AccessApp
			properties["access_app_instance"] = ScopeModel.AccessAppInstance
			properties["access_device"] = ScopeModel.AccessDevice
			properties["access_edge_app"] = ScopeModel.AccessEdgeApp
			properties["access_enterprise"] = ScopeModel.AccessEnterprise
			properties["access_storage"] = ScopeModel.AccessStorage
			properties["access_user"] = ScopeModel.AccessUser
			properties["enterprise_filter"] = ScopeModel.EnterpriseFilter
			properties["project_filter"] = ScopeModel.ProjectFilter
			properties["uuid"] = ScopeModel.UUID
			d = append(d, &properties)
		}
	}
	return
}

func ScopeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access_app": {
			Description: `Application access permission, this will be deprecated in further release - please use accessEdgeApp and accessAppInstance for granular permission access scope.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"access_app_instance": {
			Description: `Application Instance access permission`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"access_device": {
			Description: `Device access permission`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"access_edge_app": {
			Description: `Edge app access permission`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"access_enterprise": {
			Description: `Enterprise access permission`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"access_storage": {
			Description: `Storage access permission`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"access_user": {
			Description: `User access permission`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"enterprise_filter": {
			Description: `List of enterprise filters`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"project_filter": {
			Description: `List of project filters`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"uuid": {
			Description: `Unique system defined scope ID`,
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}

func GetScopePropertyFields() (t []string) {
	return []string{
		"access_app",
		"access_app_instance",
		"access_device",
		"access_edge_app",
		"access_enterprise",
		"access_storage",
		"access_user",
		"enterprise_filter",
		"project_filter",
	}
}
