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
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/projects"
	depl_client "github.com/zededa/terraform-provider-zedcloud/v2/client/deployment"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestProject_Create_RequiredOnly(t *testing.T) {
	var gotCreated, expectCreated models.Tag

	// input configs
	createPath := "project/create_required_only.tf"
	inputCreate := testhelper.MustGetTestInput(t, createPath)

	// expected output
	createPath = "project/create_required_only.yaml"
	testhelper.MustGetExpectedOutput(t, createPath, &expectCreated)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testProjectDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: inputCreate,
				Check: resource.ComposeTestCheckFunc(
					testProjectExists("zedcloud_project.test_tf_provider", &gotCreated),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "name", "test_tf_provider"),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "title", "title"),
					resource.TestMatchResourceAttr(
						"zedcloud_project.test_tf_provider",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testProjectAttributes(t, &gotCreated, &expectCreated),
				),
			},
		},
	})
}

func TestProject_Create(t *testing.T) {
	var gotCreated, expectCreated models.Tag
	// var gotUpdated, expectUpdated models.Tag

	// input configs
	createPath := "project/create.tf"
	inputCreate := testhelper.MustGetTestInput(t, createPath)
	// updatePath := "project/update.tf"
	// inputUpdate := mustGetTestInput(t, updatePath)

	// expected output
	createPath = "project/create.yaml"
	testhelper.MustGetExpectedOutput(t, createPath, &expectCreated)
	// updatePath = "project/update.yaml"
	// mustGetExpectedOutput(t, updatePath, &expectUpdated)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testProjectDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: inputCreate,
				Check: resource.ComposeTestCheckFunc(
					testProjectExists("zedcloud_project.test_tf_provider", &gotCreated),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "name", "test_tf_provider"),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "title", "title"),
					resource.TestMatchResourceAttr(
						"zedcloud_project.test_tf_provider",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testProjectAttributes(t, &gotCreated, &expectCreated),
				),
			},
			// {
			// 	Config: inputUpdate,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testProjectExists("zedcloud_project.test_tf_provider", &gotUpdated),
			// 		resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "name", "test_tf_provider"),
			// 		resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "model_id", "2f716b55-2639-486c-9a2f-55a2e94146a6"),
			// 		resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider", "title", "test_tf_provider-title"),
			// 		resource.TestMatchResourceAttr(
			// 			"zedcloud_project.test_tf_provider",
			// 			"id",
			// 			regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
			// 		),
			// 		testProjectAttributes(t, &gotUpdated, &expectUpdated),
			// 	),
			// },
		},
	})
}

func TestProject_CreateWithPNACPolicy(t *testing.T) {
	var gotCreated, expectCreated models.Tag

	createPath := "project/create_with_pnac.tf"
	inputCreate := testhelper.MustGetTestInput(t, createPath)

	createPath = "project/create_with_pnac.yaml"
	testhelper.MustGetExpectedOutput(t, createPath, &expectCreated)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testProjectDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: inputCreate,
				Check: resource.ComposeTestCheckFunc(
					testProjectExists("zedcloud_project.test_tf_provider_pnac", &gotCreated),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider_pnac", "name", "test_tf_provider_pnac"),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider_pnac", "title", "Test PNAC Project"),
					resource.TestMatchResourceAttr(
						"zedcloud_project.test_tf_provider_pnac",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testProjectPNACAttributes(t, &gotCreated, &expectCreated),
				),
			},
		},
	})
}

// TestProject_PNAC_V1_CRUD tests the v1 flow: PNAC policy inline in zedcloud_project (TAG_TYPE_PROJECT).
// Steps:
//  1. Create project with PNAC policy
//  2. Verify the referenced CEP profile cannot be deleted while in use (direct API call)
//  3. Update PNAC policy (eap_identity + title)
//  4. Destroy
func TestProject_PNAC_V1_CRUD(t *testing.T) {
	var gotCreated, gotUpdated, expectCreated, expectUpdated models.Tag

	inputCreate := testhelper.MustGetTestInput(t, "project/create_with_pnac.tf")
	inputUpdate := testhelper.MustGetTestInput(t, "project/update_with_pnac.tf")

	testhelper.MustGetExpectedOutput(t, "project/create_with_pnac.yaml", &expectCreated)
	testhelper.MustGetExpectedOutput(t, "project/update_with_pnac.yaml", &expectUpdated)

	uuidRegexp := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testProjectDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// Create project with PNAC policy (v1 flow)
			{
				Config: inputCreate,
				Check: resource.ComposeTestCheckFunc(
					testProjectExists("zedcloud_project.test_tf_provider_pnac", &gotCreated),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider_pnac", "name", "test_tf_provider_pnac"),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider_pnac", "title", "Test PNAC Project"),
					resource.TestMatchResourceAttr("zedcloud_project.test_tf_provider_pnac", "id", uuidRegexp),
					testProjectPNACAttributes(t, &gotCreated, &expectCreated),
					// Verify the CEP profile cannot be deleted while referenced by the project's PNAC policy
					testCEPDeleteRestrictedWhileReferenced("zedcloud_cep_profile.test_tf_provider_pnac_cep"),
				),
			},
			// Update PNAC policy: change title and eap_identity
			{
				Config: inputUpdate,
				Check: resource.ComposeTestCheckFunc(
					testProjectExists("zedcloud_project.test_tf_provider_pnac", &gotUpdated),
					resource.TestCheckResourceAttr("zedcloud_project.test_tf_provider_pnac", "title", "Test PNAC Project Updated"),
					resource.TestMatchResourceAttr("zedcloud_project.test_tf_provider_pnac", "id", uuidRegexp),
					testProjectPNACAttributes(t, &gotUpdated, &expectUpdated),
				),
			},
		},
	})
}

// testCEPDeleteRestrictedWhileReferenced calls the CEP profile delete API directly and
// verifies that the API rejects the request because the profile is still in use.
// It does NOT delete the profile — on a successful restriction the profile stays intact.
func testCEPDeleteRestrictedWhileReferenced(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("CEP profile resource not found in state: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("CEP profile ID is empty")
		}

		client := testProvider.Meta().(*api_client.ZedcloudAPI)
		params := cep_client.DeleteParams()
		params.ID = rs.Primary.ID

		_, err := client.CertificateEnrollmentProfile.Delete(params, nil)
		if err == nil {
			// Delete succeeded — that means the restriction is not enforced.
			return fmt.Errorf("expected CEP profile delete to be restricted while referenced by PNAC policy, but it succeeded")
		}
		// Any error from the API means the delete was correctly blocked.
		return nil
	}
}

// TestProject_PNAC_V2_Create tests the v2 flow: PNAC policy via zedcloud_deployment device_policies
// on a TAG_TYPE_DEPLOYMENT project.
func TestProject_PNAC_V2_Create(t *testing.T) {
	var gotDeployment, expectDeployment models.Deployment

	inputCreate := testhelper.MustGetTestInput(t, "project/create_with_pnac_v2.tf")
	testhelper.MustGetExpectedOutput(t, "project/create_with_pnac_v2.yaml", &expectDeployment)

	uuidRegexp := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testDeploymentPNACDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             inputCreate,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testDeploymentPNACExists("zedcloud_deployment.test_tf_provider_pnac_v2_depl", &gotDeployment),
					resource.TestCheckResourceAttr("zedcloud_deployment.test_tf_provider_pnac_v2_depl", "name", "test_tf_provider_pnac_v2_depl"),
					resource.TestCheckResourceAttr("zedcloud_deployment.test_tf_provider_pnac_v2_depl", "title", "Test PNAC v2 Deployment"),
					resource.TestMatchResourceAttr("zedcloud_deployment.test_tf_provider_pnac_v2_depl", "id", uuidRegexp),
					testDeploymentPNACAttributes(t, &gotDeployment, &expectDeployment),
				),
			},
		},
	})
}

// testDeploymentPNACExists retrieves the deployment for the v2 PNAC test.
func testDeploymentPNACExists(resourceName string, deploymentModel *models.Deployment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Deployment not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Deployment ID is not set")
		}
		projectID := rs.Primary.Attributes["project_id"]
		if projectID == "" {
			return fmt.Errorf("Project ID is not set on deployment")
		}

		client := testProvider.Meta().(*api_client.ZedcloudAPI)
		params := depl_client.NewGetByIDParams()
		params.ID = rs.Primary.ID
		params.ProjectID = projectID
		response, err := client.Deployment.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Deployment (%s): %w", rs.Primary.ID, err)
		}
		d := response.GetPayload()
		if d == nil {
			return errors.New("nil payload in deployment existence check")
		}
		*deploymentModel = *d
		return nil
	}
}

// testDeploymentPNACAttributes verifies that the PNAC device policy is correctly set in the deployment.
func testDeploymentPNACAttributes(t *testing.T, got, expected *models.Deployment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(got.DevicePolicies) == 0 {
			return fmt.Errorf("%s: expected at least one device policy, got none", t.Name())
		}
		gotPolicy := got.DevicePolicies[0]
		expPolicy := expected.DevicePolicies[0]

		if gotPolicy.PortBasedNetworkAccessControlPolicy == nil {
			return fmt.Errorf("%s: expected PNAC policy in device_policies, got nil", t.Name())
		}

		opts := cmp.Options{
			cmpopts.IgnoreFields(models.DevicePolicy{}, "MetaData"),
			cmpopts.IgnoreFields(models.DevicePortBasedNetworkAccessControlPolicy{}, "CertificateEnrollmentID"),
			cmpopts.EquateEmpty(),
		}
		if diff := cmp.Diff(*gotPolicy, *expPolicy, opts); len(diff) != 0 {
			return fmt.Errorf("%s: PNAC device policy unexpected diff:\n%s", t.Name(), diff)
		}
		return nil
	}
}

// testDeploymentPNACDestroy verifies both the deployment and project are destroyed.
func testDeploymentPNACDestroy(s *terraform.State) error {
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	for _, rs := range s.RootModule().Resources {
		switch rs.Type {
		case "zedcloud_project":
			params := config.GetByIDParams()
			params.ID = rs.Primary.ID
			resp, err := client.Project.GetByID(params, nil)
			if err == nil {
				if p := resp.GetPayload(); p != nil && p.ID == rs.Primary.ID {
					return fmt.Errorf("destroy failed: project (%s) still exists", p.ID)
				}
			}
		case "zedcloud_cep_profile":
			cepParams := cep_client.GetByIDParams()
			cepParams.ID = rs.Primary.ID
			resp, err := client.CertificateEnrollmentProfile.GetByID(cepParams, nil)
			if err == nil {
				if p := resp.GetPayload(); p != nil && p.ID == rs.Primary.ID {
					return fmt.Errorf("destroy failed: CEP profile (%s) still exists", p.ID)
				}
			}
		}
	}
	return nil
}

// testProjectPNACAttributes verifies PNAC policy is set correctly on the project.
func testProjectPNACAttributes(t *testing.T, got, expected *models.Tag) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if got.PortBasedNetworkAccessControlPolicy == nil {
			return fmt.Errorf("%s: expected PNAC policy to be set, got nil", t.Name())
		}
		if expected.PortBasedNetworkAccessControlPolicy == nil {
			return nil
		}
		opts := cmp.Options{
			cmpopts.IgnoreFields(models.Policy{}, "ID", "Revision", "Name"),
			cmpopts.IgnoreFields(models.PortBasedNetworkAccessControlPolicy{}, "CertificateEnrollmentID"),
			cmpopts.EquateEmpty(),
		}
		if diff := cmp.Diff(
			*got.PortBasedNetworkAccessControlPolicy,
			*expected.PortBasedNetworkAccessControlPolicy,
			opts,
		); len(diff) != 0 {
			return fmt.Errorf("%s: PNAC policy unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testProjectExists retrieves the Project and stores it in the provided *models.DeviceConfig.
func testProjectExists(resourceName string, projectModel *models.Tag) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Project not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Project ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the Project by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Project.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Project (%s): %w", rs.Primary.ID, err)
		}

		project := response.GetPayload()
		if project == nil {
			return errors.New("could not get response payload in Project existence test: project is nil")
		}

		*projectModel = *project
		return nil
	}
}

// testProjectAttributes verifies attributes are set correctly by Terraform
func testProjectAttributes(t *testing.T, got, expected *models.Tag) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if expected.AttestationPolicy != nil {
			ignoredFields := []string{
				"ID",
				"Revision",
				"Name",
			}
			opts := cmpopts.IgnoreFields(models.Policy{}, ignoredFields...)
			if diff := cmp.Diff(*got.AttestationPolicy, *expected.AttestationPolicy, opts); len(diff) != 0 {
				return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
			}
		}
		if expected.EdgeviewPolicy != nil {
			ignoredFields := []string{
				"ID",
				"Revision",
				"Name",
			}
			opts := cmpopts.IgnoreFields(models.Policy{}, ignoredFields...)
			if diff := cmp.Diff(*got.EdgeviewPolicy, *expected.EdgeviewPolicy, opts); len(diff) != 0 {
				return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
			}
		}

		IgnoreFields := []string{
			"ID",
			"Revision",
			"AttestationPolicy",
			"EdgeviewPolicy",
		}
		opts := cmpopts.IgnoreFields(
			models.Tag{},
			IgnoreFields...,
		)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testProjectDestroy verifies the Project has been destroyed.
func testProjectDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Project is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_project" {
			continue
		}

		// retrieve the Project by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Project.GetByID(params, nil)
		if err == nil {
			if deviceConfig := response.GetPayload(); deviceConfig != nil && deviceConfig.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Project (%s) still exists", deviceConfig.ID)
			}
			return nil
		}

		// if we use an http client with retries,
		// it overrrides ProjectsGetByIDNotFound error
		if strings.Contains(err.Error(), "unexpected HTTP status 404 Not Found") {
			return nil
		}

		// if the error is equivalent to 404 not found, the Project is destroyed
		_, ok := err.(*config.ProjectsGetByIDNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Project (%s)", params.ID)
		}
	}
	return nil
}
