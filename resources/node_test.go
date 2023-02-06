package resources

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider/client"
	config "github.com/zededa/terraform-provider/client/edge_node_configuration"
	"github.com/zededa/terraform-provider/models"
)

func TestEdgeNode_Create_RequiredAttributesOnly(t *testing.T) {
	var got models.Node

	// input config
	inputPath := "edge_node/create_required_only.tf"
	input := mustGetTestInput(t, inputPath)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { checkEnv(t) },
		CheckDestroy: testEdgeNodeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testEdgeNodeExists("zedcloud_edgenode.required_only", &got),
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
}

func TestEdgeNode_Create_AllAttributes(t *testing.T) {
	var got models.Node
	var expected models.Node

	// input config
	inputPath := "edge_node/create_complete.tf"
	input := mustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "edge_node/create_complete.yaml"
	mustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { checkEnv(t) },
		CheckDestroy: testEdgeNodeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testEdgeNodeExists("zedcloud_edgenode.complete", &got),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "name", "complete"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "model_id", "2f716b55-2639-486c-9a2f-55a2e94146a6"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "title", "complete-title"),
					resource.TestMatchResourceAttr(
						"zedcloud_edgenode.complete",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testEdgeNodeAttributes(t, &got, &expected),
				),
			},
		},
	})

	// update all fields
	// create and update fields with custom logic and separate api requests
}

// testEdgeNodeExists retrieves the EdgeNode and stores it in the provided *models.DeviceConfig.
func testEdgeNodeExists(resourceName string, edgeNodeModel *models.Node) resource.TestCheckFunc {
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
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the EdgeNode by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Node.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch EdgeNode (%s): %w", rs.Primary.ID, err)
		}

		edgeNode := response.GetPayload()
		if edgeNode == nil {
			return errors.New("could not get response payload in EdgeNode existence test: deviceConfig is nil")
		}

		*edgeNodeModel = *edgeNode
		return nil
	}
}

// testEdgeNodeAttributes verifies attributes are set correctly by Terraform
func testEdgeNodeAttributes(t *testing.T, got, expected *models.Node) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if keyGot, keyExpect := strings.Split(got.Obkey, ":")[0], strings.Split(expected.Obkey, ":")[0]; keyGot != keyExpect {
			return fmt.Errorf("expect Obkey %s but got %+v", keyExpect, keyGot)
		}
		opts := cmpopts.IgnoreFields(
			models.Node{},
			"ID",
			"ProjectID",
			"Revision",
			"ResetCounter",
			"DebugKnob",
			"Onboarding",
			"Utype",
			"Obkey",
		)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testEdgeNodeDestroy verifies the EdgeNode has been destroyed.
func testEdgeNodeDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each EdgeNode is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_edgenode" {
			continue
		}

		// retrieve the EdgeNode by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Node.GetByID(params, nil)
		if err == nil {
			if deviceConfig := response.GetPayload(); deviceConfig != nil && deviceConfig.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, EdgeNode (%s) still exists", deviceConfig.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the EdgeNode is destroyed
		_, ok := err.(*config.EdgeNodeNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for EdgeNode (%s)", params.ID)
		}
	}
	return nil
}
