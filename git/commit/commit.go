package commit

import (
	"fmt"
	"os/exec"
)

func Commit(commitMessage string) error {
	cmd := exec.Command("git", "commit", "-m", commitMessage)
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}
	return nil
}
