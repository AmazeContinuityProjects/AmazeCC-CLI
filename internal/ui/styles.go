// Package ui holds shared visuals: colors, typography, layout helpers.
// Centralizing styles keeps screens consistent and easy to retheme.
package ui

import "github.com/charmbracelet/lipgloss"

// Palette — academic / calm tones; tweak once, applies everywhere.
var (
	ColorPrimary   = lipgloss.Color("#5B8DEF")
	ColorAccent    = lipgloss.Color("#34D399")
	ColorMuted     = lipgloss.Color("#6B7280")
	ColorSurface   = lipgloss.Color("#1F2937")
	ColorBorder    = lipgloss.Color("#374151")
	ColorError     = lipgloss.Color("#F87171")
	ColorHighlight = lipgloss.Color("#FBBF24")
)

var (
	AppTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary)

	SidebarTitle = lipgloss.NewStyle().
			Foreground(ColorMuted).
			Bold(true)

	NavItem = lipgloss.NewStyle().
		PaddingLeft(1)

	NavItemActive = lipgloss.NewStyle().
			PaddingLeft(1).
			Bold(true).
			Foreground(ColorAccent).
			Background(ColorSurface)

	PanelTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#E5E7EB"))

	PanelBody = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#D1D5DB"))

	StatusOK = lipgloss.NewStyle().
			Foreground(ColorAccent)

	StatusError = lipgloss.NewStyle().
			Foreground(ColorError)

	HelpBar = lipgloss.NewStyle().
		Foreground(ColorMuted)

	Placeholder = lipgloss.NewStyle().
			Foreground(ColorMuted).
			Italic(true)
)