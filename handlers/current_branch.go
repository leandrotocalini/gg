package handlers

import (
	"github.com/leandrotocalini/gg/internal"
	"github.com/leandrotocalini/gg/ui"
)

func CurrentBranch(args []string) {
	cb := internal.GetCurrentBranch()
	ui.PrintPlain("Current branch: %s", cb)
}
