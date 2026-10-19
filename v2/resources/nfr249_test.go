package resources

// Acceptance tests for NFR-249: Edge app tags, project / edge node / deployment
// PCR policies, and in-place zedcloud_deployment updates. They need a
// controller (TF_ACC=1, TF_VAR_zedcloud_url, TF_VAR_zedcloud_token) but no
// onboarded device: the PCR policy API works on node objects that never
// checked in.
//
// The fixtures in testdata/nfr249 carry one #@<NAME>_EXTRA@ marker each; the
// steps below fill it with the attributes under test.

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	app_client "github.com/zededa/terraform-provider-zedcloud/v2/client/application"
	depl_client "github.com/zededa/terraform-provider-zedcloud/v2/client/deployment"
	node_client "github.com/zededa/terraform-provider-zedcloud/v2/client/node"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

func nfr249Config(t *testing.T, file, marker, extra string) string {
	t.Helper()
	base := testhelper.MustGetTestInput(t, "nfr249/"+file)
	if !strings.Contains(base, marker) {
		t.Fatalf("fixture nfr249/%s has no %s marker", file, marker)
	}
	return strings.Replace(base, marker, extra, 1)
}

func testClient() *api_client.ZedcloudAPI {
	return testProvider.Meta().(*api_client.ZedcloudAPI)
}

// captureID stores the resource's id from state, for PreConfig hooks that act
// on the object outside Terraform.
func captureID(resourceName string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("%s not found in state", resourceName)
		}
		*id = rs.Primary.ID
		return nil
	}
}

func sortedInt64(in []int64) []int64 {
	out := append([]int64(nil), in...)
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

// ---------------------------------------------------------------- app tags

func TestApplication_Tags(t *testing.T) {
	const res = "zedcloud_application.test_tf_provider"
	var appID string

	cfg := func(extra string) string { return nfr249Config(t, "app_tags.tf", "#@APP_EXTRA@", extra) }

	checkAPITags := func(want map[string]string) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			app, err := getApplicationByID(context.Background(), testClient(), appID)
			if err != nil {
				return err
			}
			got := app.Tags
			if len(got) == 0 && len(want) == 0 {
				return nil
			}
			if !reflect.DeepEqual(got, want) {
				return fmt.Errorf("controller tags = %v, want %v", got, want)
			}
			return nil
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testApplicationDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// invalid input fails at plan time (first, so the post-test destroy uses a valid config)
			// invalid tags fail at plan time
			{
				Config: cfg(`description = "nfr249 step 4"
	tags = {
		"a/b" = "value"
	}`),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`key must not contain`),
			},
			// 1. create with tags
			{
				Config: cfg(`description = "nfr249 step 1"
	tags = {
		customer = "acme"
		tier     = "gold"
	}`),
				Check: resource.ComposeTestCheckFunc(
					captureID(res, &appID),
					resource.TestCheckResourceAttr(res, "tags.%", "2"),
					resource.TestCheckResourceAttr(res, "tags.customer", "acme"),
					checkAPITags(map[string]string{"customer": "acme", "tier": "gold"}),
				),
			},
			// 2. change and drop a tag
			{
				Config: cfg(`description = "nfr249 step 2"
	tags = {
		customer = "acme2"
	}`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(res, "tags.%", "1"),
					checkAPITags(map[string]string{"customer": "acme2"}),
				),
			},
			// 3. remove the argument: Terraform managed the tags, so they are cleared
			{
				Config: cfg(`description = "nfr249 step 3"`),
				Check:  checkAPITags(nil),
			},
			// 4. tags set outside Terraform (UI/API) survive an unrelated apply,
			//    and cause no diff (the post-apply plan must be empty)
			{
				PreConfig: func() {
					if err := setApplicationTagsOutOfBand(appID, map[string]string{"ui-tag": "set-in-ui"}); err != nil {
						t.Fatal(err)
					}
				},
				Config: cfg(`description = "nfr249 step 4"`),
				Check: resource.ComposeTestCheckFunc(
					checkAPITags(map[string]string{"ui-tag": "set-in-ui"}),
					resource.TestCheckResourceAttr(res, "description", "nfr249 step 4"),
				),
			},
		},
	})
}

func setApplicationTagsOutOfBand(id string, tags map[string]string) error {
	client := testClient()
	app, err := getApplicationByID(context.Background(), client, id)
	if err != nil {
		return err
	}
	app.Tags = tags
	params := app_client.UpdateParams()
	params.ID = id
	params.SetBody(app)
	if _, err = client.Application.Update(params, nil); err != nil {
		return err
	}
	after, err := getApplicationByID(context.Background(), client, id)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(after.Tags, tags) {
		return fmt.Errorf("out-of-band tag update did not stick: controller has %v, want %v", after.Tags, tags)
	}
	return nil
}

// ------------------------------------------------------- project PCR policy

func TestProject_PCRPolicy(t *testing.T) {
	const res = "zedcloud_project.test_tf_provider_pcr"
	var projectID string

	cfg := func(indices string) string {
		return nfr249Config(t, "project_pcr.tf", "#@PROJECT_EXTRA@", fmt.Sprintf(`pcr_policy {
		title = "nfr249-pcr"
		type  = "POLICY_TYPE_PCR"
		pcr_policy {
			pcr_indices = %s
		}
	}`, indices))
	}

	checkAPI := func(wantIndices []int64, wantVersion int64) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			var tag models.Tag
			if err := testProjectExists(res, &tag)(s); err != nil {
				return err
			}
			if tag.PcrPolicy == nil || tag.PcrPolicy.PcrPolicy == nil {
				return fmt.Errorf("project %s has no PCR policy on the controller", projectID)
			}
			p := tag.PcrPolicy.PcrPolicy
			if !reflect.DeepEqual(sortedInt64(p.PcrIndices), wantIndices) {
				return fmt.Errorf("controller pcrIndices = %v, want %v", p.PcrIndices, wantIndices)
			}
			if p.PolicyVersion != wantVersion {
				return fmt.Errorf("controller policyVersion = %d, want %d", p.PolicyVersion, wantVersion)
			}
			return nil
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testProjectDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// invalid input fails at plan time (first, so the post-test destroy uses a valid config)
			{
				Config:      cfg("[1, 2]"),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`PCR 0`),
			},
			{
				Config:      cfg("[0, 5]"),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`volatile`),
			},
			{
				Config: cfg("[0, 1, 2, 3, 4, 6, 7]"),
				Check: resource.ComposeTestCheckFunc(
					captureID(res, &projectID),
					checkAPI([]int64{0, 1, 2, 3, 4, 6, 7}, 1),
					resource.TestCheckResourceAttr(res, "pcr_policy.0.pcr_policy.0.policy_version", "1"),
					resource.TestCheckResourceAttr(res, "pcr_policy.0.pcr_policy.0.pcr_indices.#", "7"),
				),
			},
			// same set in a different order: no change
			{
				Config:   cfg("[7, 6, 4, 3, 2, 1, 0]"),
				PlanOnly: true,
			},
			// changed set: version bumps
			{
				Config: cfg("[0, 7]"),
				Check: resource.ComposeTestCheckFunc(
					checkAPI([]int64{0, 7}, 2),
					resource.TestCheckResourceAttr(res, "pcr_policy.0.pcr_policy.0.policy_version", "2"),
				),
			},
		},
	})
}

// ----------------------------------------------------- edge node PCR policy

func TestNode_PCRPolicy(t *testing.T) {
	const res = "zedcloud_edgenode.required_only"
	var nodeID string

	cfg := func(extra string) string { return nfr249Config(t, "node_pcr.tf", "#@NODE_EXTRA@", extra) }
	pcrBlock := func(title, indices string) string {
		return fmt.Sprintf(`title = %q
	pcr_policy {
		pcr_indices = %s
	}`, title, indices)
	}

	checkAPI := func(wantIndices []int64, wantVersion int64) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			node, err := getNodeByID(nodeID)
			if err != nil {
				return err
			}
			if wantIndices == nil {
				if node.PcrPolicy != nil && node.PcrPolicy.PcrPolicy != nil && len(node.PcrPolicy.PcrPolicy.PcrIndices) > 0 {
					return fmt.Errorf("node still has a PCR policy: %v", node.PcrPolicy.PcrPolicy.PcrIndices)
				}
				return nil
			}
			if node.PcrPolicy == nil || node.PcrPolicy.PcrPolicy == nil {
				return fmt.Errorf("node has no PCR policy on the controller")
			}
			p := node.PcrPolicy.PcrPolicy
			if !reflect.DeepEqual(sortedInt64(p.PcrIndices), wantIndices) {
				return fmt.Errorf("controller pcrIndices = %v, want %v", p.PcrIndices, wantIndices)
			}
			if p.PolicyVersion != wantVersion {
				return fmt.Errorf("controller policyVersion = %d, want %d", p.PolicyVersion, wantVersion)
			}
			return nil
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testNodeDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// invalid input fails at plan time (first, so the post-test destroy uses a valid config)
			{
				Config:      cfg(pcrBlock("nfr249-step6", "[1, 2]")),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`PCR 0`),
			},
			// 1. node without a PCR policy
			{
				Config: cfg(`title = "nfr249-step1"`),
				Check: resource.ComposeTestCheckFunc(
					captureID(res, &nodeID),
					checkAPI(nil, 0),
				),
			},
			// 2. a policy set with zcli (outside Terraform) survives an unrelated
			//    apply and causes no diff
			{
				PreConfig: func() {
					if err := setNodePCRPolicyOutOfBand(nodeID, []int64{0, 7}, 1); err != nil {
						t.Fatal(err)
					}
				},
				Config: cfg(`title = "nfr249-step2"`),
				Check: resource.ComposeTestCheckFunc(
					checkAPI([]int64{0, 7}, 1),
					resource.TestCheckResourceAttr(res, "title", "nfr249-step2"),
					resource.TestCheckResourceAttr(res, "pcr_policy.#", "0"),
				),
			},
			// 3. Terraform takes over: version continues from zcli's
			{
				Config: cfg(pcrBlock("nfr249-step3", "[0, 1, 2]")),
				Check: resource.ComposeTestCheckFunc(
					checkAPI([]int64{0, 1, 2}, 2),
					resource.TestCheckResourceAttr(res, "pcr_policy.0.policy_version", "2"),
				),
			},
			// 4. same set reordered plus an unrelated change: version unchanged
			{
				Config: cfg(pcrBlock("nfr249-step4", "[2, 1, 0]")),
				Check: resource.ComposeTestCheckFunc(
					checkAPI([]int64{0, 1, 2}, 2),
					resource.TestCheckResourceAttr(res, "title", "nfr249-step4"),
				),
			},
			// 5. removing the block resets the node
			{
				Config: cfg(`title = "nfr249-step5"`),
				Check:  checkAPI(nil, 0),
			},
		},
	})
}

func getNodeByID(id string) (*models.Node, error) {
	params := node_client.GetByIDParams()
	params.ID = id
	resp, err := testClient().Node.GetByID(params, nil)
	if err != nil {
		return nil, err
	}
	return resp.GetPayload(), nil
}

// setNodePCRPolicyOutOfBand does what `zcli edge-node pcr-policy set` does: a
// read-modify-write of the whole device.
func setNodePCRPolicyOutOfBand(id string, indices []int64, version int64) error {
	node, err := getNodeByID(id)
	if err != nil {
		return err
	}
	node.PcrPolicy = &models.Policy{
		Type:      models.NewPolicyType(models.PolicyTypePOLICYTYPEPCR),
		PcrPolicy: &models.PCRPolicy{PcrIndices: indices, PolicyVersion: version},
	}
	params := node_client.UpdateParams()
	params.ID = id
	params.SetBody(node)
	_, err = testClient().Node.Update(params, nil)
	return err
}

// ------------------------------------------- deployment update + PCR policy

func TestDeployment_UpdateInPlace(t *testing.T) {
	const res = "zedcloud_deployment.test_tf_provider_deployment_pcr"
	var firstID, secondID, projectID, oobID string

	tag := func(v string) string { return "depl:nfr249-" + v + testhelper.Suffix() }
	cfg := func(title, depTag, indices string) string {
		return nfr249Config(t, "deployment_pcr.tf", "#@DEPLOYMENT_EXTRA@", fmt.Sprintf(`title          = %q
	deployment_tag = %q

	device_policies {
		meta_data {
			name  = "test_tf_provider-nfr249-pcr%s"
			title = "test_tf_provider-nfr249-pcr"
			policy_target_condition = {
				"depl" : "nfr249"
			}
		}
		policy_sub_type = "DEVICE_POLICY_TYPE_PCR"
		pcr_policy {
			pcr_indices = %s
		}
	}`, title, depTag, testhelper.Suffix(), indices))
	}

	checkAPI := func(wantTitle, wantTag string, wantIndices []int64) resource.TestCheckFunc {
		return func(s *terraform.State) error {
			var dep models.Deployment
			if err := testDeploymentExists(res, &dep)(s); err != nil {
				return err
			}
			if dep.Title != wantTitle || dep.DeploymentTag != wantTag {
				return fmt.Errorf("controller title/tag = %q/%q, want %q/%q", dep.Title, dep.DeploymentTag, wantTitle, wantTag)
			}
			for _, p := range dep.DevicePolicies {
				if p.PcrPolicy != nil {
					if !reflect.DeepEqual(sortedInt64(p.PcrPolicy.PcrIndices), wantIndices) {
						return fmt.Errorf("controller pcrIndices = %v, want %v", p.PcrPolicy.PcrIndices, wantIndices)
					}
					return nil
				}
			}
			return fmt.Errorf("deployment has no PCR device policy on the controller: %+v", dep.DevicePolicies)
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testProjectDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// invalid input fails at plan time (first, so the post-test destroy uses a valid config)
			{
				Config:      cfg("nfr249-v1", tag("v1"), "[0, 5]"),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`volatile`),
			},
			{
				Config: cfg("nfr249-v1", tag("v1"), "[0, 1, 2]"),
				Check: resource.ComposeTestCheckFunc(
					captureID(res, &firstID),
					captureID("zedcloud_project.test_tf_provider_deployment_pcr", &projectID),
					checkAPI("nfr249-v1", tag("v1"), []int64{0, 1, 2}),
					resource.TestCheckResourceAttr(res, "name", suffixed("test_tf_provider-nfr249-deployment")),
				),
			},
			// a change without a new deployment_tag is refused at plan time
			{
				Config:      cfg("nfr249-v2", tag("v1"), "[0, 7]"),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`set a new deployment_tag`),
			},
			// in-place change: a new version with a new id; the post-apply plan
			// must be empty, or every apply would create another version
			{
				Config: cfg("nfr249-v2", tag("v2"), "[0, 7]"),
				Check: resource.ComposeTestCheckFunc(
					captureID(res, &secondID),
					func(*terraform.State) error {
						if secondID == "" || secondID == firstID {
							return fmt.Errorf("expected a new deployment id after the update, still %q", firstID)
						}
						return nil
					},
					checkAPI("nfr249-v2", tag("v2"), []int64{0, 7}),
					resource.TestCheckResourceAttr(res, "name", suffixed("test_tf_provider-nfr249-deployment")),
					resource.TestCheckResourceAttr(res, "title", "nfr249-v2"),
				),
			},
			// a second in-place change on top of the first
			{
				Config: cfg("nfr249-v3", tag("v3"), "[0, 1, 7]"),
				Check:  checkAPI("nfr249-v3", tag("v3"), []int64{0, 1, 7}),
			},
			// a newer version created outside Terraform: the update is refused
			// with an explanation
			{
				PreConfig: func() {
					id := ""
					if err := latestDeploymentIDForTest(projectID, &id); err != nil {
						t.Fatal(err)
					}
					var err error
					if oobID, err = newDeploymentVersionOutOfBand(projectID, id, tag("oob")); err != nil {
						t.Fatal(err)
					}
				},
				Config:      cfg("nfr249-v4", tag("v4"), "[0, 1, 7]"),
				ExpectError: regexp.MustCompile(`not the latest version`),
			},
			// remove the out-of-band version again; the resource is unchanged
			// (the refused update did not touch it), and destroy then removes
			// every version this resource created, so the project can go too
			{
				PreConfig: func() {
					p := depl_client.NewDeleteParams()
					p.ProjectID = projectID
					p.ID = oobID
					if _, err := testClient().Deployment.Delete(p, nil); err != nil {
						t.Fatal(err)
					}
				},
				Config: cfg("nfr249-v3", tag("v3"), "[0, 1, 7]"),
				Check:  checkAPI("nfr249-v3", tag("v3"), []int64{0, 1, 7}),
			},
		},
	})
}

func latestDeploymentIDForTest(projectID string, id *string) error {
	latest, err := getLatestProjectDeployment(context.Background(), testClient(), projectID)
	if err != nil {
		return err
	}
	*id = latest
	return nil
}

func newDeploymentVersionOutOfBand(projectID, id, depTag string) (string, error) {
	client := testClient()
	get := depl_client.NewGetByIDParams()
	get.ProjectID = projectID
	get.ID = id
	resp, err := client.Deployment.GetByID(get, nil)
	if err != nil {
		return "", err
	}
	dep := resp.GetPayload()
	dep.Title = "changed-outside-terraform"
	dep.DeploymentTag = depTag
	put := depl_client.NewCreateNewVersionParams()
	put.SetProjectID(projectID)
	put.SetID(id)
	put.SetBody(dep)
	out, err := client.Deployment.CreateNewVersion(put, nil)
	if err != nil {
		return "", err
	}
	if out.GetPayload() == nil || out.GetPayload().ObjectID == "" {
		return "", fmt.Errorf("no id returned for the out-of-band deployment version")
	}
	return out.GetPayload().ObjectID, nil
}
