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
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestCEPProfile_CRUD(t *testing.T) {
	var gotCreated, gotUpdated, expectCreated, expectUpdated models.CEPCommonSCEPProfile

	// input configs
	inputCreate := testhelper.MustGetTestInput(t, "cep_profile/create.tf")
	inputUpdate := testhelper.MustGetTestInput(t, "cep_profile/update.tf")

	// expected outputs
	testhelper.MustGetExpectedOutput(t, "cep_profile/create.yaml", &expectCreated)
	testhelper.MustGetExpectedOutput(t, "cep_profile/update.yaml", &expectUpdated)

	uuidRegexp := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testCEPProfileDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// Create
			{
				Config: inputCreate,
				Check: resource.ComposeTestCheckFunc(
					testCEPProfileExists("zedcloud_cep_profile.test_tf_provider", &gotCreated),
					resource.TestCheckResourceAttr("zedcloud_cep_profile.test_tf_provider", "name", "test_tf_provider_cep"),
					resource.TestCheckResourceAttr("zedcloud_cep_profile.test_tf_provider", "title", "Test CEP Profile"),
					resource.TestMatchResourceAttr("zedcloud_cep_profile.test_tf_provider", "id", uuidRegexp),
					testCEPProfileAttributes(t, &gotCreated, &expectCreated),
				),
			},
			// Update
			{
				Config: inputUpdate,
				Check: resource.ComposeTestCheckFunc(
					testCEPProfileExists("zedcloud_cep_profile.test_tf_provider", &gotUpdated),
					resource.TestCheckResourceAttr("zedcloud_cep_profile.test_tf_provider", "name", "test_tf_provider_cep"),
					resource.TestCheckResourceAttr("zedcloud_cep_profile.test_tf_provider", "title", "Test CEP Profile Updated"),
					resource.TestMatchResourceAttr("zedcloud_cep_profile.test_tf_provider", "id", uuidRegexp),
					testCEPProfileAttributes(t, &gotUpdated, &expectUpdated),
				),
			},
			// Import
			{
				ResourceName:      "zedcloud_cep_profile.test_tf_provider",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"secret",
				},
			},
		},
	})
}

// testCEPProfileExists retrieves the CEP profile by ID and stores it in the provided model.
func testCEPProfileExists(resourceName string, cepModel *models.CEPCommonSCEPProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("CEP profile not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("CEP profile ID is not set")
		}

		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		params := cep_client.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.CertificateEnrollmentProfile.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch CEP profile (%s): %w", rs.Primary.ID, err)
		}

		profile := response.GetPayload()
		if profile == nil {
			return errors.New("could not get response payload in CEP profile existence test: profile is nil")
		}

		*cepModel = *profile
		return nil
	}
}

// testCEPProfileAttributes verifies the CEP profile attributes match the expected values.
func testCEPProfileAttributes(t *testing.T, got, expected *models.CEPCommonSCEPProfile) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"Revision",
			"EnterpriseID",
			"ProjectsList",
			"Secret",
			"CaCertPem",
			"CryptoKey",
			"EncryptedSecrets",
		}
		opts := cmp.Options{
			cmpopts.IgnoreFields(models.CEPCommonSCEPProfile{}, ignoredFields...),
			cmpopts.EquateEmpty(),
		}
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testCEPProfileDestroy verifies the CEP profile has been destroyed.
func testCEPProfileDestroy(s *terraform.State) error {
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_cep_profile" {
			continue
		}

		params := cep_client.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.CertificateEnrollmentProfile.GetByID(params, nil)
		if err == nil {
			if profile := response.GetPayload(); profile != nil && profile.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, CEP profile (%s) still exists", profile.ID)
			}
			return nil
		}

		if strings.Contains(err.Error(), "unexpected HTTP status 404 Not Found") {
			return nil
		}

		_, ok := err.(*cep_client.CEPGetByIDNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for CEP profile (%s)", params.ID)
		}
	}
	return nil
}
