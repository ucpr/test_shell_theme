package prompt

import (
	"os"
	"os/exec"
	"strings"
)

// using const
var (
	MODIFIED  = os.Getenv("HINA_GIT_MODIFIED")
	ADDED     = os.Getenv("HINA_GIT_ADDED")
	DELETED   = os.Getenv("HINA_GIT_DELETED")
	COPIED    = os.Getenv("HINA_GIT_COPIED")
	RENAMED   = os.Getenv("HINA_GIT_RENAMED")
	UNMERGED  = os.Getenv("HINA_GIT_UNMERGED")
	UNTRACKED = os.Getenv("HINA_GIT_UNTRACKED")
)

// TODO: I will back.(refactor)
func transStatusToMark(out string) string {
	var result string

	for _, row := range strings.Split(out, "\n") {
		prefix := strings.Split(strings.TrimSpace(row), " ")[0]
		switch prefix {
		case "M":
			result += MODIFIED
		case "A":
			result += ADDED
		case "D":
			result += DELETED
		case "R":
			result += COPIED
		case "C":
			result += RENAMED
		case "U":
			result += UNMERGED
		case "??":
			result += UNTRACKED
		}
	}
	return result
}

// GetGitStatus ...
func GetGitStatus() string {
	out, err := exec.Command("git", "status", "--short").Output()
	if err != nil {
		return ""
	}

	mark := transStatusToMark(string(out))
	return mark
}

// GetCurrentBranch ...
func GetCurrentBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return ""
	}

	result := strings.TrimSpace(string(out))
	return result
}
