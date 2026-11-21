
package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		want       string
		wantErr    error
	}{
		{
			name:       "valid API key",
			authHeader: "ApiKey my-secret-key",
			want:       "my-secret-key",
		},
		{
			name:    "missing authorization header",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:       "malformed authorization header",
			authHeader: "Bearer my-secret-key",
			wantErr:    nil, // Checked below as a non-nil error.
		},
		{
			name:       "missing API key",
			authHeader: "ApiKey",
			wantErr:    nil, // Checked below as a non-nil error.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			got, err := GetAPIKey(headers)

			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if tt.name != "valid API key" && err == nil {
				t.Fatal("expected an error, got nil")
			}

			if tt.name == "valid API key" && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("expected API key %q, got %q", tt.want, got)
			}
		})
	}
}
