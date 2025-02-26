package add

// import (
// 	"os/exec"
// 	"strings"
// )

// func GetStatus() ([]string, []string, error) {
// 	cmd := exec.Command("git", "status", "--porcelain")
// 	output, err := cmd.Output()

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	lines := strings.Split(string(output), "\n")

// 	changedFiles := []string{}
// 	deletedFiles := []string{}

// 	for _, line := range lines {
// 		if line == "" {
// 			continue
// 		}

// 		status := line[:2]
// 		file := line[3:]

// 		if status == "D" {
// 			deletedFiles = append(deletedFiles, file)
// 		} else {
// 			changedFiles = append(changedFiles, file)
// 		}
// 	}

// 	return changedFiles, deletedFiles, nil
// }
