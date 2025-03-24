package resources

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	app_profile_service "github.com/zededa/terraform-provider-zedcloud/v2/client/app_profile_service"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestAppProfile_Create(t *testing.T) {
	var got, expected models.AppProfileRead

	// input config
	createPath := "app_profiles/create.tf"
	input := testhelper.MustGetTestInput(t, createPath)

	// expected output
	createPath = "app_profiles/create.yaml"
	testhelper.MustGetExpectedOutput(t, createPath, &expected)

	// terraform acceptance testcase
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testhelper.CheckEnv(t)
		},
		CheckDestroy: testAppProfileDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:             input,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					testAppProfileExists("zedcloud_app_profile.test_tf_provider", &got),
					testAppProfileAttributes(t, &got, &expected),
				),
			},
		},
	})
}

// testAppProfileExists retrieves the AppProfile and stores it in the provided *models.AppProfile.
func testAppProfileExists(name string, got *models.AppProfileRead) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("AppProfile not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return errors.New("AppProfile ID not set in Terraform state")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the AppProfile by ID
		params := app_profile_service.NewAppProfileServiceGetAppProfileParams()
		params.SetID(rs.Primary.ID)
		resp, err := client.AppProfile.AppProfileServiceGetAppProfile(params, nil)
		if err != nil {
			return fmt.Errorf("error fetching AppProfile with ID %s: %s", rs.Primary.ID, err)
		}

		// store the AppProfile in the provided *models.AppProfile
		*got = *resp.GetPayload()
		return nil
	}
}

// testAppProfileAttributes verifies attributes are set correctly by Terraform
func testAppProfileAttributes(t *testing.T, got, expected *models.AppProfileRead) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(expected.AppPolicies) == 0 || len(expected.NetworkPolicies) == 0 {
			return nil
		}
		expectedAppPolicies := expected.AppPolicies[0]
		expectedNetPolicies := expected.NetworkPolicies[0]

		if len(got.AppPolicies) == 0 {
			return fmt.Errorf("expected AppPolicies to be set")
		}
		if len(got.NetworkPolicies) == 0 {
			return fmt.Errorf("expected NetworkPolicies to be set")
		}
		err := compareProfileAppPolicies(t, got.AppPolicies[0], expectedAppPolicies)
		if err != nil {
			return err
		}

		err = compareProfileNetworkPolicies(t, got.NetworkPolicies[0], expectedNetPolicies)
		if err != nil {
			return err
		}

		ignoreFields := []string{
			"ID",
			"AppPolicies",
			"NetworkPolicies",
			"VolumePolicies",
		}
		opts := cmpopts.IgnoreFields(models.AppProfileRead{}, ignoreFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("unexpected diff: \n%s", diff)
		}
		return nil
	}
}

func compareProfilePolicyMetaData(t *testing.T, got, expected *models.ProfilePolicyCommon) error {
	ignoredMetaFields := []string{
		"ID",
	}
	metaOpts := cmpopts.IgnoreFields(models.ProfilePolicyCommon{}, ignoredMetaFields...)
	if diff := cmp.Diff(*got, *expected, metaOpts); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}
	return nil
}

func compareProfileAppPolicies(t *testing.T, got, expected *models.ProfileAppPolicy) error {
	err := compareProfilePolicyMetaData(t, got.MetaData, expected.MetaData)
	if err != nil {
		return err
	}

	gotAppConfig := got.AppConfig
	gotImagesConfig := gotAppConfig.ManifestJSON.Images[0]
	ignoredImagesFields := []string{
		"Imageid",
	}
	optsImages := cmpopts.IgnoreFields(models.VMManifestImage{}, ignoredImagesFields...)
	if diff := cmp.Diff(*gotImagesConfig, *expected.AppConfig.ManifestJSON.Images[0], optsImages); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}

	ignoredManifestFields := []string{
		"Images",
	}
	optsManifest := cmpopts.IgnoreFields(models.VMManifest{}, ignoredManifestFields...)
	if diff := cmp.Diff(*gotAppConfig.ManifestJSON, *expected.AppConfig.ManifestJSON, optsManifest); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}

	ignoredAppConfigFields := []string{
		"ManifestJSON",
	}
	optsAppConfig := cmpopts.IgnoreFields(models.ProfileAppConfig{}, ignoredAppConfigFields...)
	if diff := cmp.Diff(*gotAppConfig, *expected.AppConfig, optsAppConfig); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}
	return nil
}

func compareProfileVolumePolicies(t *testing.T, got, expected *models.ProfileVolumePolicy) error {
	err := compareProfilePolicyMetaData(t, got.MetaData, expected.MetaData)
	if err != nil {
		return err
	}

	ignoredVolConfigFields := []string{
		"MetaData",
	}
	optsVolumeConfig := cmpopts.IgnoreFields(models.ProfileVolumePolicy{}, ignoredVolConfigFields...)
	if diff := cmp.Diff(*got, *expected, optsVolumeConfig); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}
	return nil
}

func compareProfileNetworkPolicies(t *testing.T, got, expected *models.ProfileNetworkPolicy) error {
	err := compareProfilePolicyMetaData(t, got.MetaData, expected.MetaData)
	if err != nil {
		return err
	}

	ignoredNetworkFields := []string{}
	optsNetworkConfig := cmpopts.IgnoreFields(models.ProfileNetworkConfig{}, ignoredNetworkFields...)
	if diff := cmp.Diff(*got.NetworkConfig, *expected.NetworkConfig, optsNetworkConfig); len(diff) != 0 {
		return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
	}
	return nil
}

// testAppProfileDestroy verifies the AppProfile is destroyed.
func testAppProfileDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each AppProfile is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_app_profile" {
			continue
		}

		// retrieve the AppProfile by ID
		params := app_profile_service.NewAppProfileServiceGetAppProfileParams()
		params.SetID(rs.Primary.ID)
		resp, err := client.AppProfile.AppProfileServiceGetAppProfile(params, nil)
		if err == nil {
			if appProfileResp := resp.GetPayload(); appProfileResp != nil && appProfileResp.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, AppProfile with ID %s still exists", appProfileResp.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the app profile is destroyed
		_, ok := err.(*app_profile_service.AppProfileServiceGetAppProfileNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for app profile (%s)", params.ID)
		}
	}
	return nil
}
