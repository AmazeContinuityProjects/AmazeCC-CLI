package app

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
)

// Init runs once at startup. Schedule background work here (API calls, timers).
func (m Model) Init() tea.Cmd {
	// WindowSize explicitly requests terminal dimensions (some Windows terminals
	// don't send WindowSizeMsg automatically on launch).
	return tea.Batch(tea.WindowSize(), m.checkHealth())
}

// Update handles every event and returns the next model + any follow-up commands.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab", "right", "l":
			m.screen = Screen((int(m.screen) + 1) % int(screenCount))
			return m, nil

		case "shift+tab", "left", "h":
			prev := int(m.screen) - 1
			if prev < 0 {
				prev = int(screenCount) - 1
			}
			m.screen = Screen(prev)
			return m, nil

		case "up", "k":
			prev := int(m.screen) - 1
			if prev < 0 {
				prev = int(screenCount) - 1
			}
			m.screen = Screen(prev)
			return m, nil

		case "down", "j":
			m.screen = Screen((int(m.screen) + 1) % int(screenCount))
			return m, nil

		case "r":
			m.loading = true
			m.apiStatus = "checking…"
			return m, m.checkHealth()
		}

	case healthCheckMsg:
		m.loading = false
		if msg.err != nil {
			m.apiConnected = false
			m.apiStatus = "error"
			return m, nil
		}
		m.apiConnected = msg.result.OK
		m.apiStatus = msg.result.Status
		return m, nil
	}

	return m, nil
}

// checkHealth returns a tea.Cmd that pings the API without blocking the UI loop.
func (m Model) checkHealth() tea.Cmd {
	return func() tea.Msg {
		result, err := m.client.Ping(context.Background())
		return healthCheckMsg{result: result, err: err}
	}
}