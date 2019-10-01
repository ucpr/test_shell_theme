package prompt

import (
	"os/exec"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

type hinaGitEnv struct {
	Modified  string `envconfig:"HINA_GIT_MODIFIED" default:"*"`
	Added     string `envconfig:"HINA_GIT_ADDED" default:"+"`
	Deleted   string `envconfig:"HINA_GIT_DELETED" default:"-"`
	Renamed   string `envconfig:"HINA_GIT_RENAMED" default:"&"`
	Copied    string `envconfig:"HINA_GIT_COPIED" default:"~"`
	Unmerged  string `envconfig:"HINA_GIT_UNMERGED" default:"="`
	Untracked string `envconfig:"HINA_GIT_UNTRACKED" default:"!"`
}

func existGit() bool {
	cmd := exec.Command("git", "--help")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// TODO: I will back.(refactor)
func transStatusToMark(out string) string {
	var env hinaGitEnv
	var flags [7]bool
	var result string

	envconfig.Process("hina_git", &env)

	for _, row := range strings.Split(out, "\n") {
		prefix := strings.Split(strings.TrimSpace(row), " ")[0]

		switch {
		case prefix == "M" && !flags[0]:
			result += env.Modified
			flags[0] = true
		case prefix == "A" && !flags[1]:
			result += env.Added
			flags[1] = true
		case prefix == "D" && !flags[2]:
			result += env.Deleted
			flags[2] = true
		case prefix == "R" && !flags[3]:
			result += env.Renamed
			flags[3] = true
		case prefix == "C" && !flags[4]:
			result += env.Copied
			flags[4] = true
		case prefix == "U" && !flags[5]:
			result += env.Unmerged
			flags[5] = true
		case prefix == "??" && !flags[6]:
			result += env.Untracked
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
