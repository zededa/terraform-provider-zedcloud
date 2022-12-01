package provider

import (
	"context"
	"fmt"
	"log"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider/client"
	"github.com/zededa/terraform-provider/resources"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigure,
		Schema: map[string]*schema.Schema{
			"zedcloud_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud url. Ex: https://zedcontrol.zededa.net",
				DefaultFunc: schema.EnvDefaultFunc("ZEDCLOUD_API_URL", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud API token",
				DefaultFunc: schema.EnvDefaultFunc("ZEDCLOUD_API_TOKEN", nil),
			},
			"user": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud user",
				DefaultFunc: schema.EnvDefaultFunc("ZEDCLOUD_USER", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ZEDCloud user password ",
				DefaultFunc: schema.EnvDefaultFunc("ZEDCLOUD_PASS", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"zedcloud_edgenode": resources.DataResourceDeviceConfig(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"zedcloud_edgenode": resources.DeviceConfig(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	zedcloudUrl := d.Get("zedcloud_url").(string)
	token := d.Get("token").(string)
	username := d.Get("user").(string)
	password := d.Get("password").(string)

	// config := client.NewConfig()

	// c := client.New(config)

	transport := httptransport.New(zedcloudUrl, "api", []string{"https"})
	transport.SetDebug(true)

	transport.DefaultAuthentication = auth

	appCli := client.New(transport, strfmt.Default)

	return c, diags
}

func NewClient(zedcloudUrl string, token, username, password string) (*Client, error) {

	// jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	// if err != nil {
	// 	return nil, err
	// }
	// transport.Jar = jar
	// client.HttpClient.Jar = jar

	// Set BearerToken
	if token != "" {
		client.Bearer = httptransport.BearerToken(token)
		client.BearerToken = token
		log.Println("[INFO] Created Bearer")
	} else {
		if username == "" || password == "" {
			err = fmt.Errorf("Token or Username/Password must be specified")
			log.Printf("***[ERROR] %s", err.Error())
			return nil, err
		}
		// Get BearerToken by doing a Login
		err = client.login(username, password)
		if err != nil {
			err = fmt.Errorf("LOGIN FAILED: %s", err.Error())
			return nil, err
		}
	}
	return client, nil
}

// Client holds metadata / config for use by Terraform resources
// type Client struct {
// 	// ZedcloudUrl Ex: "https://zedcontrol.zededa.net"
// 	ZedcloudUrl string
// 	// Token = API Token to be used
// 	//  Either Token OR Username / Password must be specified.
// 	Token    string
// 	Username string
// 	Password string
// 	Client   *zedcloudapi.Client
// }

// providerConfigure parses the config into the Terraform provider meta object
// func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

// 	if token != "" {
// 		username = ""
// 		password = ""
// 		log.Println("Token Specified. Using Token for Authentication")
// 	} else if username == "" && password == "" {
// 		err := fmt.Errorf("Neither token(%s) nor username(%s)/password(%s) specified",
// 			token, username, password)
// 		log.Printf("%s", err)
// 		return nil, diag.Errorf(err.Error())
// 	}
// 	log.Printf("providerConfigure - Configuring Provider. Creating new client")

// 	providerClient := Client{
// 		ZedcloudUrl: zedcloudUrl,
// 		Token:       token,
// 		Username:    username,
// 		Password:    password,
// 	}
// 	var err error
// 	providerClient.Client, err = zedcloudapi.NewClient(zedcloudUrl, token, username, password)
// 	if err != nil {
// 		log.Printf("Failed to create new client. err: %+v", err)
// 		return nil, diag.Errorf(err.Error())
// 	}
// 	return providerClient, diags
// }
