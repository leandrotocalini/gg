package handlers

import (
	"fmt"

	"github.com/leandrotocalini/gg/internal"
)

func Pull() {
	err := internal.RunWithConfirm("git", "pull")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
