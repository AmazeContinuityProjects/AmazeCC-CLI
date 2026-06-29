package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Sidebar renders the left navigation column.
func Sidebar(items []string, active int, height int) string {
	var b strings.Builder
	b.WriteString(SidebarTitle.Render("  NAVIGATION"))
	b.WriteString("\n\n")

	for i, item := range items {
		line := fmt.Sprintf("  %s", item)
		if i == active {
			line = NavItemActive.Render("▸ " + item)
		} else {
			line = NavItem.Render("  " + item)
		}
		b.WriteString(line)
		b.WriteString("\n")
	}

	// Pad so the sidebar fills the available height (clean split-pane look).
	linesUsed := len(items) + 3
	for linesUsed < height-2 {
		b.WriteString("\n")
		linesUsed++
	}

	return lipgloss.NewStyle().
		Width(22).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorBorder).
		Padding(0, 1).
		Render(b.String())
}

// Header renders the top app bar (title + optional subtitle).
func Header(title, subtitle string, width int) string {
	line := AppTitle.Render(title)
	if subtitle != "" {
		line += "  " + Placeholder.Render(subtitle)
	}
	return lipgloss.NewStyle().Width(width).Render(line)
}

// StatusBar shows API connection state and keyboard hints.
func StatusBar(apiStatus string, connected bool, width int) string {
	var statusStyled string
	if connected {
		statusStyled = StatusOK.Render("● " + apiStatus)
	} else {
		statusStyled = StatusError.Render("● " + apiStatus)
	}

	left := statusStyled
	right := HelpBar.Render("tab/←→ screen • j/k navigate • r refresh • q quit")

	gap := width - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 1 {
		gap = 1
	}

	return lipgloss.NewStyle().Width(width).Render(left + strings.Repeat(" ", gap) + right)
}

// PlaceholderPanel is a stub content area until real API data is wired in.
func PlaceholderPanel(title, body string, width, height int) string {
	content := PanelTitle.Render(title) + "\n\n" + PanelBody.Render(body)
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorBorder).
		Padding(1, 2).
		Render(content)
}