# qnap-certificate-uploader

[![Build Status](https://travis-ci.org/stoshiya/qnap-certificate-uploader.svg?branch=master)](https://travis-ci.org/stoshiya/qnap-certificate-uploader)
[![Coverage Status](https://coveralls.io/repos/github/stoshiya/qnap-certificate-uploader/badge.svg?branch=master)](https://coveralls.io/github/stoshiya/qnap-certificate-uploader?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/stoshiya/qnap-certificate-uploader)](https://goreportcard.com/report/github.com/stoshiya/qnap-certificate-uploader)

Let's encrypt certificates upload client for QNAP NAS.

## Build
```sh
$ go get github.com/stoshiya/qnap-certificate-uploader
$ cd $GOPATH/src/github.com/stoshiya/qnap-certificate-uploader
$ make
```

## Usage
```
Usage of qnap-certificate-uploader:
  -dir string
      Base directory for let's encrypt certificates. (default "/usr/local/etc/letsencrypt/live/")
  -password string
      NAS account password.
  -url string
      NAS protocol and host for web access. Port is optional. (example "http://nas.example.com:8080")
  -user string
      NAS account name.
  -version
      Print version information and quit.
```

## Example
```sh
$ qnap-certificate-uploader -user admin -password secret -url http://nas.example.com:8080
```

## License

MIT
