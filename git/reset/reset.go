package reset

import (
	"fmt"
	"os/exec"
)

func Reset() error {
	cmd := exec.Command("git", "reset", "--mixed", "HEAD~1")

	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to reset: %w", err)
	}
	return nil
}