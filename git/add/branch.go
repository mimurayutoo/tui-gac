package add

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"tui-gac/types"
)


// CheckBranch 現在のブランチを取得
func CheckBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	out, err := cmd.Output()
	if err != nil {
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
	for _, project := range projectConfig {
		if project.ProjectPath == currentDir {
			return true
		}
	}
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
	// ホームディレクトリを取得
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	// 設定ディレクトリのパスを作成
	configDir := filepath.Join(homeDir, ".config", "gac")

	// ディレクトリが存在しない場合は作成
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// JSONファイルのフルパスを作成
	configPath := filepath.Join(configDir, "branchIssueNum.json")

	// ファイルを作成
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	// JSONエンコーダーを設定して書き込み
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}

	return nil
}
