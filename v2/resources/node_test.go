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
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/node"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	"github.com/zededa/terraform-provider-zedcloud/v2/schemas"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestNode_Create_RequiredAttributesOnly(t *testing.T) {
	var got models.Node

	// input config
	inputPath := "node/create_required_only.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testNodeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testNodeExists("zedcloud_edgenode.required_only", &got),
					resource.TestCheckResourceAttr("zedcloud_edgenode.required_only", "name", "test_tf_provider-required_only"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.required_only", "title", "required_only-title"),
					resource.TestMatchResourceAttr(
						"zedcloud_edgenode.required_only",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
				),
			},
		},
	})
}

func TestNode_Create_AllAttributes(t *testing.T) {
	var gotCreated, gotUpdated, expectCreated, expectUpdated models.Node

	// input configs
	createPath := "node/create_all.tf"
	inputCreate := testhelper.MustGetTestInput(t, createPath)
	updatePath := "node/update_all.tf"
	inputUpdate := testhelper.MustGetTestInput(t, updatePath)

	// expected output
	createPath = "node/create_all.yaml"
	testhelper.MustGetExpectedOutput(t, createPath, &expectCreated)
	updatePath = "node/update_all.yaml"
	testhelper.MustGetExpectedOutput(t, updatePath, &expectUpdated)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testNodeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: inputCreate,
				Check: resource.ComposeTestCheckFunc(
					testNodeExists("zedcloud_edgenode.test_tf_provider", &gotCreated),
					resource.TestCheckResourceAttr("zedcloud_edgenode.test_tf_provider", "name", "test_tf_provider"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.test_tf_provider", "title", "test_tf_provider-title"),
					resource.TestMatchResourceAttr(
						"zedcloud_edgenode.test_tf_provider",
						"id",
						// Improved UUID regexp: matches any valid UUID (versions 1-5, RFC 4122)
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNodeAttributes(t, "create node", &gotCreated, &expectCreated),
				),
			},
			{
				Config: inputUpdate,
				Check: resource.ComposeTestCheckFunc(
					testNodeExists("zedcloud_edgenode.test_tf_provider", &gotUpdated),
					resource.TestCheckResourceAttr("zedcloud_edgenode.test_tf_provider", "name", "test_tf_provider"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.test_tf_provider", "title", "test_tf_provider-title"),
					resource.TestMatchResourceAttr(
						"zedcloud_edgenode.test_tf_provider",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNodeAttributes(t, "update node", &gotUpdated, &expectUpdated),
				),
			},
		},
	})
}

func TestNode_Create_WithAdapterSpecificNetwork(t *testing.T) {
	var gotCreated, gotUpdated, expectCreated, expectUpdated models.Node

	// input configs
	createPath := "node/create_with_adap_spec_net.tf"
	inputCreate := testhelper.MustGetTestInput(t, createPath)
	updatePath := "node/update_with_adap_spec_net.tf"
	inputUpdate := testhelper.MustGetTestInput(t, updatePath)

	// expected output
	createPath = "node/create_with_adap_spec_net.yaml"
	testhelper.MustGetExpectedOutput(t, createPath, &expectCreated)
	updatePath = "node/update_with_adap_spec_net.yaml"
	testhelper.MustGetExpectedOutput(t, updatePath, &expectUpdated)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testNodeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: inputCreate,
				Check: resource.ComposeTestCheckFunc(
					testNodeExists("zedcloud_edgenode.test_tf_dev_adap_spec_net", &gotCreated),
					resource.TestCheckResourceAttr("zedcloud_edgenode.test_tf_dev_adap_spec_net", "name", "test_tf_provider-dev_adap_spec_net"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.test_tf_dev_adap_spec_net", "title", "test_tf_provider-dev_adap_spec_net"),
					resource.TestMatchResourceAttr(
						"zedcloud_edgenode.test_tf_dev_adap_spec_net",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNodeAttributes(t, "create node", &gotCreated, &expectCreated),
				),
			},
			{
				Config: inputUpdate,
				Check: resource.ComposeTestCheckFunc(
					testNodeExists("zedcloud_edgenode.test_tf_dev_adap_spec_net", &gotUpdated),
					resource.TestCheckResourceAttr("zedcloud_edgenode.test_tf_dev_adap_spec_net", "name", "test_tf_provider-dev_adap_spec_net"),
					resource.TestCheckResourceAttr("zedcloud_edgenode.test_tf_dev_adap_spec_net", "title", "test_tf_dev_adap_spec_net"),
					resource.TestMatchResourceAttr(
						"zedcloud_edgenode.test_tf_dev_adap_spec_net",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNodeAttributes(t, "update node", &gotUpdated, &expectUpdated),
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
func testNodeAttributes(t *testing.T, testCase string, got, expected *models.Node) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var opts []cmp.Option
		opts = append(opts, cmpopts.IgnoreFields(
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
			"ModelID",
			// All fields bellow are ignored as the current API assigns some default values that are not
			// represented in the test YAMLs, and it is not possible to provide them,
			// because they are generated by the API and we can't control them in terraform.
			"EdgeSyncConfig",
			"Model",
			"Onboarding",
			"Project",
			"VlanAdapters",
		))
		// API and YAML unmarshal might change order of list elements so we need to ignore order when comparing
		interfacesEqual, reason := schemas.CompareSystemInterfaceList(got.Interfaces, expected.Interfaces)
		if !interfacesEqual {
			return fmt.Errorf("%s-%s:  sys interface mismatch: %s", t.Name(), testCase, reason)
		}

		if diff := cmp.Diff(*got, *expected, opts...); len(diff) != 0 {
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

		// if we use an http client with retries,
		// it overrrides EdgeNodeNotFound error
		if strings.Contains(err.Error(), "unexpected HTTP status 404 Not Found") {
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
