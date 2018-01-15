package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
)

const (
	ExitCodeOK int = iota
	ExitCodeParseError
	ExitCodeUsage
	ExitCodeError
	ExitCodeURLParseError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var (
		err                                  error
		version                              bool
		username, password, baseUrl, baseDir string
	)
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.BoolVar(&version, "version", false, "Print version information and quit.")
	flags.StringVar(&username, "user", "", "NAS account name.")
	flags.StringVar(&password, "password", "", "NAS account password.")
	flags.StringVar(&baseUrl, "url", "", "NAS protocol and host for web access. Port is optional. (example \"http://nas.example.com:8080\")")
	flags.StringVar(&baseDir, "dir", "/usr/local/etc/letsencrypt/live/", "Base directory for let's encrypt certificates.")

	if err = flags.Parse(args[1:]); err != nil {
		return ExitCodeParseError
	}

	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if username == "" || password == "" || baseUrl == "" {
		flags.Usage()
		return ExitCodeUsage
	}

	u, err := url.Parse(baseUrl)
	if err != nil {
		fmt.Fprintf(cli.errStream, "%v\n", err)
		return ExitCodeURLParseError
	}

	sid, err := Auth(baseUrl, username, password)
	if err != nil {
		fmt.Fprintf(cli.errStream, "%v\n", err)
		return ExitCodeError
	}

	if err = Upload(baseUrl, sid, filepath.Join(baseDir, u.Hostname())); err != nil {
		fmt.Fprintf(cli.errStream, "%v\n", err)
		return ExitCodeError
	}
	return ExitCodeOK
}
