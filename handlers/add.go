package handlers

import (
	"fmt"

	"github.com/leandrotocalini/gg/internal"
)

func Add(args []string) {
	if internal.IsProtectedBranch() {
		fmt.Println("❌ Writing to 'main' or 'master' is not allowed.")
		return
	}
	internal.RunWithConfirm("git", append([]string{"add"}, args...)...)
}
