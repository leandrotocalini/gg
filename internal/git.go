package internal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func RunWithConfirm(name string, args ...string) error {
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

	return Run(name, args...)
}

func GetCurrentBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func IsProtectedBranch() bool {
	branch := strings.ToLower(strings.TrimSpace(GetCurrentBranch()))
	fmt.Println("Current branch:", branch)
	return branch == "main" || branch == "master"
}
