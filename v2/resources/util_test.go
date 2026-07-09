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
