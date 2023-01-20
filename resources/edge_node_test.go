package resources

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiclient "github.com/zededa/terraform-provider/client"
	"github.com/zededa/terraform-provider/client/edge_node_configuration"
	"github.com/zededa/terraform-provider/models"
)

func TestAccGetEdgeNode(t *testing.T) {
	// var edgeNodeBefore models.DeviceConfig
	// var edgeNodeAfter models.DeviceConfig

	// // create with required and computed fields
	// requiredOnlyConfigPath := "edge_node/required_only.tf"
	// requiredOnlyConfig, err := getTestConfig(requiredOnlyConfigPath)
	// if err != nil {
	// 	t.Fatal(fmt.Sprintf("could not get testdata for %s", requiredOnlyConfigPath))
	// }
	// resource.Test(t, resource.TestCase{
	// 	PreCheck:     func() { testAccPreCheck(t) },
	// 	CheckDestroy: testEdgeNodeDestroy,
	// 	Providers:    testAccProviders,
	// 	Steps: []resource.TestStep{
	// 		{
	// 			Config: requiredOnlyConfig,
	// 			Check: resource.ComposeTestCheckFunc(
	// 				testEdgeNodeExists("zedcloud_edgenode.required_only", &edgeNodeBefore),
	// 				resource.TestCheckResourceAttr("zedcloud_edgenode.required_only", "name", "required_only"),
	// 				resource.TestCheckResourceAttr("zedcloud_edgenode.required_only", "model_id", "2f716b55-2639-486c-9a2f-55a2e94146a6"),
	// 				resource.TestCheckResourceAttr("zedcloud_edgenode.required_only", "title", "required_only-title"),
	// 				resource.TestMatchResourceAttr(
	// 					"zedcloud_edgenode.required_only",
	// 					"id",
	// 					regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
	// 				),
	// 			),
	// 		},
	// 	},
	// })

	// create with all fields
	var edgeNodeComplete models.DeviceConfig
	// var edgeNodeAfter models.DeviceConfig
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
				),
			},
		},
	})

	// update all fields
	// create and update fields with custom logic and separate api requests
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
