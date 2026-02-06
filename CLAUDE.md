# Termfolio

A terminal-based portfolio application built with Go and Bubble Tea.

## Tech Stack

- **Language**: Go 1.25+
- **TUI Framework**: [Bubble Tea](https://github.com/charmbracelet/bubbletea) (Elm architecture)
- **Styling**: [Lipgloss](https://github.com/charmbracelet/lipgloss)

### Libraries

| Library | Purpose | Install |
|---------|---------|---------|
| `bubbletea` | TUI framework (core) | `go get github.com/charmbracelet/bubbletea` |
| `lipgloss` | Styling, colors, layout | `go get github.com/charmbracelet/lipgloss` |
| `bubbles` | Pre-built components (textinput, viewport, list, spinner) | `go get github.com/charmbracelet/bubbles` |
| `harmonica` | Spring-based smooth animations | `go get github.com/charmbracelet/harmonica` |
| `bubblezone` | Mouse event tracking | `go get github.com/lrstanley/bubblezone` |
| `ntcharts` | Terminal charts and graphs | `go get github.com/NimbleMarkets/ntcharts` |
| `wish` | SSH server support | `go get github.com/charmbracelet/wish` |

## Project Structure

```
termfolio/
├── main.go           # Main application entry point (model, view, update)
├── data/
│   └── resume.go     # Resume data structures and content
├── go.mod            # Go module definition
├── go.sum            # Dependency checksums
└── planning.md       # Sprint planning document
```

## Build & Run

```bash
# Run the application
go run main.go

# Build binary
go build -o termfolio

# Run binary
./termfolio
```

## Architecture

The app follows the **Elm architecture** (Model-Update-View):

- **Model**: Application state (dimensions, cursor position, language)
- **Update**: Handle keyboard/window events
- **View**: Render the UI as a string

### Key Components

- **Navigation**: Horizontal menu (Experience, Projects, Skills)
- **Panels**: Left panel (profile info, bio) + Right panel (content based on selected tab)
- **i18n**: Bilingual support (EN/FR) via `translations` map
- **Keybindings**: `q` quit, `←→` navigate tabs, `Tab` toggle language

## Code Conventions

- Colors defined as `lipgloss.Color` variables at top of file
- Styles defined as `lipgloss.NewStyle()` variables
- Translations use a `map[Lang]map[string]string` structure
- Helper `t(key)` method on model for translations
- Manual frame/border rendering for custom box-drawing layouts

## Resume Data

Resume content lives in `data/resume.go` with these types:
- `Contact`, `Experience`, `Project`, `Education`, `Resume`
- `GetResume()` returns the populated data

## Development Notes

- Minimum terminal size: 80x20
- Uses OSC 8 hyperlinks for clickable URLs in supported terminals
- ASCII art name banner at top of profile
- Frame borders use Unicode box-drawing characters
