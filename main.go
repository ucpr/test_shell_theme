package main

import (
	"fmt"

	"github.com/ucpr/hina/prompt"
)

func main() {
	k8sContext := prompt.GetK8sContext()
	promptLine := prompt.GetPromptLine()

	// output prompt line
	fmt.Println()
	fmt.Println(promptLine)
	fmt.Println(k8sContext + "$ ")
}
