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
	var flags [7]bool
	var result string

	for _, row := range strings.Split(out, "\n") {
		prefix := strings.Split(strings.TrimSpace(row), " ")[0]

		switch {
		case prefix == "M" && !flags[0]:
			result += MODIFIED
			flags[0] = true
		case prefix == "A" && !flags[1]:
			result += ADDED
			flags[1] = true
		case prefix == "D" && !flags[2]:
			result += DELETED
			flags[2] = true
		case prefix == "R" && !flags[3]:
			result += COPIED
			flags[3] = true
		case prefix == "C" && !flags[4]:
			result += RENAMED
			flags[4] = true
		case prefix == "U" && !flags[5]:
			result += UNMERGED
			flags[5] = true
		case prefix == "??" && !flags[6]:
			result += UNTRACKED
			flags[6] = true
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
