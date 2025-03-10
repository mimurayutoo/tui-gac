package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"tui-gac/model"
	"tui-gac/types"

	tea "github.com/charmbracelet/bubbletea"
)

var projectConfig []types.ProjectInfo

func main() {
	// jsonファイルを開く
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	configFile, err := os.Open(filepath.Join(homeDir, ".config", "gac", "branchIssueNum.json"))
	if err != nil {
		projectConfig = []types.ProjectInfo{}
	} else {
		defer configFile.Close()
		byteValue, err := io.ReadAll(configFile)
		if err != nil {
			return
		}

		if err := json.Unmarshal(byteValue, &projectConfig); err != nil {
			return
		}
	}

	model := model.InitModel(projectConfig)
	_, err = os.Getwd()
	if err != nil {
		return
	}

	// アプリケーションを初期化
	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
