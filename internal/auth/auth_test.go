package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
		err   error
	}{
		"simple": {
			input: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey test_key")
				return h
			}(),
			want: "test_key",
			err:  nil,
		},
		"missing auth header": {
			input: func() http.Header {
				h := http.Header{}
				return h
			}(),
			want: "",
			err:  ErrNoAuthHeaderIncluded,
		},
		"not formatted correctly": {
			input: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "test_key")
				return h
			}(),
			want: "",
			err:  ErrMalformedAuthHeader,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if err != tc.err {
				t.Fatalf("expected error %v, got %v", tc.err, err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
