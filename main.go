package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"tui-gac/model"
	"tui-gac/types"

	// "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// type state int

// const (
// 	getBranch state = iota
// 	selectFiles
// 	selectFixOverView
// 	inputCommitMessage
// 	Push
// )

// type BranchIssue struct {
// 	BranchName  string `json:"branch"`
// 	IssueNumber int    `json:"issueNum"`
// }

// type ProjectInfo struct {
// 	ProjectName string        `json:"dir"`
// 	Branches    []BranchIssue `json:"branchInfo"`
// }

// type model struct {
// 	cursor        int
// 	changedFiles  []string
// 	selectedFiles []string
// 	input         textinput.Model
// 	currentState  state
// 	isDone        bool
// 	projectConfig []ProjectInfo
// 	currentDir string
// }

// todo
// func (m model.Model) Init() tea.Cmd {
// 	return textinput.Blink
// }

// func (m model.Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch m.GetCurrentState() {
// 		case model.GetBranch:
// 			branch, issueNum, err := add.CheckInfo(projectConfig)
// 			log.Println(branch, num, err)
// 			// branch, err := add.GetBranch()
// 			if err != nil {
// 				log.Println(err)
// 				return m, nil
// 			}
// 			log.Println(branch)
// 			switch msg.String() {
// 			case "q", "ctrl+c":
// 				return m, tea.Quit
// 			default:
// 				return m, nil
// 			}
// 		case model.SelectFiles:
// 			return m, nil
// 		default:
// 			return m, nil
// 		}
// 	default:
// 		return m, nil
// 	}
// }

// func (m model) View() string {
// 	changedFiles := ""
// 	// 現在のディレクトリを取得
// 	currentDir, err := os.Getwd()
// 	if err != nil {
// 		changedFiles += "current dir: error\n"
// 	} else {
// 		changedFiles += "current dir: " + path.Base(currentDir) + "\n"
// 	}
// 	changedFiles += "Select Branch\n"
// 	for _, branch := range m.getBranch() {
// 		branchName := strings.TrimSpace(branch)
// 		if branchName == "" {
// 			continue
// 		}
// 		changedFiles += branchName + "\n"
// 	}
// 	return changedFiles
// }

// func (m model) getBranch() []string {
// 	cmd := exec.Command("git", "branch", "--sort=-committerdate")
// 	out, err := cmd.Output()
// 	if err != nil {
// 		return []string{}
// 	}
// 	return strings.Split(string(out), "\n")
// }
// todo

// git status --porcelain を実行して、変更されたファイルを取得
// todo あとでpackageを分けてそれぞれのコマンドで実装する。
// func (m model) status() []string {
// 	cmd := exec.Command("git", "status", "--porcelain")
// 	out, err := cmd.Output()
// 	if err != nil {
// 		return []string{}
// 	}

// 	return strings.Split(string(out), "\n")
// }

var projectConfig []types.ProjectInfo

func main() {
	// jsonファイルを開く
	configFile, err := os.Open("branchIssueNum.json")
	if err != nil {
		// ファイルが存在しない場合は空の配列で初期化
		projectConfig = []types.ProjectInfo{}
	} else {
		defer configFile.Close()
		// jsonファイルを読み込む
		byteValue, err := io.ReadAll(configFile)
		if err != nil {
			log.Println("Failed to read JSON file:", err)
			return
		}

		// jsonファイルをパースする
		if err := json.Unmarshal(byteValue, &projectConfig); err != nil {
			log.Println("Failed to parse JSON:", err)
			return
		}
		log.Printf("main Initial ProjectConfig: %+v\n", projectConfig)
	}

	model := model.InitModel(projectConfig)
	dir, err := os.Getwd()
	if err != nil {
		log.Println("Failed to get current directory:", err)
		return
	}
	log.Println(dir)
	// model.CurrentDir = dir

	// アプリケーションを初期化
	p := tea.NewProgram(model)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
