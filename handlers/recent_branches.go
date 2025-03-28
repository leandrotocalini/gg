package handlers

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func RecentBranches() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("Not a Git repository: %v\n", err)
	}

	refs, err := repo.Branches()
	if err != nil {
		log.Fatalf("Error getting branches: %v\n", err)
	}

	type branchInfo struct {
		Name    string
		Date    time.Time
		Message string
	}

	var branches []branchInfo

	err = refs.ForEach(func(ref *plumbing.Reference) error {
		commit, err := repo.CommitObject(ref.Hash())
		if err != nil {
			return nil // skip
		}

		msg := strings.Split(commit.Message, "\n")[0] // solo la primera línea
		branches = append(branches, branchInfo{
			Name:    ref.Name().Short(),
			Date:    commit.Author.When,
			Message: msg,
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(branches, func(i, j int) bool {
		return branches[i].Date.After(branches[j].Date)
	})

	fmt.Println("Recent branches:")
	max := 10
	for i, b := range branches {
		if i >= max {
			break
		}
		fmt.Printf("- %s: %s - %s\n",
			b.Name,
			b.Message,
			b.Date.Format("2006-01-02 15:04"),
		)
	}
}
