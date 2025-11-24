package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func GitRepoAuthConfigModel(d *schema.ResourceData) *models.GitRepoAuthConfig {
	encodedCertificates, _ := d.Get("encoded_certificates").(string)
	var git *models.GitAuthConfig // GitAuthConfig
	gitInterface, gitIsSet := d.GetOk("git")
	if gitIsSet && gitInterface != nil {
		gitMap := gitInterface.([]interface{})
		if len(gitMap) > 0 {
			git = GitAuthConfigModelFromMap(gitMap[0].(map[string]interface{}))
		}
	}
	var helm *models.HelmAuthConfig // HelmAuthConfig
	helmInterface, helmIsSet := d.GetOk("helm")
	if helmIsSet && helmInterface != nil {
		helmMap := helmInterface.([]interface{})
		if len(helmMap) > 0 {
			helm = HelmAuthConfigModelFromMap(helmMap[0].(map[string]interface{}))
		}
	}
	insecureSkipTLSVerify, _ := d.Get("insecure_skip_tls_verify").(bool)
	return &models.GitRepoAuthConfig{
		EncodedCertificates:   encodedCertificates,
		Git:                   git,
		Helm:                  helm,
		InsecureSkipTLSVerify: insecureSkipTLSVerify,
	}
}

func GitRepoAuthConfigModelFromMap(m map[string]interface{}) *models.GitRepoAuthConfig {
	encodedCertificates := m["encoded_certificates"].(string)
	var git *models.GitAuthConfig // GitAuthConfig
	gitInterface, gitIsSet := m["git"]
	if gitIsSet && gitInterface != nil {
		gitMap := gitInterface.([]interface{})
		if len(gitMap) > 0 {
			git = GitAuthConfigModelFromMap(gitMap[0].(map[string]interface{}))
		}
	}
	//
	var helm *models.HelmAuthConfig // HelmAuthConfig
	helmInterface, helmIsSet := m["helm"]
	if helmIsSet && helmInterface != nil {
		helmMap := helmInterface.([]interface{})
		if len(helmMap) > 0 {
			helm = HelmAuthConfigModelFromMap(helmMap[0].(map[string]interface{}))
		}
	}
	//
	insecureSkipTLSVerify := m["insecure_skip_tls_verify"].(bool)
	return &models.GitRepoAuthConfig{
		EncodedCertificates:   encodedCertificates,
		Git:                   git,
		Helm:                  helm,
		InsecureSkipTLSVerify: insecureSkipTLSVerify,
	}
}

func SetGitRepoAuthConfigResourceData(d *schema.ResourceData, m *models.GitRepoAuthConfig) {
	d.Set("encoded_certificates", m.EncodedCertificates)
	d.Set("git", SetGitAuthConfigSubResourceData([]*models.GitAuthConfig{m.Git}))
	d.Set("helm", SetHelmAuthConfigSubResourceData([]*models.HelmAuthConfig{m.Helm}))
	d.Set("insecure_skip_tls_verify", m.InsecureSkipTLSVerify)
}

func SetGitRepoAuthConfigSubResourceData(m []*models.GitRepoAuthConfig) (d []*map[string]interface{}) {
	for _, GitRepoAuthConfigModel := range m {
		if GitRepoAuthConfigModel != nil {
			properties := make(map[string]interface{})
			properties["encoded_certificates"] = GitRepoAuthConfigModel.EncodedCertificates
			properties["git"] = SetGitAuthConfigSubResourceData([]*models.GitAuthConfig{GitRepoAuthConfigModel.Git})
			properties["helm"] = SetHelmAuthConfigSubResourceData([]*models.HelmAuthConfig{GitRepoAuthConfigModel.Helm})
			properties["insecure_skip_tls_verify"] = GitRepoAuthConfigModel.InsecureSkipTLSVerify
			d = append(d, &properties)
		}
	}
	return
}

func GitRepoAuthConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"encoded_certificates": {
			Description: `Base64 encoded certificates for TLS authentication`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"git": {
			Description: `Git-specific authentication configuration`,
			Type:        schema.TypeList, //GoType: GitAuthConfig
			Elem: &schema.Resource{
				Schema: GitAuthConfigSchema(),
			},
			Optional: true,
		},

		"helm": {
			Description: `Helm-specific authentication configuration`,
			Type:        schema.TypeList, //GoType: HelmAuthConfig
			Elem: &schema.Resource{
				Schema: HelmAuthConfigSchema(),
			},
			Optional: true,
		},

		"insecure_skip_tls_verify": {
			Description: `Skip TLS certificate verification`,
			Type:        schema.TypeBool,
			Optional:    true,
		},
	}
}

func GetGitRepoAuthConfigPropertyFields() (t []string) {
	return []string{
		"encoded_certificates",
		"git",
		"helm",
		"insecure_skip_tls_verify",
	}
}
