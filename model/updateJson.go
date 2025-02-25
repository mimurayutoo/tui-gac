package model

import (
	"tui-gac/git/add"
	"tui-gac/types"
)


func (m Model)UpdateJson(projectConfig []types.ProjectInfo, currentDir string, currentBranch string, issueNum string) ([]types.ProjectInfo, error) {
	// スライスへのissue番号の追加
	for i := range m.ProjectConfig {
		if m.ProjectConfig[i].ProjectPath == m.CurrentDir {
			// ブランチが存在することが前提なので、該当ブランチのissue番号を更新
			for j := range m.ProjectConfig[i].Branches {
				if m.ProjectConfig[i].Branches[j].BranchName == m.CurrentBranch {
					m.ProjectConfig[i].Branches[j].IssueNumber = m.IssueNum
					// 更新後すぐにjsonファイルに保存
					if err := add.SaveProjectConfig(m.ProjectConfig); err != nil {
					}
					break
				}
			}
			break
		}
	}
	return m.ProjectConfig, nil
}
