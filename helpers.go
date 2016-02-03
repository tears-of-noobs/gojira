package gojira

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func handleJiraError(body []byte) error {
	errorAnswer := ApiError{}
	err := json.Unmarshal(body, &errorAnswer)
	if err != nil {
		return err
	}
	return errors.New(errorAnswer.String())
}

func updateLabelsHelper(labels []string, issueKey string) error {
	updateParams := []byte(fmt.Sprintf(`{ "update": { "labels": [ %s ] } }`,
		strings.Join(labels, ", ")))
	url := fmt.Sprintf("%s/issue/%s", BaseURL, issueKey)
	code, body := execRequest("PUT", url, bytes.NewBuffer(updateParams))
	if code == http.StatusNoContent {
		return nil
	}
	return handleJiraError(body)
}

func getUserHelper(code int, body []byte) (*User, error) {
	if code == http.StatusOK {
		var jiraUser User
		err := json.Unmarshal(body, &jiraUser)
		if err != nil {
			return nil, err
		}
		return &jiraUser, nil
	}
	return nil, handleJiraError(body)
}
