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
	gitBranch := colors.Brown(GetCurrentBranch()).String()
	gitMark := GetGitStatus()
	return currentDir + " " + gitBranch + gitMark + " "
}
