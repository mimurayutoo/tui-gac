package add

import (
	"fmt"
	"os/exec"
)

// 選択されたファイルを追加する
func AddSelectedFile(deletedFiles []string, changedFiles []string, addFile []bool) error {
	if len(addFile) != len(changedFiles) {
		return fmt.Errorf("addFile slice and changedFiles slice have different lengths")
	}

	// changedFilesの処理
	for i, add := range addFile {
		if add {
			cmd := exec.Command("git", "add", changedFiles[i])
			_, err := cmd.Output()
			if err != nil {
				return fmt.Errorf("failed to add file %s: %w", changedFiles[i], err)
			}
		}
	}

	// deletedFilesの処理
	for _, deletedFile := range deletedFiles {
		cmd := exec.Command("git", "add", deletedFile)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("failed to add deleted file %s: %w", deletedFile, err)
		}
	}

	return nil
}

// 全てのファイルを追加する
func AddAll(changedFiles []string, deletedFiles []string) error {
	cmd := exec.Command("git", "add", "-A")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add all files: %w", err)
	}
	return nil
}
