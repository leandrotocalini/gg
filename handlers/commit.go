package handlers

import (
	"strings"

	"github.com/leandrotocalini/gg/internal"
)

func Commit(args []string) {
	msg := strings.Join(args, " ")
	internal.RunWithConfirm("git", "commit", "-am", msg)
}
