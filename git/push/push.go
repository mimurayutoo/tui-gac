package push

import "os/exec"

func Push(currentBranch string) error {
	cmd := exec.Command("git", "push", "origin", currentBranch)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
