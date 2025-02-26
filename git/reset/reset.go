package reset

import (
	"fmt"
	"os/exec"
)

// ResetStaging ステージングされた変更を取り消す
func ResetStaging() error {
	cmd := exec.Command("git", "reset", "HEAD", ".")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to reset staging: %w\nOutput: %s", err, string(output))
	}
	return nil
}

// ResetLastCommit 最後のコミットを取り消す
func ResetLastCommit() error {
	cmd := exec.Command("git", "reset", "--soft", "HEAD^")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to reset last commit: %w\nOutput: %s", err, string(output))
	}
	return nil
}

// ResetAdd ステージングエリアの変更を取り消す
func ResetAdd() error {
	cmd := exec.Command("git", "reset", "HEAD", ".")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to reset add: %w\nOutput: %s", err, string(output))
	}
	return nil
}

// ResetCommit 直前のコミットを取り消す
func ResetCommit() error {
	cmd := exec.Command("git", "reset", "--soft", "HEAD~1")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to reset commit: %w\nOutput: %s", err, string(output))
	}
	return nil
}
