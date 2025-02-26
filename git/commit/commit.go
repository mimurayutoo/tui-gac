package commit

import (
	"fmt"
	"os/exec"
)

// ユーザーが入力したメッセージを使用してコミットする。
func Commit(commitMessage string) error {
	cmd := exec.Command("git", "commit", "-m", commitMessage)
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}
	return nil
}
