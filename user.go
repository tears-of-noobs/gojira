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

// Return User structure of logged in Jira user
func Myself() (*User, error) {
	url := fmt.Sprintf("%s/myself", BaseUrl)
	return getUserHelper(execRequest("GET", url, nil))
}

// Get user structure by username
func GetUser(name string) (*User, error) {
	url := fmt.Sprintf("%s/user?%s", BaseUrl, name)
	return getUserHelper(execRequest("GET", url, nil))

}

func getUserHelper(code int, body []byte) (*User, error) {
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
