package handlers

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
)

func Status() {
	r, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("Not a Git repository: %v\n", err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	status, err := w.Status()
	if err != nil {
		log.Fatal(err)
	}

	if len(status) == 0 {
		fmt.Println("Working tree clean.")
		return
	}

	fmt.Println("Changes not staged for commit:")
	for file, s := range status {
		fmt.Printf("  %s: %s\n", s.Worktree, file)
	}
}
