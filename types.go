package gojira

import "strings"

type ApiError struct {
	ErrorMessages []string    `json:"errorMessages"`
	Errors        interface{} `json:"errors"`
}

func (e ApiError) String() string {
	return strings.Join(e.ErrorMessages, " ")
}

type JiraSearchIssues struct {
	SearchHead
	Issues []Issue `json:"issues"`
}
type Comments struct {
	SearchHead
	Comments []Comment `json:"comments"`
}

type Worklogs struct {
	SearchHead
	Worklogs []Worklog `json:worklogs"`
}

type SearchHead struct {
	StartAt    int `json:"startAt"`
	MaxResults int `json:"maxResults"`
	Total      int `json:"total"`
}

type Comment struct {
	BaseFields
	Author       IssueFieldCreator `json:"author"`
	Body         string            `json:"body"`
	UpdateAuthor IssueFieldCreator `json:"updateAuthor"`
	Created      string            `json:"created"`
	Updated      string            `json:"updated"`
	Visibility   map[string]string `json:"visibility"`
}

type Worklog struct {
	BaseFields
	Author           IssueFieldCreator `json:"author"`
	UpdateAuthor     IssueFieldCreator `json:"updateAuthor"`
	Comment          string            `json:"comment"`
	Visibility       map[string]string `json:"visibility"`
	Started          string            `json:"started"`
	TimeSpent        string            `json:"timeSpent"`
	TimeSpentSeconds int64             `json:"timeSpentSeconds"`
}

type BaseFields struct {
	Id   string `json:"id"`
	Self string `json:"self"`
}

type JiraProject struct {
	BaseFields
}

type IssueLink struct {
	BaseFields
	Type         map[string]string `json:"type"`
	InwardIssue  Issue             `json:"inwardIssue"`
	OutwardIssue Issue             `json:"outwardIssue"`
}

type Transitions struct {
	Expand      string       `json:"expand"`
	Transitions []Transition `json:"transitions"`
}

type Transition struct {
	BaseFields
	Name string           `json:"name"`
	To   TransitionFields `json:"to"`
}

type TransitionFields struct {
	BaseFields
	Description    string                 `json:"description"`
	IconUrl        string                 `json:"iconURL"`
	Name           string                 `json:"name"`
	StatusCategory map[string]interface{} `json:"statusCategory"`
}
