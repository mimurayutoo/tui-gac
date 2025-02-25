package model

import (
	"fmt"
	"strings"
	"tui-gac/model/styles"
)

func (m Model) View() string {
	var s strings.Builder

	// Header section with app title and branch info
	s.WriteString(styles.TitleStyle.Render("🚀 Git Commit Assistant") + "\n")
	s.WriteString(styles.BranchStyle.Render("🌿 Current Branch: "+m.CurrentBranch) + "\n\n")

	// Main content
	switch m.CurrentState {
	case GetBranch:
		if m.IssueNum == "" {
			s.WriteString(styles.SubtitleStyle.Render("📝 Enter Issue Number") + "\n")
			s.WriteString(styles.InputStyle.Render(m.InputIssueNum.View()) + "\n\n")
			s.WriteString(styles.HelpStyle.Render("• Enter: Confirm\n• Ctrl+C: Exit"))
		} else {
			s.WriteString(styles.StatusStyle.Render("✓ Issue: "+m.IssueNum) + "\n\n")
			s.WriteString(styles.HelpStyle.Render("• Enter: Next\n• c: Modify Issue Number\n• Ctrl+C: Exit"))
		}

	case CheckBranchAndIssueNum:
		s.WriteString(styles.SubtitleStyle.Render("🔍 Confirm Branch Information") + "\n\n")
		s.WriteString(styles.StatusStyle.Render("Branch: "+m.CurrentBranch) + "\n")
		s.WriteString(styles.StatusStyle.Render("Issue: "+m.IssueNum) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: Next\n• c: Modify Issue Number\n• Ctrl+C: Exit"))

	case FixIssueNumber:
		s.WriteString(styles.SubtitleStyle.Render("✏️  Modify Issue Number") + "\n\n")
		s.WriteString(styles.InputStyle.Render(m.InputIssueNum.View()) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: Confirm\n• Ctrl+C: Exit"))

	case InputIssueNum:
		s.WriteString(styles.SubtitleStyle.Render("📎 Enter Issue Number") + "\n")
		s.WriteString(styles.InputStyle.Render(m.InputIssueNum.View()) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: Confirm\n• Ctrl+C: Exit"))

	case AddAllOrSelect:
		s.WriteString(styles.SubtitleStyle.Render("📁 Select Files to Stage") + "\n\n")

		if len(m.ChangedFiles) > 0 {
			s.WriteString(styles.ItemStyle.Render("Modified Files:") + "\n")
			for _, file := range m.ChangedFiles {
				s.WriteString(styles.ItemStyle.Render("  ↳ "+file) + "\n")
			}
		}

		if len(m.DeletedFiles) > 0 {
			if len(m.ChangedFiles) > 0 {
				s.WriteString("\n")
			}
			s.WriteString(styles.WarningStyle.Render("Deleted Files:") + "\n")
			for _, file := range m.DeletedFiles {
				s.WriteString(styles.WarningStyle.Render("  ⨯ "+file) + "\n")
			}
		}
		if len(m.ChangedFiles) > 0 || len(m.DeletedFiles) > 0 {
			s.WriteString("\n" + styles.HelpStyle.Render("• y: Stage All Files\n• n: Select Individual Files\n• Ctrl+C: Exit"))
		}

	case AddSelectedFiles:
		s.WriteString(styles.SubtitleStyle.Render("🔍 Select Files") + "\n\n")
		for i, file := range m.ChangedFiles {
			prefix := "   "
			if i == m.Cursor {
				prefix = " ➜ "
			}
			style := styles.ItemStyle
			if i == m.Cursor {
				style = style.Inherit(styles.SelectedStyle)
			}
			status := "○"
			if m.AddFile[i] {
				status = "●"
			}
			s.WriteString(style.Render(prefix+status+" "+file) + "\n")
		}
		s.WriteString("\n" + styles.HelpStyle.Render("• ↑/↓: Navigate\n• y: Select\n• n: Deselect\n• Enter: Confirm"))

		
	case SelectFixOverView:
		s.WriteString(styles.SubtitleStyle.Render("📋 Select Commit Type") + "\n\n")
		for i, fix := range m.FixOverView {
			prefix := "   "
			if i == m.Cursor {
				prefix = " ➜ "
			}
			style := styles.ItemStyle
			if i == m.Cursor {
				style = style.Inherit(styles.SelectedStyle)
			}
			commitInfo := fmt.Sprintf("%s %s - %s", fix.Icon, fix.Label, fix.Desc)
			s.WriteString(style.Render(prefix+commitInfo) + "\n")
		}
		s.WriteString("\n" + styles.HelpStyle.Render("• ↑/↓: Navigate\n• Enter: Select"))

	case InputCommitMessage:
		s.WriteString(styles.SubtitleStyle.Render("✍️  Enter Commit Message") + "\n\n")
		s.WriteString(styles.StatusStyle.Render("Issue: "+m.IssueNum) + "\n")
		s.WriteString(styles.StatusStyle.Render("Type: "+m.FixOverView[m.Cursor].Icon+" "+m.FixOverView[m.Cursor].Label) + "\n\n")
		s.WriteString(styles.InputStyle.Render("Message: "+m.InputCommitMessage.View()) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: Confirm\n• Ctrl+C: Exit"))

	case Commit:
		s.WriteString(styles.SubtitleStyle.Render("👀 Confirm Commit Message") + "\n\n")
		s.WriteString(styles.StatusStyle.Render("Message: "+m.CommitMessage) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: Execute Commit\n• Ctrl+C: Cancel"))

	case Push:
		s.WriteString(styles.SubtitleStyle.Render("🚀 Ready to Push") + "\n\n")
		s.WriteString(styles.StatusStyle.Render("Commit Message: "+m.CommitMessage) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: Push\n• Ctrl+C: Exit"))

	case Error:
		s.WriteString(styles.SubtitleStyle.Render("❌ Error Occurred") + "\n\n")
		s.WriteString(styles.ErrorStyle.Render(m.ErrorMsg) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• r: Retry\n• q/Ctrl+C: Exit"))
	}

	return styles.BaseStyle.Render(s.String())
}
