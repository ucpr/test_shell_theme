package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	colors "github.com/logrusorgru/aurora"
	"github.com/ucpr/hina/prompt"
)

type hinaK8sEnv struct {
	K8sOn string `envconfig:"HINA_K8S" default:"off"`
}

func wrapParenthesis(s string) string {
	return "(" + s + ") "
}

func getK8sContext() string {
	var env hinaK8sEnv
	envconfig.Process("hina_k8s", &env)
	if env.K8sOn != "on" {
		return ""
	}

	var k8sContext string
	k8sContext = colors.Cyan(prompt.GetCurrentContext()).String()
	return wrapParenthesis(k8sContext)
}

func main() {
	k8sContext := getK8sContext()
	promptLine := prompt.GetPromptLine()

	// output prompt line
	fmt.Println()
	fmt.Println(promptLine)
	fmt.Println(k8sContext + "$ ")
}
