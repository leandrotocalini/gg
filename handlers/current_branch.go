package handlers

import (
	"github.com/leandrotocalini/gg/internal"
	"github.com/leandrotocalini/gg/ui"
)

func CurrentBranch(args []string) {
	cb := internal.GetCurrentBranch()
	ui.PrintPlainf("Current branch: %s", cb)
}
