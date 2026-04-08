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
	cep_client "github.com/zededa/terraform-provider-zedcloud/v2/client/certificate_enrollment_profile"
	hw_client "github.com/zededa/terraform-provider-zedcloud/v2/client/hardware_model"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/node"
	project_client "github.com/zededa/terraform-provider-zedcloud/v2/client/projects"
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
			"BondAdapter",
		))
		// API and YAML unmarshal might change order of list elements so we need to ignore order when comparing
		interfacesEqual, reason := schemas.CompareSystemInterfaceList(got.Interfaces, expected.Interfaces)
		if !interfacesEqual {
			return fmt.Errorf("%s-%s:  sys interface mismatch: %s", t.Name(), testCase, reason)
		}
		bondAdaptersEqual, bondReason := schemas.CompareBondAdapterLists(got.BondAdapter, expected.BondAdapter)
		if !bondAdaptersEqual {
			return fmt.Errorf("%s-%s: bond_adapter mismatch: %s", t.Name(), testCase, bondReason)
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

// TestNode_PNAC_Onboard tests onboarding an edge node into a PNAC-enabled project.
// Steps:
//  1. Create a PNAC project (with CEP profile) and an edge node with eth0 having PNAC enabled.
//     Cross-verify that the CEP profile referenced by the project is valid (scep_url and csr_profile set).
//  2. Update: disable PNAC on eth0 and verify the interface flag is cleared.
func TestNode_PNAC_Onboard(t *testing.T) {
	var gotCreated, gotUpdated, expectCreated, expectUpdated models.Node

	inputCreate := testhelper.MustGetTestInput(t, "node/create_with_pnac.tf")
	inputUpdate := testhelper.MustGetTestInput(t, "node/update_with_pnac.tf")

	testhelper.MustGetExpectedOutput(t, "node/create_with_pnac.yaml", &expectCreated)
	testhelper.MustGetExpectedOutput(t, "node/update_with_pnac.yaml", &expectUpdated)

	uuidRegexp := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testNodePNACDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// Step 1: Create node in PNAC project with eth0 PNAC enabled.
			{
				Config: inputCreate,
				Check: resource.ComposeTestCheckFunc(
					testNodeExists("zedcloud_edgenode.test_tf_provider_node_pnac", &gotCreated),
					resource.TestCheckResourceAttr("zedcloud_edgenode.test_tf_provider_node_pnac", "name", "test_tf_provider_node_pnac"),
					resource.TestMatchResourceAttr("zedcloud_edgenode.test_tf_provider_node_pnac", "id", uuidRegexp),
					testNodeAttributes(t, "create node pnac", &gotCreated, &expectCreated),
					// Cross-verify: CEP profile referenced by the project is valid.
					testNodePNACCEPProfileValid("zedcloud_cep_profile.test_tf_provider_node_pnac_cep"),
				),
			},
			// Step 2: Update node — disable PNAC on eth0.
			{
				Config: inputUpdate,
				Check: resource.ComposeTestCheckFunc(
					testNodeExists("zedcloud_edgenode.test_tf_provider_node_pnac", &gotUpdated),
					resource.TestCheckResourceAttr("zedcloud_edgenode.test_tf_provider_node_pnac", "name", "test_tf_provider_node_pnac"),
					resource.TestMatchResourceAttr("zedcloud_edgenode.test_tf_provider_node_pnac", "id", uuidRegexp),
					testNodeAttributes(t, "update node pnac", &gotUpdated, &expectUpdated),
				),
			},
		},
	})
}

// testNodePNACCEPProfileValid fetches the CEP profile from the API and verifies it has a non-empty
// scep_url and csr_profile, confirming the profile is properly configured for PNAC use.
func testNodePNACCEPProfileValid(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("CEP profile resource not found in state: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("CEP profile ID is empty")
		}

		client := testProvider.Meta().(*api_client.ZedcloudAPI)
		params := cep_client.GetByIDParams()
		params.ID = rs.Primary.ID

		resp, err := client.CertificateEnrollmentProfile.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch CEP profile (%s): %w", rs.Primary.ID, err)
		}
		cep := resp.GetPayload()
		if cep == nil {
			return fmt.Errorf("CEP profile payload is nil for ID %s", rs.Primary.ID)
		}
		if cep.ScepURL == nil || *cep.ScepURL == "" {
			return fmt.Errorf("CEP profile %s has empty scep_url", rs.Primary.ID)
		}
		if cep.CsrProfile == nil {
			return fmt.Errorf("CEP profile %s has nil csr_profile", rs.Primary.ID)
		}
		if cep.CsrProfile.CommonName == nil || *cep.CsrProfile.CommonName == "" {
			return fmt.Errorf("CEP profile %s csr_profile has empty common_name", rs.Primary.ID)
		}
		return nil
	}
}

// testNodePNACDestroy verifies the node, project, CEP profile, brand, and model are all destroyed.
func testNodePNACDestroy(s *terraform.State) error {
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	for _, rs := range s.RootModule().Resources {
		switch rs.Type {
		case "zedcloud_edgenode":
			params := config.GetByIDParams()
			params.ID = rs.Primary.ID
			resp, err := client.Node.GetByID(params, nil)
			if err == nil {
				if n := resp.GetPayload(); n != nil && n.ID == rs.Primary.ID {
					return fmt.Errorf("destroy failed: node (%s) still exists", n.ID)
				}
			}
		case "zedcloud_project":
			params := project_client.GetByIDParams()
			params.ID = rs.Primary.ID
			resp, err := client.Project.GetByID(params, nil)
			if err == nil {
				if p := resp.GetPayload(); p != nil && p.ID == rs.Primary.ID {
					return fmt.Errorf("destroy failed: project (%s) still exists", p.ID)
				}
			}
		case "zedcloud_cep_profile":
			params := cep_client.GetByIDParams()
			params.ID = rs.Primary.ID
			resp, err := client.CertificateEnrollmentProfile.GetByID(params, nil)
			if err == nil {
				if p := resp.GetPayload(); p != nil && p.ID == rs.Primary.ID {
					return fmt.Errorf("destroy failed: CEP profile (%s) still exists", p.ID)
				}
			}
		case "zedcloud_brand":
			params := hw_client.NewHardwareModelGetHardwareBrandParams()
			params.ID = rs.Primary.ID
			resp, err := client.HardwareModel.HardwareModelGetHardwareBrand(params, nil)
			if err == nil {
				if b := resp.GetPayload(); b != nil && b.ID == rs.Primary.ID {
					return fmt.Errorf("destroy failed: brand (%s) still exists", b.ID)
				}
			}
		case "zedcloud_model":
			params := hw_client.NewHardwareModelGetHardwareModelParams()
			params.ID = rs.Primary.ID
			resp, err := client.HardwareModel.HardwareModelGetHardwareModel(params, nil)
			if err == nil {
				if m := resp.GetPayload(); m != nil && m.ID == rs.Primary.ID {
					return fmt.Errorf("destroy failed: model (%s) still exists", m.ID)
				}
			}
		}
	}
	return nil
}
