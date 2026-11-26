package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// validateAppNamingScheme rejects an app_config naming_scheme that is not one of
// the controller's AppNamingScheme values, at plan time. The controller answers an
// unknown value with a 400 only when the request is sent (ENG-3119). An empty
// string is the same as not set.
var validateAppNamingScheme = validation.ToDiagFunc(validation.StringInSlice([]string{
	"",
	string(models.AppNamingSchemeAPPNAMINGSCHEMEUNSPECIFIED),
	string(models.AppNamingSchemeAPPNAMINGSCHEMEDEVICE),
	string(models.AppNamingSchemeAPPNAMINGSCHEMEAPPDEVICE),
	string(models.AppNamingSchemeAPPNAMINGSCHEMEPROJECTDEVICE),
	string(models.AppNamingSchemeAPPNAMINGSCHEMEPROJECTAPPDEVICE),
}, false))

// validateAppNamingSchemeV2 is the same check for app_inst_config.naming_scheme,
// which the controller types as AppNamingSchemeV2 (APP_NAMING_SCHEMEV2_*, not
// APP_NAMING_SCHEME_*). Using the V1 list there rejects every valid value.
var validateAppNamingSchemeV2 = validation.ToDiagFunc(validation.StringInSlice([]string{
	"",
	string(models.AppNamingSchemeV2APPNAMINGSCHEMEV2UNSPECIFIED),
	string(models.AppNamingSchemeV2APPNAMINGSCHEMEV2DEVICE),
	string(models.AppNamingSchemeV2APPNAMINGSCHEMEV2APPDEVICE),
	string(models.AppNamingSchemeV2APPNAMINGSCHEMEV2PROJECTDEVICE),
	string(models.AppNamingSchemeV2APPNAMINGSCHEMEV2PROJECTAPPDEVICE),
}, false))
