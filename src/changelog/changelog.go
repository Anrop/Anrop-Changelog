package changelog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

func outputChangelog(w http.ResponseWriter) error {
	client := github.NewClient(nil)
	ctx := context.Background()
	org := "Anrop"

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(ctx, org, opt)

	if err != nil {
		return err
	}

	allCommits := []*github.RepositoryCommit{}

	for _, repo := range repos {
		commits, _, err := client.Repositories.ListCommits(ctx, org, *repo.Name, nil)

		if err != nil {
			continue
		}

		for _, commit := range commits {
			allCommits = append(allCommits, commit)
		}
	}

	return json.NewEncoder(w).Encode(allCommits)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := outputChangelog(w); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "Error reading changelog: %v", err)
	}
}
