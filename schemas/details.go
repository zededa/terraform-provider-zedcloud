package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

func DetailsModel(d *schema.ResourceData) *models.Details {
	agreementList := map[string]string{}
	agreementListInterface, agreementListIsSet := d.GetOk("agreementList")
	if agreementListIsSet {
		agreementListMap := agreementListInterface.(map[string]interface{})
		for k, v := range agreementListMap {
			if v == nil {
				continue
			}
			agreementList[k] = v.(string)
		}
	}

	var appCategory *models.AppCategory // AppCategory
	appCategoryInterface, appCategoryIsSet := d.GetOk("app_category")
	if appCategoryIsSet {
		appCategoryModel := appCategoryInterface.(string)
		appCategory = models.NewAppCategory(models.AppCategory(appCategoryModel))
	}
	category, _ := d.Get("category").(string)
	licenseList := map[string]string{}
	licenseListInterface, licenseListIsSet := d.GetOk("licenseList")
	if licenseListIsSet {
		licenseListMap := licenseListInterface.(map[string]interface{})
		for k, v := range licenseListMap {
			if v == nil {
				continue
			}
			licenseList[k] = v.(string)
		}
	}

	logo := map[string]string{}
	logoInterface, logoIsSet := d.GetOk("logo")
	if logoIsSet {
		logoMap := logoInterface.(map[string]interface{})
		for k, v := range logoMap {
			if v == nil {
				continue
			}
			logo[k] = v.(string)
		}
	}

	os, _ := d.Get("os").(string)
	screenshotList := map[string]string{}
	screenshotListInterface, screenshotListIsSet := d.GetOk("screenshotList")
	if screenshotListIsSet {
		screenshotListMap := screenshotListInterface.(map[string]interface{})
		for k, v := range screenshotListMap {
			if v == nil {
				continue
			}
			screenshotList[k] = v.(string)
		}
	}

	support, _ := d.Get("support").(string)
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
	agreementList := map[string]string{}
	agreementListInterface, agreementListIsSet := m["agreement_list"]
	if agreementListIsSet {
		agreementListMap := agreementListInterface.(map[string]interface{})
		for k, v := range agreementListMap {
			if v == nil {
				continue
			}
			agreementList[k] = v.(string)
		}
	}

	var appCategory *models.AppCategory // AppCategory
	appCategoryInterface, appCategoryIsSet := m["app_category"]
	if appCategoryIsSet {
		appCategoryModel := appCategoryInterface.(string)
		appCategory = models.NewAppCategory(models.AppCategory(appCategoryModel))
	}
	category := m["category"].(string)
	licenseList := map[string]string{}
	licenseListInterface, licenseListIsSet := m["license_list"]
	if licenseListIsSet {
		licenseListMap := licenseListInterface.(map[string]interface{})
		for k, v := range licenseListMap {
			if v == nil {
				continue
			}
			licenseList[k] = v.(string)
		}
	}

	logo := map[string]string{}
	logoInterface, logoIsSet := m["logo"]
	if logoIsSet {
		logoMap := logoInterface.(map[string]interface{})
		for k, v := range logoMap {
			if v == nil {
				continue
			}
			logo[k] = v.(string)
		}
	}

	os := m["os"].(string)
	screenshotList := map[string]string{}
	screenshotListInterface, screenshotListIsSet := m["screenshot_list"]
	if screenshotListIsSet {
		screenshotListMap := screenshotListInterface.(map[string]interface{})
		for k, v := range screenshotListMap {
			if v == nil {
				continue
			}
			screenshotList[k] = v.(string)
		}
	}

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

// Iterate through and update the Details resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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
			Type:        schema.TypeString,
			Required:    true,
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
