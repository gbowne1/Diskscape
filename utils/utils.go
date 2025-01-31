package utils

import (
	"fmt"
	"os/exec"
)

func RunCommand(command string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %w", err)
	}
	return output, nil
}
