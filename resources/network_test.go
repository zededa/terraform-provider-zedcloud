package resources

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider/client"
	config "github.com/zededa/terraform-provider/client/edge_network_configuration"
	"github.com/zededa/terraform-provider/models"
)

func TestCreateNetwork_RequiredAttributesOnly(t *testing.T) {
	var got models.Network
	var expected models.Network

	// input config
	inputPath := "network/create_required_only.tf"
	input := mustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "network/create_required_only_expected.yaml"
	mustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { checkEnv(t) },
		CheckDestroy: testNetworkDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testNetworkExists("zedcloud_network.required_only", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_network.required_only",
						"project_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_network.required_only",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNetworkAttributes(t, &got, &expected),
				),
			},
		},
	})
}

// func TestCreateNetwork_AllAttributes_WithProxy(t *testing.T) {
// 	var got models.Network
// 	var expected models.Network

// 	// input config
// 	inputPath := "network/create_complete_with_proxy.tf"
// 	input := mustGetTestInput(t, inputPath)

// 	// expected output
// 	expectedPath := "network/create_complete_with_proxy.yaml"
// 	mustGetExpectedOutput(t, expectedPath, &expected)

// 	// terraform acceptance test case
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { checkEnv(t) },
// 		CheckDestroy: testNetworkDestroy,
// 		Providers:    testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: input,
// 				Check: resource.ComposeTestCheckFunc(
// 					testNetworkExists("zedcloud_network.complete_with_proxy", &got),
// 					resource.TestCheckResourceAttr("zedcloud_network.complete_with_proxy", "name", "zedcloud_network.complete_with_proxy.name"),
// 					resource.TestCheckResourceAttr("zedcloud_network.complete_with_proxy", "title", "zedcloud_network.complete_with_proxy.title"),
// 					resource.TestMatchResourceAttr(
// 						"zedcloud_network.complete_with_proxy",
// 						"project_id",
// 						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
// 					),
// 					resource.TestMatchResourceAttr(
// 						"zedcloud_network.complete_with_proxy",
// 						"id",
// 						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
// 					),
// 					testNetworkAttributes(t, &got, &expected),
// 				),
// 			},
// 		},
// 	})
// }

// func TestCreateNetwork_AllAttributes_WithPac(t *testing.T) {
// 	var got models.Network
// 	var expected models.Network

// 	// input config
// 	inputPath := "network/create_complete_with_pac.tf"
// 	input := mustGetTestInput(t, inputPath)

// 	// expected output
// 	expectedPath := "network/create_complete_with_pac.yaml"
// 	mustGetExpectedOutput(t, expectedPath, &expected)

// 	// terraform acceptance test case
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { checkEnv(t) },
// 		CheckDestroy: testNetworkDestroy,
// 		Providers:    testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: input,
// 				Check: resource.ComposeTestCheckFunc(
// 					testNetworkExists("zedcloud_network.complete_with_pac", &got),
// 					resource.TestMatchResourceAttr(
// 						"zedcloud_network.complete_with_pac",
// 						"project_id",
// 						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
// 					),
// 					resource.TestMatchResourceAttr(
// 						"zedcloud_network.complete_with_pac",
// 						"id",
// 						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
// 					),
// 					testNetworkAttributes(t, &got, &expected),
// 				),
// 			},
// 		},
// 	})
// }

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
func testNetworkAttributes(t *testing.T, got, expected *models.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"Revision",
		}
		if expected.Proxy != nil && expected.Proxy.NetworkProxyCerts == nil {
			ignoredFields = append(ignoredFields, "Proxy.NetworkProxyCerts")
		}
		opts := cmpopts.IgnoreFields(models.Network{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
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
