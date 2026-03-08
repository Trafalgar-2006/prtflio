package main

import (
	"github.com/charmbracelet/lipgloss"
)

// Color palette — inspired by the Mori Liu aesthetic: cyan accents on dark bg
var (
	ColorCyan      = lipgloss.Color("#00DFDF")
	ColorMagenta   = lipgloss.Color("#FF6AC1")
	ColorDimWhite  = lipgloss.Color("#888888")
	ColorWhite     = lipgloss.Color("#E0E0E0")
	ColorBrightWht = lipgloss.Color("#FFFFFF")
	ColorDarkBg    = lipgloss.Color("#0A0A0A")
	ColorGold      = lipgloss.Color("#FFD700")
	ColorGreen     = lipgloss.Color("#50FA7B")
	ColorPurple    = lipgloss.Color("#BD93F9")
	ColorOrange    = lipgloss.Color("#FFB86C")
	ColorRed       = lipgloss.Color("#FF5555")
	ColorPink      = lipgloss.Color("#FF79C6")
)

// Styles
var (
	// Tab styles
	ActiveTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorCyan)

	InactiveTabStyle = lipgloss.NewStyle().
				Foreground(ColorDimWhite)

	TabSeparator = lipgloss.NewStyle().
			Foreground(ColorDimWhite).
			SetString("    ")

	// Content styles
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorCyan).
			MarginBottom(1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(ColorMagenta).
			Italic(true)

	BodyStyle = lipgloss.NewStyle().
			Foreground(ColorWhite)

	DimStyle = lipgloss.NewStyle().
			Foreground(ColorDimWhite).
			Italic(true)

	// Project card styles
	ProjectTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorGold)

	ProjectDescStyle = lipgloss.NewStyle().
				Foreground(ColorWhite)

	TagStyle = lipgloss.NewStyle().
			Foreground(ColorPurple)

	AccuracyStyle = lipgloss.NewStyle().
			Foreground(ColorGreen).
			Bold(true)

	// Contact styles
	ContactIconStyle = lipgloss.NewStyle().
				Foreground(ColorCyan)

	ContactLinkStyle = lipgloss.NewStyle().
				Foreground(ColorWhite).
				Underline(true)

	// ASCII art style
	AsciiArtStyle = lipgloss.NewStyle().
			Foreground(ColorCyan)

	// Star decorations
	StarStyle = lipgloss.NewStyle().
			Foreground(ColorDimWhite)

	BrightStarStyle = lipgloss.NewStyle().
				Foreground(ColorCyan)

	// Navigation hint
	HintStyle = lipgloss.NewStyle().
			Foreground(ColorDimWhite).
			Italic(true)

	// Selected item in lists
	SelectedStyle = lipgloss.NewStyle().
			Foreground(ColorCyan).
			Bold(true)

	UnselectedStyle = lipgloss.NewStyle().
				Foreground(ColorDimWhite)

	// Borders
	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorCyan).
			Padding(1, 2)

	// Bio section
	BioKeyStyle = lipgloss.NewStyle().
			Foreground(ColorCyan).
			Bold(true)

	BioValueStyle = lipgloss.NewStyle().
			Foreground(ColorWhite)

	SectionDivider = lipgloss.NewStyle().
				Foreground(ColorDimWhite).
				SetString("─────────────────────────────────────────")
)
