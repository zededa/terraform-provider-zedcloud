package resources

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	api_client "github.com/zededa/terraform-provider-zedcloud/v2/client"
	"github.com/zededa/terraform-provider-zedcloud/v2/client/artifact"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
	testhelper "github.com/zededa/terraform-provider-zedcloud/v2/testing"
)

const artifactResourceName = "zedcloud_artifact.test_tf_provider"

// artifactTestInput renders a fixture with __SUFFIX__ and __TESTDATA__, the
// absolute path of ./testdata (terraform runs in a temp dir during tests).
func artifactTestInput(t *testing.T, path string) string {
	t.Helper()
	dir, err := filepath.Abs("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	return testhelper.MustGetTestInputWithVars(t, path, map[string]string{
		"SUFFIX":   testhelper.Suffix(),
		"TESTDATA": filepath.ToSlash(dir),
	})
}

// artifactIDRegexp matches "<uuid>_<name>" for the given name.
func artifactIDRegexp(name string) *regexp.Regexp {
	return regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}_` + regexp.QuoteMeta(name) + `$`)
}

func TestArtifact_CRUD(t *testing.T) {
	var firstID string
	name := suffixed("test_tf_provider-logo") + ".png"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: testArtifactDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// Create from a file
			{
				Config: artifactTestInput(t, "artifact/create.tf"),
				Check: resource.ComposeTestCheckFunc(
					testArtifactExists(artifactResourceName, &firstID),
					resource.TestCheckResourceAttr(artifactResourceName, "name", name),
					resource.TestMatchResourceAttr(artifactResourceName, "id", artifactIDRegexp(name)),
					resource.TestCheckResourceAttr(artifactResourceName, "size", "69"),
					resource.TestCheckResourceAttr(artifactResourceName, "retain_on_destroy", "false"),
				),
			},
			// Replace: new content, new id; the old artifact is deleted
			{
				Config: artifactTestInput(t, "artifact/replace.tf"),
				Check: resource.ComposeTestCheckFunc(
					testArtifactReplaced(artifactResourceName, &firstID),
					resource.TestMatchResourceAttr(artifactResourceName, "id", artifactIDRegexp(name)),
					resource.TestMatchResourceAttr(artifactResourceName, "content_base64", regexp.MustCompile(`^sha256:[0-9a-f]{64}$`)),
				),
			},
			// Import: the API returns no content, so content inputs can't be verified
			{
				ResourceName:            artifactResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"content_base64", "source", "source_hash", "size"},
			},
		},
	})
}

func TestArtifact_RetainOnDestroy(t *testing.T) {
	var id string

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testhelper.CheckEnv(t) },
		Providers: testAccProviders,
		// Destroy must leave the artifact on the controller; then clean it up.
		CheckDestroy: func(s *terraform.State) error {
			if id == "" {
				return errors.New("artifact id was never recorded")
			}
			client := testProvider.Meta().(*api_client.ZedcloudAPI)
			_, err := client.Artifact.GetSignedURL(context.Background(), id)
			if derr := client.Artifact.Delete(context.Background(), id); derr != nil {
				t.Logf("cleanup of retained artifact %s: %s", id, derr)
			}
			if err != nil {
				return fmt.Errorf("retain_on_destroy: artifact %s should still exist after destroy: %w", id, err)
			}
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: artifactTestInput(t, "artifact/retain.tf"),
				Check: resource.ComposeTestCheckFunc(
					testArtifactExists(artifactResourceName, &id),
					resource.TestCheckResourceAttr(artifactResourceName, "retain_on_destroy", "true"),
				),
			},
		},
	})
}

func TestApplication_CreateWithLogo(t *testing.T) {
	var app models.Application

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testhelper.CheckEnv(t) },
		CheckDestroy: resource.ComposeTestCheckFunc(testApplicationDestroy, testArtifactDestroy),
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: artifactTestInput(t, "application/create_with_logo.tf"),
				Check: resource.ComposeTestCheckFunc(
					testApplicationExists("zedcloud_application.test_tf_provider", &app),
					resource.TestCheckResourceAttrPair(
						"zedcloud_application.test_tf_provider", "manifest.0.desc.0.logo.logo",
						"zedcloud_artifact.test_tf_provider_logo", "id",
					),
					testApplicationLogoIs(&app, "zedcloud_artifact.test_tf_provider_logo"),
				),
			},
		},
	})
}

func TestApplication_LogoURLRejected(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testhelper.CheckEnv(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testhelper.MustGetTestInput(t, "application/logo_url_rejected.tf"),
				PlanOnly:    true,
				ExpectError: regexp.MustCompile(`App logo must be an artifact id, not a URL`),
			},
		},
	})
}

// testApplicationLogoIs checks, against the API, that the app's
// manifestJSON.desc.logo is exactly {"logo": <artifact id>}, the shape the UI writes.
func testApplicationLogoIs(app *models.Application, artifactResource string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[artifactResource]
		if !ok {
			return fmt.Errorf("not found: %s", artifactResource)
		}
		if app.Manifest == nil || app.Manifest.Desc == nil {
			return errors.New("application has no manifest.desc")
		}
		logo := app.Manifest.Desc.Logo
		if len(logo) != 1 || logo["logo"] != rs.Primary.ID {
			return fmt.Errorf("manifestJSON.desc.logo = %v, want {logo: %s}", logo, rs.Primary.ID)
		}
		return nil
	}
}

// testArtifactExists checks the artifact in state is available on the
// controller and records its id.
func testArtifactExists(resourceName string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("artifact not found in state: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return errors.New("artifact ID is not set")
		}
		client := testProvider.Meta().(*api_client.ZedcloudAPI)
		a, err := client.Artifact.GetSignedURL(context.Background(), rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("artifact %s is not available: %w", rs.Primary.ID, err)
		}
		if a.SignedURL == "" {
			return fmt.Errorf("artifact %s: empty signed URL", rs.Primary.ID)
		}
		*id = rs.Primary.ID
		return nil
	}
}

// testArtifactReplaced checks the resource got a new id and the previous
// artifact is gone from the controller.
func testArtifactReplaced(resourceName string, previousID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var id string
		if err := testArtifactExists(resourceName, &id)(s); err != nil {
			return err
		}
		if id == *previousID {
			return fmt.Errorf("artifact was not replaced: id is still %s", id)
		}
		client := testProvider.Meta().(*api_client.ZedcloudAPI)
		if _, err := client.Artifact.GetSignedURL(context.Background(), *previousID); !errors.Is(err, artifact.ErrNotFound) {
			return fmt.Errorf("replaced artifact %s should be deleted, got err=%v", *previousID, err)
		}
		*previousID = id
		return nil
	}
}

func testArtifactDestroy(s *terraform.State) error {
	client := testProvider.Meta().(*api_client.ZedcloudAPI)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zedcloud_artifact" {
			continue
		}
		_, err := client.Artifact.GetSignedURL(context.Background(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("destroy failed, artifact %s still exists", rs.Primary.ID)
		}
		if !errors.Is(err, artifact.ErrNotFound) {
			return fmt.Errorf("checking artifact %s after destroy: %w", rs.Primary.ID, err)
		}
	}
	return nil
}

// ---- unit tests (no controller needed) ----

func TestArtifactNameFromID(t *testing.T) {
	for id, want := range map[string]string{
		"07a928b1-c1ae-11f1-b06c-ca49049db012_logo.png": "logo.png",
		"07a928b1-c1ae-11f1-b06c-ca49049db012_a_b":      "a_b",
		"no-underscore": "",
		"trailing_":     "",
	} {
		if got := artifactNameFromID(id); got != want {
			t.Errorf("artifactNameFromID(%q) = %q, want %q", id, got, want)
		}
	}
}

// fakeArtifacts is a ClientService whose GetSignedURL reports not found for
// the first `pending` calls.
type fakeArtifacts struct {
	pending int
	err     error
	calls   int
}

func (f *fakeArtifacts) Create(context.Context, string) (*artifact.Artifact, error) { return nil, nil }
func (f *fakeArtifacts) Upload(context.Context, string, []byte) error               { return nil }
func (f *fakeArtifacts) Delete(context.Context, string) error                       { return nil }
func (f *fakeArtifacts) SetTransport(runtime.ClientTransport)                       {}
func (f *fakeArtifacts) GetSignedURL(context.Context, string) (*artifact.Artifact, error) {
	f.calls++
	if f.err != nil {
		return nil, f.err
	}
	if f.calls <= f.pending {
		return nil, artifact.ErrNotFound
	}
	return &artifact.Artifact{SignedURL: "https://s3/x"}, nil
}

func TestWaitForArtifact(t *testing.T) {
	old := artifactPollInterval
	artifactPollInterval = time.Millisecond
	defer func() { artifactPollInterval = old }()

	f := &fakeArtifacts{pending: 3}
	if err := waitForArtifact(context.Background(), f, "id", time.Second); err != nil || f.calls != 4 {
		t.Errorf("eventually available: err=%v calls=%d", err, f.calls)
	}

	f = &fakeArtifacts{pending: 1 << 30}
	if err := waitForArtifact(context.Background(), f, "id", 20*time.Millisecond); err == nil {
		t.Error("never available: expected a timeout error")
	}

	f = &fakeArtifacts{err: errors.New("boom")}
	if err := waitForArtifact(context.Background(), f, "id", time.Second); err == nil || f.calls != 1 {
		t.Errorf("hard error: err=%v calls=%d, want immediate failure", err, f.calls)
	}
}

// TestArtifactResource_InternalValidate catches schema mistakes (ForceNew
// without Update, validation on computed fields, ...) without a controller.
func TestArtifactResource_InternalValidate(t *testing.T) {
	if err := ArtifactResource().InternalValidate(nil, true); err != nil {
		t.Fatal(err)
	}
	if err := ApplicationResource().InternalValidate(nil, true); err != nil {
		t.Fatal(err)
	}
}
