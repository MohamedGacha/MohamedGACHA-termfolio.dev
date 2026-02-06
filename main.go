package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Colors
var (
	primaryColor   = lipgloss.Color("#0174DF") // Blue accent
	secondaryColor = lipgloss.Color("#08298A") // Dark blue
	textColor      = lipgloss.Color("#FFFFFF")
	mutedColor     = lipgloss.Color("#666666")
)

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	contentStyle = lipgloss.NewStyle().
			Foreground(textColor)

	mutedStyle = lipgloss.NewStyle().
			Foreground(mutedColor)

	footerStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			MarginTop(1)
)

// Minimum dimensions
const (
	minWidth  = 80
	minHeight = 20
)

// Language type
type Lang string

const (
	EN Lang = "en"
	FR Lang = "fr"
)

// Translations
var translations = map[Lang]map[string]string{
	EN: {
		"subtitle":       "Computer Science & Networks Engineer",
		"about_title":    "About Me",
		"bio_title":      "Biography",
		"about_1":        "Backend Developer & Computer Science Engineer",
		"about_2":        "passionate about building robust APIs and",
		"about_3":        "scalable systems.",
		"about_4":        "Currently working on cybersecurity solutions",
		"about_5":        "at Gatewatcher, contributing to open-source",
		"about_6":        "projects and developing CLI tools.",
		"exp_title":      "Experience",
		"exp_scroll":     "(↑↓ to scroll)",
		"mission":        "Mission:",
		"stack":          "Stack:",
		"challenge":      "Challenge:",
		"feedback":       "Feedback:",
		"gw_role":        "Cybersecurity Engineer | Paris | Oct 2025 - Sept 2026",
		"gw_mission_1":   "Working on GCap NDR solution. Managing cloud",
		"gw_mission_2":   "infrastructure (Proxmox, Nutanix, Azure, AWS). Building",
		"gw_mission_3":   "developer tools to automate repetitive tasks for the team.",
		"gw_challenge":   "Learning cybersecurity from scratch with the team.",
		"gw_feedback_1":  "Loved the challenge & growth! Great AI philosophy",
		"gw_feedback_2":  "- using AI to sharpen skills while ensuring top quality.",
		"etifak_role":    "Main Backend Developer | Remote | Sept 2024 - Sept 2025",
		"etifak_mission_1": "Led backend development for B2B marketplace",
		"etifak_mission_2": "at an early-stage startup. Built REST APIs, unit tests,",
		"etifak_mission_3": "and managed AWS deployment.",
		"etifak_challenge": "Self-taught many concepts as a junior dev.",
		"etifak_feedback":  "Great team communication made workflow smooth!",
		"suez_role":      "Data Engineer Intern | Paris | Jun - Aug 2024",
		"suez_mission_1": "Built intranet web app for data quality & cleaning.",
		"suez_mission_2": "Implemented deduplication using Levenshtein distance.",
		"suez_feedback":  "First internship. Great Scrum team dynamics!",
		"proj_title":     "Projects",
		"sls_desc":       "Free setups for Assetto Corsa Competizione",
		"mealpass_desc":  "Food distribution app for students",
		"skills_title":   "Skills",
		"contact_title":  "Contact",
		"contact_reach":  "Feel free to reach out!",
		"contact_open":   "Open to opportunities in:",
		"contact_backend": "Backend Development",
		"contact_cyber":  "Cybersecurity",
		"contact_devops": "DevOps",
		"footer":         "q: Quit • ←→: Navigate • Tab: Language",
		"email":          "Email",
		"location":       "Location",
		"linkedin":       "LinkedIn",
		"github":         "GitHub",
	},
	FR: {
		"subtitle":       "Ingénieur en informatique et réseaux",
		"about_title":    "À propos",
		"bio_title":      "Biographie",
		"about_1":        "Développeur Backend & Ingénieur Informatique",
		"about_2":        "passionné par la création d'APIs robustes et",
		"about_3":        "de systèmes évolutifs.",
		"about_4":        "Actuellement en cybersécurité chez Gatewatcher,",
		"about_5":        "contribution open-source et développement",
		"about_6":        "d'outils CLI.",
		"exp_title":      "Expériences",
		"exp_scroll":     "(↑↓ pour défiler)",
		"mission":        "Mission:",
		"stack":          "Stack:",
		"challenge":      "Défi:",
		"feedback":       "Avis:",
		"gw_role":        "Ingénieur Cybersécurité | Paris | Oct 2025 - Sept 2026",
		"gw_mission_1":   "Travail sur la solution NDR GCap. Gestion de",
		"gw_mission_2":   "l'infrastructure cloud (Proxmox, Nutanix, Azure, AWS).",
		"gw_mission_3":   "Création d'outils pour automatiser les tâches de l'équipe.",
		"gw_challenge":   "Apprentissage de la cybersécurité avec l'équipe.",
		"gw_feedback_1":  "J'ai adoré le défi et la croissance! Philosophie IA",
		"gw_feedback_2":  "- utiliser l'IA pour améliorer les compétences.",
		"etifak_role":    "Développeur Backend Principal | Remote | Sept 2024 - Sept 2025",
		"etifak_mission_1": "Développement backend pour marketplace B2B",
		"etifak_mission_2": "dans une startup. APIs REST, tests unitaires,",
		"etifak_mission_3": "et déploiement AWS.",
		"etifak_challenge": "Auto-apprentissage en tant que dev junior.",
		"etifak_feedback":  "Excellente communication d'équipe!",
		"suez_role":      "Stagiaire Data Engineer | Paris | Juin - Août 2024",
		"suez_mission_1": "Application intranet pour qualité des données.",
		"suez_mission_2": "Déduplication avec distance de Levenshtein.",
		"suez_feedback":  "Premier stage. Super dynamique Scrum!",
		"proj_title":     "Projets",
		"sls_desc":       "Setups gratuits pour Assetto Corsa Competizione",
		"mealpass_desc":  "App de distribution alimentaire pour étudiants",
		"skills_title":   "Compétences",
		"contact_title":  "Contact",
		"contact_reach":  "N'hésitez pas à me contacter!",
		"contact_open":   "Ouvert aux opportunités en:",
		"contact_backend": "Développement Backend",
		"contact_cyber":  "Cybersécurité",
		"contact_devops": "DevOps",
		"footer":         "q: Quitter • ←→: Naviguer • Tab: Langue",
		"email":          "Email",
		"location":       "Lieu",
		"linkedin":       "LinkedIn",
		"github":         "GitHub",
	},
}

// ASCII name
var name = `╔╦╗╔═╗╦ ╦╔═╗╔╦╗╔═╗╔╦╗  ╔═╗╔═╗╔═╗╦ ╦╔═╗
║║║║ ║╠═╣╠═╣║║║║╣  ║║  ║ ╦╠═╣║  ╠═╣╠═╣
╩ ╩╚═╝╩ ╩╩ ╩╩ ╩╚═╝═╩╝  ╚═╝╩ ╩╚═╝╩ ╩╩ ╩`

type model struct {
	width  int
	height int
	cursor int
	lang   Lang
}

func initialModel() model {
	return model{
		width:  80,
		height: 24,
		cursor: 0,
		lang:   EN,
	}
}

// Helper to get translation
func (m model) t(key string) string {
	if val, ok := translations[m.lang][key]; ok {
		return val
	}
	return key
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left", "h":
			if m.cursor > 0 {
				m.cursor--
			}
		case "right", "l":
			if m.cursor < 2 {
				m.cursor++
			}
		case "tab":
			// Toggle language
			if m.lang == EN {
				m.lang = FR
			} else {
				m.lang = EN
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	}

	return m, nil
}

func (m model) View() string {
	// Calculate dimensions with minimums
	width := m.width
	height := m.height

	if width < minWidth {
		width = minWidth
	}
	if height < minHeight {
		height = minHeight
	}

	// Frame dimensions (with margins)
	innerWidth := width - 4

	// Menu items horizontal
	menuItems := []string{"Experience", "Projects", "Skills"}
	var menuDisplay []string
	for i, item := range menuItems {
		if i == m.cursor {
			menuDisplay = append(menuDisplay, titleStyle.Render(" "+item+" "))
		} else {
			menuDisplay = append(menuDisplay, contentStyle.Render(" "+item+" "))
		}
	}

	// Navbar (no border, sits on top of frame)
	navbarStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#1a1a1a")).
		Padding(0, 1)

	navbar := navbarStyle.Render(lipgloss.JoinHorizontal(lipgloss.Center, menuDisplay...))

	// Content
	subtitle := lipgloss.NewStyle().
		Foreground(mutedColor).
		Render(m.t("subtitle"))
	title := titleStyle.Render(name)

	mainContent := lipgloss.JoinVertical(
		lipgloss.Left,
		subtitle,
		title,
	)

	// Build frame manually
	borderColor := lipgloss.NewStyle().Foreground(primaryColor)

	// Panel widths for vertical split (30% left, 70% right)
	leftPanelWidth := innerWidth * 30 / 100
	rightPanelWidth := innerWidth - leftPanelWidth - 1 // -1 for the middle │

	// Calculate navbar position (top right of overall view)
	navbarLen := lipgloss.Width(navbar)
	navbarRightPad := 1 // minimal padding from right edge
	navbarPadding := rightPanelWidth - navbarLen - 2 - navbarRightPad // -2 for ┤ and ├
	if navbarPadding < 0 {
		navbarPadding = 0
	}

	// Site text box (top left)
	siteText := " termfolio.dev "
	siteLen := lipgloss.Width(siteText)
	siteLeftPad := 1 // minimal dashes before site text
	siteRightPad := leftPanelWidth - siteLeftPad - 2 - siteLen // -2 for ┤ and ├
	if siteRightPad < 0 {
		siteRightPad = 0
	}

	// Site box borders
	siteBoxTop := borderColor.Render("╭" + repeatString("─", siteLen) + "╮")
	siteBoxBottom := borderColor.Render("╰" + repeatString("─", siteLen) + "╯")

	// Navbar box top border
	navbarBoxTop := borderColor.Render("╭" + repeatString("─", navbarLen) + "╮")
	navbarBoxBottom := borderColor.Render("╰" + repeatString("─", navbarLen) + "╯")

	// Frame top line 1: site box top on left, navbar box top on right
	// siteBoxTop's ╭ aligns with ┤ on line 2 (position siteLeftPad+1)
	frameTopLine1 := repeatString(" ", siteLeftPad+1) +
		siteBoxTop +
		repeatString(" ", siteRightPad+navbarPadding+1) +
		navbarBoxTop +
		repeatString(" ", navbarRightPad+1)

	// Frame top line 2: left border + site text + ┬ + right dashes with navbar embedded
	frameTopLine2 := borderColor.Render("╭"+repeatString("─", siteLeftPad)+"┤") +
		contentStyle.Render(siteText) +
		borderColor.Render("├"+repeatString("─", siteRightPad)+"┬"+repeatString("─", navbarPadding)+"┤") +
		navbar +
		borderColor.Render("├"+repeatString("─", navbarRightPad)+"╮")

	// Frame top line 3: site box bottom on left + spaces + middle │ + navbar bottom + right │
	// siteBoxBottom's ╰ aligns with ┤ on line 2 (position siteLeftPad+1)
	frameTopLine3 := borderColor.Render("│") +
		repeatString(" ", siteLeftPad) +
		siteBoxBottom +
		repeatString(" ", siteRightPad) +
		borderColor.Render("│") +
		repeatString(" ", navbarPadding) +
		navbarBoxBottom +
		repeatString(" ", navbarRightPad) +
		borderColor.Render("│")

	frameTop := frameTopLine1 + "\n" + frameTopLine2 + "\n" + frameTopLine3

	// Personal info
	infoLabel := lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	infoValue := lipgloss.NewStyle().Foreground(textColor)

	// Horizontal separator line
	separator := borderColor.Render(repeatString("─", leftPanelWidth-2))

	// Helper for clickable links (OSC 8 hyperlinks)
	link := func(url, text string) string {
		return fmt.Sprintf("\x1b]8;;%s\x07%s\x1b]8;;\x07", url, text)
	}

	// Fixed width for labels to align values
	labelWidth := 10
	formatLabel := func(label string) string {
		padding := labelWidth - len(label)
		if padding < 0 {
			padding = 0
		}
		return infoLabel.Render(label) + repeatString(" ", padding)
	}

	personalInfo := lipgloss.JoinVertical(
		lipgloss.Left,
		separator,
		formatLabel(m.t("email"))+infoValue.Render(link("mailto:simogacha@gmail.com", "simogacha@gmail.com")),
		formatLabel(m.t("location"))+infoValue.Render("France/Paris"),
		formatLabel(m.t("linkedin"))+infoValue.Render(link("https://linkedin.com/in/mohamed-gacha", "linkedin.com/in/mohamed-gacha")),
		formatLabel(m.t("github"))+infoValue.Render(link("https://github.com/MohamedGacha", "github.com/MohamedGacha")),
	)

	// Biography
	bio := lipgloss.JoinVertical(
		lipgloss.Left,
		"",
		titleStyle.Render(m.t("bio_title")),
		contentStyle.Render(m.t("about_1")),
		contentStyle.Render(m.t("about_2")),
		contentStyle.Render(m.t("about_3")),
		"",
		contentStyle.Render(m.t("about_4")),
		contentStyle.Render(m.t("about_5")),
		contentStyle.Render(m.t("about_6")),
	)

	// Left panel content (name, subtitle, personal info, and bio)
	leftPanelContent := lipgloss.JoinVertical(
		lipgloss.Left,
		mainContent,
		personalInfo,
		bio,
	)

	leftContent := lipgloss.NewStyle().
		PaddingLeft(1).
		Width(leftPanelWidth).
		Render(leftPanelContent)

	// Right panel content based on selected tab
	var rightPanelContent string

	switch m.cursor {
	case 0: // Experience
		rightPanelContent = lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(m.t("exp_title"))+" "+mutedStyle.Render(m.t("exp_scroll")),
			"",
			titleStyle.Render("━━━ Gatewatcher ━━━"),
			contentStyle.Render(m.t("gw_role")),
			mutedStyle.Render(m.t("mission"))+" "+contentStyle.Render(m.t("gw_mission_1")),
			contentStyle.Render(m.t("gw_mission_2")),
			contentStyle.Render(m.t("gw_mission_3")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Python, Ansible, Docker, Bash, Linux, CI/CD"),
			mutedStyle.Render(m.t("challenge"))+" "+contentStyle.Render(m.t("gw_challenge")),
			mutedStyle.Render(m.t("feedback"))+" "+contentStyle.Render(m.t("gw_feedback_1")),
			contentStyle.Render(m.t("gw_feedback_2")),
			"",
			titleStyle.Render("━━━ Etifak ━━━"),
			contentStyle.Render(m.t("etifak_role")),
			mutedStyle.Render(m.t("mission"))+" "+contentStyle.Render(m.t("etifak_mission_1")),
			contentStyle.Render(m.t("etifak_mission_2")),
			contentStyle.Render(m.t("etifak_mission_3")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Python, Django, PostgreSQL, Docker, AWS EC2"),
			mutedStyle.Render(m.t("challenge"))+" "+contentStyle.Render(m.t("etifak_challenge")),
			mutedStyle.Render(m.t("feedback"))+" "+contentStyle.Render(m.t("etifak_feedback")),
			"",
			titleStyle.Render("━━━ Suez Digital Solutions ━━━"),
			contentStyle.Render(m.t("suez_role")),
			mutedStyle.Render(m.t("mission"))+" "+contentStyle.Render(m.t("suez_mission_1")),
			contentStyle.Render(m.t("suez_mission_2")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Python, Streamlit, Pandas, HTML/CSS, Azure"),
			mutedStyle.Render(m.t("feedback"))+" "+contentStyle.Render(m.t("suez_feedback")),
		)

	case 1: // Projects
		rightPanelContent = lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(m.t("proj_title")),
			"",
			titleStyle.Render("SimplyLovelySetups.com"),
			contentStyle.Render("  "+m.t("sls_desc")),
			contentStyle.Render("  FastAPI, MongoDB, NextJS, Docker"),
			"",
			titleStyle.Render("MealPass"),
			contentStyle.Render("  "+m.t("mealpass_desc")),
			contentStyle.Render("  Django, PostgreSQL, Flutter"),
		)

	case 2: // Skills
		rightPanelContent = lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(m.t("skills_title")),
			"",
			titleStyle.Render("Backend")+"  "+contentStyle.Render("Python, Django, FastAPI, PostgreSQL"),
			titleStyle.Render("Frontend")+" "+contentStyle.Render("JavaScript, ReactJS, NextJS, HTML/CSS"),
			titleStyle.Render("DevOps")+"   "+contentStyle.Render("Docker, AWS, Azure, CI/CD, Linux"),
			titleStyle.Render("Tools")+"    "+contentStyle.Render("Git, Suricata, Streamlit, Pandas"),
			titleStyle.Render("Other")+"    "+contentStyle.Render("Java, C, C#, F#, Bash"),
		)

	}

	rightContent := lipgloss.NewStyle().
		PaddingLeft(1).
		Width(rightPanelWidth).
		Render(rightPanelContent)

	// Split content into lines
	leftLines := splitLines(leftContent)
	rightLines := splitLines(rightContent)

	contentHeight := height - 7 // Leave space for top (3 lines) + bottom (3 lines) + margin
	var framedContent string
	for i := 0; i < contentHeight; i++ {
		leftLine := ""
		rightLine := ""
		if i < len(leftLines) {
			leftLine = leftLines[i]
		}
		if i < len(rightLines) {
			rightLine = rightLines[i]
		}
		// Pad lines to correct width
		leftLineWidth := lipgloss.Width(leftLine)
		if leftLineWidth < leftPanelWidth {
			leftLine += repeatString(" ", leftPanelWidth-leftLineWidth)
		}
		rightLineWidth := lipgloss.Width(rightLine)
		if rightLineWidth < rightPanelWidth {
			rightLine += repeatString(" ", rightPanelWidth-rightLineWidth)
		}
		framedContent += borderColor.Render("│") + leftLine + borderColor.Render("│") + rightLine + borderColor.Render("│") + "\n"
	}

	// Language selector box (bottom left, integrated into border)
	var langEN, langFR string
	if m.lang == EN {
		langEN = titleStyle.Render(" EN ")
		langFR = mutedStyle.Render(" FR ")
	} else {
		langEN = mutedStyle.Render(" EN ")
		langFR = titleStyle.Render(" FR ")
	}

	// Lang selector parts (visual width = 11)
	// ╭────┬────╮ = 11 chars
	// ┤ EN │ FR ├ = 11 chars (┤ + 4 + │ + 4 + ├)
	// ╰────┴────╯ = 11 chars
	langSelectorTop := "╭────┬────╮"
	langSelectorBot := "╰────┴────╯"
	langWidth := 11

	// Position lang selector at bottom left (minimal padding)
	langLeftPad := 1
	langRightPadLeft := leftPanelWidth - langWidth - langLeftPad
	if langRightPadLeft < 0 {
		langRightPadLeft = 0
	}

	// Bottom line 1: │ + spaces + lang box top + spaces + │ + right spaces + │
	frameBottomLine1 := borderColor.Render("│") +
		repeatString(" ", langLeftPad) +
		borderColor.Render(langSelectorTop) +
		repeatString(" ", langRightPadLeft) +
		borderColor.Render("│") +
		repeatString(" ", rightPanelWidth) +
		borderColor.Render("│")

	// Bottom line 2: main bottom border with lang selector embedded
	// ╰ + dashes + ┤ + EN + │ + FR + ├ + dashes + ┴ + dashes + ╯
	frameBottomLine2 := borderColor.Render("╰"+repeatString("─", langLeftPad)+"┤") +
		langEN +
		borderColor.Render("│") +
		langFR +
		borderColor.Render("├"+repeatString("─", langRightPadLeft)+"┴"+repeatString("─", rightPanelWidth)+"╯")

	// Bottom line 3: lang selector bottom box on left + footer centered
	footer := lipgloss.NewStyle().
		Foreground(mutedColor).
		Render(m.t("footer"))

	// Calculate total frame width
	totalFrameWidth := leftPanelWidth + rightPanelWidth + 3 // +3 for │ characters

	// Lang selector bottom takes: langLeftPad + 1 (for initial space) + langWidth
	langSelectorEnd := langLeftPad + 1 + langWidth

	// Center footer in the remaining space (or full width)
	footerWidth := lipgloss.Width(footer)
	footerStart := (totalFrameWidth - footerWidth) / 2

	// If footer would overlap with lang selector, place it after
	if footerStart < langSelectorEnd + 2 {
		footerStart = langSelectorEnd + 2
	}

	// Build line 3
	frameBottomLine3 := repeatString(" ", langLeftPad+1) +
		borderColor.Render(langSelectorBot) +
		repeatString(" ", footerStart-langSelectorEnd) +
		footer +
		repeatString(" ", totalFrameWidth-footerStart-footerWidth)

	frameBottom := frameBottomLine1 + "\n" + frameBottomLine2 + "\n" + frameBottomLine3

	// Combine
	fullView := frameTop + "\n" + framedContent + frameBottom

	// Center horizontally
	return lipgloss.PlaceHorizontal(width, lipgloss.Center, fullView)
}

func repeatString(s string, count int) string {
	if count <= 0 {
		return ""
	}
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

func splitLines(s string) []string {
	var lines []string
	current := ""
	for _, r := range s {
		if r == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(r)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
