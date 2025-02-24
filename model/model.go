package model

import (
	"os"
	"tui-gac/git/add"
	"tui-gac/git/status"
	"tui-gac/types"

	"github.com/charmbracelet/bubbles/textinput"
)

type state int

const (
	GetBranch state = iota
	CheckBranchInfo
	ChangeIssueNumber
	AddAllOrSelect
	Add
	AddAll
	AddSelectedFiles
	SelectFixOverView
	InputCommitMessage
	Commit
	Push
)

// それぞれのフィールドを大文字にして外部からのアクセスができるようにして
type Model struct {
	Cursor             int
	ChangedFiles       []string
	DeletedFiles       []string
	SelectedFiles      []string
	CurrentState       state
	IsDone             bool
	ProjectConfig      []types.ProjectInfo
	CurrentBranch      string
	IssueNum           string // 存在しない場合には""存在している場合には"#111"
	InputIssueNum      textinput.Model
	InputCommitMessage textinput.Model
	IsExistDir         bool
	IsExistBranch      bool
	IsExistIssueNum    bool
	CurrentDir         string
	FixOverView        []string
	AddFile            []bool
	UserIntention      bool
	StagedFiles        []string
	CommitMessage      string
}

// モデルの初期化
func InitModel(projectConfig []types.ProjectInfo) Model {
	ti := textinput.New()
	m := Model{
		Cursor:             0,
		ChangedFiles:       []string{},
		SelectedFiles:      []string{},
		CurrentState:       GetBranch,
		IsDone:             false,
		ProjectConfig:      projectConfig,
		IssueNum:           "",
		InputIssueNum:      ti,
		InputCommitMessage: ti,
		FixOverView: []string{
			"fix",
			"add",
			"update",
			"refactor",
			"style",
			"remove",
		},
		StagedFiles:   []string{},
		CommitMessage: "",
	}
	currentBranch, err := add.CheckBranch()
	if err != nil {
		return m
	}
	m.CurrentBranch = currentBranch

	m.CurrentDir, err = os.Getwd()
	if err != nil {
		return m
	}

	m.IsExistDir = add.SearchDir(m.ProjectConfig, m.CurrentDir)
	if !m.IsExistDir {
		// jsonファイルにプロジェクトを追加する関数を実装。参照ではなく、ポインタを渡す。
		add.WriteDir(m.CurrentDir, &m.ProjectConfig)
	}

	// jsonファイルにディレクトリに対応するブランチが存在しているかのフラグ
	m.IsExistBranch = add.SearchBranch(m.ProjectConfig, m.CurrentDir, m.CurrentBranch)
	if !m.IsExistBranch {
		// ブランチが存在していない場合にはissue番号も存在していない。
		m.IssueNum = ""
		m.IsExistIssueNum = false
		// jsonファイルに現在のブランチを追加する関数を実装。
		updatedConfig, err := add.WriteBranch(m.CurrentDir, m.CurrentBranch, m.ProjectConfig)
		if err != nil {
			return m
		}
		m.ProjectConfig = updatedConfig
	} else {
		m.IssueNum = add.GetIssueNumber(m.ProjectConfig, m.CurrentDir, m.CurrentBranch)
	}
	//　 ここに全てのデータを表示してデバックする。
	changedFiles, deletedFiles, err := status.GetStatus()
	if err != nil {
		return m
	}
	m.ChangedFiles = changedFiles
	m.DeletedFiles = deletedFiles
	m.AddFile = make([]bool, len(changedFiles))
	for i := range m.AddFile {
		m.AddFile[i] = false
	}
	m.UserIntention = false
	return m
}
