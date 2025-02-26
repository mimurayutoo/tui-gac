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


// ResetStaging はステージングされた変更を取り消します
func ResetStaging() error {
	cmd := exec.Command("git", "reset")
	return cmd.Run()
}

// ResetLastCommit は最後のコミットを取り消します
func ResetLastCommit() error {
	cmd := exec.Command("git", "reset", "--soft", "HEAD^")
	return cmd.Run()
}

// ResetHard は全ての変更を破棄します（注意: 取り消せない操作です）
func ResetHard() error {
	cmd := exec.Command("git", "reset", "--hard", "HEAD")
	return cmd.Run()
}
