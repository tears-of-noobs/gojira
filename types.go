package gojira

type JiraSearchResult struct {
	StartAt    int         `json:"startAt"`
	MaxResults int         `json:"maxResults"`
	Total      int         `json:"total"`
	Issues     []JiraIssue `json:"issues"`
}

type JiraIssue struct {
	BaseFields
	Key    string          `json:"key"`
	Fields JiraIssueFields `json:fields"`
}

type BaseFields struct {
	Id   string `json:"id"`
	Self string `json:"self"`
}

type JiraIssueFields struct {
	Summary        string                 `json:"summary"`
	Progress       JiraIssueFieldProgress `json:"progress"`
	IssueType      JiraIssueType          `json:"issuetype"`
	ResolutionDate interface{}            `json:"resolutiondate"`
	Timespent      interface{}            `json:"timespent"`
	Creator        JiraIssueFieldCreator  `json:"creator"`
	Created        string                 `json:"created"`
	Updated        string                 `json:"updated"`
	Description    interface{}            `json:"description"`
	IssueLinks     []JiraIssueLink        `json:"issueLinks"`
	Status         JiraIssueStatus        `json:"status"`
}

type JiraIssueFieldProgress struct {
	Progress int `json:"progress"`
	Total    int `json:"total"`
}

type JiraIssueFieldCreator struct {
	Self         string            `json:"self"`
	Name         string            `json:"name"`
	EmailAddress string            `json:"emailAddress"`
	AvatarUrls   map[string]string `json:"avatarUrls"`
	DisplayName  string            `json:"displayName"`
	Active       bool              `json:"active"`
}

type JiraIssueType struct {
	BaseFields
	Description string `json:"description"`
	IconUrl     string `json:"iconURL"`
	Name        string `json:"name"`
	Subtask     bool   `json:"subtask"`
}

type JiraProject struct {
	BaseFields
}

type JiraIssueLink struct {
	BaseFields
	Type         map[string]string `json:"type"`
	InwardIssue  JiraIssue         `json:"inwardIssue"`
	OutwardIssue JiraIssue         `json:"outwardIssue"`
}

type JiraIssueStatus struct {
	BaseFields
	Name string `json:"name"`
}

type JiraTransition struct {
	BaseFields
	Name string               `json:"name"`
	To   JiraTransitionFields `json:"to"`
}

type JiraTransitionFields struct {
	BaseFields
	Description    string                 `json:"description"`
	IconUrl        string                 `json:"iconURL"`
	Name           string                 `json:"name"`
	StatusCategory map[string]interface{} `json:"statusCategory"`
}
