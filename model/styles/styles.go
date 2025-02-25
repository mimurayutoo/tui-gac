package styles

import "github.com/charmbracelet/lipgloss"

var (
	// 基本色の定義
	subtle    = lipgloss.AdaptiveColor{Light: "#666666", Dark: "#999999"}
	highlight = lipgloss.AdaptiveColor{Light: "#2AC3DE", Dark: "#89DCEB"} // 爽やかな水色
	special   = lipgloss.AdaptiveColor{Light: "#40A02B", Dark: "#A6E3A1"} // 明るい緑
	warning   = lipgloss.AdaptiveColor{Light: "#FE640B", Dark: "#FFA066"} // オレンジ

	// Base style - 最小限の余白
	BaseStyle = lipgloss.NewStyle().
			PaddingLeft(1)

	// Title style - 爽やかな印象に
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#2AC3DE")). // 鮮やかな水色
			PaddingBottom(1)

	// Branch style - 緑で安定感を表現
	BranchStyle = lipgloss.NewStyle().
			Foreground(special).
			PaddingLeft(2)

	// Input style - 水色で視認性を確保
	InputStyle = lipgloss.NewStyle().
			Foreground(highlight).
			PaddingLeft(2)

	// Item style - 通常のテキスト
	ItemStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	// Selected style - 選択項目を水色で強調
	SelectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight)

	// Help style - 控えめなグレー
	HelpStyle = lipgloss.NewStyle().
			Foreground(subtle).
			Italic(true)

	// Status style - 成功を表す緑
	StatusStyle = lipgloss.NewStyle().
			Foreground(special)

	// Warning style - 警告を表すオレンジ
	WarningStyle = lipgloss.NewStyle().
			Foreground(warning)

	// Error style - エラーを表す赤
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E64553"))

	// Subtitle style - 水色でセクションを区切り
	SubtitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight).
			PaddingBottom(1)
)
