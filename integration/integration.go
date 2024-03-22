package integration

import (
	"os"
	"os/exec"
	"path/filepath"
)

func RunIntegration(args ...string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	mainPath := filepath.Join(wd, "..", "main.go")
	args = append([]string{"run", mainPath}, args...)
	cmd := exec.Command("go", args...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}
