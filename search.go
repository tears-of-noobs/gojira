package gojira

import (
	"fmt"
	"net/http"
)

func RawSearch(jql string) ([]byte, error) {
	url := fmt.Sprintf("%s/search?jql=%s", BaseUrl, jql)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		return body, nil
	} else {
		return nil, handleJiraError(body)
	}

}
