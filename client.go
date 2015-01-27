package gojira

import (
	"io"
	"io/ioutil"
	"net/http"
)

var BaseUrl string
var Username string
var Password string

func execRequest(requestType, requestUrl string,
	data io.Reader) (int, []byte) {

	client := &http.Client{}
	req, err := http.NewRequest(requestType, requestUrl, data)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(Username, Password)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return resp.StatusCode, body
}
