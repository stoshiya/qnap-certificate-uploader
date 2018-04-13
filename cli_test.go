package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestCLI_RunParseError(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{
		outStream: outStream,
		errStream: errStream,
	}
	args := strings.Split("qnap-certificate-uploader -foo", " ")

	status := cli.Run(args)
	if status != ExitCodeParseError {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeParseError)
	}
}

func TestCLI_RunVersion(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{
		outStream: outStream,
		errStream: errStream,
	}
	args := strings.Split("qnap-certificate-uploader -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}

	expected := fmt.Sprintf("%s version %s\n", Name, Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("Output=%q, want %q", errStream.String(), expected)
	}
}

func TestCLI_RunUsage(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{
		outStream: outStream,
		errStream: errStream,
	}
	args := strings.Split("qnap-certificate-uploader", " ")

	status := cli.Run(args)
	if status != ExitCodeUsage {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeUsage)
	}
}

func TestCLI_RunURLParseError(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{
		outStream: outStream,
		errStream: errStream,
	}
	args := strings.Split("qnap-certificate-uploader -user root -password password -url ://localhost", " ")

	status := cli.Run(args)
	if status != ExitCodeURLParseError {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeURLParseError)
	}
}

func TestCLI_RunAuthError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{
		outStream: outStream,
		errStream: errStream,
	}
	args := strings.Split("qnap-certificate-uploader -user root -password password -url http://localhost", " ")

	status := cli.Run(args)
	if status != ExitCodeAuthError {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeAuthError)
	}
}

func TestCLI_RunUploadError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<QDocRoot><xml:authSid>" + random() + "</xml:authSid></QDocRoot>"))
	}))
	defer server.Close()

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{
		outStream: outStream,
		errStream: errStream,
	}
	args := strings.Split("qnap-certificate-uploader -user root -password password -url " + server.URL , " ")

	status := cli.Run(args)
	if status != ExitCodeUploadError {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeUploadError)
	}
}
