package gojira

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Project struct {
	BaseFields
	Key             string            `json:"key"`
	Name            string            `json:"name"`
	AvatarUrls      map[string]string `json:"avatarUrls"`
	ProjectCategory ProjectCategory   `json:"projectCategory"`
	Description     string            `json:"description"`
	Lead            IssueFieldCreator `json:"lead"`
	Components      []interface{}     `json:"components"`
	IssueTypes      []IssueType       `json:"issueTypes"`
	Url             string            `json:"url"`
	Email           string            `json:"email"`
	AssigneeType    string            `json:"assigneeType"`
	Roles           map[string]string `json:"roles"`
}

type ProjectCategory struct {
	BaseFields
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetProjects() ([]*Project, error) {
	url := fmt.Sprintf("%s/project", BaseUrl)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var projects []*Project
		err := json.Unmarshal(body, &projects)
		if err != nil {
			return nil, err
		}
		return projects, nil
	} else {
		return nil, handleJiraError(body)
	}
}

func GetProject(projectKey string) (*Project, error) {
	url := fmt.Sprintf("%s/project/%s", BaseUrl, projectKey)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var project Project
		err := json.Unmarshal(body, &project)
		if err != nil {
			return nil, err
		}
		return &project, nil
	} else {
		return nil, handleJiraError(body)
	}

}
