package resources

import (
	"errors"
	"testing"
)

type codedErr struct {
	code int
	msg  string
}

func (e *codedErr) Error() string { return e.msg }
func (e *codedErr) Code() int     { return e.code }

func TestIsStatusNotFound(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil error", nil, false},
		{"typed 404", &codedErr{code: 404, msg: "not found"}, true},
		{"typed 500", &codedErr{code: 500, msg: "boom"}, false},
		{"retryablehttp wrapped 404", errors.New(`giving up after 5 attempt(s): unexpected HTTP status 404 Not Found`), true},
		{"unrelated error", errors.New("connection reset"), false},
		{"500 substring not matched", errors.New("500 Internal Server Error"), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isStatusNotFound(tc.err); got != tc.want {
				t.Errorf("isStatusNotFound(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}

// NFR-165 §4.1. `terraform import` takes one opaque string, and the IAM
// resources accept either the system ID or the object's name. This is the
// discriminator; if it misjudges, an import either queries a nonexistent name
// or a nonexistent ID.
func TestLooksLikeObjectID(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		want bool
	}{
		{"real user ID", "AAGFABAEqnH4je5PHZTXSmHOs-XC", true},
		{"ID with underscore and equals", "AAGFABAEqnH4je5PHZTXSmH_s=XC", true},
		{"all digits, 28 long", "1234567890123456789012345678", true},

		{"email username", "alice@corp.com", false},
		{"role name", "readonly-operators", false},
		{"empty", "", false},
		{"27 characters", "AAGFABAEqnH4je5PHZTXSmHOs-X", false},
		{"29 characters", "AAGFABAEqnH4je5PHZTXSmHOs-XCD", false},
		// 28 characters, but '@' and '.' are outside the ID alphabet, so a
		// long email address is still read as a name.
		{"28-character email", "alice.mcname@corporate.co.uk", false},
		{"leading space", " AGFABAEqnH4je5PHZTXSmHOs-XC", false},
		{"embedded newline", "AAGFABAEqnH4je5PHZTXSmHOs-X\n", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := looksLikeObjectID(tc.raw); got != tc.want {
				t.Errorf("looksLikeObjectID(%q) = %v, want %v", tc.raw, got, tc.want)
			}
		})
	}
}
