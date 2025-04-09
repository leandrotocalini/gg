package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/leandrotocalini/gg/handlers"
)

type Command struct {
	Name        string
	Aliases     []string
	Description string
	Handler     func(args []string)
}

var commands = []Command{
	{Name: "c", Description: "Commit with message", Handler: handlers.Commit},
	{Name: "p", Aliases: []string{"push"}, Description: "Push current branch", Handler: func(_ []string) { handlers.Push() }},
	{Name: "pl", Aliases: []string{"pull"}, Description: "Pull from remote", Handler: func(_ []string) { handlers.Pull() }},
	{Name: "s", Aliases: []string{"status"}, Description: "Show status", Handler: func(_ []string) { handlers.Status() }},
	{Name: "a", Aliases: []string{"add"}, Description: "Add files", Handler: handlers.Add},
	{Name: "co", Aliases: []string{"checkout"}, Description: "Checkout branch", Handler: handlers.Checkout},
	{Name: "cb", Aliases: []string{"current-branch"}, Description: "Current branch", Handler: handlers.CurrentBranch},
	{Name: "nb", Aliases: []string{"new-branch"}, Description: "Create new branch", Handler: handlers.NewBranch},
	{Name: "l", Aliases: []string{"log"}, Description: "Show recent commits", Handler: func(_ []string) { handlers.Log() }},
	{Name: "rb", Aliases: []string{"recent"}, Description: "Show recent branches", Handler: func(_ []string) { handlers.RecentBranches() }},
}

var commandMap = make(map[string]Command)

func main() {
	if len(os.Args) < 2 {
		printAvailableCommands()
		return
	}
	cmd := os.Args[1]
	args := os.Args[2:]
	handleCommand(cmd, args)
}

func init() {
	for _, cmd := range commands {
		commandMap[cmd.Name] = cmd
		for _, alias := range cmd.Aliases {
			commandMap[alias] = cmd
		}
	}
}

func printAvailableCommands() {
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		aliasStr := ""
		if len(cmd.Aliases) > 0 {
			aliasStr = fmt.Sprintf(" (aliases: %s)", strings.Join(cmd.Aliases, ", "))
		}
		fmt.Printf("  %-4s %s%s\n", cmd.Name, cmd.Description, aliasStr)
	}
}

func handleCommand(cmd string, args []string) {
	if command, ok := commandMap[cmd]; ok {
		command.Handler(args)
	} else {
		fmt.Printf("Unknown command: %s\n", cmd)
		printAvailableCommands()
	}
}
