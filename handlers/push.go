package handlers

import (
	"bufio"
	"os"
	"strings"

	"github.com/leandrotocalini/gg/internal"
	"github.com/leandrotocalini/gg/ui"
)

func Push() {
	if internal.IsProtectedBranch() {
		ui.PrintErrorf("❌ Writing to 'main' or 'master' is not allowed.")
		return
	}
	err := internal.RunWithConfirm("git", "push")
	if err != nil {
		// Detect specific upstream error
		if strings.Contains(err.Error(), "exit status 128") {
			// Detect current branch
			currentBranch := internal.GetCurrentBranch()
			if currentBranch != "" {
				ui.PrintWarningf("⚠️ No upstream branch set.")
				ui.PrintErrorf("Suggesting: git push --set-upstream origin %s\n", currentBranch)
				ui.PrintPlainf("Run this now? [y/N]: ")

				reader := bufio.NewReader(os.Stdin)
				resp, _ := reader.ReadString('\n')
				resp = strings.TrimSpace(strings.ToLower(resp))

				if resp == "y" || resp == "yes" {
					internal.Run("git", "push", "--set-upstream", "origin", currentBranch)
				} else {
					ui.PrintPlainf("You can run it manually:")
					ui.PrintWarningf("git push --set-upstream origin %s\n", currentBranch)
				}
			}
		} else {
			ui.PrintErrorf("Error: %s", err)
		}
	}
}
