package handlers

import (
	"github.com/leandrotocalini/gg/ui"

	"github.com/leandrotocalini/gg/internal"
)

func Add(args []string) {
	if internal.IsProtectedBranch() {
		ui.PrintErrorf("Writing to 'main' or 'master' is not allowed.")
		return
	}
	internal.RunWithConfirm("git", append([]string{"add"}, args...)...)
}
