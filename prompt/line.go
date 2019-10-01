package prompt

import (
	"os"

	colors "github.com/logrusorgru/aurora"
)

func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}

// GetPromptLine ...
func GetPromptLine() string {
	currentDir := colors.Magenta(getCurrentDir()).String()
	var gitBranch string
	var gitMark string

	if existGit() {
		gitBranch = colors.Brown(getCurrentBranch()).String()
		gitMark = getGitStatus()
	}
	return currentDir + " " + gitBranch + gitMark + " "
}
