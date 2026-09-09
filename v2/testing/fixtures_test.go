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
