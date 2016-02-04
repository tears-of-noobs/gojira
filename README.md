# gojira                                                                                                                             
Simple Atlassian JIRA API client implementation                                                                                      
                                                                                                                                     
# Usage                                                                                                                              
* Install package                                                                                                               
``` sh                                                                                                                               
go get github.com/tears-of-noobs/gojira                                                                                              
```                                                                                                                                  
* Import in you source code                                                                                                     
```go                                                                                                                                
import "github.com/tears-of-noobs/gojira"                                                                                            
```                                                                                                                                  
* Set initial parameters                                                                                                        
```go                                                                                                                                
gojira.Username = "USERNAME"                                                                                                         
gojira.Password = "PASSWORD"                                                                                                         
gojira.BaseUrl = "http://JIRA_URL/rest/api/2"                                                                                        
```                                                                                                                                  
                                                                                                                                     
# Examples                                                                                                                           
                                                                                                                                     
#### Searching

Search issues in project TEST and assignee to USERNAME, order by JIRA key
```go
jqlString := "project%20%3D%20TEST"%20AND%20assignee%20%3D%20USERNAME%20order%20by%20key%20DESC"
result, err := gojira.RawSearch(jqlString)
if err != nil {
    fmt.Println(err)
    }
var jiraIssues gojira.JiraSearchIssue
err = json.Unmarshal(result, &jiraIssues)
if err != nil {
    fmt.Println(err)
    }
for _, issue := range jiraIssues.Issues {
    fmt.Printf("%15s %13s %s\n", issue.Key,
         issue.Fields.Status.Name, issue.Fields.Summary)
    }

```

#### Get issue by issue key and view all comments:                                                                                   
```go
issue, err := gojira.GetIssue("TEST-123")                                                                                            
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
comments, err := issue.GetComments()                                                                                                 
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    }                                                                                                                                
for _, comment := range comments.Comments {                                                                                          
    fmt.Println(comment.Comment)                                                                                                     
}                                                                                                                                    
```                                                                                                                                  
                                                                                                                                     
#### Get comment by comment id                                                                                                       
```go                                                                                                                                 
issue, err := gojira.GetIssue("TEST-123")                                                                                            
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
comment, err := issue.GetComment(1000)                                                                                               
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
fmt.Printf("%v\n", comment)                                                                                                          
```                                                                                                                                  
                                                                                                                                     
#### Write comment to issue                                                                                                          
                                                                                                                                     
```go                                                                                                                                 
issue, err := gojira.GetIssue("TEST-123")                                                                                            
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
// Prepare comment                                                                                                                   
newComment := &gojira.Comment{Body: "test comment"} 

comment, err := issue.SetComment(newComment)                                                                                        
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
}                                                                                                    
```

#### Update existing comment in issue

// You also may use UpdateComment(id int, comment io.Reader) to update exist comment by id                                           
```go                                                                                                                                 
issue, err := gojira.GetIssue("TEST-123")                                                                                            
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
// Prepare comment                                                                                                                   
newComment := &gojira.Comment{Body: "test comment"} 

comment, err := issue.UpdateComment(123456, newComment)                                                                                        
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
fmt.Printf("%v\n", comment.Comment)                                                                                                  
```                                                                                                                                  
                                                                                                                                     
#### Delete comment                                                                                                                  
```go                                                                                                                                 
issue, err := gojira.GetIssue("TEST-123")                                                                                            
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
err := issue.DeleteComment(1000)                                                                                                     
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
fmt.Println("Comment deleted")                                                                                                       
```                                                                                                                                  


#### Logging work
```go
issue, err := gojira.GetIssue("TEST-123")                                                                                            
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
err := issue.SetWorklog("1h 30m", "I`m working")
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
```
                                                                                                                                     
# TODO's                                                                                                                             
                                                                                                                                     
 - Documentaton                                                                                                                  
 - Write tests                                                                                                                   
 - Huge coverage                                                                                                                 
                                              
