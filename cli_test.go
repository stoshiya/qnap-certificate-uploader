package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
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
