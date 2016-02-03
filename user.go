package gojira

import "fmt"

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
	url := fmt.Sprintf("%s/myself", BaseURL)
	return getUserHelper(execRequest("GET", url, nil))
}

// Get user structure by username
func GetUser(name string) (*User, error) {
	url := fmt.Sprintf("%s/user?%s", BaseURL, name)
	return getUserHelper(execRequest("GET", url, nil))

}
