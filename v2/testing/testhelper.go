package testing

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ghodss/yaml"
)

// Suffix returns the per-run suffix appended to every object name the fixtures
// create, or "" when ZEDCLOUD_TEST_SUFFIX is unset.
//
// Deliberately NOT generated when absent, unlike RealNode's fallback: with an
// empty suffix the rendered fixtures are byte-identical to the pre-tokenisation
// ones, which is what makes the rollout a provable no-op. CI sets it (see
// .github/workflows/e2e.yml); `make test` leaves it empty and behaves as before.
//
// Object names are validated by the controller against
// [a-zA-Z0-9][a-zA-Z0-9_.-]+ with maxLength 256, so a suffix may only use
// letters, digits, underscore, dot and hyphen. The longest name in testdata is
// 55 characters, leaving ample room.
func Suffix() string {
	return os.Getenv(EnvTestSuffix)
}

// defaultVars are the placeholders every fixture may use without the test
// having to pass anything.
func defaultVars() map[string]string {
	return map[string]string{"SUFFIX": Suffix()}
}

// expandTokens replaces __TOKEN__ placeholders in a fixture and hard-fails on
// one that was left behind.
//
// An unexpanded placeholder is fatal rather than ignored: applying a config
// containing a literal __TOKEN__ would create garbage objects on a shared
// controller and fail later with a confusing error. A fixture that needs a token
// this loader does not know about (__NODE_ID__, say) must be read through the
// WithVars variant.
func expandTokens(t *testing.T, path, in string, vars map[string]string) string {
	t.Helper()

	out := in
	for k, v := range vars {
		out = strings.ReplaceAll(out, "__"+k+"__", v)
	}

	if i := strings.Index(out, "__"); i >= 0 {
		if j := strings.Index(out[i+2:], "__"); j >= 0 {
			t.Fatalf("fixture %s still contains an unexpanded placeholder %q; known vars: %v",
				path, out[i:i+2+j+2], keysOf(vars))
		}
	}

	return out
}

func mustReadTestdata(t *testing.T, path string) string {
	t.Helper()

	testdataDir, err := filepath.Abs("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(filepath.Join(testdataDir, path))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

// MustGetTestInput reads ./testdata/<path> and expands the __TOKEN__
// placeholders every fixture may use -- today just __SUFFIX__, from
// ZEDCLOUD_TEST_SUFFIX.
//
// Expansion happens here rather than at the ~47 call sites so a new test cannot
// forget it, and so a fixture that grows a name can never be applied with a
// literal placeholder.
func MustGetTestInput(t *testing.T, path string) string {
	t.Helper()

	return expandTokens(t, path, mustReadTestdata(t, path), defaultVars())
}

// MustGetExpectedOutput reads the golden ./testdata/<path> and unmarshals it
// into i, expanding the same placeholders as MustGetTestInput.
//
// The goldens carry __SUFFIX__ too: names ARE compared (few of the ~49
// cmpopts.IgnoreFields lists ignore Name/Title), so tokenising them keeps that
// assertion instead of weakening it to make room for the suffix.
func MustGetExpectedOutput(t *testing.T, path string, i interface{}) {
	t.Helper()

	MustGetExpectedOutputWithVars(t, path, defaultVars(), i)
}

// MustGetExpectedOutputWithVars is MustGetExpectedOutput with an explicit
// placeholder set, for goldens that need more than the suite-wide defaults.
func MustGetExpectedOutputWithVars(t *testing.T, path string, vars map[string]string, i interface{}) {
	t.Helper()

	expanded := expandTokens(t, path, mustReadTestdata(t, path), vars)

	jsonBytes, err := yaml.YAMLToJSON([]byte(expanded))
	if err != nil {
		t.Logf("unexpected error: %s", err)
		return
	}
	if err := json.Unmarshal(jsonBytes, i); err != nil {
		t.Fatal(err)
	}
}

func ToYAML(path string, i interface{}) {
	testdataDir, err := filepath.Abs("./testdata")
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(filepath.Join(testdataDir, path), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	yamlBytes, err := yaml.Marshal(i)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(yamlBytes)
	if err != nil {
		panic(err)
	}
}

// CheckEnv validates the necessary test API keys exist in the testing environment.
func CheckEnv(t *testing.T) {
	if v := os.Getenv("TF_CLI_CONFIG_FILE"); v == "" {
		if _, err := os.Stat("dev.tfrc"); errors.Is(err, os.ErrNotExist) {
			t.Fatal("TF_CLI_CONFIG_FILE must be set or dev.tfrc file must be present for acceptance tests, it should contain the dev_overrides config that points to local instance of the provider")
		}
	}

	if v := os.Getenv("TF_VAR_zedcloud_token"); v == "" {
		t.Fatal("TF_VAR_zedcloud_token must be set for acceptance tests to access the zedcloud API")
	}
}
