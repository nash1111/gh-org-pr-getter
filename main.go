package main

import (
	"fmt"
	"os"
	"sort"

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

	var allPRs []PullRequest
	for _, repo := range repos {
		prs, err := getPullRequests(*client, org, repo.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		allPRs = append(allPRs, prs...)
	}

	sort.Slice(allPRs, func(i, j int) bool {
		return allPRs[i].CreatedAt.Before(allPRs[j].CreatedAt)
	})

	for _, pr := range allPRs {
		fmt.Printf("CreatedAt: %s\n", pr.CreatedAt)
	}
}
