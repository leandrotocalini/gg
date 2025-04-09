package handlers

import (
	"fmt"

	"github.com/leandrotocalini/gg/internal"
)

func CurrentBranch(args []string) {
	cb := internal.GetCurrentBranch()
	fmt.Println("Current branch:", cb)
}
