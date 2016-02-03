package gojira

import (
	"io"
	"io/ioutil"
	"net/http"
)

var BaseURL string
var Username string
var Password string

func execRequest(
	requestType, requestUrl string, data io.Reader,
) (int, []byte) {

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

// RAW request to Jira API. The function takes full URL of resource,
// type of request (GET, POST, PUT, ...) and body with required params.
// Return values are - responce body and HTTP status code, so you cam manually
// handle the answer from Jira and decode body to you own data type.
func RawRequest(url, requestType string, body io.Reader) (int, []byte) {
	return execRequest(requestType, url, body)

}
