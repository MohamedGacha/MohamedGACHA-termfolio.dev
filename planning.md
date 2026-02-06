# Termfolio Sprint Plan

A terminal-based portfolio application built with Go and Bubble Tea.

---

## General Sprint Overview

| Day | Focus | Goal |
|-----|-------|------|
| 1 | Learning | Learn Go basics and Bubble Tea library |
| 2 | Architecture | Design app structure and components |
| 3 | Core UI | Build main navigation and layout |
| 4 | Sections | Implement About, Skills, Projects sections |
| 5 | Styling | Add colors, animations, and polish |
| 6 | Testing | Test, debug, and refine |
| 7 | Deployment | Build, document, and release |

---

## Day 1: Learning Go & Libraries

### Objectives
- Understand Go fundamentals
- Learn Bubble Tea TUI framework
- Build a working prototype

### Tasks

- [x] Set up Go project with `go mod init`
- [x] Install Bubble Tea: `go get github.com/charmbracelet/bubbletea`
- [x] Learn the Elm architecture (Model, Update, View)
- [x] Build basic interactive example (todo list selector)
- [ ] Explore Lip Gloss for styling
- [ ] Explore Bubbles components (textinput, viewport, list, spinner)
- [ ] Explore Harmonica for animations
- [ ] Explore BubbleZone for mouse support
- [ ] Explore ntcharts for potential skill visualizations
- [ ] Study reference projects (Glow, gh-dash, Superfile)
- [ ] Read Bubble Tea documentation and examples

### Key Concepts Learned

**Bubble Tea Architecture:**
```
Model (state) -> Update (handle input) -> View (render UI)
```

**Bubble Tea Ecosystem:**

| Library | Purpose | Install |
|---------|---------|---------|
| `bubbletea` | TUI framework (core) | `go get github.com/charmbracelet/bubbletea` |
| `lipgloss` | Styling, colors, layout | `go get github.com/charmbracelet/lipgloss` |
| `bubbles` | Pre-built components (textinput, viewport, list, spinner, etc.) | `go get github.com/charmbracelet/bubbles` |
| `harmonica` | Spring-based smooth animations | `go get github.com/charmbracelet/harmonica` |
| `bubblezone` | Mouse event tracking | `go get github.com/lrstanley/bubblezone` |
| `ntcharts` | Terminal charts and graphs | `go get github.com/NimbleMarkets/ntcharts` |

**Reference Projects (Staff Favourites):**
- [Glow](https://github.com/charmbracelet/glow) - Markdown reader/browser
- [Huh?](https://github.com/charmbracelet/huh) - Interactive prompts and forms
- [gh-dash](https://github.com/dlvhdr/gh-dash) - GitHub CLI dashboard for PRs/issues
- [Superfile](https://github.com/yorukot/superfile) - Terminal file manager
- [circumflex](https://github.com/bensadeh/circumflex) - Hacker News TUI
- [Tetrigo](https://github.com/Broderick-Westrope/tetrigo) - Tetris in terminal

### Notes
- Current `main.go` contains a basic todo selector example
- Uses vim-style keybindings (j/k for navigation)
- Space/Enter to select items, q to quit

---

## Day 2: Architecture (Upcoming)

### Planned Tasks
- [ ] Define app sections (About, Skills, Projects, Contact)
- [ ] Create component structure
- [ ] Design navigation system
- [ ] Plan data models for portfolio content

---

## Day 3: Core UI (Upcoming)

### Planned Tasks
- [ ] Implement main menu/navigation
- [ ] Create tab or sidebar navigation
- [ ] Build responsive layout
- [ ] Add header/footer components

---

## Day 4: Sections (Upcoming)

### Planned Tasks
- [ ] **About Me** - Name, title, bio, photo placeholder
- [ ] **Experience** - 3 positions (Gatewatcher, Etifak, Suez)
- [ ] **Projects** - SimplyLovelySetups, MealPass
- [ ] **Skills** - 26 skills with visual representation
- [ ] **Education** - ENSISA, Carnot prépas, Baccalauréat
- [ ] **Languages** - FR, EN, AR, ES with levels
- [ ] **Contact** - Email, phone, LinkedIn, GitHub
- [ ] **Interests** - F1, Gaming, Astrophysics, Aeronautics

### Data Source
Resume data stored in `data/resume.go`

---

## Day 5: Styling (Upcoming)

### Planned Tasks
- [ ] Define color scheme
- [ ] Add Lip Gloss styling
- [ ] Implement smooth transitions
- [ ] Add ASCII art or banner

---

## Day 6: Testing (Upcoming)

### Planned Tasks
- [ ] Test on different terminal sizes
- [ ] Handle edge cases
- [ ] Performance optimization
- [ ] Code cleanup and refactoring

---

## Day 7: Deployment (Upcoming)

### Planned Tasks
- [ ] Build release binary
- [ ] Write README documentation
- [ ] Add installation instructions
- [ ] Publish to GitHub
