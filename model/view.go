package model

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FF00")).
			MarginBottom(1)

	branchStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF00FF")).
			PaddingLeft(2)

	issueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFFF")).
			PaddingLeft(2)

	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFF00")).
			PaddingLeft(2)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.Color("#FFFFFF"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00")).
			Bold(true).
			Background(lipgloss.Color("#333333"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Italic(true).
			MarginTop(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFFF")).
			Bold(true).
			MarginTop(1).
			MarginBottom(1)
)

func (m Model) View() string {
	message := ""
	if m.IsDone {
		return "Goodbye! 👋\n"
	}

	switch m.CurrentState {
	case GetBranch:
		message += titleStyle.Render("🌳 Branch Information") + "\n"
		message += branchStyle.Render("Branch: "+m.CurrentBranch) + "\n"
		if m.IssueNum == "" {
			message += inputStyle.Render("Issue Number: ") + m.InputIssueNum.View() + "\n"
			message += helpStyle.Render("Enter issue number and press Enter")
		} else {
			message += issueStyle.Render("Issue: "+m.IssueNum) + "\n"
			message += helpStyle.Render("Press Enter to continue")
		}

	case CheckBranchInfo:
		message += titleStyle.Render("✓ Confirmation") + "\n"
		message += branchStyle.Render("Branch: "+m.CurrentBranch) + "\n"
		message += issueStyle.Render("Issue: "+m.IssueNum) + "\n"
		message += helpStyle.Render("Press Enter to select files or enter c to change issue number")

	case ChangeIssueNumber:
		message += titleStyle.Render("🔄 Change Issue Number") + "\n"
		message += inputStyle.Render("Issue Number: ") + m.InputIssueNum.View() + "\n"
		message += helpStyle.Render("Enter issue number and press Enter")

	case AddAllOrSelect:
		message += titleStyle.Render("📁 Select Files") + "\n"
		message += helpStyle.Render("y: add all • n: add selected files")

	case AddSelectedFiles:
		message += titleStyle.Render("📁 Select Files") + "\n"
		for i, file := range m.ChangedFiles {
			cursor := "×"
			if m.AddFile[i] {
				cursor = "✓"
			}
			style := itemStyle
			if i == m.Cursor {
				style = style.Inherit(selectedStyle)
			}
			message += style.Render(cursor+" "+file) + "\n"
		}
		message += helpStyle.Render("↑/↓: move • y: select • n: deselect • Enter: continue")

	case SelectFixOverView:
		message += titleStyle.Render("📝 Commit Overview") + "\n\n"

		// Staged Files Section
		message += subtitleStyle.Render("Staged Files:") + "\n"
		if len(m.DeletedFiles) == 0 && len(m.ChangedFiles) == 0 {
			message += itemStyle.Render("No files staged") + "\n"
		} else {
			if len(m.ChangedFiles) > 0 {
				for _, file := range m.ChangedFiles {
					message += itemStyle.Render("✓ "+file) + "\n"
				}
			}
			if len(m.DeletedFiles) > 0 {
				for _, file := range m.DeletedFiles {
					message += itemStyle.Render("✗ "+file) + "\n"
				}
			}
		}
		message += "\n"

		// Fix Overview Section
		message += subtitleStyle.Render("Select Commit Type:") + "\n"
		for i, fixOverview := range m.FixOverView {
			style := itemStyle
			if i == m.Cursor {
				style = style.Inherit(selectedStyle)
			}
			message += style.Render("• "+fixOverview) + "\n"
		}
		message += "\n"
		message += helpStyle.Render("↑/↓: move • Enter: select commit type")

	case InputCommitMessage:
		message += titleStyle.Render("📝 Commit Message") + "\n\n"
		message += inputStyle.Render("Commit Message: ") + m.InputCommitMessage.View() + "\n"
		message += helpStyle.Render("Enter commit message and press Enter")
	}

	return message
}
