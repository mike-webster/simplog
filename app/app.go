package app

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mike-webster/simplog/env"
)

func Search(cfg *env.Config, text string) ([]string, error) {
	dir := cfg.Logs.Directory
	path := cfg.Logs.Filename
	if len(dir) > 0 {
		if strings.HasSuffix(dir, "/") {
			path = fmt.Sprint(dir, path)
		} else {
			path = fmt.Sprint(dir, "/", path)
		}
	}
	cmd := exec.Command("grep", "-i", text, path)
	std, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(std), "\n"), nil
}
