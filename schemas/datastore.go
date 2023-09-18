package schemas

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/models"
)

// Note, when providing secrets as variable instead of hard-coding them, the fields api_key, ds_err and enterprise_id
// reports a diff in a subsequent plan command. This seems to be a bug in TF SDKv2 transition of nil values or similar.
// So these fields are removed (commented) from the schema.

func DatastoreModel(d *schema.ResourceData) *models.Datastore {
	aPIKey, _ := d.Get("api_key").(string)
	var certificateChain *models.CertificateChain // CertificateChain
	certificateChainInterface, certificateChainIsSet := d.GetOk("certificate_chain")
	if certificateChainIsSet && certificateChainInterface != nil {
		certificateChainMap := certificateChainInterface.([]interface{})
		if len(certificateChainMap) > 0 {
			certificateChain = CertificateChainModelFromMap(certificateChainMap[0].(map[string]interface{}))
		}
	}
	description, _ := d.Get("description").(string)
	dsFQDN, _ := d.Get("ds_fqdn").(string)
	dsPath, _ := d.Get("ds_path").(string)
	var dsStatus *models.DatastoreStatus // DatastoreStatus
	dsStatusInterface, dsStatusIsSet := d.GetOk("ds_status")
	if dsStatusIsSet {
		dsStatusModel := dsStatusInterface.(string)
		dsStatus = models.NewDatastoreStatus(models.DatastoreStatus(dsStatusModel))
	}
	var dsType *models.DatastoreType // DatastoreType
	dsTypeInterface, dsTypeIsSet := d.GetOk("ds_type")
	if dsTypeIsSet {
		dsTypeModel := dsTypeInterface.(string)
		dsType = models.NewDatastoreType(models.DatastoreType(dsTypeModel))
	}
	// enterpriseID, _ := d.Get("enterprise_id").(string)
	id, _ := d.Get("id").(string)
	name, _ := d.Get("name").(string)
	needClearText, _ := d.Get("need_clear_text").(bool)
	var projectAccessList []string
	projectAccessListInterface, projectAccessListIsSet := d.GetOk("project_access_list")
	if projectAccessListIsSet {
		var items []interface{}
		if listItems, isList := projectAccessListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = projectAccessListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			projectAccessList = append(projectAccessList, v.(string))
		}
	}
	region, _ := d.Get("region").(string)
	var secret *models.DatastoreInfoSecrets // DatastoreInfoSecrets
	secretInterface, secretIsSet := d.GetOk("secret")
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = DatastoreSecretsModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	title, _ := d.Get("title").(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := d.GetOk("revision")
	if revisionIsSet {
		revisionMap := revisionInterface.([]interface{})[0].(map[string]interface{})
		revision = ObjectRevisionModelFromMap(revisionMap)
	}
	return &models.Datastore{
		APIKey:           aPIKey,
		CertificateChain: certificateChain,
		Description:      description,
		DsFQDN:           &dsFQDN,
		DsPath:           &dsPath,
		DsStatus:         dsStatus,
		DsType:           dsType,
		// EnterpriseID:      enterpriseID,
		ID:                id,
		Name:              &name,
		NeedClearText:     needClearText,
		ProjectAccessList: projectAccessList,
		Region:            region,
		Secret:            secret,
		Title:             &title,
		Revision:          revision,
	}
}

func DatastoreModelFromMap(m map[string]interface{}) *models.Datastore {
	aPIKey := m["api_key"].(string)
	var certificateChain *models.CertificateChain // CertificateChain
	certificateChainInterface, certificateChainIsSet := m["certificate_chain"]
	if certificateChainIsSet && certificateChainInterface != nil {
		certificateChainMap := certificateChainInterface.([]interface{})
		if len(certificateChainMap) > 0 {
			certificateChain = CertificateChainModelFromMap(certificateChainMap[0].(map[string]interface{}))
		}
	}
	//
	description := m["description"].(string)
	dsFQDN := m["ds_fqdn"].(string)
	dsPath := m["ds_path"].(string)
	var dsStatus *models.DatastoreStatus // DatastoreStatus
	dsStatusInterface, dsStatusIsSet := m["ds_status"]
	if dsStatusIsSet {
		dsStatusModel := dsStatusInterface.(string)
		dsStatus = models.NewDatastoreStatus(models.DatastoreStatus(dsStatusModel))
	}
	var dsType *models.DatastoreType // DatastoreType
	dsTypeInterface, dsTypeIsSet := m["ds_type"]
	if dsTypeIsSet {
		dsTypeModel := dsTypeInterface.(string)
		dsType = models.NewDatastoreType(models.DatastoreType(dsTypeModel))
	}
	// enterpriseID := m["enterprise_id"].(string)
	id := m["id"].(string)
	name := m["name"].(string)
	needClearText := m["need_clear_text"].(bool)
	var projectAccessList []string
	projectAccessListInterface, projectAccessListIsSet := m["project_access_list"]
	if projectAccessListIsSet {
		var items []interface{}
		if listItems, isList := projectAccessListInterface.([]interface{}); isList {
			items = listItems
		} else {
			items = projectAccessListInterface.(*schema.Set).List()
		}
		for _, v := range items {
			if v == nil {
				continue
			}
			projectAccessList = append(projectAccessList, v.(string))
		}
	}
	region := m["region"].(string)
	var secret *models.DatastoreInfoSecrets // DatastoreInfoSecrets
	secretInterface, secretIsSet := m["secret"]
	if secretIsSet && secretInterface != nil {
		secretMap := secretInterface.([]interface{})
		if len(secretMap) > 0 {
			secret = DatastoreSecretsModelFromMap(secretMap[0].(map[string]interface{}))
		}
	}
	//
	title := m["title"].(string)
	var revision *models.ObjectRevision // ObjectRevision
	revisionInterface, revisionIsSet := m["revision"]
	if revisionIsSet && revisionInterface != nil {
		revisionMap := revisionInterface.([]interface{})
		if len(revisionMap) > 0 {
			revision = ObjectRevisionModelFromMap(revisionMap[0].(map[string]interface{}))
		}
	}
	return &models.Datastore{
		APIKey:           aPIKey,
		CertificateChain: certificateChain,
		Description:      description,
		DsFQDN:           &dsFQDN,
		DsPath:           &dsPath,
		DsStatus:         dsStatus,
		DsType:           dsType,
		// EnterpriseID:      enterpriseID,
		ID:                id,
		Name:              &name,
		NeedClearText:     needClearText,
		ProjectAccessList: projectAccessList,
		Region:            region,
		Secret:            secret,
		Title:             &title,
		Revision:          revision,
	}
}

func SetDatastoreResourceData(d *schema.ResourceData, m *models.Datastore) {
	d.Set("api_key", m.APIKey)
	d.Set("certificate_chain", SetCertificateChainSubResourceData([]*models.CertificateChain{m.CertificateChain}))
	d.Set("crypto_key", m.CryptoKey)
	d.Set("description", m.Description)
	// d.Set("ds_err", m.DsErr)
	d.Set("ds_fqdn", m.DsFQDN)
	d.Set("ds_path", m.DsPath)
	d.Set("ds_status", m.DsStatus)
	d.Set("ds_type", m.DsType)
	d.Set("encrypted_secrets", m.EncryptedSecrets)
	// d.Set("enterprise_id", m.EnterpriseID)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("need_clear_text", m.NeedClearText)
	d.Set("origin_type", m.OriginType)
	d.Set("project_access_list", m.ProjectAccessList)
	d.Set("region", m.Region)
	d.Set("revision", SetObjectRevisionSubResourceData([]*models.ObjectRevision{m.Revision}))
	d.Set("secret", SetDatastoreSecretsSubResourceData([]*models.DatastoreInfoSecrets{m.Secret}))
	d.Set("title", m.Title)
}

func SetDatastoreInfoSubResourceData(m []*models.Datastore) (d []*map[string]interface{}) {
	for _, DatastoreInfoModel := range m {
		if DatastoreInfoModel != nil {
			properties := make(map[string]interface{})
			properties["api_key"] = DatastoreInfoModel.APIKey
			properties["certificate_chain"] = SetCertificateChainSubResourceData([]*models.CertificateChain{DatastoreInfoModel.CertificateChain})
			properties["crypto_key"] = DatastoreInfoModel.CryptoKey
			properties["description"] = DatastoreInfoModel.Description
			// properties["ds_err"] = DatastoreInfoModel.DsErr
			properties["ds_fqdn"] = DatastoreInfoModel.DsFQDN
			properties["ds_path"] = DatastoreInfoModel.DsPath
			properties["ds_status"] = DatastoreInfoModel.DsStatus
			properties["ds_type"] = DatastoreInfoModel.DsType
			properties["encrypted_secrets"] = DatastoreInfoModel.EncryptedSecrets
			// properties["enterprise_id"] = DatastoreInfoModel.EnterpriseID
			properties["id"] = DatastoreInfoModel.ID
			properties["name"] = DatastoreInfoModel.Name
			properties["need_clear_text"] = DatastoreInfoModel.NeedClearText
			properties["origin_type"] = DatastoreInfoModel.OriginType
			properties["project_access_list"] = DatastoreInfoModel.ProjectAccessList
			properties["region"] = DatastoreInfoModel.Region
			properties["revision"] = SetObjectRevisionSubResourceData([]*models.ObjectRevision{DatastoreInfoModel.Revision})
			properties["secret"] = SetDatastoreSecretsSubResourceData([]*models.DatastoreInfoSecrets{DatastoreInfoModel.Secret})
			properties["title"] = DatastoreInfoModel.Title
			d = append(d, &properties)
		}
	}
	return
}

func Datastore() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_key": {
			Description: ``,
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
		},

		"certificate_chain": {
			Description: `Certificate chain of the certificate`,
			Type:        schema.TypeList, //GoType: CertificateChain
			Elem: &schema.Resource{
				Schema: CertificateChainSchema(),
			},
			Optional:  true,
			Computed:  true,
			Sensitive: true,
		},

		"crypto_key": {
			Description: `Internal - Encryption Key context`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"description": {
			Description: `Detailed description of the datastore.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		// "ds_err": {
		// 	Description: `Datastore validation detailed error/status message`,
		// 	Type:        schema.TypeString,
		// 	Computed:    true,
		// },

		"ds_fqdn": {
			Description: `Datastore Fully Qualified Domain Name`,
			Type:        schema.TypeString,
			Optional:    true,
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				if oldValue == fmt.Sprintf("https://%s", newValue) {
					return true
				}
				return false
			},
		},

		"ds_path": {
			Description: `Datastore relative path w.r.t. Datastore root`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"ds_status": {
			Description: `Datastore status`,
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
		},

		"ds_type": {
			Description: `Datastore type`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"encrypted_secrets": {
			Description: `Internal - Encrypted sensitive data`,
			Type:        schema.TypeMap, //GoType: map[string]string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed:  true,
			Sensitive: true,
		},

		// "enterprise_id": {
		// 	Description: ``,
		// 	Type:        schema.TypeString,
		// 	Optional:    true,
		// },

		"id": {
			Description: `System defined universally unique Id of the datastore.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"name": {
			Description: `User defined name of the datastore, unique across the enterprise. Once datastore is created, name can’t be changed.`,
			Type:        schema.TypeString,
			Required:    true,
			ForceNew: true,
		},

		"need_clear_text": {
			Description: `knob for sending creds in clear text`,
			Type:        schema.TypeBool,
			Optional:    true,
			DiffSuppressFunc: func(k, oldValue, newValue string, d *schema.ResourceData) bool {
				return true
			},
		},

		"origin_type": {
			Description: `Origin type of datastore.`,
			Type:        schema.TypeString,
			Computed:    true,
		},

		"project_access_list": {
			Description: `project access list of the datastore`,
			Type:        schema.TypeList, //GoType: []string
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"region": {
			Description: `Datastore region - valid for AWS S3 and Azure BlobStorage`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"revision": {
			Description: `Object revision details`,
			Type:        schema.TypeList, //GoType: ObjectRevision
			Elem: &schema.Resource{
				Schema: ObjectRevision(),
			},
			Optional: true,
			Computed: true,
		},

		"secret": {
			Description: `Plain-text sensitive data`,
			Type:        schema.TypeList, //GoType: DatastoreInfoSecrets
			Elem: &schema.Resource{
				Schema: DatastoreSecrets(),
			},
			Optional:  true,
			Sensitive: true,
		},

		"title": {
			Description: `User defined title of the datastore. Title can be changed at any time.`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the DatastoreInfo resource
func GetDatastoreInfoPropertyFields() (t []string) {
	return []string{
		"api_key",
		"certificate_chain",
		"description",
		"ds_fqdn",
		"ds_path",
		"ds_status",
		"ds_type",
		"enterprise_id",
		"id",
		"name",
		"need_clear_text",
		"project_access_list",
		"region",
		"secret",
		"title",
	}
}
