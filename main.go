package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gg [command]")
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	var gitCmd []string

	switch cmd {
	case "c": // Commit
		msg := strings.Join(args, " ")
		gitCmd = []string{"commit", "-am", msg}

	case "p": // Push
		gitCmd = []string{"push"}

	case "s": // Status
		gitCmd = []string{"status"}

	case "a": // Add
		gitCmd = append([]string{"add"}, args...)

	case "co": // Checkout
		gitCmd = append([]string{"checkout"}, args...)

	case "nb": // New branch
		gitCmd = append([]string{"checkout", "-b"}, args...)

	default:
		fmt.Println("Unknown command:", cmd)
		return
	}

	// Show the command and ask for confirmation
	fullCmd := "git " + strings.Join(gitCmd, " ")
	fmt.Println("Command to execute:", fullCmd)
	fmt.Print("Proceed? [y/N]: ")

	reader := bufio.NewReader(os.Stdin)
	resp, _ := reader.ReadString('\n')
	resp = strings.TrimSpace(strings.ToLower(resp))

	if resp != "y" && resp != "yes" {
		fmt.Println("Aborted.")
		return
	}

	run("git", gitCmd...)
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
