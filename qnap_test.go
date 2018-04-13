package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"
)

func random() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func TestAuthOK(t *testing.T) {
	expected := random()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<QDocRoot><xml:authSid>" + expected + "</xml:authSid></QDocRoot>"))
	}))
	defer server.Close()

	actual, err := Auth(server.URL, "user", "password")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Error("Sid is not match.")
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

func TestUploadOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	dir := "testdata"

	err := Upload(server.URL, "sid", dir)
	if err != nil {
		t.Error(err)
	}
}
