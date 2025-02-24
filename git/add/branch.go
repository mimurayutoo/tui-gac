package add

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"tui-gac/types"
)

// CheckBranch 現在のブランチを取得
func CheckBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return "", err
	}
	currentBranch := strings.TrimSpace(string(out))
	return currentBranch, nil
}

// ブランチが存在しているのか確認。
func SearchBranch(projectConfig []types.ProjectInfo, currentDir string, currentBranch string) bool {
	branchExist := false

	for _, project := range projectConfig {
		if project.ProjectPath == currentDir {
			for _, branch := range project.Branches {
				if branch.BranchName == currentBranch {
					branchExist = true
					return branchExist
				}
			}
		}
	}
	return branchExist
}

// ディレクトリブランチ名からissue番号を取得する。issue番号が存在しない場合にはnilを返す。
func GetIssueNumber(projectConfig []types.ProjectInfo, currentDir string, currentBranch string) string {
	for _, project := range projectConfig {
		if project.ProjectPath == currentDir {
			for _, branch := range project.Branches {
				if branch.BranchName == currentBranch {
					return branch.IssueNumber
				}
			}
		}
	}
	return ""
}

// jsonファイルにプロジェクトが存在しているのか確認
func SearchDir(projectConfig []types.ProjectInfo, currentDir string) bool {
	log.Printf("SearchDir - Current Directory: %s\n", currentDir)
	for _, project := range projectConfig {
		log.Printf("SearchDir - Comparing with Project Path: %s\n", project.ProjectPath)
		if project.ProjectPath == currentDir {
			log.Println("SearchDir - Match found!")
			return true
		}
	}
	log.Println("SearchDir - No match found")
	return false
}

func WriteDir(currentDir string, projectConfig *[]types.ProjectInfo) {
	*projectConfig = append(*projectConfig, types.ProjectInfo{
		ProjectPath: currentDir,
		Branches:    []types.BranchIssue{},
	})
}

// ブランチが存在していない場合にのみブランチを追加する。
func WriteBranch(currentDir string, currentBranch string, projectConfig []types.ProjectInfo) ([]types.ProjectInfo, error) {
	for i := range projectConfig {
		if projectConfig[i].ProjectPath == currentDir {
			// インデックスを使用して直接スライスを更新
			projectConfig[i].Branches = append(projectConfig[i].Branches, types.BranchIssue{
				BranchName:  currentBranch,
				IssueNumber: "",
			})
		}
	}
	return projectConfig, nil
}

func SaveProjectConfig(config []types.ProjectInfo) error {
	file, err := os.Create("branchIssueNum.json")
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}
	return nil
}
