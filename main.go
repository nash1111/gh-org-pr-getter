package main

import (
	"fmt"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <organization>")
		return
	}
	org := os.Args[1]

	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	var repos []struct {
		Name string `json:"name"`
	}

	err = client.Get(fmt.Sprintf("orgs/%s/repos", org), &repos)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Repositories in organization %s:\n", org)
	for _, repo := range repos {
		fmt.Println(repo.Name)
		prs, err := getPullRequests(*client, org, repo.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Pull Requests for %s:\n", repo.Name)
		for _, pr := range prs {
			// fmt.Println(pr.Title)
			fmt.Println(pr)
		}
	}
}

type PullRequest struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	Description string `json:"description"`
}

func getPullRequests(client api.RESTClient, org, repo string) ([]PullRequest, error) {
	var prs []PullRequest
	err := client.Get(fmt.Sprintf("repos/%s/%s/pulls", org, repo), &prs)
	if err != nil {
		return nil, err
	}
	return prs, nil
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
