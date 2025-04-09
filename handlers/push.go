package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/leandrotocalini/gg/internal"
)

func Push() {
	if internal.IsProtectedBranch() {
		fmt.Println("❌ Writing to 'main' or 'master' is not allowed.")
		return
	}
	err := internal.RunWithConfirm("git", "push")
	if err != nil {
		// Detect specific upstream error
		if strings.Contains(err.Error(), "exit status 128") {
			// Detect current branch
			currentBranch := internal.GetCurrentBranch()
			if currentBranch != "" {
				fmt.Println("⚠️ No upstream branch set.")
				fmt.Printf("Suggesting: git push --set-upstream origin %s\n", currentBranch)
				fmt.Print("Run this now? [y/N]: ")

				reader := bufio.NewReader(os.Stdin)
				resp, _ := reader.ReadString('\n')
				resp = strings.TrimSpace(strings.ToLower(resp))

				if resp == "y" || resp == "yes" {
					internal.Run("git", "push", "--set-upstream", "origin", currentBranch)
				} else {
					fmt.Println("You can run it manually:")
					fmt.Printf("git push --set-upstream origin %s\n", currentBranch)
				}
			}
		} else {
			fmt.Println("Error:", err)
		}
	}
}
