package app

import (
	"AmazeCLI/internal/api"
	"AmazeCLI/internal/config"
)

// Screen identifies which main view is active.
// Add new screens here as you map FastAPI endpoint groups to TUI pages.
type Screen int

const (
	ScreenDashboard Screen = iota
	ScreenCourses
	ScreenGrades
	ScreenProgress
	ScreenProfile
	screenCount
)

// screenNames line up 1:1 with Screen constants — used for sidebar labels.
var screenNames = []string{
	"Dashboard",
	"Courses",
	"Grades",
	"Progress",
	"Profile",
}

func (s Screen) String() string {
	if int(s) < 0 || int(s) >= len(screenNames) {
		return "Unknown"
	}
	return screenNames[s]
}

// Model is the single source of truth for all TUI state.
type Model struct {
	client *api.Client
	config config.Config

	// Layout
	width  int
	height int

	// Navigation
	screen Screen

	// API
	apiConnected bool
	apiStatus    string
	loading      bool
}

// New builds the root Bubble Tea model with injected dependencies.
func New(client *api.Client, cfg config.Config) Model {
	return Model{
		client:    client,
		config:    cfg,
		screen:    ScreenDashboard,
		apiStatus: "checking…",
		loading:   true,
	}
}