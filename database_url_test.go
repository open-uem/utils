package utils

import (
	"net/url"
	"testing"
)

func TestPostgresDatabaseURLPasswordEscaping(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{"invalid escape percent", "2%6ZyFm6N3%2%oN8G6C9coUPj8Ck!pGN*Bde$Upt3Cp6e5VL&8u3BnnAiDX*nUBC"},
		{"at sign", "p@ss@word"},
		{"colon", "pa:ss:word"},
		{"slash and question", "pa/?ss#word"},
		{"space", "pa ss word"},
		{"plus", "pa+ss+word"},
		{"percent only", "100%"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dbURL := postgresDatabaseURL("openuem", tc.password, "localhost", "5432", "openuem")

			u, err := url.Parse(dbURL)
			if err != nil {
				t.Fatalf("could not parse generated URL %q: %v", dbURL, err)
			}

			got, set := u.User.Password()
			if !set {
				t.Fatalf("password was not set in parsed URL %q", dbURL)
			}
			if got != tc.password {
				t.Errorf("password does not round-trip: want %q, got %q (url: %s)", tc.password, got, dbURL)
			}
			if u.User.Username() != "openuem" {
				t.Errorf("username does not round-trip: got %q", u.User.Username())
			}
			if u.Host != "localhost:5432" || u.Path != "/openuem" {
				t.Errorf("host/path parsed incorrectly: %q", dbURL)
			}
		})
	}
}
