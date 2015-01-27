package gojira

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func handleJiraError(body []byte) error {
	errorAnswer := ApiError{}
	err := json.Unmarshal(body, &errorAnswer)
	if err != nil {
		return err
	}
	return errors.New(strings.Join(errorAnswer.ErrorMessages, " "))
}

func CreateIssue(params io.Reader) (*Issue, error) {
	url := fmt.Sprintf("%s/issue", BaseUrl)
	code, body := execRequest("POST", url, params)
	if code == http.StatusCreated {
		answer := make(map[string]string)
		err := json.Unmarshal(body, &answer)
		if err != nil {
			return nil, err
		}

		return GetIssue(answer["key"])
	} else {
		return nil, handleJiraError(body)
	}
}

func GetIssue(issueKey string) (*Issue, error) {
	url := fmt.Sprintf("%s/issue/%s", BaseUrl, issueKey)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var issue Issue
		err := json.Unmarshal(body, &issue)
		if err != nil {
			return nil, err
		}
		return &issue, nil
	} else {
		return nil, handleJiraError(body)
	}
}

func (issue *Issue) Assignee(name string) error {
	params := make(map[string]string)
	params["name"] = name
	b, err := json.Marshal(params)
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(b)
	url := fmt.Sprintf("%s/issue/%s/assignee", BaseUrl, issue.Key)
	code, body := execRequest("PUT", url, buff)
	if code == http.StatusNoContent {
		return nil
	} else {
		return handleJiraError(body)
	}
}

func (issue *Issue) GetWorklog() (*Worklogs, error) {
	url := fmt.Sprintf("%s/issue/%s/worklog", BaseUrl, issue.Key)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraWorklogs Worklogs
		err := json.Unmarshal(body, &jiraWorklogs)
		if err != nil {
			return nil, err
		}
		return &jiraWorklogs, nil
	} else {
		return nil, handleJiraError(body)
	}
}

func (issue *Issue) GetTransitions() (*Transitions, error) {
	url := fmt.Sprintf("%s/issue/%s/transitions", BaseUrl, issue.Key)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraTransitions Transitions
		err := json.Unmarshal(body, &jiraTransitions)
		if err != nil {
			return nil, err
		}
		return &jiraTransitions, nil
	} else {
		return nil, handleJiraError(body)
	}
}

func (issue *Issue) SetTransition(transition io.Reader) error {
	url := fmt.Sprintf("%s/issue/%s/transitions", BaseUrl, issue.Key)
	code, body := execRequest("POST", url, transition)
	if code == http.StatusNoContent {
		return nil
	} else {
		return handleJiraError(body)
	}

}

func (issue *Issue) GetComments() (*Comments, error) {
	url := fmt.Sprintf("%s/issue/%s/comment", BaseUrl, issue.Key)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraComments Comments
		err := json.Unmarshal(body, &jiraComments)
		if err != nil {
			return nil, err
		}
		return &jiraComments, nil
	} else {
		return nil, handleJiraError(body)
	}
}

func (issue *Issue) GetComment(id int) (*Comment, error) {
	url := fmt.Sprintf("%s/issue/%s/comment/%d", BaseUrl, issue.Key, id)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraComment Comment
		err := json.Unmarshal(body, &jiraComment)
		if err != nil {
			return nil, err
		}
		return &jiraComment, nil
	} else {
		return nil, handleJiraError(body)
	}
}

func (issue *Issue) SetComment(comment io.Reader) (*Comment, error) {
	url := fmt.Sprintf("%s/issue/%s/comment", BaseUrl, issue.Key)
	code, body := execRequest("POST", url, comment)
	if code == http.StatusCreated {
		var jiraComment Comment
		err := json.Unmarshal(body, &jiraComment)
		if err != nil {
			return nil, err
		}
		return &jiraComment, nil
	} else {
		return nil, handleJiraError(body)
	}
}

func (issue *Issue) UpdateComment(id int, comment io.Reader) (*Comment, error) {
	url := fmt.Sprintf("%s/issue/%s/comment/%d", BaseUrl, issue.Key, id)
	code, body := execRequest("PUT", url, comment)
	if code == http.StatusOK {
		var jiraComment Comment
		err := json.Unmarshal(body, &jiraComment)
		if err != nil {
			return nil, err
		}
		return &jiraComment, nil
	} else {
		return nil, handleJiraError(body)
	}
}

func (issue *Issue) DeleteComment(id int) error {
	url := fmt.Sprintf("%s/issue/%s/comment/%d", BaseUrl, issue.Key, id)
	code, body := execRequest("DELETE", url, nil)
	if code == http.StatusNoContent {
		return nil
	} else {
		return handleJiraError(body)
	}
}

/*
func UpdateIssue(params io.Reader, issueTag string) (int, []byte) {
	url := fmt.Sprintf("%s/issue/%s", BaseUrl, issueTag)
	return execRequest("PUT", url, params)
}

func DeleteIssue(issueTag string) (int, []byte) {
	url := fmt.Sprintf("%s/issue/%s", BaseUrl, issueTag)
	return execRequest("DELETE", url, nil)
}

*/
