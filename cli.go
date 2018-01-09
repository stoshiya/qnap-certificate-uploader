package main

import (
	"fmt"
	"io"
	"flag"
)

const (
	ExitCodeOK int = iota
	ExitCodeParseFlagError
	ExitCodeUsage
	ExitCodeError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var (
		err error
		version bool
		user, password, host, dir string
	)
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.BoolVar(&version, "version", false, "print version information and quit")
	flags.StringVar(&user, "user", "", "NAS account name")
	flags.StringVar(&password, "password", "", "NAS account password")
	flags.StringVar(&host, "host", "", "NAS Hostname (example \"nas.example.com\")")
	flags.StringVar(&dir, "dir", "/usr/local/etc/letsencrypt/live/", "Base directory for let's encrypt")

	if err = flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if user == "" || password == "" || host == "" {
		flags.Usage()
		return ExitCodeUsage
	}

	sid, err := Auth(host, user, password)
	if err != nil {
		fmt.Fprintf(cli.errStream, "%v", err)
		return ExitCodeError
	}
	fmt.Println(sid)

	if err := Upload(host, sid, dir + host); err != nil {
		fmt.Fprintf(cli.errStream, "%v", err)
		return ExitCodeError
	}
	return ExitCodeOK
}

