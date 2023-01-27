package resources

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider/client"
	config "github.com/zededa/terraform-provider/client/edge_network_configuration"
	"github.com/zededa/terraform-provider/models"
)

func TestAccGetNetwork(t *testing.T) {
	var networkBefore models.Network
	// var networkAfter models.DeviceConfig

	// create with required and computed fields
	requiredOnlyConfigPath := "network/required_only.tf"
	requiredOnlyConfig, err := getTestConfig(requiredOnlyConfigPath)
	if err != nil {
		t.Fatal(fmt.Sprintf("could not get testdata for %s", requiredOnlyConfigPath))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testNetworkDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: requiredOnlyConfig,
				Check: resource.ComposeTestCheckFunc(
					testNetworkExists("zedcloud_network.required_only", &networkBefore),
					resource.TestCheckResourceAttr("zedcloud_network.required_only", "name", "zedcloud_network.required_only.name"),
					resource.TestCheckResourceAttr("zedcloud_network.required_only", "title", "zedcloud_network.required_only.title"),
					resource.TestCheckResourceAttr("zedcloud_network.required_only", "project_id", "2f716b55-2639-486c-9a2f-55a2e94146a6"),
					resource.TestMatchResourceAttr(
						"zedcloud_network.required_only",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
				),
			},
		},
	})

	// create with all fields
	var networkComplete models.Network
	completeConfigPath := "network/complete.tf"
	completeConfig, err := getTestConfig(completeConfigPath)
	if err != nil {
		t.Fatal(fmt.Sprintf("could not get testdata for %s", completeConfigPath))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testNetworkDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: completeConfig,
				Check: resource.ComposeTestCheckFunc(
					testNetworkExists("zedcloud_network.complete", &networkComplete),
					resource.TestCheckResourceAttr("zedcloud_network.complete", "name", "zedcloud_network.complete.name"),
					resource.TestCheckResourceAttr("zedcloud_network.complete", "title", "zedcloud_network.complete.title"),
					resource.TestCheckResourceAttr("zedcloud_network.complete", "project_id", "2f716b55-2639-486c-9a2f-55a2e94146a6"),
					resource.TestMatchResourceAttr(
						"zedcloud_network.complete",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNetworkAttributes(&networkComplete),
				),
			},
		},
	})

	// update all fields
	// create and update fields with custom logic and separate api requests
}

// testNetworAccPreCheck validates the necessary test API keys exist in the testing environment
func testNetworAccPreCheck(t *testing.T) {
	if v := os.Getenv("TF_CLI_CONFIG_FILE"); v == "" {
		t.Fatal("TF_CLI_CONFIG_FILE must be set for acceptance tests, it should contain the dev_overrides config that points to local instance of the provider")
	}
	if v := os.Getenv("TF_VAR_zedcloud_token"); v == "" {
		t.Fatal("TF_VAR_zedcloud_token must be set for acceptance tests to access the zedcloud API")
	}
}

// testNetworkExists retrieves the Network and stores it in the provided *models.DeviceConfig.
func testNetworkExists(resourceName string, networkModel *models.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Network not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Network ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the Network by referencing its state ID for API lookup
		// params := config.GetByIDParams()
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Network.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Network (%s): %w", rs.Primary.ID, err)
		}

		network := response.GetPayload()
		if network == nil {
			return errors.New("could not get response payload in Network existence test: network is nil")
		}

		// store the resulting Network config in the *models.DeviceConfig variable
		*networkModel = *network
		return nil
	}
}

// testNetworkAttributes verifies attributes are set correctly by Terraform
func testNetworkAttributes(device *models.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return nil
	}
}

// testNetworkDestroy verifies the Network has been destroyed.
func testNetworkDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Network is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_network" {
			continue
		}

		// retrieve the Network by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Network.GetByID(params, nil)
		if err == nil {
			if network := response.GetPayload(); network != nil && network.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Network (%s) still exists", network.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the Network is destroyed
		_, ok := err.(*config.NetworkNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Network (%s)", params.ID)
		}
	}
	return nil
}
