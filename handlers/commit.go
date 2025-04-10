package handlers

import (
	"strings"

	"github.com/leandrotocalini/gg/internal"
	"github.com/leandrotocalini/gg/ui"
)

func Commit(args []string) {
	msg := strings.Join(args, " ")
	if internal.IsProtectedBranch() {
		ui.PrintError("❌ Writing to 'main' or 'master' is not allowed.")
		return
	}
	internal.RunWithConfirm("git", "commit", "-am", msg)
}
