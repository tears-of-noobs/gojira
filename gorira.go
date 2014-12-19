package gojira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type JiraClient struct {
	BaseUrl   string
	Username  string
	Password  string
	ProjectId string
}

func NewJiraClient(baseUrl, username, password string) *JiraClient {
	return &JiraClient{
		BaseUrl:  baseUrl,
		Username: username,
		Password: password}
}

func (this *JiraClient) SetProjectId(projectId string) {
	this.ProjectId = projectId
}

func (this *JiraClient) execRequest(requestType, requestUrl string,
	data io.Reader) (int, []byte) {

	client := &http.Client{}
	req, err := http.NewRequest(requestType, requestUrl, data)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(this.Username, this.Password)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return resp.StatusCode, body
}

func (this *JiraClient) Search(jql string, fields []string) (int, []byte) {
	fld := strings.Join(fields, ",")
	url := fmt.Sprintf("%s/search?jql=%s&fields=%s",
		this.BaseUrl, jql, fld)
	return this.execRequest("GET", url, nil)
}

func (this *JiraClient) CreateIssue(params io.Reader) (int, []byte) {
	url := fmt.Sprintf("%s/issue", this.BaseUrl)
	return this.execRequest("POST", url, params)
}

func (this *JiraClient) UpdateIssue(params io.Reader, issueTag string) (int, []byte) {
	url := fmt.Sprintf("%s/issue/%s", this.BaseUrl, issueTag)
	return this.execRequest("PUT", url, params)
}

func (this *JiraClient) DeleteIssue(issueTag string) (int, []byte) {
	url := fmt.Sprintf("%s/issue/%s", this.BaseUrl, issueTag)
	return this.execRequest("DELETE", url, nil)
}

func (this *JiraClient) GetIssue(issueTag string) (int, []byte) {
	url := fmt.Sprintf("%s/issue/%s", this.BaseUrl, issueTag)
	return this.execRequest("GET", url, nil)
}

func (this *JiraClient) Transitions(issueTag string) (int, []byte) {
	url := fmt.Sprintf("%s/issue/%s/transitions", this.BaseUrl, issueTag)
	return this.execRequest("GET", url, nil)
}

func (this *JiraClient) SetTransition(params io.Reader, issueTag string) (int, []byte) {
	url := fmt.Sprintf("%s/issue/%s/transitions", this.BaseUrl, issueTag)
	return this.execRequest("POST", url, params)
}

func (this *JiraClient) CreateLink(inwardIssue, outwardIssue string,
	typeParams map[string]string) (int, []byte) {
	inw := map[string]string{
		"key": inwardIssue}
	outw := map[string]string{
		"key": outwardIssue}
	reqBody := struct {
		Type         map[string]string `json:"type"`
		InwardIssue  map[string]string `json:"inwardIssue"`
		OutwardIssue map[string]string `json:"outwardIssue"`
	}{typeParams, inw, outw}

	b, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	buff := bytes.NewBuffer(b)

	url := fmt.Sprintf("%s/issueLink", this.BaseUrl)
	return this.execRequest("POST", url, buff)
}

func (this *JiraClient) DeleteLink(linkId string) (int, []byte) {
	url := fmt.Sprintf("%s/issueLink/%s", this.BaseUrl, linkId)
	return this.execRequest("DELETE", url, nil)
}
