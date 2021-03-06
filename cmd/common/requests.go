package common

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"net/http"
)



var client = &http.Client{}




func Execute(req *http.Request) ([]byte) {
	resp, err := client.Do(req)
	defer resp.Body.Close()
	PanicIfNonEmpty(err,resp)
	body, err := ioutil.ReadAll(resp.Body)
	PanicIfNonEmpty(err, resp)
	return body
}

func BuildPostRequest(path string,body io.Reader) (*http.Request) {
	req := buildRequest(path,"POST",body)
	req.Header.Add("Content-Type", "application/json")
	return req
}

func BuilGetRequest(path string) (*http.Request) {
	return buildRequest(path,"GET",nil)
}

func buildRequest(path string, method string, body io.Reader) *http.Request {
	domain := viper.GetString("jira_domain")
	user := viper.GetString("jira_user")
	pass := viper.GetString("jira_pass")
	url := fmt.Sprintf("https://%s%s", domain, path)
	req, err := http.NewRequest(method, url, body)
	PanicIfNonEmpty(err, nil)
	creds := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pass)))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", creds))
	return req
}
