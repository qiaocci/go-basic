package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {
	res, err := SearchIssues([]string{"python", "java", "中文"})
	if err != nil {
		fmt.Println(err)
		return
	}

	pastDay := make([]*Issue, 0)
	pastMonth := make([]*Issue, 0)
	now := time.Now()
	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)

	for _, item := range res.Items {
		switch {
		case item.CreatedAt.After(day):
			pastDay = append(pastDay, item)
		case item.CreatedAt.After(month):
			pastMonth = append(pastMonth, item)
		}
	}

	if len(pastDay) > 0 {
		fmt.Println("Past Day:")
		for _, item := range pastDay {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastMonth) > 0 {
		fmt.Println("Past Month:")
		for _, item := range pastMonth {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Printf("q=%v", q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
