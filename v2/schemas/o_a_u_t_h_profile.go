package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func OAUTHProfileModel(d *schema.ResourceData) *models.OAUTHProfile {
	oIDCEndPoint, _ := d.Get("o_id_c_end_point").(string)
	additionalParameters := stringMap(d.Get("additional_parameters"))
	clientID, _ := d.Get("client_id").(string)
	clientSecret, _ := d.Get("client_secret").(string)
	cryptoKey, _ := d.Get("crypto_key").(string)
	encryptedSecrets := map[string]string{}
	encryptedSecretsInterface, encryptedSecretsIsSet := d.GetOk("encryptedSecrets")
	if encryptedSecretsIsSet {
		encryptedSecretsMap := encryptedSecretsInterface.(map[string]interface{})
		for k, v := range encryptedSecretsMap {
			if v == nil {
				continue
			}
			encryptedSecrets[k] = v.(string)
		}
	}

	idpID, _ := d.Get("idp_id").(string)
	var jwtAuthProfile *models.OAUTHProfileJWTAuthProfile // OAUTHProfileJWTAuthProfile
	jwtAuthProfileInterface, jwtAuthProfileIsSet := d.GetOk("jwt_auth_profile")
	if jwtAuthProfileIsSet && jwtAuthProfileInterface != nil {
		jwtAuthProfileMap := jwtAuthProfileInterface.([]interface{})
		if len(jwtAuthProfileMap) > 0 {
			jwtAuthProfile = OAUTHProfileJWTAuthProfileModelFromMap(jwtAuthProfileMap[0].(map[string]interface{}))
		}
	}
	roleScope, _ := d.Get("role_scope").(string)
	return &models.OAUTHProfile{
		OIDCEndPoint:         oIDCEndPoint,
		AdditionalParameters: additionalParameters,
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		CryptoKey:            cryptoKey,
		EncryptedSecrets:     encryptedSecrets,
		IdpID:                idpID,
		JwtAuthProfile:       jwtAuthProfile,
		RoleScope:            roleScope,
	}
}

func OAUTHProfileModelFromMap(m map[string]interface{}) *models.OAUTHProfile {
	oIDCEndPoint := m["o_id_c_end_point"].(string)
	additionalParameters := stringMap(m["additional_parameters"])
	clientID := m["client_id"].(string)
	clientSecret := m["client_secret"].(string)
	cryptoKey := m["crypto_key"].(string)
	encryptedSecrets := map[string]string{}
	encryptedSecretsInterface, encryptedSecretsIsSet := m["encrypted_secrets"]
	if encryptedSecretsIsSet {
		encryptedSecretsMap := encryptedSecretsInterface.(map[string]interface{})
		for k, v := range encryptedSecretsMap {
			if v == nil {
				continue
			}
			encryptedSecrets[k] = v.(string)
		}
	}

	idpID := m["idp_id"].(string)
	var jwtAuthProfile *models.OAUTHProfileJWTAuthProfile // OAUTHProfileJWTAuthProfile
	jwtAuthProfileInterface, jwtAuthProfileIsSet := m["jwt_auth_profile"]
	if jwtAuthProfileIsSet && jwtAuthProfileInterface != nil {
		jwtAuthProfileMap := jwtAuthProfileInterface.([]interface{})
		if len(jwtAuthProfileMap) > 0 {
			jwtAuthProfile = OAUTHProfileJWTAuthProfileModelFromMap(jwtAuthProfileMap[0].(map[string]interface{}))
		}
	}
	//
	roleScope := m["role_scope"].(string)
	return &models.OAUTHProfile{
		OIDCEndPoint:         oIDCEndPoint,
		AdditionalParameters: additionalParameters,
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		CryptoKey:            cryptoKey,
		EncryptedSecrets:     encryptedSecrets,
		IdpID:                idpID,
		JwtAuthProfile:       jwtAuthProfile,
		RoleScope:            roleScope,
	}
}

func SetOAUTHProfileResourceData(d *schema.ResourceData, m *models.OAUTHProfile) {
	d.Set("o_id_c_end_point", m.OIDCEndPoint)
	d.Set("additional_parameters", emptyMapToNil(m.AdditionalParameters))
	d.Set("client_id", m.ClientID)
	d.Set("client_secret", m.ClientSecret)
	d.Set("crypto_key", m.CryptoKey)
	d.Set("encrypted_secrets", m.EncryptedSecrets)
	d.Set("idp_id", m.IdpID)
	d.Set("jwt_auth_profile", SetOAUTHProfileJWTAuthProfileSubResourceData([]*models.OAUTHProfileJWTAuthProfile{m.JwtAuthProfile}))
	d.Set("role_scope", m.RoleScope)
}

func SetOAUTHProfileSubResourceData(m []*models.OAUTHProfile) (d []*map[string]interface{}) {
	for _, OAUTHProfileModel := range m {
		if OAUTHProfileModel != nil {
			properties := make(map[string]interface{})
			properties["o_id_c_end_point"] = OAUTHProfileModel.OIDCEndPoint
			// UE-168: the gateway always emits this map, as {} when unset.
			// Writing an empty map into state where the configuration omits
			// the attribute leaves a diff that never settles, so normalise
			// empty to absent.
			properties["additional_parameters"] = emptyMapToNil(OAUTHProfileModel.AdditionalParameters)
			properties["client_id"] = OAUTHProfileModel.ClientID
			properties["client_secret"] = OAUTHProfileModel.ClientSecret
			properties["crypto_key"] = OAUTHProfileModel.CryptoKey
			properties["encrypted_secrets"] = OAUTHProfileModel.EncryptedSecrets
			properties["idp_id"] = OAUTHProfileModel.IdpID
			properties["jwt_auth_profile"] = SetOAUTHProfileJWTAuthProfileSubResourceData([]*models.OAUTHProfileJWTAuthProfile{OAUTHProfileModel.JwtAuthProfile})
			properties["role_scope"] = OAUTHProfileModel.RoleScope
			d = append(d, &properties)
		}
	}
	return
}

func OAUTHProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"o_id_c_end_point": {
			Description: `OIDC endpoint for oauth validation`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"additional_parameters": {
			Description: `pass additional url parameters during the exchange and authorization process`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"client_id": {
			Description: `OAUTH client ID`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		// UE-168: secret material. The controller encrypts client_secret on
		// write and blanks all three on every read (fillAuthProfile in
		// zedcloud srvs/indusv2/authprofileproc.go), so they are never
		// returned and an imported profile carries none of them — but
		// whatever the operator puts in configuration does land in state,
		// hence Sensitive.
		"client_secret": {
			Description: `OAUTH client secret. Write-only: the API never returns it, so it ` +
				`must be supplied in configuration after importing a profile.`,
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},

		"crypto_key": {
			Description: `Key used to encrypt the stored secrets. Set by the controller; never returned.`,
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
		},

		"encrypted_secrets": {
			Description: `Controller-encrypted secrets. Set by the controller; never returned.`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:  true,
			Sensitive: true,
		},

		"idp_id": {
			Description: `id for Vmware IDP`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"jwt_auth_profile": {
			Description: `Config for JWT based authentication, jwks_uri is derived from OIDC Well Known Endpoints`,
			Type:        schema.TypeList, //GoType: OAUTHProfileJWTAuthProfile
			Elem: &schema.Resource{
				Schema: OAUTHProfileJWTAuthProfileSchema(),
			},
			Optional: true,
		},

		"role_scope": {
			Description: `OIDC scope to fetch application role`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

func GetOAUTHProfilePropertyFields() (t []string) {
	return []string{
		"o_id_c_end_point",
		"additional_parameters",
		"client_id",
		"client_secret",
		"crypto_key",
		"encrypted_secrets",
		"idp_id",
		"jwt_auth_profile",
		"role_scope",
	}
}
