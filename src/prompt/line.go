package prompt

// GetPromptLine ...
func GetPromptLine() string {
	gitBranch := GetCurrentBranch()
	return gitBranch + " $ "
}
