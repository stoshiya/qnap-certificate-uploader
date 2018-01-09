package main

import (
	"os"
)

const (
	Version = "0.1.0"
	Name    = "qnap-certificate-uploader"
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
