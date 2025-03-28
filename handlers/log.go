package handlers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func Log() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("Not a Git repository: %v\n", err)
	}

	ref, err := repo.Head()
	if err != nil {
		log.Fatalf("Error getting HEAD: %v\n", err)
	}

	iter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		log.Fatalf("Error getting commit log: %v\n", err)
	}

	fmt.Println("Recent commits:")
	err = iter.ForEach(func(c *object.Commit) error {
		fmt.Printf("\nAuthor:  %s\n", c.Author.Name)
		fmt.Printf("Date:    %s\n", c.Author.When.Format(time.RFC1123))
		fmt.Printf("Message: %s\n", strings.TrimSpace(c.Message))
		fmt.Printf("Commit:  %s\n", c.Hash.String()[:7])

		stats, err := getDiffStats(c)
		if err == nil {
			fmt.Printf("Changes: +%d -%d\n", stats.added, stats.deleted)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

type diffStats struct {
	added   int
	deleted int
}

func getDiffStats(c *object.Commit) (diffStats, error) {
	var stats diffStats

	if c.NumParents() == 0 {
		return stats, nil // root commit, no parent to compare
	}

	parent, err := c.Parents().Next()
	if err != nil {
		return stats, err
	}

	currentTree, err := c.Tree()
	if err != nil {
		return stats, err
	}

	parentTree, err := parent.Tree()
	if err != nil {
		return stats, err
	}

	patch, err := parentTree.Patch(currentTree)
	if err != nil {
		return stats, err
	}

	for _, fileStat := range patch.Stats() {
		stats.added += fileStat.Addition
		stats.deleted += fileStat.Deletion
	}

	return stats, nil
}
