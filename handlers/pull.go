package handlers

import (
	"github.com/leandrotocalini/gg/internal"
	"github.com/leandrotocalini/gg/ui"
)

func Pull() {
	err := internal.RunWithConfirm("git", "pull")
	if err != nil {
		ui.PrintErrorf("Error: %s", err)
	}
}
