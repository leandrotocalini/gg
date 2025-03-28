package main

import (
	"fmt"
	"os"

	"github.com/leandrotocalini/gg/handlers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gg [command]")
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	handleCommand(cmd, args)
}

func handleCommand(cmd string, args []string) {
	switch cmd {
	case "c":
		handlers.Commit(args)
	case "p":
		handlers.Push()
	case "s":
		handlers.Status()
	case "a":
		handlers.Add(args)
	case "co":
		handlers.Checkout(args)
	case "nb":
		handlers.NewBranch(args)
	case "l":
		handlers.Log()
	case "rb", "recent":
		handlers.RecentBranches()
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
