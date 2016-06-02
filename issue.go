package gojira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Issue struct {
	BaseFields
	Key    string      `json:"key"`
	Fields IssueFields `json:fields"`
}

type IssueFields struct {
	Summary        string             `json:"summary"`
	Progress       IssueFieldProgress `json:"progress"`
	IssueType      IssueType          `json:"issuetype"`
	ResolutionDate interface{}        `json:"resolutiondate"`
	Timespent      interface{}        `json:"timespent"`
	Creator        IssueFieldCreator  `json:"creator"`
	Created        string             `json:"created"`
	Updated        string             `json:"updated"`
	Labels         []string           `json:"labels"`
	Assignee       IssueFieldCreator  `json:"assignee"`
	Description    interface{}        `json:"description"`
	IssueLinks     []IssueLink        `json:"issueLinks"`
	Status         IssueStatus        `json:"status"`
}

type IssueFieldProgress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

type IssueFieldCreator struct {
	Self         string            `json:"self"`
	Name         string            `json:"name"`
	EmailAddress string            `json:"emailAddress"`
	AvatarUrls   map[string]string `json:"avatarUrls"`
	DisplayName  string            `json:"displayName"`
	Active       bool              `json:"active"`
}

type IssueType struct {
	BaseFields
	Description string `json:"description"`
	IconUrl     string `json:"iconURL"`
	Name        string `json:"name"`
	Subtask     bool   `json:"subtask"`
}

type IssueStatus struct {
	BaseFields
	Name string `json:"name"`
}

func CreateIssue(params io.Reader) (*Issue, error) {
	url := fmt.Sprintf("%s/issue", BaseURL)
	code, body := execRequest("POST", url, params)
	if code == http.StatusCreated {
		response := make(map[string]string)
		err := json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}

		return GetIssue(response["key"])
	}
	return nil, handleJiraError(body)
}

// GetIssue - return issue by key
func GetIssue(issueKey string) (*Issue, error) {
	url := fmt.Sprintf("%s/issue/%s", BaseURL, issueKey)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var issue Issue
		err := json.Unmarshal(body, &issue)
		if err != nil {
			return nil, err
		}
		return &issue, nil
	}
	return nil, handleJiraError(body)
}

// GetLabels - return labels from issue
func (issue *Issue) GetLabels() []string {
	return issue.Fields.Labels
}

// AddLabel - add label to issue
func (issue *Issue) AddLabel(labels []string) error {
	for i, val := range labels {
		labels[i] = fmt.Sprintf(`{"add": "%s"}`, val)
	}
	return updateLabelsHelper(labels, issue.Key)
}

// Remove labels to issue
func (issue *Issue) RemoveLabel(labels []string) error {
	for i, val := range labels {
		labels[i] = fmt.Sprintf(`{"remove": "%s"}`, val)
	}

	return updateLabelsHelper(labels, issue.Key)
}

// Assignee - assign issue to another name
func (issue *Issue) Assignee(name string) error {
	encodedParams, err := json.Marshal(map[string]string{"name": name})
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/issue/%s/assignee", BaseURL, issue.Key)
	code, body := execRequest("PUT", url, bytes.NewBuffer(encodedParams))
	if code == http.StatusNoContent {
		return nil
	}
	return handleJiraError(body)
}

// GetWorklogs - get all worklogs from issue
func (issue *Issue) GetWorklogs() (*Worklogs, error) {
	url := fmt.Sprintf("%s/issue/%s/worklog", BaseURL, issue.Key)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraWorklogs Worklogs
		err := json.Unmarshal(body, &jiraWorklogs)
		if err != nil {
			return nil, err
		}
		return &jiraWorklogs, nil
	}
	return nil, handleJiraError(body)
}

// GetWorklog - return worklog from issue by ID
func (issue *Issue) GetWorklog(id int) (*Worklog, error) {
	url := fmt.Sprintf("%s/issue/%s/worklog/%d", BaseURL, issue.Key, id)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraWorklog Worklog
		err := json.Unmarshal(body, &jiraWorklog)
		if err != nil {
			return nil, err
		}
		return &jiraWorklog, nil
	}
	return nil, handleJiraError(body)
}

// SetWorklog - logging work in issue with comment
func (issue *Issue) SetWorklog(timeSpent, comment string) error {
	worklog := map[string]string{
		"timeSpent": timeSpent,
		"comment":   comment,
	}
	encodedParams, err := json.Marshal(worklog)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/issue/%s/worklog", BaseURL, issue.Key)
	code, body := execRequest("POST", url, bytes.NewBuffer(encodedParams))
	if code == http.StatusCreated {
		return nil
	}
	return handleJiraError(body)
}

// GetTransitions - return available transitions for issue
func (issue *Issue) GetTransitions() (*Transitions, error) {
	url := fmt.Sprintf("%s/issue/%s/transitions", BaseURL, issue.Key)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraTransitions Transitions
		err := json.Unmarshal(body, &jiraTransitions)
		if err != nil {
			return nil, err
		}
		return &jiraTransitions, nil
	}
	return nil, handleJiraError(body)
}

func (issue *Issue) SetTransition(transition io.Reader) error {
	url := fmt.Sprintf("%s/issue/%s/transitions", BaseURL, issue.Key)
	code, body := execRequest("POST", url, transition)
	if code == http.StatusNoContent {
		return nil
	}
	return handleJiraError(body)
}

// GetComments - return all comments for issue
func (issue *Issue) GetComments() (*Comments, error) {
	url := fmt.Sprintf("%s/issue/%s/comment", BaseURL, issue.Key)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraComments Comments
		err := json.Unmarshal(body, &jiraComments)
		if err != nil {
			return nil, err
		}
		return &jiraComments, nil
	}
	return nil, handleJiraError(body)
}

// GetComment - return one comment by ID
func (issue *Issue) GetComment(id int) (*Comment, error) {
	url := fmt.Sprintf("%s/issue/%s/comment/%d", BaseURL, issue.Key, id)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraComment Comment
		err := json.Unmarshal(body, &jiraComment)
		if err != nil {
			return nil, err
		}
		return &jiraComment, nil
	}
	return nil, handleJiraError(body)
}

// SetComment - add comment in issue
func (issue *Issue) SetComment(comment *Comment) (*Comment, error) {
	url := fmt.Sprintf("%s/issue/%s/comment", BaseURL, issue.Key)
	encodedParams, err := json.Marshal(comment)
	if err != nil {
		return nil, err
	}

	code, body := execRequest("POST", url, bytes.NewBuffer(encodedParams))
	if code == http.StatusCreated {
		var jiraComment Comment
		err := json.Unmarshal(body, &jiraComment)
		if err != nil {
			return nil, err
		}
		return &jiraComment, nil
	}
	return nil, handleJiraError(body)
}

// UpdateComment - update existing comment by id
func (issue *Issue) UpdateComment(
	id int, comment *Comment,
) (*Comment, error) {

	url := fmt.Sprintf("%s/issue/%s/comment/%d", BaseURL, issue.Key, id)
	encodedParams, err := json.Marshal(comment)
	if err != nil {
		return nil, err
	}
	code, body := execRequest("PUT", url, bytes.NewBuffer(encodedParams))
	if code == http.StatusOK {
		var jiraComment Comment
		err := json.Unmarshal(body, &jiraComment)
		if err != nil {
			return nil, err
		}
		return &jiraComment, nil
	}
	return nil, handleJiraError(body)
}

// DeleteComment - remove comment from issue
func (issue *Issue) DeleteComment(id int64) error {
	url := fmt.Sprintf("%s/issue/%s/comment/%d", BaseURL, issue.Key, id)
	code, body := execRequest("DELETE", url, nil)
	if code == http.StatusNoContent {
		return nil
	}
	return handleJiraError(body)
}

func (issue *Issue) Delete() error {
	url := fmt.Sprintf("%s/issue/%s", BaseURL, issue.Key)
	code, body := execRequest("DELETE", url, nil)
	if code != http.StatusNoContent {
		return handleJiraError(body)
	}
	return nil
}
