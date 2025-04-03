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
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/deployment"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestDeployment_Create(t *testing.T) {
	var gotCreated, expectCreated models.Deployment

	// input configs
	createPath := "deployment/create.tf"
	inputCreate := testhelper.MustGetTestInput(t, createPath)

	// expected output
	createPath = "deployment/create.yaml"
	testhelper.MustGetExpectedOutput(t, createPath, &expectCreated)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testProjectDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             inputCreate,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testDeploymentExists("zedcloud_deployment.tf_deployment", &gotCreated),
					resource.TestCheckResourceAttr("zedcloud_deployment.tf_deployment", "name", "test_tf_provider-deployment"),
					resource.TestCheckResourceAttr("zedcloud_deployment.tf_deployment", "title", "test_tf_provider-deployment"),
					resource.TestMatchResourceAttr(
						"zedcloud_deployment.tf_deployment",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testDeploymentAttributes(t, &gotCreated, &expectCreated),
				),
			},
		},
	})
}

// testProjectExists retrieves the Project and stores it in the provided *models.DeviceConfig.
func testDeploymentExists(resourceName string, deploymentModel *models.Deployment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Deployment not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Deployment ID is not set")
		}

		projectID := rs.Primary.Attributes["project_id"]
		if projectID == "" {
			return fmt.Errorf("Project ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the Project by referencing its state ID for API lookup
		params := config.NewGetByIDParams()
		params.ID = rs.Primary.ID
		params.ProjectID = projectID
		response, err := client.Deployment.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Deployment (%s) for Project(%s): %w", rs.Primary.ID, projectID, err)
		}

		deployment := response.GetPayload()
		if deployment == nil {
			return errors.New("could not get response payload in Deployment existence test: deployment is nil")
		}

		*deploymentModel = *deployment
		return nil
	}
}

// testProjectAttributes verifies attributes are set correctly by Terraform
func testDeploymentAttributes(t *testing.T, got, expected *models.Deployment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// expected values are defined in the testdata/deployment/create.yaml file
		// and we always expect the first element in the list of each type.
		expectedDevicePolicy := expected.DevicePolicies[0]
		expectedVolumeInstPolicy := expected.VolumeInstPolicies[0]
		expectedAppInstPolicy := expected.AppInstPolicies[0]
		expectedNetworkInstPolicy := expected.NetworkInstPolicies[0]

		if len(got.DevicePolicies) == 0 {
			return fmt.Errorf("DevicePolicy not found in the got Deployment object")
		}

		if len(got.VolumeInstPolicies) == 0 {
			return fmt.Errorf("VolumeInstPolicy not found in the got Deployment object")
		}

		if len(got.AppInstPolicies) == 0 {
			return fmt.Errorf("AppInstPolicy not found in the got Deployment object")
		}

		if expected.EdgeviewPolicy == nil {
			return fmt.Errorf("EdgeviewPolicy not found in the expected Deployment object")
		}

		if len(got.NetworkInstPolicies) == 0 {
			return fmt.Errorf("NetworkInstPolicy not found in the got Deployment object")
		}

		// check device policy
		err := compareDevicePolicies(t, got.DevicePolicies[0], expectedDevicePolicy)
		if err != nil {
			return err
		}
		// check volume instance policy
		err = compareVolumeInstPolicies(t, got.VolumeInstPolicies[0], expectedVolumeInstPolicy)
		if err != nil {
			return err
		}

		// check app instance policy
		err = compareAppInstPolicies(t, got.AppInstPolicies[0], expectedAppInstPolicy)
		if err != nil {
			return err
		}

		// check edgeview policy
		err = compareEdgeViewPolicies(t, got.EdgeviewPolicy, expected.EdgeviewPolicy)
		if err != nil {
			return err
		}

		//check network instance policy
		err = compareNetworkInstPolicies(t, got.NetworkInstPolicies[0], expectedNetworkInstPolicy)

		// check deployment
		IgnoreFields := []string{
			"ID",
			"ProjectID",
			"Revision",
			"VolumeInstPolicies",
			"AppInstPolicies",
			"DevicePolicies",
			"EdgeviewPolicy",
			"NetworkInstPolicies",
		}
		opts := cmpopts.IgnoreFields(
			models.Deployment{},
			IgnoreFields...,
		)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testProjectDestroy verifies the Project has been destroyed.
func testDeploymentDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Project is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_deployment" {
			continue
		}

		projectID := rs.Primary.Attributes["project_id"]
		if projectID == "" {
			return fmt.Errorf("Project ID is not set")
		}

		// retrieve the Project by referencing it's state ID for API lookup
		params := config.NewGetByIDParams()
		params.ID = rs.Primary.ID
		params.ProjectID = projectID
		response, err := client.Deployment.GetByID(params, nil)
		if err == nil {
			if deviceConfig := response.GetPayload(); deviceConfig != nil && deviceConfig.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Project (%s) still exists", deviceConfig.ID)
			}
			return nil
		}

		// if we use an http client with retries,
		// it overrrides GetByIDNotFound error
		if strings.Contains(err.Error(), "unexpected HTTP status 404 Not Found") {
			return nil
		}

		// if the error is equivalent to 404 not found, the Project is destroyed
		_, ok := err.(*config.GetByIDNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Project (%s)", params.ID)
		}
	}
	return nil
}

func comparePolicyMetaData(t *testing.T, got, expected *models.PolicyCommon) error {
	ignoredMetaFields := []string{
		"ID",
		"Revision",
	}
	metaOpts := cmpopts.IgnoreFields(models.PolicyCommon{}, ignoredMetaFields...)
	if diff := cmp.Diff(*got, *expected, metaOpts); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}
	return nil
}

func compareDevicePolicies(t *testing.T, got, expected *models.DevicePolicy) error {
	err := comparePolicyMetaData(t, got.MetaData, expected.MetaData)
	if err != nil {
		return err
	}

	ignoredPolicyFields := []string{
		"MetaData",
	}
	optsDevicePolicy := cmpopts.IgnoreFields(models.DevicePolicy{}, ignoredPolicyFields...)
	if diff := cmp.Diff(*got, *expected, optsDevicePolicy); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}
	return nil
}

func compareVolumeInstPolicies(t *testing.T, got, expected *models.VolumeInstPolicy) error {
	err := comparePolicyMetaData(t, got.MetaData, expected.MetaData)
	if err != nil {
		return err
	}

	ignoredVolInstConfigFields := []string{
		"MetaData",
	}
	optsVolumeInst := cmpopts.IgnoreFields(models.VolumeInstPolicy{}, ignoredVolInstConfigFields...)
	if diff := cmp.Diff(*got, *expected, optsVolumeInst); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}
	return nil
}

func compareAppInstPolicies(t *testing.T, got, expected *models.AppInstPolicy) error {
	err := comparePolicyMetaData(t, got.MetaData, expected.MetaData)
	if err != nil {
		return err
	}

	gotAppInstConfig := got.AppInstConfig
	gotImagesConfig := gotAppInstConfig.ManifestJSON.Images[0]
	ignoredImagesFields := []string{
		"Imageid",
	}
	optsImages := cmpopts.IgnoreFields(models.VMManifestImage{}, ignoredImagesFields...)
	if diff := cmp.Diff(*gotImagesConfig, *expected.AppInstConfig.ManifestJSON.Images[0], optsImages); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}

	ignoredManifestFields := []string{
		"Images",
	}
	optsManifest := cmpopts.IgnoreFields(models.VMManifest{}, ignoredManifestFields...)
	if diff := cmp.Diff(*gotAppInstConfig.ManifestJSON, *expected.AppInstConfig.ManifestJSON, optsManifest); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}

	ignoredAppInstConfigFields := []string{
		"BundleID",
		"ManifestJSON",
	}
	optsAppInst := cmpopts.IgnoreFields(models.AppInstConfig{}, ignoredAppInstConfigFields...)
	if diff := cmp.Diff(*gotAppInstConfig, *expected.AppInstConfig, optsAppInst); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}
	return nil
}

func compareEdgeViewPolicies(t *testing.T, got, expected *models.EdgeViewPolicy) error {
	err := comparePolicyMetaData(t, got.MetaData, expected.MetaData)
	if err != nil {
		return err
	}

	ignoredEdgeViewFields := []string{
		"MetaData",
	}
	optsEdgeView := cmpopts.IgnoreFields(models.EdgeViewPolicy{}, ignoredEdgeViewFields...)
	if diff := cmp.Diff(*got, *expected, optsEdgeView); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}
	return nil
}

func compareNetworkInstPolicies(t *testing.T, got, expected *models.NetworkInstPolicy) error {
	err := comparePolicyMetaData(t, got.MetaData, expected.MetaData)
	if err != nil {
		return err
	}

	ignoredNetworkInstFields := []string{
		"ID",
		"ProjectID",
		"Revision",
	}
	optsNetworkInst := cmpopts.IgnoreFields(models.NetworkInstance{}, ignoredNetworkInstFields...)
	if diff := cmp.Diff(*got.NetInstConfig, *expected.NetInstConfig, optsNetworkInst); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}
	return nil
}
