package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//IssuesURL link
const IssuesURL = "https://api.github.com/search/issues"

//User details
type User struct {
	Login   string
	HTMLURL string
}

//Issues tracjed
type Issues struct {
	Number    int
	HTMLURL   string `json:"Html_Url"`
	Title     string
	State     string
	User      *User
	Createdat time.Time
	Body      string
}

//Issueresult informations
type Issueresult struct {
	Totalcount int
	Items      []*Issues
}

func main() {

}

// SearchIssues function
func SearchIssues(terms []string) (*Issueresult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search query failed: %s", resp.Status)
	}
	var result Issueresult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
