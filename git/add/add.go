package add

import (
	"fmt"
	"os/exec"
)

func AddSelectedFile(deletedFiles []string, changedFiles []string, addFile []bool) error {
	if len(addFile) != len(changedFiles) {
		return fmt.Errorf("addFile slice and changedFiles slice have different lengths")
	}

	// まずchangedFilesの処理
	for i, add := range addFile {
		if add {
			cmd := exec.Command("git", "add", changedFiles[i])
			_, err := cmd.Output()
			if err != nil {
				return fmt.Errorf("failed to add file %s: %w", changedFiles[i], err)
			}
		}
	}

	// 次にdeletedFilesの処理
	for _, deletedFile := range deletedFiles {
		cmd := exec.Command("git", "add", deletedFile)
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("failed to add deleted file %s: %w", deletedFile, err)
		}
	}

	return nil
}

func AddAll(changedFiles []string, deletedFiles []string) error {
	cmd := exec.Command("git", "add", "-A")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add all files: %w", err)
	}
	return nil
}
