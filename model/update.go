package model

import (
	"strings"
	"tui-gac/git/add"
	"tui-gac/git/commit"
	"tui-gac/git/push"
	"tui-gac/git/reset"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.CurrentState {
		case GetBranch:
			if m.ChangedFiles == nil {
				return m, tea.Quit
			}
			if m.CurrentBranch == "" {
				return m, tea.Quit
			}
			if m.IssueNum == "" {
				m.CurrentState = InputIssueNum
			}
			if m.IssueNum != "" && m.CurrentState == GetBranch {
				m.CurrentState = FixIssueNumber
			}
			switch msg.String() {
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			case "enter":
				m.CurrentState = CheckBranchAndIssueNum
			case "c":
				m.CurrentState = FixIssueNumber
			}

		case InputIssueNum:
			if !m.InputIssueNum.Focused() {
				m.InputIssueNum.Focus()
			}
			switch msg.String() {
			case "enter":
				input := m.InputIssueNum.Value()
				if input != "" {
					// #で始まる場合はそのまま、そうでない場合は#を追加
					if !strings.HasPrefix(input, "#") {
						input = "#" + input
					}
					m.IssueNum = input
					m.InputIssueNum.Reset()
					m.InputIssueNum.Blur()
					m.UpdateJson(m.ProjectConfig, m.CurrentDir, m.CurrentBranch, m.IssueNum)
					m.CurrentState = CheckBranchAndIssueNum
				}
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			}
			var cmd tea.Cmd
			m.InputIssueNum, cmd = m.InputIssueNum.Update(msg)
			return m, cmd

		case CheckBranchAndIssueNum:
			switch msg.String() {
			case "enter":
				m.CurrentState = AddAllOrSelect
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			case "c":
				m.CurrentState = FixIssueNumber
			}

		case FixIssueNumber:
			switch msg.String() {
			case "enter":
				input := m.InputIssueNum.Value()
				if input != "" {
					// "#"がない場合は追加
					if !strings.HasPrefix(input, "#") {
						input = "#" + input
					}
					m.IssueNum = input
					m.InputIssueNum.Reset()
					m.InputIssueNum.Blur()
					m.UpdateJson(m.ProjectConfig, m.CurrentDir, m.CurrentBranch, m.IssueNum)
				}
				m.CurrentState = AddAllOrSelect
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			}
			if !m.InputIssueNum.Focused() {
				m.InputIssueNum.Focus()
			}
			var cmd tea.Cmd
			m.InputIssueNum, cmd = m.InputIssueNum.Update(msg)
			return m, cmd

		case AddAllOrSelect:
			switch msg.String() {
			case "y":
				if err := add.AddAll(m.ChangedFiles, m.DeletedFiles); err != nil {
					m.ErrorMsg = "ファイルのステージングに失敗しました: " + err.Error()
					m.CurrentState = Error
					return m, nil
				}
				m.CurrentState = SelectFixOverView
			case "n":
				// 全てのファイルを追加するように初期化
				for i := range m.AddFile {
					m.AddFile[i] = true
				}
				m.Cursor = 0
				m.CurrentState = AddSelectedFiles
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			}

		case AddAll:
			switch msg.String() {
			case "enter":
				m.CurrentState = InputCommitMessage
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			}

		case AddSelectedFiles:
			switch msg.String() {
			case "enter":
				if err := add.AddSelectedFile(m.DeletedFiles, m.ChangedFiles, m.AddFile); err != nil {
					m.ErrorMsg = "選択したファイルのステージングに失敗しました: " + err.Error()
					m.CurrentState = Error
					return m, nil
				}
				m.CurrentState = SelectFixOverView
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			case "up":
				if m.Cursor > 0 {
					m.Cursor--
				}
			case "down":
				if m.Cursor < len(m.ChangedFiles)-1 {
					m.Cursor++
				}
			case "y":
				m.AddFile[m.Cursor] = true
			case "n":
				m.AddFile[m.Cursor] = false
			}

		case Add:
			switch msg.String() {
			case "enter":
				if err := add.AddSelectedFile(m.DeletedFiles, m.ChangedFiles, m.AddFile); err != nil {
					m.IsDone = true
					return m, tea.Quit
				}
				m.CurrentState = SelectFixOverView
				m.Cursor = 0
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			}
		case SelectFixOverView:
			switch msg.String() {
			case "up":
				if m.Cursor > 0 {
					m.Cursor--
				}
			case "down":
				if m.Cursor < len(m.FixOverView)-1 {
					m.Cursor++
				}
			case "enter":
				// 修正の概要を保存（アイコンを含める）
				m.CommitMessage = m.IssueNum + " " + m.FixOverView[m.Cursor].Icon + " " + m.FixOverView[m.Cursor].Label
				m.CurrentState = InputCommitMessage
			case "ctrl+c", "q":
				if err := reset.ResetAdd(); err != nil {
					m.ErrorMsg = "ステージングのリセットに失敗しました: " + err.Error()
					return m, nil
				}
				m.IsDone = true
				return m, tea.Quit
			}
		case InputCommitMessage:
			switch msg.String() {
			case "enter":
				input := m.InputCommitMessage.Value()
				if input != "" {
					// ユーザーの入力を追加
					m.CommitMessage = m.CommitMessage + " " + input
					m.InputCommitMessage.Reset()
					m.InputCommitMessage.Blur()
					m.CurrentState = Commit
				}
			case "ctrl+c", "q":
				if err := reset.ResetAdd(); err != nil {
					m.ErrorMsg = "ステージングのリセットに失敗しました: " + err.Error()
					return m, nil
				}
				m.IsDone = true
				return m, tea.Quit
			}
			if !m.InputCommitMessage.Focused() {
				m.InputCommitMessage.Focus()
			}
			var cmd tea.Cmd
			m.InputCommitMessage, cmd = m.InputCommitMessage.Update(msg)
			return m, cmd

		case Commit:
			switch msg.String() {
			case "enter":
				if err := commit.Commit(m.CommitMessage); err != nil {
					m.ErrorMsg = "コミットに失敗しました: " + err.Error()
					m.PreviousState = m.CurrentState
					m.CurrentState = Error
					return m, nil
				}
				m.CurrentState = Push
			case "ctrl+c", "q":
				if err := reset.ResetAdd(); err != nil {
					m.ErrorMsg = "ステージングのリセットに失敗しました: " + err.Error()
					return m, nil
				}
				m.IsDone = true
				return m, tea.Quit
			}
		case Push:
			switch msg.String() {
			case "enter":
				if err := push.Push(m.CurrentBranch); err != nil {
					m.ErrorMsg = "プッシュに失敗しました: " + err.Error()
					m.PreviousState = m.CurrentState
					m.CurrentState = Error
					return m, nil
				}
				m.IsDone = true
				return m, tea.Quit
			case "ctrl+c", "q":
				if err := reset.ResetCommit(); err != nil {
					m.ErrorMsg = "コミットのリセットに失敗しました: " + err.Error()
					return m, nil
				}
				m.IsDone = true
				return m, tea.Quit
			}
		case Error:
			switch msg.String() {
			case "r": // リトライ
				switch m.PreviousState {
				case AddAllOrSelect, AddSelectedFiles:
					if err := reset.ResetStaging(); err != nil {
						m.ErrorMsg = "ステージングのリセットに失敗しました: " + err.Error()
						return m, nil
					}
					m.CurrentState = m.PreviousState
				case Commit:
					if err := reset.ResetLastCommit(); err != nil {
						m.ErrorMsg = "コミットのリセットに失敗しました: " + err.Error()
						return m, nil
					}
					m.CurrentState = m.PreviousState
				case Push:
					m.CurrentState = m.PreviousState
				}
			case "q", "ctrl+c":
				m.IsDone = true
				return m, tea.Quit
			}
		}
	}
	return m, nil
}
