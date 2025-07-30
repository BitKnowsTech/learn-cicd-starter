package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		name   string
		input  http.Header
		want   string
		errMsg string
	}

	headerWithBadAuth := http.Header{}
	headerWithBadAuth.Add("Authorization", "bad")

	headerWithOKAuth := http.Header{}
	headerWithOKAuth.Add("Authorization", "ApiKey TestKey")

	tests := []test{
		{
			"No headers",
			http.Header{},
			"",
			"no authorization header included",
		},
		{
			"Bad Auth",
			headerWithBadAuth,
			"",
			"malformed authorization header",
		},
		{
			"ok",
			headerWithOKAuth,
			"TestKey",
			"",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if tc.errMsg == "" {
				if err != nil {
					t.Errorf("Expected no error, got: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error: %q, got no error", tc.errMsg)
				} else if err.Error() != tc.errMsg {
					t.Errorf("Expected error message: %q, got %q", tc.errMsg, err.Error())
				}
			}

			if got != tc.want {
				t.Errorf("Expected: %q, got %q", tc.want, got)
			}
		})
	}
}
