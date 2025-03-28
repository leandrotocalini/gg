package handlers

import "github.com/leandrotocalini/gg/internal"

func Checkout(args []string) {
	internal.RunWithConfirm("git", append([]string{"checkout"}, args...)...)
}
