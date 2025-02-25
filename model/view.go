package model

import (
	"strings"
	"tui-gac/model/styles"
)

func (m Model) View() string {
	var s strings.Builder

	// ヘッダー部分
	s.WriteString(styles.TitleStyle.Render("🛠  Git Commit Assistant") + "\n")
	s.WriteString(styles.BranchStyle.Render("Branch: "+m.CurrentBranch) + "\n\n")

	// メインコンテンツ
	switch m.CurrentState {
	case GetBranch:
		if m.IssueNum == "" {
			s.WriteString(styles.SubtitleStyle.Render("Issue番号の入力") + "\n")
			s.WriteString(styles.InputStyle.Render(m.InputIssueNum.View()) + "\n\n")
			s.WriteString(styles.HelpStyle.Render("• Enter: 確定\n• Ctrl+C: 終了"))
		} else {
			s.WriteString(styles.StatusStyle.Render("✓ Issue: "+m.IssueNum) + "\n\n")
			s.WriteString(styles.HelpStyle.Render("• Enter: 次へ\n• c: Issue番号を修正\n• Ctrl+C: 終了"))
		}

	case CheckBranchAndIssueNum:
		s.WriteString(styles.SubtitleStyle.Render("ブランチ情報の確認") + "\n\n")
		s.WriteString(styles.StatusStyle.Render("Branch: "+m.CurrentBranch) + "\n")
		s.WriteString(styles.StatusStyle.Render("Issue: "+m.IssueNum) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: 次へ\n• c: Issue番号を修正\n• Ctrl+C: 終了"))

	case FixIssueNumber:
		s.WriteString(styles.SubtitleStyle.Render("Issue番号の修正") + "\n\n")
		s.WriteString(styles.InputStyle.Render(m.InputIssueNum.View()) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: 確定\n• Ctrl+C: 終了"))

	case InputIssueNum:
		s.WriteString(styles.SubtitleStyle.Render("Issue番号の入力") + "\n")
		s.WriteString(styles.InputStyle.Render(m.InputIssueNum.View()) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: 確定\n• Ctrl+C: 終了"))

	case AddAllOrSelect:
		s.WriteString(styles.SubtitleStyle.Render("ステージングするファイルの選択") + "\n\n")

		if len(m.ChangedFiles) > 0 {
			s.WriteString(styles.ItemStyle.Render("変更されたファイル:") + "\n")
			for _, file := range m.ChangedFiles {
				s.WriteString(styles.ItemStyle.Render("  • "+file) + "\n")
			}
		}

		if len(m.DeletedFiles) > 0 {
			if len(m.ChangedFiles) > 0 {
				s.WriteString("\n")
			}
			s.WriteString(styles.WarningStyle.Render("削除されたファイル:") + "\n")
			for _, file := range m.DeletedFiles {
				s.WriteString(styles.WarningStyle.Render("  • "+file) + "\n")
			}
		}
		if len(m.ChangedFiles) > 0 || len(m.DeletedFiles) > 0 {
			s.WriteString("\n" + styles.HelpStyle.Render("• y: 全てのファイルを追加\n• n: 個別に選択\n• Ctrl+C: 終了"))
		}

	case AddSelectedFiles:
		s.WriteString(styles.SubtitleStyle.Render("ファイルの選択") + "\n\n")
		for i, file := range m.ChangedFiles {
			prefix := "   "
			if i == m.Cursor {
				prefix = " ➜ "
			}
			style := styles.ItemStyle
			if i == m.Cursor {
				style = style.Inherit(styles.SelectedStyle)
			}
			status := " "
			if m.AddFile[i] {
				status = "✓"
			}
			s.WriteString(style.Render(prefix+status+" "+file) + "\n")
		}
		s.WriteString("\n" + styles.HelpStyle.Render("• ↑/↓: 移動\n• y: 選択\n• n: 選択解除\n• Enter: 確定"))

	case SelectFixOverView:
		s.WriteString(styles.SubtitleStyle.Render("コミットタイプの選択") + "\n\n")
		for i, fix := range m.FixOverView {
			prefix := "   "
			if i == m.Cursor {
				prefix = " ➜ "
			}
			style := styles.ItemStyle
			if i == m.Cursor {
				style = style.Inherit(styles.SelectedStyle)
			}
			s.WriteString(style.Render(prefix+fix) + "\n")
		}
		s.WriteString("\n" + styles.HelpStyle.Render("• ↑/↓: 移動\n• Enter: 選択"))

	case InputCommitMessage:
		s.WriteString(styles.SubtitleStyle.Render("コミットメッセージの入力") + "\n\n")
		s.WriteString(styles.StatusStyle.Render("Issue: "+m.IssueNum) + "\n")
		s.WriteString(styles.StatusStyle.Render("Type: "+m.FixOverView[m.Cursor]) + "\n\n")
		s.WriteString(styles.InputStyle.Render("Message: "+m.InputCommitMessage.View()) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: 確定\n• Ctrl+C: 終了"))

	case Commit:
		s.WriteString(styles.SubtitleStyle.Render("コミットメッセージの確認") + "\n\n")
		s.WriteString(styles.StatusStyle.Render("コミットメッセージ: "+m.CommitMessage) + "\n\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: コミットを実行\n• Ctrl+C: キャンセル"))

	case Push:
		s.WriteString(styles.SubtitleStyle.Render("以下の内容でpushします") + "\n\n")
		s.WriteString(styles.StatusStyle.Render("コミットメッセージ: "+m.CommitMessage) + "\n")
		s.WriteString(styles.HelpStyle.Render("• Enter: プッシュ\n• Ctrl+C: 終了"))
	}

	return styles.BaseStyle.Render(s.String())
}
