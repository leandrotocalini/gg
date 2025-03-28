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

// Commands without confirmation

func handleStatus() {
	run("git", "status")
}

// Utility functions

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
