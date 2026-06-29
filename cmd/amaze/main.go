// Entry point — keep this file thin.
// All TUI logic lives under internal/app; wire dependencies here only.
package main

import (
	"fmt"
	"os"

	"AmazeCLI/internal/api"
	"AmazeCLI/internal/app"
	"AmazeCLI/internal/config"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cfg := config.Load()

	// API client is injected so screens never hard-code URLs or HTTP details.
	client := api.NewClient(cfg.APIBaseURL)

	p := tea.NewProgram(
		app.New(client, cfg),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "amaze: %v\n", err)
		os.Exit(1)
	}
}