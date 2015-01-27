package gojira

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	IssueFieldCreator
	TimeZone string `json:"timeZone"`
	Groups   Groups `json:"groups"`
	Expand   string `json:"expand"`
}

type Groups struct {
	Size  int                 `json:"size"`
	Items []map[string]string `json:"items"`
}

func Myself() (*User, error) {
	url := fmt.Sprintf("%s/myself", BaseUrl)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraUser User
		err := json.Unmarshal(body, &jiraUser)
		if err != nil {
			return nil, err
		}
		return &jiraUser, nil
	} else {
		return nil, handleJiraError(body)
	}
}
