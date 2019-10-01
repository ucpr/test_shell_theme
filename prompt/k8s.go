package prompt

import (
	"os/exec"

	"github.com/kelseyhightower/envconfig"
	colors "github.com/logrusorgru/aurora"
	"gopkg.in/yaml.v2"
)

type kubeConfig struct {
	CurrentContext string `yaml:"current-context"`
}

type hinaK8sEnv struct {
	K8sOn string `envconfig:"HINA_K8S" default:"off"`
}

func wrapParenthesis(s string) string {
	return "(" + s + ") "
}

func getCurrentContext() string {
	var conf kubeConfig

	out, err := exec.Command("kubectl", "config", "view", "--raw").Output()
	if err != nil {
		return ""
	}

	err = yaml.Unmarshal(out, &conf)
	if err != nil {
		return ""
	}

	return conf.CurrentContext
}

// GetK8sContext ...
func GetK8sContext() string {
	var env hinaK8sEnv
	envconfig.Process("hina_k8s", &env)
	if env.K8sOn != "on" {
		return ""
	}

	var k8sContext string
	k8sContext = colors.Cyan(getCurrentContext()).String()
	return wrapParenthesis(k8sContext)
}
