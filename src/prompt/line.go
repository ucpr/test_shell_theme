package prompt

import (
	"os"
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
	gitBranch := GetCurrentBranch()
	gitMark := GetGitStatus()
	return getCurrentDir() + " " + gitBranch + gitMark +  "\n\n$ "
}
