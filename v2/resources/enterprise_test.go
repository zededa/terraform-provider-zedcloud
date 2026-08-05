package resources

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/identity_access_management"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestEnterprise_Create(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping enterprise test for CI environment")
	}

	var got models.Enterprise
	var expected models.Enterprise

	// input config
	inputPath := "iam/enterprise.create.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "iam/enterprise.create.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testEnterpriseDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testEnterpriseExists("zedcloud_enterprise.test_tf_provider", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_enterprise.test_tf_provider",
						"id",
						regexp.MustCompile("^[0-9A-Za-z_=-]{28}$"),
					),
					testEnterpriseAttributes(t, &got, &expected),
				),
			},
		},
	})
}

// TestEnterpriseResource_InternalValidate checks the enterprise schema against the
// SDK's own rules. It runs without a controller, unlike the acceptance tests below,
// so it is the only automated guard on wiring like the ConflictsWith pair between
// white_labeling and attributes - a bad reference there panics at provider startup.
func TestEnterpriseResource_InternalValidate(t *testing.T) {
	if err := EnterpriseResource().InternalValidate(nil, true); err != nil {
		t.Errorf("enterprise resource schema is invalid: %v", err)
	}
}

// TestEnterprise_WhiteLabeling verifies the white_labeling block reaches the API as
// the enterprise attribute keys the console reads, and that removing the block clears
// them again. The attribute keys are the actual contract - the API rejects any key
// outside its allowlist, and the console looks up exactly these four - so this asserts
// the API-side map rather than only the Terraform state.
func TestEnterprise_WhiteLabeling(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping enterprise test for CI environment")
	}

	var got models.Enterprise

	resourceName := "zedcloud_enterprise.test_tf_provider_white_labeling"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testEnterpriseDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testhelper.MustGetTestInput(t, "iam/enterprise.white_labeling.tf"),
				Check: resource.ComposeTestCheckFunc(
					testEnterpriseExists(resourceName, &got),
					resource.TestCheckResourceAttr(resourceName, "white_labeling.0.primary_color", "#0A2540"),
					resource.TestCheckResourceAttr(resourceName, "white_labeling.0.secondary_color", "#00B3A4"),
					resource.TestCheckResourceAttr(resourceName, "white_labeling.0.logo_url", "https://acme.example.com/logo.svg"),
					resource.TestCheckResourceAttr(resourceName, "white_labeling.0.product_name", "Acme Edge"),
					// the block owns these keys, so they must not also show up in the
					// deprecated raw attributes map, where they would fight over the
					// same value on every plan
					resource.TestCheckNoResourceAttr(resourceName, "attributes.$ztag.entp.zui.ux.logo"),
					resource.TestCheckNoResourceAttr(resourceName, "attributes.$ztag.entp.zui.ux.product.name"),
					testEnterpriseWhiteLabelAttributes(t, &got, map[string]string{
						"$ztag.entp.zui.ux.color.primary":   "#0A2540",
						"$ztag.entp.zui.ux.color.secondary": "#00B3A4",
						"$ztag.entp.zui.ux.logo":            "https://acme.example.com/logo.svg",
						"$ztag.entp.zui.ux.product.name":    "Acme Edge",
					}),
				),
			},
			{
				Config: testhelper.MustGetTestInput(t, "iam/enterprise.white_labeling.cleared.tf"),
				Check: resource.ComposeTestCheckFunc(
					testEnterpriseExists(resourceName, &got),
					resource.TestCheckResourceAttr(resourceName, "white_labeling.#", "0"),
					testEnterpriseWhiteLabelAttributes(t, &got, map[string]string{}),
				),
			},
		},
	})
}

// TestEnterprise_WhiteLabelingControllerHostURL covers the host mapping that makes
// white-labeling actually reach the console: it is what GET /api/v1/cloud/environment
// matches the request host against to pick an enterprise, including a child one.
//
// controllerHostURL is write-only - the API stores it but never returns it - so
// without special handling state reads back empty and every plan wants to re-apply a
// value that is already stored. The framework fails a step whose apply leaves a
// non-empty plan, so this is what proves the field is usable at all.
func TestEnterprise_WhiteLabelingControllerHostURL(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping enterprise test for CI environment")
	}

	var got models.Enterprise

	resourceName := "zedcloud_enterprise.test_tf_provider_wl_host"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testEnterpriseDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testhelper.MustGetTestInput(t, "iam/enterprise.white_labeling.host.tf"),
				Check: resource.ComposeTestCheckFunc(
					testEnterpriseExists(resourceName, &got),
					// the configured host survives the read even though the API omits it
					resource.TestCheckResourceAttr(resourceName, "controller_host_url", "tf-wl-host.local.zededa.net"),
					resource.TestCheckResourceAttr(resourceName, "white_labeling.0.primary_color", "#112233"),
					testEnterpriseWhiteLabelAttributes(t, &got, map[string]string{
						"$ztag.entp.zui.ux.color.primary": "#112233",
						"$ztag.entp.zui.ux.product.name":  "Host Brand",
					}),
				),
			},
		},
	})
}

// TestEnterprise_WhiteLabelingLegacyAttributes protects configs written before the
// white_labeling block existed, which set the $ztag keys straight into attributes.
// The values must stay in attributes rather than migrating into the block: the test
// framework fails a step whose apply leaves a non-empty plan, so this is what proves
// such a config does not get a diff that never converges.
func TestEnterprise_WhiteLabelingLegacyAttributes(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping enterprise test for CI environment")
	}

	var got models.Enterprise

	resourceName := "zedcloud_enterprise.test_tf_provider_white_labeling_legacy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testEnterpriseDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testhelper.MustGetTestInput(t, "iam/enterprise.white_labeling.legacy.tf"),
				Check: resource.ComposeTestCheckFunc(
					testEnterpriseExists(resourceName, &got),
					// the values stay where the config put them
					resource.TestCheckResourceAttr(resourceName, "attributes.$ztag.entp.zui.ux.color.primary", "#0A2540"),
					resource.TestCheckResourceAttr(resourceName, "attributes.$ztag.entp.zui.ux.product.name", "Acme Edge"),
					// and are not also projected into the block, which would make both
					// fields claim the same value
					resource.TestCheckResourceAttr(resourceName, "white_labeling.#", "0"),
					testEnterpriseWhiteLabelAttributes(t, &got, map[string]string{
						"$ztag.entp.zui.ux.color.primary": "#0A2540",
						"$ztag.entp.zui.ux.product.name":  "Acme Edge",
					}),
				),
			},
		},
	})
}

// testEnterpriseWhiteLabelAttributes asserts the white-label entries of the
// enterprise attribute map as stored by the API.
func testEnterpriseWhiteLabelAttributes(t *testing.T, got *models.Enterprise, expected map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		whiteLabelKeys := []string{
			"$ztag.entp.zui.ux.color.primary",
			"$ztag.entp.zui.ux.color.secondary",
			"$ztag.entp.zui.ux.logo",
			"$ztag.entp.zui.ux.product.name",
		}
		for _, key := range whiteLabelKeys {
			if got.Attributes[key] != expected[key] {
				return fmt.Errorf(
					"%s: enterprise attribute %q: expected %q, got %q",
					t.Name(), key, expected[key], got.Attributes[key],
				)
			}
		}
		return nil
	}
}

// testEnterpriseExists retrieves the Enterprise and stores it in the provided *models.DeviceConfig.
func testEnterpriseExists(resourceName string, enterpriseModel *models.Enterprise) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Enterprise not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Enterprise ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the ApplicationInstance by referencing its state ID for API lookup
		params := config.NewIdentityAccessManagementGetEnterpriseParams()
		params.ID = rs.Primary.ID
		response, err := client.IdentityAccessManagement.IdentityAccessManagementGetEnterprise(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Enterprise (%s): %w", rs.Primary.ID, err)
		}

		enterprise := response.GetPayload()
		if enterprise == nil {
			return errors.New("could not get response payload in Enterprise existence test: enterprise is nil")
		}

		*enterpriseModel = *enterprise
		return nil
	}
}

// testEnterpriseAttributes verifies attributes are set correctly by Terraform
func testEnterpriseAttributes(t *testing.T, got, expected *models.Enterprise) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"HubspotID",
			"SfdcID",
			"AzureSubID",
			"ParentEntpID",
			"Revision",
		}
		opts := cmpopts.IgnoreFields(models.Enterprise{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testEnterpriseDestroy verifies the Enterprise has been destroyed.
func testEnterpriseDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Enterprise is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_enterprise" {
			continue
		}

		// retrieve the Enterprise by referencing its state ID for API lookup
		params := config.NewIdentityAccessManagementGetEnterpriseParams()
		params.ID = rs.Primary.ID
		response, err := client.IdentityAccessManagement.IdentityAccessManagementGetEnterprise(params, nil)
		if err == nil {
			if enterprise := response.GetPayload(); enterprise != nil && enterprise.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Enterprise (%s) still exists", enterprise.ID)
			}
			return nil
		}

		// if we use an http client with retries,
		// it overrrides IdentityAccessManagementGetEnterpriseNotFound error
		if strings.Contains(err.Error(), "unexpected HTTP status 404 Not Found") {
			return nil
		}

		// if the error is equivalent to 404 not found, the ApplicationInstance is destroyed
		_, ok := err.(*config.IdentityAccessManagementGetEnterpriseNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Enterprise (%s)", params.ID)
		}
	}
	return nil
}
