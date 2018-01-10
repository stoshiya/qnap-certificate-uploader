# qnap-certificate-uploader

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
