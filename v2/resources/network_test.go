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
	config "github.com/zededa/terraform-provider-zedcloud/v2/client/network"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	"github.com/zededa/terraform-provider-zedcloud/v2/schemas"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func TestNetwork_Create_RequiredAttributesOnly(t *testing.T) {
	var got models.Network
	var expected models.Network

	// input config
	inputPath := "network/create_required_only.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "network/create_required_only_expected.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testNetworkDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testNetworkExists("zedcloud_network.required_only", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_network.required_only",
						"project_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_network.required_only",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNetworkAttributes(t, &got, &expected),
				),
			},
		},
	})
}

func TestNetwork_Create_AllAttributes_WithProxy(t *testing.T) {
	var got models.Network
	var expected models.Network

	// input config
	inputPath := "network/create_complete_with_proxy.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "network/create_complete_with_proxy.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testNetworkDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testNetworkExists("zedcloud_network.complete_with_proxy", &got),
					resource.TestCheckResourceAttr("zedcloud_network.complete_with_proxy", "name", "zedcloud_network.complete_with_proxy.name"),
					resource.TestCheckResourceAttr("zedcloud_network.complete_with_proxy", "title", "zedcloud_network.complete_with_proxy.title"),
					resource.TestMatchResourceAttr(
						"zedcloud_network.complete_with_proxy",
						"project_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_network.complete_with_proxy",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNetworkAttributes(t, &got, &expected),
				),
			},
		},
	})
}

func TestNetwork_Create_AllAttributes_WithPac(t *testing.T) {
	var got models.Network
	var expected models.Network

	// input config
	inputPath := "network/create_complete_with_pac.tf"
	input := testhelper.MustGetTestInput(t, inputPath)

	// expected output
	expectedPath := "network/create_complete_with_pac.yaml"
	testhelper.MustGetExpectedOutput(t, expectedPath, &expected)

	// terraform acceptance test case
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testNetworkDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: input,
				Check: resource.ComposeTestCheckFunc(
					testNetworkExists("zedcloud_network.complete_with_pac", &got),
					resource.TestMatchResourceAttr(
						"zedcloud_network.complete_with_pac",
						"project_id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					resource.TestMatchResourceAttr(
						"zedcloud_network.complete_with_pac",
						"id",
						regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"),
					),
					testNetworkAttributes(t, &got, &expected),
				),
			},
		},
	})
}

func testNetworkExists(resourceName string, networkModel *models.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// retrieve the resource by name from state
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Network not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Network ID is not set")
		}

		// retrieve the client established in Provider configuration
		client := testProvider.Meta().(*api_client.ZedcloudAPI)

		// retrieve the Network by referencing its state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Network.GetByID(params, nil)
		if err != nil {
			return fmt.Errorf("could not fetch Network (%s): %w", rs.Primary.ID, err)
		}

		network := response.GetPayload()
		if network == nil {
			return errors.New("could not get response payload in Network existence test: network is nil")
		}

		*networkModel = *network
		return nil
	}
}

// testNetworkAttributes verifies attributes are set correctly by Terraform
func testNetworkAttributes(t *testing.T, got, expected *models.Network) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ignoredFields := []string{
			"ID",
			"Revision",
			"ProjectID",
		}
		if expected.Proxy != nil && expected.Proxy.NetworkProxyCerts == nil {
			ignoredFields = append(ignoredFields, "Proxy.NetworkProxyCerts")
		}
		// API and YAML unmarshal might change order of list elements so we need a special comparison
		if !schemas.CompareDNSLists(got.DNSList, expected.DNSList) {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), cmp.Diff(got.DNSList, expected.DNSList))
		}
		// API and YAML unmarshal might change order of list elements so we need a special comparison
		if !schemas.CompareProxyLists([]*models.Proxy{got.Proxy}, []*models.Proxy{expected.Proxy}) {
			return fmt.Errorf("%s: unexpected diff in proxy: \n%s", t.Name(), cmp.Diff(got.Proxy, expected.Proxy))
		}
		if expected.Proxy != nil && expected.Proxy.Proxies != nil {
			ignoredFields = append(ignoredFields, "Proxy.Proxies")
		}
		opts := cmpopts.IgnoreFields(models.Network{}, ignoredFields...)
		if diff := cmp.Diff(*got, *expected, opts); len(diff) != 0 {
			return fmt.Errorf("%s: unexpected diff: \n%s", t.Name(), diff)
		}
		return nil
	}
}

// testNetworkInstanceDestroy verifies the Network has been destroyed.
func testNetworkDestroy(s *terraform.State) error {
	// retrieve the client established in Provider configuration
	client := testProvider.Meta().(*api_client.ZedcloudAPI)

	// loop through the resources in state, verifying each Network is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_network" {
			continue
		}

		// retrieve the Network by referencing it's state ID for API lookup
		params := config.GetByIDParams()
		params.ID = rs.Primary.ID
		response, err := client.Network.GetByID(params, nil)
		if err == nil {
			if network := response.GetPayload(); network != nil && network.ID == rs.Primary.ID {
				return fmt.Errorf("destroy failed, Network (%s) still exists", network.ID)
			}
			return nil
		}

		// if we use an http client with retries,
		// it overrrides NetworkNotFound error
		if strings.Contains(err.Error(), "unexpected HTTP status 404 Not Found") {
			return nil
		}

		// if the error is equivalent to 404 not found, the Network is destroyed
		_, ok := err.(*config.NetworkNotFound)
		if !ok {
			return fmt.Errorf("destroy failed, expect status code 404 for Network (%s)", params.ID)
		}
	}
	return nil
}
