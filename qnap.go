package main

import (
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type QDocRoot struct {
	AuthSid string `xml:"authSid"`
}

func Auth(host string, user string, password string) (string, error) {
	values := url.Values{}
	values.Add("user", user)
	values.Add("pwd", base64.StdEncoding.EncodeToString([]byte(password)))

	req, err := http.NewRequest("POST","https://" + host + "/cgi-bin/authLogin.cgi", strings.NewReader(values.Encode()))
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	data := new(QDocRoot)
	if err := xml.Unmarshal(b, data); err != nil {
		return "", err
	}
	return data.AuthSid, nil
}

func Upload(host string, sid string, dir string) error {
	cert, err := ioutil.ReadFile(dir + "/cert.pem")
	if err != nil {
		return err
	}
	key, err := ioutil.ReadFile(dir + "/privkey.pem")
	if err != nil {
		return err
	}
	chain, err := ioutil.ReadFile(dir + "/chain.pem")
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Add("sid", sid)
	values.Add("certificate_content", string(cert))
	values.Add("key_content", string(key))
	values.Add("ic_update", "update")
	values.Add("issuer_certificate_content", string(chain))

	req, err := http.NewRequest("POST","https://" + host + "/cgi-bin/sys/sysRequest.cgi?&subfunc=security&apply=1&action=ssl&todo=upload", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	defer req.Body.Close()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
