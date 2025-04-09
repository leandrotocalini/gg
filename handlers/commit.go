package handlers

import (
	"fmt"
	"strings"

	"github.com/leandrotocalini/gg/internal"
)

func Commit(args []string) {
	msg := strings.Join(args, " ")
	if internal.IsProtectedBranch() {
		fmt.Println("❌ Writing to 'main' or 'master' is not allowed.")
		return
	}
	internal.RunWithConfirm("git", "commit", "-am", msg)
}
