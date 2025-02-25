package model

import (
	"strings"
	"tui-gac/git/add"
	"tui-gac/git/commit"

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
				if !m.InputIssueNum.Focused() {
					m.InputIssueNum.Focus()
				}
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
				}
			}
			if m.IssueNum != "" && m.CurrentState == GetBranch {
				m.CurrentState = CheckBranchInfo
			}
			switch msg.String() {
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			}
			var cmd tea.Cmd
			m.InputIssueNum, cmd = m.InputIssueNum.Update(msg)
			return m, cmd

		case CheckBranchInfo:
			switch msg.String() {
			case "enter":
				m.CurrentState = AddAllOrSelect
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			case "c":
				m.CurrentState = ChangeIssueNumber
			}

		case ChangeIssueNumber:
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
				m.CurrentState = AddAll
			case "n":
				// 全てのファイルを追加
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
				if err := add.AddAll(m.ChangedFiles, m.DeletedFiles); err != nil {
					m.IsDone = true
					return m, tea.Quit
				}
				m.CurrentState = InputCommitMessage
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			}

		case AddSelectedFiles:
			switch msg.String() {
			case "enter":
				if err := add.AddSelectedFile(m.DeletedFiles, m.ChangedFiles, m.AddFile); err != nil {
					m.IsDone = true
					return m, tea.Quit
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
				m.CommitMessage += m.IssueNum + " " + m.FixOverView[m.Cursor]
				m.CurrentState = InputCommitMessage
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			}
		case InputCommitMessage:
			switch msg.String() {
			case "enter":
				input := m.InputCommitMessage.Value()
				if input != "" {
					m.CommitMessage += " " + input
					m.InputCommitMessage.Reset()
					m.InputCommitMessage.Blur()
					m.CurrentState = Commit
				}
				m.CurrentState = Commit
			case "ctrl+c", "q":
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
					m.IsDone = true
					return m, tea.Quit
				}
				m.CurrentState = Push
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			}
		case Push:
			switch msg.String() {
			case "enter":
				m.CurrentState = Push
			case "ctrl+c", "q":
				m.IsDone = true
				return m, tea.Quit
			}
		}
	}
	return m, nil
}
