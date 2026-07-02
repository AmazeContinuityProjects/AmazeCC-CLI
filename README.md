# AmazeCC CLI

AmazeCC CLI is a terminal user interface for viewing academic-progress data from an AmazeCC/FastAPI backend. It is built in Go with Bubble Tea and Lip Gloss.

The current version is a clean TUI scaffold: it starts a full-screen terminal app, checks API reachability, and provides placeholder screens for dashboard, courses, grades, progress, and profile data.

## Features

- Full-screen terminal UI powered by Bubble Tea
- Sidebar navigation for academic-progress sections
- API health check on startup and manual refresh
- Configurable backend URL through an environment variable
- Shared UI components and styles for future screens

## Project Status

This repository currently contains the TUI foundation and API client plumbing. The screen content is still placeholder text until the backend endpoints are wired in.

Planned endpoint areas are visible in the app:

- `GET /courses`
- `GET /grades`
- `GET /progress`
- `GET /me`

Only the configured API root is called by the current code.

## Requirements

- Go `1.26.1` or newer, as declared in `go.mod`
- A terminal that supports ANSI escape sequences
- Optional: access to the AmazeCC API

## Installation

Clone the repository:

```bash
git clone https://github.com/AmazeContinuityProjects/AmazeCC-CLI.git
cd AmazeCC-CLI
```

Install dependencies:

```bash
go mod download
```

Build the CLI:

```bash
go build -o amazecc ./cmd/amaze
```

Run it:

```bash
./amazecc
```

You can also run it without creating a binary:

```bash
go run ./cmd/amaze
```

## Configuration

By default, the CLI connects to:

```text
https://api.amazecc.com
```

Set `AMAZE_API_URL` to point the CLI at another backend:

```bash
AMAZE_API_URL="http://localhost:8000" ./amazecc
```

The app checks the configured API root:

```text
GET https://api.amazecc.com
```

If the endpoint returns a `2xx` response, the status bar shows `connected`. Otherwise, it shows the HTTP status or `unreachable`.

## Keyboard Shortcuts

| Key | Action |
| --- | --- |
| `tab`, `right`, `l` | Next screen |
| `shift+tab`, `left`, `h` | Previous screen |
| `down`, `j` | Move down in navigation |
| `up`, `k` | Move up in navigation |
| `r` | Refresh API health status |
| `q`, `ctrl+c` | Quit |

## Repository Layout

```text
.
├── cmd/amaze/main.go          # Application entry point
├── internal/api/client.go     # HTTP client and health check
├── internal/app/              # Bubble Tea model, update loop, messages, views
├── internal/config/config.go  # Environment-based runtime config
├── internal/ui/               # Reusable Lip Gloss components and styles
├── go.mod
└── go.sum
```

## Development

Format the code:

```bash
gofmt -w cmd internal
```

Run tests:

```bash
go test ./...
```

Build:

```bash
go build -o amazecc ./cmd/amaze
```

## Extending the App

To add a real data screen:

1. Add an API method in `internal/api/client.go`.
2. Add any async message type in `internal/app/messages.go`.
3. Trigger the request from `internal/app/update.go`.
4. Render the result in `internal/app/view.go`.
5. Reuse shared layout and styles from `internal/ui`.

## Tech Stack

- Go
- Bubble Tea
- Lip Gloss

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).
