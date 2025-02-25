package styles

import "github.com/charmbracelet/lipgloss"

var (
	// モダンでクリーンな配色定義
	subtle    = lipgloss.AdaptiveColor{Light: "#767C9DB2", Dark: "#A1AACDB2"} // よりソフトなグレー
	highlight = lipgloss.AdaptiveColor{Light: "#0EA5E9", Dark: "#38BDF8"}     // 明るい青
	special   = lipgloss.AdaptiveColor{Light: "#10B981", Dark: "#34D399"}     // 鮮やかな緑
	warning   = lipgloss.AdaptiveColor{Light: "#F59E0B", Dark: "#FBBF24"}     // 温かみのあるオレンジ
	error     = lipgloss.AdaptiveColor{Light: "#EF4444", Dark: "#F87171"}     // 目立つ赤

	// Base style - クリーンな余白
	BaseStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			PaddingRight(2)

	// Title style - 大きく目立つヘッダー
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderBottom(true).
			MarginBottom(1).
			PaddingLeft(2).
			PaddingRight(2)

	// Branch style - 目立つが控えめな表示
	BranchStyle = lipgloss.NewStyle().
			Foreground(special).
			PaddingLeft(2).
			MarginBottom(1)

	// Input style - 入力欄を強調
	InputStyle = lipgloss.NewStyle().
			Foreground(highlight).
			PaddingLeft(4).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderLeft(true)

	// Item style - 整理されたリスト表示
	ItemStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.AdaptiveColor{Light: "#334155", Dark: "#CBD5E1"})

	// Selected style - 選択項目を明確に
	SelectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight).
			Background(lipgloss.AdaptiveColor{Light: "#E0F2FE", Dark: "#0C4A6E"})

	// Help style - 読みやすいヘルプテキスト
	HelpStyle = lipgloss.NewStyle().
			Foreground(subtle).
			Italic(true).
			PaddingTop(1).
			PaddingBottom(1)

	// Status style - 成功状態の表示
	StatusStyle = lipgloss.NewStyle().
			Foreground(special).
			PaddingLeft(4).
			Bold(true)

	// Warning style - 警告の表示
	WarningStyle = lipgloss.NewStyle().
			Foreground(warning).
			PaddingLeft(4).
			Bold(true)

	// Error style - エラーメッセージ
	ErrorStyle = lipgloss.NewStyle().
			Foreground(error).
			Bold(true).
			PaddingLeft(4).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderLeft(true).
			BorderForeground(error)

	// Subtitle style - セクションタイトル
	SubtitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(highlight).
			PaddingBottom(1).
			PaddingLeft(2).
			MarginTop(1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderLeft(true)
)
