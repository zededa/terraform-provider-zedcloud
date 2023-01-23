package resources

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiclient "github.com/zededa/terraform-provider/client"
	"github.com/zededa/terraform-provider/client/edge_node_configuration"
	"github.com/zededa/terraform-provider/models"
)

func TestAccGetEdgeNode(t *testing.T) {
	var edgeNodeBefore models.DeviceConfig
	// var edgeNodeAfter models.DeviceConfig

	// create with required and computed fields
	requiredOnlyConfigPath := "edge_node/required_only.tf"
	requiredOnlyConfig, err := getTestConfig(requiredOnlyConfigPath)
	if err != nil {
		t.Fatal(fmt.Sprintf("could not get testdata for %s", requiredOnlyConfigPath))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testEdgeNodeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: requiredOnlyConfig,
				Check: resource.ComposeTestCheckFunc(
					testEdgeNodeExists("zedcloud_edgenode.required_only", &edgeNodeBefore),
					resource.TestCheckResourceAttr("zedcloud_edgenode.required_only", "name", "required_only"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.required_only", "model_id", "2f716b55-2639-486c-9a2f-55a2e94146a6"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.required_only", "title", "required_only-title"),
					resource.TestMatchResourceAttr(
						"zedcloud_edgenode.required_only",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
				),
			},
		},
	})

	// create with all fields
	var edgeNodeComplete models.DeviceConfig
	completeConfigPath := "edge_node/complete.tf"
	completeConfig, err := getTestConfig(completeConfigPath)
	if err != nil {
		t.Fatal(fmt.Sprintf("could not get testdata for %s", completeConfigPath))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testEdgeNodeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: completeConfig,
				Check: resource.ComposeTestCheckFunc(
					testEdgeNodeExists("zedcloud_edgenode.complete", &edgeNodeComplete),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "name", "complete"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "model_id", "2f716b55-2639-486c-9a2f-55a2e94146a6"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "title", "complete-title"),
					resource.TestMatchResourceAttr(
						"zedcloud_edgenode.complete",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testEdgeNodeAttributes(&edgeNodeComplete),
				),
			},
		},
	})

	// update all fields
	// create and update fields with custom logic and separate api requests
}

// testAccPreCheck validates the necessary test API keys exist
// in the testing environment
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TF_CLI_CONFIG_FILE"); v == "" {
		t.Fatal("TF_CLI_CONFIG_FILE must be set for acceptance tests, it should contain the dev_overrides config that points to local instance of the provider")
	}
	if v := os.Getenv("TF_VAR_zedcloud_token"); v == "" {
		t.Fatal("TF_VAR_zedcloud_token must be set for acceptance tests to access the zedcloud API")
	}
}

// testEdgeNodeExists retrieves the EdgeNode and stores it in the provided *models.DeviceConfig.
func testEdgeNodeExists(resourceName string, device *models.DeviceConfig) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("EdgeNode not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("EdgeNode ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testAccProvider.Meta().(*apiclient.Zedcloudapi)

		// retrieve the EdgeNode by referencing its state ID for API lookup
		params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeNodeParams()
		params.ID = rs.Primary.ID
		response, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeNode(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch EdgeNode (%s): %w", rs.Primary.ID, err)
		}

		deviceConfig := response.GetPayload()
		if deviceConfig == nil {
			return errors.New("could not get response payload in EdgeNode existence test: deviceConfig is nil")
		}

		// store the resulting EdgeNode config in the *models.DeviceConfig variable
		*device = *deviceConfig
		return nil
	}
}

// testEdgeNodeAttributes verifies attributes are set correctly by Terraform
func testEdgeNodeAttributes(device *models.DeviceConfig) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if string(*device.AdminState) != "ADMIN_STATE_ACTIVE" {
			return fmt.Errorf("expect device.AdminState == ADMIN_STATE_ACTIVE but got %+v", device.AdminState)
		}
		if device.AssetID != "asset_id" {
			return fmt.Errorf("expect device.AssetId == asset_id but got %+v", device.AssetID)
		}
		if len(device.BaseImage) != 0 {
			return fmt.Errorf("expect device.BaseImage == [] but got %+v", device.BaseImage)
		}
		if device.ConfigItem[0].BoolValue != true {
			return fmt.Errorf("expect device.ConfigItem[0].- Boolvalue == true but got %+v", device.ConfigItem[0].BoolValue)
		}
		if device.ConfigItem[0].FloatValue != 1.0 {
			return fmt.Errorf("expect device.ConfigItem[0].Floatvalue == 1 but got %+v", device.ConfigItem[0].FloatValue)
		}
		if device.ConfigItem[0].Key != "key" {
			return fmt.Errorf("expect device.ConfigItem[0].Key == key but got %+v", device.ConfigItem[0].Key)
		}
		if device.ConfigItem[0].StringValue != "string" {
			return fmt.Errorf("expect device.ConfigItem[0].Stringvalue == string but got %+v", device.ConfigItem[0].StringValue)
		}
		if device.ConfigItem[0].Uint32Value != 32 {
			return fmt.Errorf("expect device.ConfigItem[0].Uint32value == 32 but got %+v", device.ConfigItem[0].Uint32Value)
		}
		if device.ConfigItem[0].Uint64Value != "64" {
			return fmt.Errorf("expect device.ConfigItem[0].Uint64value == 64 but got %+v", device.ConfigItem[0].Uint64Value)
		}
		if device.ConfigItem[0].ValueType != "value type" {
			return fmt.Errorf("expect device.ConfigItem[0].Valuetype == value type but got %+v", device.ConfigItem[0].ValueType)
		}
		if device.DeploymentTag != "depl_tag" {
			return fmt.Errorf("expect device.DeploymentTag == depl_tag but got %+v", device.DeploymentTag)
		}
		if device.Description != "description" {
			return fmt.Errorf("expect device.Description == description but got %+v", device.Description)
		}
		if device.DevLocation.City != "berlin" {
			return fmt.Errorf("expect device.DevLocation[0].City == berlin but got %+v", device.DevLocation.City)
		}
		if device.DevLocation.Country != "germany" {
			return fmt.Errorf("expect device.DevLocation[0].Country == germany but got %+v", device.DevLocation.Country)
		}
		if device.DevLocation.Freeloc != "freeloc" {
			return fmt.Errorf("expect device.DevLocation[0].Freeloc == freeloc but got %+v", device.DevLocation.Freeloc)
		}
		if device.DevLocation.Hostname != "hostname" {
			return fmt.Errorf("expect device.DevLocation[0].Hostname == hostname but got %+v", device.DevLocation.Hostname)
		}
		if device.DevLocation.Latlong != "" {
			return fmt.Errorf("expect device.DevLocation[0].Latlong == \"\" but got %+v", device.DevLocation.Latlong)
		}
		if device.DevLocation.Loc != "52.520008, 13.404954" {
			return fmt.Errorf("expect device.DevLocation[0].Loc == 52.520008, 13.404954 but got %+v", device.DevLocation.Loc)
		}
		if device.DevLocation.Org != "zededa" {
			return fmt.Errorf("expect device.DevLocation[0].Org == zededa but got %+v", device.DevLocation.Org)
		}
		if device.DevLocation.Postal != "10115" {
			return fmt.Errorf("expect device.DevLocation[0].Postal == \"10115\" but got %+v", device.DevLocation.Postal)
		}
		if device.DevLocation.Region != "europe/west" {
			return fmt.Errorf("expect device.DevLocation[0].Region == europe/west but got %+v", device.DevLocation.Region)
		}
		if device.DevLocation.UnderlayIP != "" {
			return fmt.Errorf("expect device.DevLocation[0].Underlayip == \"\" but got %+v", device.DevLocation.UnderlayIP)
		}
		if device.Interfaces[0].Cost != 0 {
			return fmt.Errorf("expect device.Interfaces[0].Cost == 0 but got %+v", device.Interfaces[0].Cost)
		}
		if string(*device.Interfaces[0].IntfUsage) != "ADAPTER_USAGE_UNSPECIFIED" {
			return fmt.Errorf("expect device.Interfaces[0].Intfusage == ADAPTER_USAGE_UNSPECIFIED but got %+v", device.Interfaces[0].IntfUsage)
		}
		if device.Interfaces[0].Intfname != "eth0" {
			return fmt.Errorf("expect device.Interfaces[0].Intfname == eth0 but got %+v", device.Interfaces[0].Intfname)
		}
		if device.Interfaces[0].Ipaddr != "" {
			return fmt.Errorf("expect device.Interfaces[0].Ipaddr == \"\" but got %+v", device.Interfaces[0].Ipaddr)
		}
		if device.Interfaces[0].Macaddr != "" {
			return fmt.Errorf("expect device.Interfaces[0].Macaddr == \"\" but got %+v", device.Interfaces[0].Macaddr)
		}
		if device.Interfaces[0].Netname != "" {
			return fmt.Errorf("expect device.Interfaces[0].Netname == \"\" but got %+v", device.Interfaces[0].Netname)
		}
		if len(device.Interfaces[0].Tags) != 0 {
			return fmt.Errorf("expect len(device.Interfaces[0].Tags) == 0 but got %+v", len(device.Interfaces[0].Tags))
		}
		if device.Interfaces[1].Cost != 255 {
			return fmt.Errorf("expect device.Interfaces[1].Cost == 255 but got %+v", device.Interfaces[1].Cost)
		}
		if string(*device.Interfaces[1].IntfUsage) != "ADAPTER_USAGE_MANAGEMENT" {
			return fmt.Errorf("expect device.Interfaces[1].IntfUsage == ADAPTER_USAGE_MANAGEMENT but got %+v", device.Interfaces[1].IntfUsage)
		}
		if device.Interfaces[1].Intfname != "defaultIPv4" {
			return fmt.Errorf("expect device.Interfaces[1].Intfname == defaultIPv4 but got %+v", device.Interfaces[1].Intfname)
		}
		if device.Interfaces[1].Ipaddr != "127.0.0.1" {
			return fmt.Errorf("expect device.Interfaces[1].Ipaddr == 127.0.0.1 but got %+v", device.Interfaces[1].Ipaddr)
		}
		if device.Interfaces[1].Macaddr != "00:00:00:00:00:00" {
			return fmt.Errorf("expect device.Interfaces[1].Macaddr == \"00:00:00:00:00:00\" but got %+v", device.Interfaces[1].Macaddr)
		}
		if device.Interfaces[1].Netname != "" {
			return fmt.Errorf("expect device.Interfaces[1].Netname == \"\" but got %+v", device.Interfaces[1].Netname)
		}
		if device.Interfaces[1].Tags["system_interface_1_key"] != "system_interface_1_value" {
			return fmt.Errorf("expect device.Interfaces[1].Tags[\"system_interface_1_key\"] == system_interface_1_value but got %+v", device.Interfaces[1].Tags["system_interface_1_key"])
		}
		if *device.ModelID != "2f716b55-2639-486c-9a2f-55a2e94146a6" {
			return fmt.Errorf("expect device.ModelId == 2f716b55-2639-486c-9a2f-55a2e94146a6 but got %+v", *device.ModelID)
		}
		if *device.Name != "complete" {
			return fmt.Errorf("expect device.Name == complete but got %+v", device.Name)
		}
		if strings.Split(device.Obkey, ":")[0] != "5d0767ee-0547-4569-b530-387e526f8cb9" {
			return fmt.Errorf("expect device.Obkey == 5d0767ee-0547-4569-b530-387e526f8cb9:* but got %+v", device.Obkey)
		}
		if *device.ProjectID != "4754cd0f-82d7-4e06-a68f-ff9e23e75ccf" {
			return fmt.Errorf("expect device.ProjectId == 4754cd0f-82d7-4e06-a68f-ff9e23e75ccf but got %+v", *device.ProjectID)
		}
		if device.Serialno != "2293dbe8-29ce-420c-8264-962857efc46b" {
			return fmt.Errorf("expect device.Serialno == 2293dbe8-29ce-420c-8264-962857efc46b but got %+v", device.Serialno)
		}
		if len(device.SitePictures) != 0 {
			return fmt.Errorf("expect device.SitePictures == [] but got %+v", device.SitePictures)
		}
		if device.Tags["tag-key-1"] != "tag-value-1" {
			return fmt.Errorf("expect device.Tags[\"tag-key-1\"] == tag-value-1 but got %+v", device.Tags["tag-key-1"])
		}
		if device.Tags["tag-key-2"] != "tag-value-2" {
			return fmt.Errorf("expect device.Tags[\"tag-key-2\"] == tag-value-2 but got %+v", device.Tags["tag-key-2"])
		}
		if *device.Title != "complete-title" {
			return fmt.Errorf("expect device.Title == complete-title but got %+v", device.Title)
		}
		if device.Token != "token" {
			return fmt.Errorf("expect device.Token == token but got %+v", device.Token)
		}
		return nil
	}
}

// testEdgeNodeDestroy verifies the EdgeNode has been destroyed.
func testEdgeNodeDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testAccProvider.Meta().(*apiclient.Zedcloudapi)

	// loop through the resources in state, verifying each EdgeNode is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_edgenode" {
			continue
		}

		// retrieve the EdgeNode by referencing it's state ID for API lookup
		params := edge_node_configuration.NewEdgeNodeConfigurationGetEdgeNodeParams()
		params.ID = rs.Primary.ID
		response, err := client.EdgeNodeConfiguration.EdgeNodeConfigurationGetEdgeNode(params, nil)
		if err == nil {
			if deviceConfig := response.GetPayload(); deviceConfig != nil && deviceConfig.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, EdgeNode (%s) still exists", deviceConfig.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the EdgeNode is destroyed
		_, ok := err.(*edge_node_configuration.EdgeNodeConfigurationGetEdgeNodeNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for EdgeNode (%s)", params.ID)
		}
	}
	return nil
}
