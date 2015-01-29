package gojira

import (
	"encoding/json"
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

type Filter struct {
	SearchUrl string `json:"searchUrl"`
}

func FilterSearch(id int) ([]byte, error) {
	url := fmt.Sprintf("%s/filter/%d", BaseUrl, id)
	code, body := execRequest("GET", url, nil)
	if code != http.StatusOK {
		return nil, handleJiraError(body)
	}
	filter := Filter{}
	err := json.Unmarshal(body, &filter)
	if err != nil {
		return nil, fmt.Errorf("error parsing filter: %s", err.Error())
	}

	code, body = execRequest("GET", filter.SearchUrl, nil)
	if code == http.StatusOK {
		return body, nil
	} else {
		return nil, handleJiraError(body)
	}
}
