package resources

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider/client"
	config "github.com/zededa/terraform-provider/client/edge_network_instance_configuration"
	"github.com/zededa/terraform-provider/models"
	"github.com/zededa/terraform-provider/schemas"
)

func TestNetworkInstance_Create_RequiredAttributesOnly(t *testing.T) {
	var got models.NetworkInstance
	var expected models.NetworkInstance

	// input config
	inputPath := "network_instance/create_required_only.tf"
	input := mustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "network_instance/create_required_only_expected.yaml"
	mustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { checkEnv(t) },
		CheckDestroy: testNetworkInstanceDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testNetworkInstanceExists("zedcloud_network_instance.required_only", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_network_instance.required_only",
						"project_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_network_instance.required_only",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNetworkInstanceAttributes(t, &got, &expected),
				),
			},
		},
	})
}

func TestNetworkInstance_Create_Complete(t *testing.T) {
	var got models.NetworkInstance
	var expected models.NetworkInstance

	// input config
	inputPath := "network_instance/create_complete.tf"
	input := mustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "network_instance/create_complete_expected.yaml"
	mustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { checkEnv(t) },
		CheckDestroy: testNetworkInstanceDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testNetworkInstanceExists("zedcloud_network_instance.complete", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_network_instance.complete",
						"project_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_network_instance.complete",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_network_instance.complete",
						"device_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNetworkInstanceAttributes(t, &got, &expected),
				),
			},
		},
	})
}

func testNetworkInstanceExists(resourceName string, networkModel *models.NetworkInstance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Network Instance not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Network Instance ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the Network by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.NetworkInstance.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Network Instance (%s): %w", rs.Primary.ID, err)
		}

		network := response.GetPayload()
		if network == nil {
			return errors.New("could not get response payload in Network Instance test: network is nil")
		}

		*networkModel = *network
		return nil
	}
}

// testNetworkInstanceAttributes verifies attributes are set correctly by Terraform
func testNetworkInstanceAttributes(t *testing.T, got, expected *models.NetworkInstance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"Revision",
			"IP",
			"DNSList",
			"DeviceID",
			"ProjectID",
		}
		// API and YAML unmarshal might change order of list elements so we need to use a compare function
		if !schemas.CompareDNSLists(got.DNSList, expected.DNSList) {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), cmp.Diff(got.DNSList, expected.DNSList))
		}
		opts := cmpopts.IgnoreFields(models.NetworkInstance{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testNetworkInstanceDestroy verifies the Network has been destroyed.
func testNetworkInstanceDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Network is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_network_instance" {
			continue
		}

		spew.Dump(rs)

		// retrieve the Network by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.NetworkInstance.GetByID(params, nil)
		if err == nil {
			if network := response.GetPayload(); network != nil && network.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Network Instance (%s) still exists", network.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the Network is destroyed
		_, ok := err.(*config.GetEdgeNetworkInstanceNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Network Instance (%s)", params.ID)
		}
	}
	return nil
}
