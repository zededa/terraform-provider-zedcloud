package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func DetailsModel(d *schema.ResourceData) *models.Details {
	agreementList := map[string]string{}
	agreementListInterface, agreementListIsSet := d.GetOk("agreement_list")
	if agreementListIsSet {
		agreementListMap, _ := agreementListInterface.(map[string]interface{})
		for k, v := range agreementListMap {
			if v == nil {
				continue
			}
			agreementList[k], _ = v.(string)
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
	licenseListInterface, licenseListIsSet := d.GetOk("license_list")
	if licenseListIsSet {
		licenseListMap, _ := licenseListInterface.(map[string]interface{})
		for k, v := range licenseListMap {
			if v == nil {
				continue
			}
			licenseList[k], _ = v.(string)
		}
	}

	logo := map[string]string{}
	logoInterface, logoIsSet := d.GetOk("logo")
	if logoIsSet {
		logoMap, _ := logoInterface.(map[string]interface{})
		for k, v := range logoMap {
			if v == nil {
				continue
			}
			logo[k], _ = v.(string)
		}
	}

	os, _ := d.Get("os").(string)
	screenshotList := map[string]string{}
	screenshotListInterface, screenshotListIsSet := d.GetOk("screenshot_list")
	if screenshotListIsSet {
		screenshotListMap, _ := screenshotListInterface.(map[string]interface{})
		for k, v := range screenshotListMap {
			if v == nil {
				continue
			}
			screenshotList[k], _ = v.(string)
		}
	}

	support, _ := d.Get("support").(string)
	return &models.Details{
		AgreementList:  agreementList,
		AppCategory:    appCategory,
		Category:       &category, // string
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
		agreementListMap, _ := agreementListInterface.(map[string]interface{})
		for k, v := range agreementListMap {
			if v == nil {
				continue
			}
			agreementList[k], _ = v.(string)
		}
	}

	var appCategory *models.AppCategory // AppCategory
	appCategoryInterface, appCategoryIsSet := m["app_category"]
	if appCategoryIsSet {
		if appCategoryModel, ok := appCategoryInterface.(string); ok {
			appCategory = models.NewAppCategory(models.AppCategory(appCategoryModel))
		}
	}
	category, _ := m["category"].(string)
	licenseList := map[string]string{}
	licenseListInterface, licenseListIsSet := m["license_list"]
	if licenseListIsSet {
		licenseListMap, _ := licenseListInterface.(map[string]interface{})
		for k, v := range licenseListMap {
			if v == nil {
				continue
			}
			licenseList[k], _ = v.(string)
		}
	}

	logo := map[string]string{}
	logoInterface, logoIsSet := m["logo"]
	if logoIsSet {
		logoMap, _ := logoInterface.(map[string]interface{})
		for k, v := range logoMap {
			if v == nil {
				continue
			}
			logo[k], _ = v.(string)
		}
	}

	os, _ := m["os"].(string)
	screenshotList := map[string]string{}
	screenshotListInterface, screenshotListIsSet := m["screenshot_list"]
	if screenshotListIsSet {
		screenshotListMap, _ := screenshotListInterface.(map[string]interface{})
		for k, v := range screenshotListMap {
			if v == nil {
				continue
			}
			screenshotList[k], _ = v.(string)
		}
	}

	support, _ := m["support"].(string)
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

func Details() map[string]*schema.Schema {
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
			Description: "Licenses, keyed by license name. A value is either a URL or, for keys `CUSTOM_UPLOAD`, " +
				"`CUSTOM_UPLOAD_2`, ..., the id of an uploaded file (see `zedcloud_artifact`). " +
				"UI map: AppMarketplacePage:AppCard:License, AppEditPage:IdentityPane:License, AppDetailsPage:IdentityPane:License",
			Type: schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"logo": {
			Description: "App logo, as `{ logo = <artifact id> }`. Upload the image with `zedcloud_artifact` and use its " +
				"`id`. The key must be `logo` and the value must be an artifact id: the UI reads only one entry, " +
				"preferring the `logo` key, and does not render URLs. The UI accepts PNG or JPEG images up to 5 MB. " +
				"UI map: AppEditPage:IdentityPane:Logo, AppDetailsPage:IdentityPane:Logo",
			Type: schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:         true,
			ValidateDiagFunc: ValidateAppLogo,
		},

		"os": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"screenshot_list": {
			Description: "Screenshots, as artifact ids (see `zedcloud_artifact`). The current UI does not display them. " +
				"UI map: AppEditPage:IdentityPane:Screenshot_Fields, AppDetailsPage:IdentityPane:Screenshot_Fields",
			Type: schema.TypeMap, //GoType: map[string]string
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
