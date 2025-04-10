package internal

import (
	"os"
	"os/exec"
	"strings"

	"github.com/leandrotocalini/gg/ui"
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
	ui.PrintWarningf("Command to execute: %s", fullCmd)

	if !ConfirmProceed() {
		ui.PrintErrorf("Aborted.")
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
	ui.PrintPlainf("Current branch: %s", branch)
	return branch == "main" || branch == "master"
}
