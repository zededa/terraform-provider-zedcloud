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
	config "github.com/zededa/terraform-provider-zedcloud/client/image"
	"github.com/zededa/terraform-provider-zedcloud/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/testing"
)

func TestImage_Create(t *testing.T) {
	var got models.Image
	var expected models.Image

	// input config
	inputPath := "image/create.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "image/create.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testImageDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testImageExists("zedcloud_image.test_image", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_image.test_image",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testImageAttributes(t, &got, &expected),
				),
			},
		},
	})
}

func testImageExists(resourceName string, imageModel *models.Image) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Image not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Image ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the Image by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Image.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Image (%s): %w", rs.Primary.ID, err)
		}

		image := response.GetPayload()
		if image == nil {
			return errors.New("could not get response payload in Image existence test: image is nil")
		}

		*imageModel = *image
		return nil
	}
}

// testImageAttributes verifies attributes are set correctly by Terraform
func testImageAttributes(t *testing.T, got, expected *models.Image) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"Revision",
			"ImageError",
		}
		opts := cmpopts.IgnoreFields(models.Image{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testImageInstanceDestroy verifies the Image has been destroyed.
func testImageDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Image is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_image" {
			continue
		}

		// retrieve the Image by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Image.GetByID(params, nil)
		if err == nil {
			if image := response.GetPayload(); image != nil && image.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Image (%s) still exists", image.ID)
			}
			return nil
		}

		// if the error is equivalent to 404 not found, the Image is destroyed
		_, ok := err.(*config.GetImageNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Image (%s)", params.ID)
		}
	}
	return nil
}
