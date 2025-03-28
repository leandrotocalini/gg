package handlers

import "github.com/leandrotocalini/gg/internal"

func Add(args []string) {
	internal.RunWithConfirm("git", append([]string{"add"}, args...)...)
}
