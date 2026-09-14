// Copyright (c) Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package testing

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// fixtureRoot is the acceptance-suite testdata, which lives next to the tests
// that use it rather than next to these helpers.
const fixtureRoot = "../resources/testdata"

// knownTokens is every placeholder a loader can expand.
//
// SUFFIX comes from ZEDCLOUD_TEST_SUFFIX via MustGetTestInput; NODE_ID is
// injected by the real-node test through MustGetTestInputWithVars.
var knownTokens = map[string]string{
	"SUFFIX":  "MustGetTestInput / MustGetExpectedOutput",
	"NODE_ID": "MustGetTestInputWithVars, from testhelper.RealNode",
}

var tokenRE = regexp.MustCompile(`__[A-Z][A-Z0-9_]*__`)

// TestFixtureTokensAreKnown fails on a placeholder no loader will ever expand.
//
// Without this, a typo like __SUFIX__ surfaces only as a mid-apply failure
// against a shared controller -- and only for whoever happens to run that one
// test. The loaders hard-fail on an unexpanded token at read time, but that is
// per-test and lazy; this walks every fixture in one cheap, cluster-free test.
func TestFixtureTokensAreKnown(t *testing.T) {
	seen := map[string]int{}

	err := filepath.Walk(fixtureRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		switch filepath.Ext(path) {
		case ".tf", ".yaml":
		default:
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		for _, tok := range tokenRE.FindAllString(string(b), -1) {
			name := strings.Trim(tok, "_")
			seen[name]++
			if _, ok := knownTokens[name]; !ok {
				t.Errorf("%s: unknown placeholder %s; no loader expands it", path, tok)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking %s: %s", fixtureRoot, err)
	}

	// A zero count would mean the rollout was reverted or the walk found
	// nothing, which should not pass quietly.
	if seen["SUFFIX"] == 0 {
		t.Errorf("no __SUFFIX__ placeholders found under %s; object names would collide "+
			"between concurrent runs", fixtureRoot)
	}
	t.Logf("placeholders: %v", seen)
}

// goldenDenyKeys are golden fields that legitimately hold an object's name
// without the token.
//
// Several fixtures set `description` to the same string as the object's name.
// Descriptions have no uniqueness constraint so the fixtures leave them alone,
// which means the golden copy must stay untokenised too. nameAppPart is the
// same: deployment/create.tf sets name_app_part as a naming-scheme component,
// not as a reference to the app, so the controller echoes it unsuffixed.
//
// Add a key here only after confirming against a controller that the field
// really does come back without the suffix.
var goldenDenyKeys = map[string]bool{"description": true, "nameAppPart": true}

var (
	// name/title/serialno/token = "value__SUFFIX__" in a .tf fixture.
	fixtureTokenised = regexp.MustCompile(
		`(?m)^\s*(?:name|title|serialno|token)\s*=\s*"([^"]*)__SUFFIX__"\s*$`)
	// `key: value` in a golden, including the first key of a list item.
	goldenScalar = regexp.MustCompile(`^\s*(?:-\s+)?([A-Za-z_][A-Za-z0-9_]*):\s*"?([^"\n]*?)"?\s*$`)
)

// TestGoldensTokeniseEveryFixtureName fails when a golden asserts an object name
// that the fixtures suffix, without the token.
//
// This is the one failure mode the mechanism has that nothing else catches
// cheaply: the fixture creates `foo_gh123` and the golden still expects `foo`,
// so the suite passes with an empty suffix and fails only in CI, where a suffix
// is set. It was found the expensive way -- two golden `imagename:` references
// written as list items (`- imagename: x`) were missed by the rewrite, and only
// a live suffixed run against a controller showed it.
//
// Goldens name objects under many keys (`imagename`, `netname`, `projectName`,
// `nameAppPart`, ...), so this matches on the value rather than the key.
func TestGoldensTokeniseEveryFixtureName(t *testing.T) {
	tokenised := map[string]string{} // base name -> fixture that declares it

	err := filepath.Walk(fixtureRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".tf" {
			return err
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for _, m := range fixtureTokenised.FindAllStringSubmatch(string(b), -1) {
			if m[1] != "" {
				tokenised[m[1]] = path
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking %s: %s", fixtureRoot, err)
	}
	if len(tokenised) == 0 {
		t.Fatalf("no tokenised names found under %s", fixtureRoot)
	}

	err = filepath.Walk(fixtureRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".yaml" {
			return err
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for i, line := range strings.Split(string(b), "\n") {
			m := goldenScalar.FindStringSubmatch(line)
			if m == nil || goldenDenyKeys[m[1]] {
				continue
			}
			if src, ok := tokenised[m[2]]; ok {
				t.Errorf("%s:%d: %s: %q is suffixed in %s but not here; "+
					"append __SUFFIX__ (or add the key to goldenDenyKeys if the fixture "+
					"really does send it unsuffixed)",
					path, i+1, m[1], m[2], src)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking %s: %s", fixtureRoot, err)
	}
}

// TestSuffixIsEnvVerbatim pins the property the whole rollout rests on: with
// ZEDCLOUD_TEST_SUFFIX unset the fixtures render exactly as they did before
// tokenisation, so a green suite proves the refactor changed nothing.
func TestSuffixIsEnvVerbatim(t *testing.T) {
	t.Setenv(EnvTestSuffix, "")
	if got := Suffix(); got != "" {
		t.Errorf("Suffix() with %s empty = %q, want %q (a generated fallback would "+
			"make the no-op proof impossible)", EnvTestSuffix, got, "")
	}

	t.Setenv(EnvTestSuffix, "_gh123")
	if got, want := Suffix(), "_gh123"; got != want {
		t.Errorf("Suffix() = %q, want %q", got, want)
	}
}
