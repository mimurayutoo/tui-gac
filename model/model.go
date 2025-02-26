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
	GetBranch              state = iota // ブランチ名を取得する
	InputIssueNum                       // イシュー番号を入力する
	CheckBranchAndIssueNum              // ブランチ名とイシュー番号を確認する
	FixIssueNumber                      // イシュー番号を修正する
	AddAllOrSelect                      // 全てを追加するか、選択したファイルを追加するかを選択する
	Add                                 // 選択したファイルを追加する
	AddAll                              // 全てを追加する
	AddSelectedFiles                    // 選択したファイルを追加する
	SelectFixOverView                   // コミットタイプを選択する
	InputCommitMessage                  // コミットメッセージを入力する
	Commit                              // コミット
	Push                                // プッシュ
	Error                               // エラーが発生した場合
)

type CommitType struct {
	Label string
	Icon  string
	Desc  string
}

type Model struct {
	Cursor             int                 // カーソルの位置
	ChangedFiles       []string            // 変更されたファイル
	DeletedFiles       []string            // 削除されたファイル
	SelectedFiles      []string            // ステージングするように選択されたファイル
	CurrentState       state               // 現在の状態
	IsDone             bool                // 終了したかどうか
	ProjectConfig      []types.ProjectInfo // jsonファイルに保存されている情報
	CurrentBranch      string              // 現在のブランチ
	IssueNum           string              // 存在しない場合には""存在している場合には"#~~~"
	InputIssueNum      textinput.Model     // イシュー番号を入力する
	InputCommitMessage textinput.Model     // コミットメッセージを入力する
	IsExistDir         bool                // ディレクトリが存在しているかどうか
	IsExistBranch      bool                // ブランチが存在しているかどうか
	IsExistIssueNum    bool                // イシュー番号が存在しているかどうか
	CurrentDir         string              // 現在のディレクトリ
	FixOverView        []CommitType        // コミットタイプ
	AddFile            []bool              // 追加するファイル
	UserIntention      bool                // ユーザーの意図
	StagedFiles        []string            // ステージングされたファイル
	CommitMessage      string              // コミットメッセージ
	ErrorMsg           string              // エラーメッセージ
	PreviousState      state               // エラー発生前の状態を保存
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
		FixOverView: []CommitType{
			{
				Label: "FIX",
				Icon:  "🛠️",
				Desc:  "Bug fixes and patches",
			},
			{
				Label: "ADD",
				Icon:  "✨",
				Desc:  "New features and additions",
			},
			{
				Label: "UPDATE",
				Icon:  "🚀",
				Desc:  "Updates and improvements",
			},
			{
				Label: "REFACTOR",
				Icon:  "🔄",
				Desc:  "Code refactoring",
			},
			{
				Label: "STYLE",
				Icon:  "🎨",
				Desc:  "Style and formatting",
			},
			{
				Label: "REMOVE",
				Icon:  "🗑️",
				Desc:  "Removing code or files",
			},
			{
				Label: "REVIEW_FIX",
				Icon:  "📝",
				Desc:  "Fixes based on code review",
			},
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
	// プロジェクトが存在していない場合にはディレクトリ追加する。
	if !m.IsExistDir {
		add.WriteDir(m.CurrentDir, &m.ProjectConfig)
	}

	// jsonファイルにディレクトリに対応するブランチが存在しているかのフラグ
	m.IsExistBranch = add.SearchBranch(m.ProjectConfig, m.CurrentDir, m.CurrentBranch)
	if !m.IsExistBranch {
		// ブランチが存在していない場合にはissue番号も存在していない。
		m.IssueNum = ""
		m.IsExistIssueNum = false
		m.CurrentState = InputIssueNum
		// jsonファイルに現在のブランチを追加する。
		updatedConfig, err := add.WriteBranch(m.CurrentDir, m.CurrentBranch, m.ProjectConfig)
		if err != nil {
			return m
		}
		m.ProjectConfig = updatedConfig
	} else {
		// ブランチが存在している場合にはissue番号を取得する。
		m.IssueNum = add.GetIssueNumber(m.ProjectConfig, m.CurrentDir, m.CurrentBranch)
	}

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
