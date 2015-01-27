package gojira

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func handleJiraError(body []byte) error {
	errorAnswer := ApiError{}
	err := json.Unmarshal(body, &errorAnswer)
	if err != nil {
		return err
	}
	return errors.New(errorAnswer.String())
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

// Assign issue to another name
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

// Get all worklogs of issue
func (issue *Issue) GetWorklogs() (*Worklogs, error) {
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

// Get worklog of issue by ID
func (issue *Issue) GetWorklog(id int) (*Worklog, error) {
	url := fmt.Sprintf("%s/issue/%s/worklog/%d", BaseUrl, issue.Key, id)
	code, body := execRequest("GET", url, nil)
	if code == http.StatusOK {
		var jiraWorklog Worklog
		err := json.Unmarshal(body, &jiraWorklog)
		if err != nil {
			return nil, err
		}
		return &jiraWorklog, nil
	} else {
		return nil, handleJiraError(body)
	}
}

// Logging work in issue with comment
func (issue *Issue) SetWorklog(timeSpent, comment string) error {
	worklog := make(map[string]string)
	worklog["timeSpent"] = timeSpent
	worklog["comment"] = comment
	b, err := json.Marshal(worklog)
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(b)
	url := fmt.Sprintf("%s/issue/%s/worklog", BaseUrl, issue.Key)
	code, body := execRequest("POST", url, buff)
	if code == http.StatusCreated {
		return nil
	} else {
		return handleJiraError(body)
	}
}

// Return available transitions for issue
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

// Return all comments for issue
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

// Return one comment by ID
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

// Add comment in issue
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

// Update existing comment by id
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

// Remove comment from issue
func (issue *Issue) DeleteComment(id int) error {
	url := fmt.Sprintf("%s/issue/%s/comment/%d", BaseUrl, issue.Key, id)
	code, body := execRequest("DELETE", url, nil)
	if code == http.StatusNoContent {
		return nil
	} else {
		return handleJiraError(body)
	}
}
