package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
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
	case "rb", "recent":
		handleRecentBranches()
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

func handleCommit(args []string) {
	msg := strings.Join(args, " ")
	runWithConfirm("git", "commit", "-am", msg)
}

func handlePush() {
	err := runWithConfirm("git", "push")
	if err != nil {
		// Detect specific upstream error
		if strings.Contains(err.Error(), "exit status 128") {
			// Detect current branch
			currentBranch := getCurrentBranch()
			if currentBranch != "" {
				fmt.Println("⚠️ No upstream branch set.")
				fmt.Printf("Suggesting: git push --set-upstream origin %s\n", currentBranch)
				fmt.Print("Run this now? [y/N]: ")

				reader := bufio.NewReader(os.Stdin)
				resp, _ := reader.ReadString('\n')
				resp = strings.TrimSpace(strings.ToLower(resp))

				if resp == "y" || resp == "yes" {
					run("git", "push", "--set-upstream", "origin", currentBranch)
				} else {
					fmt.Println("You can run it manually:")
					fmt.Printf("git push --set-upstream origin %s\n", currentBranch)
				}
			}
		} else {
			fmt.Println("Error:", err)
		}
	}
}

func getCurrentBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func handleAdd(args []string) {
	runWithConfirm("git", append([]string{"add"}, args...)...)
}

func handleCheckout(args []string) {
	runWithConfirm("git", append([]string{"checkout"}, args...)...)
}

func handleNewBranch(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: gg nb <branch-name>")
		return
	}

	raw := strings.Join(args, " ")
	cleaned := cleanBranchName(raw)

	fmt.Println("Branch to create:", cleaned)
	fmt.Print("Proceed? [y/N]: ")

	reader := bufio.NewReader(os.Stdin)
	resp, _ := reader.ReadString('\n')
	resp = strings.TrimSpace(strings.ToLower(resp))

	if resp != "y" && resp != "yes" {
		fmt.Println("Aborted.")
		return
	}

	run("git", "checkout", "-b", cleaned)
}

func cleanBranchName(input string) string {
	// Trim spaces and lowercase
	input = strings.ToLower(strings.TrimSpace(input))

	// If there's a slash, split and clean each part separately
	if strings.Contains(input, "/") {
		parts := strings.SplitN(input, "/", 2)
		return cleanSegment(parts[0]) + "/" + cleanSegment(parts[1])
	}

	return cleanSegment(input)
}

func cleanSegment(s string) string {
	s = strings.TrimSpace(s)
	// Replace all underscores and spaces with hyphens
	s = strings.ReplaceAll(s, "_", "-")
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, "-")
	s = regexp.MustCompile(`-+`).ReplaceAllString(s, "-") // no double dashes
	s = strings.Trim(s, "-")
	return s
}

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

func handleRecentBranches() {
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

func runWithConfirm(name string, args ...string) error {
	fullCmd := name + " " + strings.Join(args, " ")
	fmt.Println("Command to execute:", fullCmd)
	fmt.Print("Proceed? [y/N]: ")

	reader := bufio.NewReader(os.Stdin)
	resp, _ := reader.ReadString('\n')
	resp = strings.TrimSpace(strings.ToLower(resp))

	if resp != "y" && resp != "yes" {
		fmt.Println("Aborted.")
		return nil
	}

	return run(name, args...)
}

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
