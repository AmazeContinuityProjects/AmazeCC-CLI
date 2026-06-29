package app

import "AmazeCLI/internal/api"

// Custom messages decouple async work from tea.KeyMsg / tea.WindowSizeMsg.
// Bubble Tea delivers these back into Update() after a tea.Cmd finishes.

// healthCheckMsg carries the result of an API ping.
type healthCheckMsg struct {
	result api.HealthResult
	err    error
}