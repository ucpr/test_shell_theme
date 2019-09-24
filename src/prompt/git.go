package prompt

import (
	"os/exec"
	"strings"
)

// GetCurrentBranch ...
func GetCurrentBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return ""
	}

	result := strings.TrimSpace(string(out))
	return result
}
