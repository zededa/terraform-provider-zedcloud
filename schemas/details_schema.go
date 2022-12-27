package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Function to perform the following actions:
// (1) Translate Details resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DetailsModel(d *schema.ResourceData) *models.Details {
	agreementList := d.Get("agreement_list").(map[string]string) // map[string]string
	appCategory := d.Get("app_category").(*models.AppCategory)   // AppCategory
	category := d.Get("category").(string)
	licenseList := d.Get("license_list").(map[string]string) // map[string]string
	logo := d.Get("logo").(map[string]string)                // map[string]string
	os := d.Get("os").(string)
	screenshotList := d.Get("screenshot_list").(map[string]string) // map[string]string
	support := d.Get("support").(string)
	return &models.Details{
		AgreementList:  agreementList,
		AppCategory:    appCategory,
		Category:       &category, // string true false false
		LicenseList:    licenseList,
		Logo:           logo,
		Os:             os,
		ScreenshotList: screenshotList,
		Support:        support,
	}
}

func DetailsModelFromMap(m map[string]interface{}) *models.Details {
	agreementList := m["agreement_list"].(map[string]string)
	appCategory := m["app_category"].(*models.AppCategory) // AppCategory
	category := m["category"].(string)
	licenseList := m["license_list"].(map[string]string)
	logo := m["logo"].(map[string]string)
	os := m["os"].(string)
	screenshotList := m["screenshot_list"].(map[string]string)
	support := m["support"].(string)
	return &models.Details{
		AgreementList:  agreementList,
		AppCategory:    appCategory,
		Category:       &category,
		LicenseList:    licenseList,
		Logo:           logo,
		Os:             os,
		ScreenshotList: screenshotList,
		Support:        support,
	}
}

// Update the underlying Details resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDetailsResourceData(d *schema.ResourceData, m *models.Details) {
	d.Set("agreement_list", m.AgreementList)
	d.Set("app_category", m.AppCategory)
	d.Set("category", m.Category)
	d.Set("license_list", m.LicenseList)
	d.Set("logo", m.Logo)
	d.Set("os", m.Os)
	d.Set("screenshot_list", m.ScreenshotList)
	d.Set("support", m.Support)
}

// Iterate throught and update the Details resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDetailsSubResourceData(m []*models.Details) (d []*map[string]interface{}) {
	for _, DetailsModel := range m {
		if DetailsModel != nil {
			properties := make(map[string]interface{})
			properties["agreement_list"] = DetailsModel.AgreementList
			properties["app_category"] = DetailsModel.AppCategory
			properties["category"] = DetailsModel.Category
			properties["license_list"] = DetailsModel.LicenseList
			properties["logo"] = DetailsModel.Logo
			properties["os"] = DetailsModel.Os
			properties["screenshot_list"] = DetailsModel.ScreenshotList
			properties["support"] = DetailsModel.Support
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the Details resource defined in the Terraform configuration
func DetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"agreement_list": {
			Description: `UI map: AppEditPage:DeveloperPane:Developer_Agreement_Field, AppDetailsPage:DeveloperPane:Developer_Agreement_Field`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"app_category": {
			Description: ``,
			// We assume it's an enum type
			Type:     schema.TypeString,
			Required: true,
		},

		"category": {
			Description: `UI map: AppMarketplacePage:AppCard:DescriptionField, AppEditPage:IdentityPane:CategoryField, AppDetailsPage:IdentityPane:CategoryField`,
			Type:        schema.TypeString,
			Default:     "All",
			Optional:    true,
		},

		"license_list": {
			Description: `UI map: AppMarketplacePage:AppCard:License, AppEditPage:IdentityPane:License, AppDetailsPage:IdentityPane:License`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"logo": {
			Description: `UI map: AppEditPage:IdentityPane:Logo, AppDetailsPage:IdentityPane:Logo`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"os": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"screenshot_list": {
			Description: `UI map: AppEditPage:IdentityPane:Screenshot_Fields, AppDetailsPage:IdentityPane:Screenshot_Fields`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"support": {
			Description: `UI map: AppEditPage:DeveloperPane:Support_Description_Field, AppDetailsPage:DeveloperPane:Support_Description_Field`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the Details resource
func GetDetailsPropertyFields() (t []string) {
	return []string{
		"agreement_list",
		"app_category",
		"category",
		"license_list",
		"logo",
		"os",
		"screenshot_list",
		"support",
	}
}
