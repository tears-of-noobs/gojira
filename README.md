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
                                                                                                                                     
#### Get issue by issue key and view all comments:                                                                                   
```                                                                                                                                  
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
```                                                                                                                                  
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
                                                                                                                                     
```                                                                                                                                  
issue, err := gojira.GetIssue("TEST-123")                                                                                            
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
// Prepare comment                                                                                                                   
var b = []byte(`{                                                                                                                    
    "body": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque eget venenatis elit. Duis eu justo eget augue iaculis fermentum. Sed semper quam laoreet nisi egestas at posuere augue semper.",
    "visibility": {                                                                                                                  
        "type": "role",                                                                                                              
        "value": "Administrators"                                                                                                    
    }                                                                                                                                
}`)                                                                                                                                  
newComment := bytes.NewBuffer(b)                                                                                                     
comment, err := issue.SetComments(newComment)                                                                                        
// You also may use UpdateComment(id int, comment io.Reader) to update exist comment by id                                           
if err != nil {                                                                                                                      
    fmt.Printf("%s\n", err)                                                                                                          
    os.Exit(1)                                                                                                                       
    }                                                                                                                                
fmt.Printf("%v\n", comment.Comment)                                                                                                  
```                                                                                                                                  
                                                                                                                                     
#### Delete comment                                                                                                                  
```                                                                                                                                  
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
                                                                                                                                     
# TODO's                                                                                                                             
                                                                                                                                     
 - Documentaton                                                                                                                  
 - Write tests                                                                                                                   
 - Huge coverage                                                                                                                 
                                              
