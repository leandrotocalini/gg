package handlers

import (
	"regexp"
	"strings"

	"github.com/leandrotocalini/gg/internal"
	"github.com/leandrotocalini/gg/ui"
)

func cleanBranchName(input string) string {
	// Trim spaces and lowercase
	input = strings.ToLower(strings.TrimSpace(input))

	// If there's a slash, split and clean each part separately
	if strings.Contains(input, "/") {
		parts := strings.SplitN(input, "/", 2)
		return cleanSegment(parts[0]) + "/" + cleanSegment(parts[1])
	}

	return cleanSegment(input)
}

func cleanSegment(s string) string {
	s = strings.TrimSpace(s)
	// Replace all underscores and spaces with hyphens
	s = strings.ReplaceAll(s, "_", "-")
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, "-")
	s = regexp.MustCompile(`-+`).ReplaceAllString(s, "-") // no double dashes
	s = strings.Trim(s, "-")
	return s
}

func NewBranch(args []string) {
	if len(args) == 0 {
		ui.PrintErrorf("Usage: gg nb <branch-name>")
		return
	}

	raw := strings.Join(args, " ")
	cleaned := cleanBranchName(raw)

	ui.PrintWarningf("Branch to create:", cleaned)

	if !internal.ConfirmProceed() {
		ui.PrintErrorf("Aborted.")
		return
	}

	internal.Run("git", "checkout", "-b", cleaned)
}
