package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gg [command]")
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	handleCommand(cmd, args)
}

func handleCommand(cmd string, args []string) {
	switch cmd {
	case "c":
		handleCommit(args)
	case "p":
		handlePush()
	case "s":
		handleStatus()
	case "a":
		handleAdd(args)
	case "co":
		handleCheckout(args)
	case "nb":
		handleNewBranch(args)
	case "l":
		handleLog()
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

// Commands with confirmation

func handleCommit(args []string) {
	msg := strings.Join(args, " ")
	runWithConfirm("git", "commit", "-am", msg)
}

func handlePush() {
	runWithConfirm("git", "push")
}

func handleAdd(args []string) {
	runWithConfirm("git", append([]string{"add"}, args...)...)
}

func handleCheckout(args []string) {
	runWithConfirm("git", append([]string{"checkout"}, args...)...)
}

func handleNewBranch(args []string) {
	runWithConfirm("git", append([]string{"checkout", "-b"}, args...)...)
}

// No confirmation

func handleStatus() {
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

func handleLog() {
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

// Utils

func runWithConfirm(name string, args ...string) {
	fullCmd := name + " " + strings.Join(args, " ")
	fmt.Println("Command to execute:", fullCmd)
	fmt.Print("Proceed? [y/N]: ")

	reader := bufio.NewReader(os.Stdin)
	resp, _ := reader.ReadString('\n')
	resp = strings.TrimSpace(strings.ToLower(resp))

	if resp != "y" && resp != "yes" {
		fmt.Println("Aborted.")
		return
	}

	run(name, args...)
}

func run(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
