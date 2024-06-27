package main

import (
	"fmt"
	"time"

	"github.com/cli/go-gh/v2/pkg/api"
)

type PullRequest struct {
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Url         string    `json:"html_url"`
}

func getPullRequests(client api.RESTClient, org, repo string) ([]PullRequest, error) {
	var prs []PullRequest
	err := client.Get(fmt.Sprintf("repos/%s/%s/pulls", org, repo), &prs)
	if err != nil {
		return nil, err
	}
	return prs, nil
}
