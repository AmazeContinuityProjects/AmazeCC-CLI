package app

import (
	"fmt"
	"strings"

	"AmazeCLI/internal/ui"

	"github.com/charmbracelet/lipgloss"
)

// defaultTermSize is used when the terminal hasn't reported dimensions yet.
const defaultTermWidth, defaultTermHeight = 80, 24

// View renders the current frame. Pure function — no I/O, no mutation.
func (m Model) View() string {
	width, height := m.width, m.height
	if width == 0 {
		width = defaultTermWidth
	}
	if height == 0 {
		height = defaultTermHeight
	}

	const (
		headerLines = 2
		statusLines = 1
		gap         = 1
	)

	contentHeight := height - headerLines - statusLines - gap
	if contentHeight < 8 {
		contentHeight = 8
	}

	sidebarW := 24
	mainW := width - sidebarW - 1
	if mainW < 20 {
		mainW = 20
	}

	var b strings.Builder

	b.WriteString(ui.Header("Amaze", "academic progress viewer", width))
	b.WriteString("\n\n")

	sidebar := ui.Sidebar(screenNames, int(m.screen), contentHeight)
	main := m.renderScreen(mainW, contentHeight)

	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, sidebar, main))
	b.WriteString("\n")
	b.WriteString(ui.StatusBar(m.apiStatus, m.apiConnected, width))

	return b.String()
}

func (m Model) renderScreen(width, height int) string {
	if m.loading {
		return ui.PlaceholderPanel(
			m.screen.String(),
			"Loading…",
			width, height,
		)
	}

	// Each case will eventually call api.Client methods and render real data.
	switch m.screen {
	case ScreenDashboard:
		return ui.PlaceholderPanel(
			"Dashboard",
			"Overview of GPA, active courses, and upcoming deadlines.\n\n"+
				fmt.Sprintf("API: %s (%s)", m.config.APIBaseURL, m.apiStatus),
			width, height,
		)

	case ScreenCourses:
		return ui.PlaceholderPanel(
			"Courses",
			"Enrolled courses, credits, and instructors.\n\n"+
				"Wire to: GET /courses (or your FastAPI route)",
			width, height,
		)

	case ScreenGrades:
		return ui.PlaceholderPanel(
			"Grades",
			"Assignments, exams, and term breakdowns.\n\n"+
				"Wire to: GET /grades",
			width, height,
		)

	case ScreenProgress:
		return ui.PlaceholderPanel(
			"Progress",
			"Degree completion, requirements, and milestones.\n\n"+
				"Wire to: GET /progress",
			width, height,
		)

	case ScreenProfile:
		return ui.PlaceholderPanel(
			"Profile",
			"Student ID, program, and account settings.\n\n"+
				"Wire to: GET /me",
			width, height,
		)

	default:
		return ui.PlaceholderPanel("?", "Unknown screen", width, height)
	}
}