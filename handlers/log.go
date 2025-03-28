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
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
