package gojira

import (
	"encoding/json"
	"fmt"
	"io"
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

type ProjectComponents struct {
	BaseFields
}

func requestHelper(requestType, url string, data io.Reader) (interface{}, error) {
	code, body := execRequest(requestType, url, data)
	if code == http.StatusOK {
		var result interface{}
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	} else {
		return nil, handleJiraError(body)
	}
}

func GetProjects() (*[]Project, error) {
	url := fmt.Sprintf("%s/project", BaseUrl)
	result, err := requestHelper("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return result.(*[]Project), nil
}

func GetProject(projectKey string) (*Project, error) {
	url := fmt.Sprintf("%s/project/%s", BaseUrl, projectKey)
	result, err := requestHelper("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return result.(*Project), nil
}

/*
func (project *Project) GetComponents() (*[]ProjectComponents, error) {
	url := fmt.Sprintf("%s/project/%s/components", BaseUrl, project.Key)
	result, err := requestHelper("GET", url, nil)
}
*/
