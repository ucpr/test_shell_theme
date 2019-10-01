package prompt

import (
	"os/exec"
	"gopkg.in/yaml.v2"
)

type kubeConfig struct {
	CurrentContext string `yaml:"current-context"`
}

// GetCurrentContext ...
func GetCurrentContext() string {
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
