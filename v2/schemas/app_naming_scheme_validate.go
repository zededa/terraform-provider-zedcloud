package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// validateAppNamingScheme rejects a naming_scheme that is not one of the
// controller's AppNamingScheme values, at plan time. The controller answers an
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
