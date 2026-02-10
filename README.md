# termfolio.dev

A terminal-based portfolio built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## Features

- Interactive TUI with keyboard navigation
- Bilingual support (English/French)
- Clickable hyperlinks (OSC 8) in supported terminals
- Responsive layout with split panels
- ASCII art header

## Quick Start

```bash
# Clone the repo
git clone https://github.com/MohamedGacha/MohamedGACHA-termfolio.dev.git
cd MohamedGACHA-termfolio.dev

# Run
go run main.go

# Or build and run
go build -o termfolio && ./termfolio
```

## Controls

| Key | Action |
|-----|--------|
| `q` | Quit |
| `←` `→` | Navigate tabs |
| `Tab` | Toggle language (EN/FR) |

## Tech Stack

- **Framework**: [Bubble Tea](https://github.com/charmbracelet/bubbletea) (Elm architecture)
- **Styling**: [Lipgloss](https://github.com/charmbracelet/lipgloss)

## License

MIT
