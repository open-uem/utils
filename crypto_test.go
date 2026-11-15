package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writePEM(t *testing.T, name string, block *pem.Block) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := pem.Encode(f, block); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestReadPEMPrivateKey(t *testing.T) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	ecKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	pkcs8DER, err := x509.MarshalPKCS8PrivateKey(rsaKey)
	if err != nil {
		t.Fatal(err)
	}
	ecDER, err := x509.MarshalPKCS8PrivateKey(ecKey)
	if err != nil {
		t.Fatal(err)
	}
	ecSEC1DER, err := x509.MarshalECPrivateKey(ecKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("PKCS1 RSA", func(t *testing.T) {
		path := writePEM(t, "key.pem", &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(rsaKey)})
		got, err := ReadPEMPrivateKey(path)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.Equal(rsaKey) {
			t.Error("parsed key differs from generated key")
		}
	})

	t.Run("PKCS8 RSA", func(t *testing.T) {
		path := writePEM(t, "key.pem", &pem.Block{Type: "PRIVATE KEY", Bytes: pkcs8DER})
		got, err := ReadPEMPrivateKey(path)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.Equal(rsaKey) {
			t.Error("parsed key differs from generated key")
		}
	})

	t.Run("PKCS8 EC rejected with clear error", func(t *testing.T) {
		path := writePEM(t, "key.pem", &pem.Block{Type: "PRIVATE KEY", Bytes: ecDER})
		if _, err := ReadPEMPrivateKey(path); err == nil || !strings.Contains(err.Error(), "RSA private key is required") {
			t.Fatalf("expected RSA-required error, got: %v", err)
		}
	})

	t.Run("SEC1 EC rejected with clear error", func(t *testing.T) {
		path := writePEM(t, "key.pem", &pem.Block{Type: "EC PRIVATE KEY", Bytes: ecSEC1DER})
		if _, err := ReadPEMPrivateKey(path); err == nil || !strings.Contains(err.Error(), "EC private key") {
			t.Fatalf("expected EC error, got: %v", err)
		}
	})

	t.Run("non-PEM file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "key.pem")
		if err := os.WriteFile(path, []byte("not a pem file"), 0600); err != nil {
			t.Fatal(err)
		}
		if _, err := ReadPEMPrivateKey(path); err == nil || !strings.Contains(err.Error(), "does not content a private key") {
			t.Fatalf("expected not-a-private-key error, got: %v", err)
		}
	})

	t.Run("certificate instead of key", func(t *testing.T) {
		path := writePEM(t, "key.pem", &pem.Block{Type: "CERTIFICATE", Bytes: pkcs8DER})
		if _, err := ReadPEMPrivateKey(path); err == nil || !strings.Contains(err.Error(), "does not content a private key") {
			t.Fatalf("expected not-a-private-key error, got: %v", err)
		}
	})
}
