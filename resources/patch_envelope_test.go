package resources

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/client"
	config "github.com/zededa/terraform-provider-zedcloud/client/patch_envelope"
	"github.com/zededa/terraform-provider-zedcloud/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/testing"
)

func TestPatchEnvelope_Create(t *testing.T) {
	var got models.PatchEnvelope
	var expected models.PatchEnvelope

	// input config
	inputPath := "patch_envelope/create.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "patch_envelope/create.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testPatchEnvelopeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testPatchEnvelopeExists("zedcloud_patch_envelope.test_tf_provider", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_patch_envelope.test_tf_provider",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_patch_envelope.test_tf_provider",
						"project_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testPatchEnvelopeAttributes(t, &got, &expected),
				),
			},
		},
	})
}

// testPatchEnvelopeExists retrieves the PatchEnvelope and stores it in the provided *models.PatchEnvelope.
func testPatchEnvelopeExists(resourceName string, patchEnvelopeModel *models.PatchEnvelope) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("PatchEnvelope not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("PatchEnvelope ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the PatchEnvelope by referencing its state ID for API lookup
		params := config.NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.PatchEnvelope.PatchEnvelopeConfigurationGetPatchEnvelopeByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch PatchEnvelope (%s): %w", rs.Primary.ID, err)
		}

		patchEnv := response.GetPayload()
		if patchEnv == nil {
			return errors.New("could not get response payload in PatchEnvelope existence test: patch_envelope is nil")
		}

		*patchEnvelopeModel = *patchEnv
		return nil
	}
}

// testPatchEnvelopeAttributes verifies attributes are set correctly by Terraform
func testPatchEnvelopeAttributes(t *testing.T, got, expected *models.PatchEnvelope) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		ignoredFields := []string{
			"ID",
			"ProjectID",
			"Artifacts",
			"Revision",
		}
		opts := cmpopts.IgnoreFields(models.PatchEnvelope{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}

		ignoredFields = []string{
			"ID",
			"BinaryArtifact",
			"Base64Artifact",
		}
		opts = cmpopts.IgnoreFields(models.BinaryArtifact{}, ignoredFields...)
		if diff := cmp.Diff(*got.Artifacts[0], *expected.Artifacts[0], opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}

		ignoredFields = []string{}
		opts = cmpopts.IgnoreFields(models.InlineOpaqueBase64Data{}, ignoredFields...)
		if diff := cmp.Diff(*got.Artifacts[0].Base64Artifact, *expected.Artifacts[0].Base64Artifact, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		ignoredFields = []string{
			"ImageID",
		}
		opts = cmpopts.IgnoreFields(models.ExternalOpaqueBinaryBlob{}, ignoredFields...)
		if diff := cmp.Diff(*got.Artifacts[1].BinaryArtifact, *expected.Artifacts[1].BinaryArtifact, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}

		return nil
	}
}

// testPatchEnvelopeInstanceDestroy verifies the PatchEnvelope has been destroyed.
func testPatchEnvelopeDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each PatchEnvelope is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_patch_envelope" {
			continue
		}

		// retrieve the PatchEnvelope by referencing it's state ID for API lookup
		params := config.NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.PatchEnvelope.PatchEnvelopeConfigurationGetPatchEnvelopeByID(params, nil)
		if err == nil {
			if patchEnv := response.GetPayload(); patchEnv != nil && patchEnv.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, PatchEnvelope (%s) still exists", patchEnv.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the PatchEnvelope is destroyed
		_, ok := err.(*config.PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for PatchEnvelope (%s)", params.ID)
		}
	}
	return nil
}
