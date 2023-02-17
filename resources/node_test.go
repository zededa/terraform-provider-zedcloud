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
	"github.com/zededa/terraform-provider/schemas"
)

func TestNode_Create_RequiredAttributesOnly(t *testing.T) {
	var got models.Node

	// input config
	inputPath := "node/create_required_only.tf"
	input := mustGetTestInput(t, inputPath)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { checkEnv(t) },
		CheckDestroy: testNodeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testNodeExists("zedcloud_edgenode.required_only", &got),
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

func TestNode_Create_AllAttributes(t *testing.T) {
	var got, expectCreated, expectUpdated models.Node

	// input configs
	createPath := "node/create_complete.tf"
	inputCreate := mustGetTestInput(t, createPath)
	updatePath := "node/update_complete.tf"
	inputUpdate := mustGetTestInput(t, updatePath)

	// expected output
	createPath = "node/create_complete.yaml"
	mustGetExpectedOutput(t, createPath, &expectCreated)
	updatePath = "node/update_complete.yaml"
	mustGetExpectedOutput(t, updatePath, &expectUpdated)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { checkEnv(t) },
		CheckDestroy: testNodeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: inputCreate,
				Check: resource.ComposeTestCheckFunc(
					testNodeExists("zedcloud_edgenode.complete", &got),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "name", "complete"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "model_id", "2f716b55-2639-486c-9a2f-55a2e94146a6"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "title", "complete-title"),
					resource.TestMatchResourceAttr(
						"zedcloud_edgenode.complete",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNodeAttributes(t, &got, &expectCreated),
				),
			},
			{
				Config: inputUpdate,
				Check: resource.ComposeTestCheckFunc(
					testNodeExists("zedcloud_edgenode.complete", &got),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "name", "complete"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "model_id", "2f716b55-2639-486c-9a2f-55a2e94146a6"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.complete", "title", "complete-title"),
					resource.TestMatchResourceAttr(
						"zedcloud_edgenode.complete",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNodeAttributes(t, &got, &expectUpdated),
				),
			},
		},
	})
}

// testNodeExists retrieves the Node and stores it in the provided *models.DeviceConfig.
func testNodeExists(resourceName string, edgeNodeModel *models.Node) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Node not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Node ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the Node by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Node.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Node (%s): %w", rs.Primary.ID, err)
		}

		edgeNode := response.GetPayload()
		if edgeNode == nil {
			return errors.New("could not get response payload in Node existence test: deviceConfig is nil")
		}

		*edgeNodeModel = *edgeNode
		return nil
	}
}

// testNodeAttributes verifies attributes are set correctly by Terraform
func testNodeAttributes(t *testing.T, got, expected *models.Node) resource.TestCheckFunc {
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
			"Interfaces",
		)
		// API and YAML unmarshal might change order of list elements so we need to ignore order when comparing
		if !schemas.CompareSystemInterfaceList(got.Interfaces, expected.Interfaces) {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), cmp.Diff(got.Interfaces, expected.Interfaces))
		}
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testNodeDestroy verifies the Node has been destroyed.
func testNodeDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Node is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_edgenode" {
			continue
		}

		// retrieve the Node by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Node.GetByID(params, nil)
		if err == nil {
			if deviceConfig := response.GetPayload(); deviceConfig != nil && deviceConfig.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Node (%s) still exists", deviceConfig.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the Node is destroyed
		_, ok := err.(*config.EdgeNodeNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Node (%s)", params.ID)
		}
	}
	return nil
}
