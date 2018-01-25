package main

import (
	"fmt"
	"net/http/httptest"
	"net/http"
	"path/filepath"
	"testing"
)

func TestAuth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	_, err := Auth(server.URL, "user", "password")
	if err == nil {
		t.Error("error not occurs.")
	}
}

func TestUploadCertNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	dir := filepath.Join("testdata", "cert-not-found")

	err := Upload(server.URL, "sid", dir)
	if err.Error() != fmt.Sprintf("open %s/cert.pem: no such file or directory", dir) {
		t.Error("error not match.")
	}
}

func TestUploadPrivateKeyNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	dir := filepath.Join("testdata", "privkey-not-found")

	err := Upload(server.URL, "sid", dir)
	if err.Error() != fmt.Sprintf("open %s/privkey.pem: no such file or directory", dir) {
		t.Error("error not match.")
	}
}

func TestUploadChainNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	dir := filepath.Join("testdata", "chain-not-found")

	err := Upload(server.URL, "sid", dir)
	if err.Error() != fmt.Sprintf("open %s/chain.pem: no such file or directory", dir) {
		t.Error("error not match.")
	}
}
