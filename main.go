package main

import (
	"fmt"
	"os"
	"time"

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

// Screen enum
type Screen int

const (
	WelcomeScreen   Screen = iota
	PortfolioScreen
)

// Animation phase enum
type AnimPhase int

const (
	PhaseInitialBlink AnimPhase = iota // cursor blinks before anything starts
	PhaseTypingDomain
	PhaseBlinkFirst
	PhaseCursorHome  // cursor jumps to start, pause before typing prefix
	PhasePrepend
	PhasePauseAfterPrepend // pause before cursor jumps to end
	PhaseBlinkSecond
	PhaseCursorEnd   // cursor jumps to end, pause before deleting
	PhaseDelete
	PhaseRestart
)

// Animation timing
const (
	tickInterval  = 80 * time.Millisecond
	blinkInterval = 10  // ticks between blink toggles (~480ms)
	blinkDuration = 62 // ticks for 5s of blinking
	deleteChars   = 1  // characters deleted per tick (faster than typing)
	restartPause  = 8  // ticks to pause before restarting
	handMovePause    = 10 // ticks to pause after cursor jump (~800ms for hand repositioning)
	initialBlinkTime = 12 // ticks for initial cursor blink (~1s)
)

// Animation text constants
const (
	domainText  = "termfolio.dev"
	prependText = "Mohamed.Gacha@"
)

// Tick message for animation
type tickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

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
		"about_1":        "Engineer passionate about building",
		"about_2":        "robust, meaningful systems — from scalable",
		"about_3":        "backend APIs to cybersecurity solutions.",
		"about_4":        "Driven by curiosity and impact, with a focus",
		"about_5":        "on system security and software quality.",
		"about_6":        "Also into F1, gaming, and astrophysics.",
		"exp_title":      "Experience",
		"exp_scroll":     "(↑↓ to scroll)",
		"mission":        "Mission:",
		"stack":          "Stack:",
		"challenge":      "Challenge:",
		"feedback":       "Feedback:",
		"gw_role":        "Paris | Oct 2025 - Sept 2026",
		"gw_mission_1":   "Developing and maintaining GCap, Gatewatcher's NDR",
		"gw_mission_2":   "(Network Detection & Response) solution, while working",
		"gw_mission_3":   "on open-source security projects in parallel.",
		"gw_mission_4":   "Managing cloud infrastructure (Proxmox, Nutanix, Azure,",
		"gw_mission_5":   "AWS) and building internal tooling for the team.",
		"gw_challenge":   "Ramped up on cybersecurity from the ground up.",
		"gw_feedback_1":  "Thriving in a fast-paced environment with a strong",
		"gw_feedback_2":  "engineering culture and thoughtful use of AI.",
		"gw_title":       "Cybersecurity Engineer",
		"etifak_role":    "Remote | Sept 2024 - Sept 2025",
		"etifak_title":   "Backend Developer",
		"etifak_mission_1": "Led backend development for B2B marketplace",
		"etifak_mission_2": "at an early-stage startup. Built REST APIs, unit tests,",
		"etifak_mission_3": "and managed AWS deployment.",
		"etifak_challenge": "Self-taught many concepts as a junior dev.",
		"etifak_feedback":  "Great team communication made workflow smooth!",
		"suez_role":      "Paris | Jun - Aug 2024",
		"suez_title":     "Data Engineer",
		"suez_mission_1": "Internship: Built intranet web app for data quality & cleaning.",
		"suez_mission_2": "Implemented deduplication using Levenshtein distance.",
		"suez_feedback":  "First internship. Great Scrum team dynamics!",
		"proj_title":     "Projects",
		"sls_desc_1":     "Community-driven platform providing free car setups",
		"sls_desc_2":     "for Assetto Corsa Competizione. Full-stack project",
		"sls_desc_3":     "with REST API, database management, and modern frontend.",
		"sls_status":     "Status: Live & actively maintained",
		"mealpass_desc_1": "Mobile app managing food distributions for students",
		"mealpass_desc_2": "in need. QR-code based check-in system with real-time",
		"mealpass_desc_3": "tracking and admin dashboard.",
		"mealpass_status": "Status: Completed — volunteer project for an association",
		"termfolio_desc_1": "Terminal-based portfolio built with Go and Bubble Tea.",
		"termfolio_desc_2": "Interactive TUI with bilingual support, SSH access,",
		"termfolio_desc_3": "and a custom-rendered UI with box-drawing characters.",
		"termfolio_status":    "Status: Live & actively maintained — ",
		"termfolio_highlight": "you're looking at it",
		"vectorart_desc_1": "Illustrator-like vector drawing application with basic",
		"vectorart_desc_2": "shape tools, layers, and export features.",
		"vectorart_status": "Status: Completed — 1st year project at ENSISA",
		"discordclone_desc_1": "Real-time messaging platform inspired by Discord",
		"discordclone_desc_2": "with channels, user auth, and live chat.",
		"discordclone_status": "Status: Completed — 1st year project at ENSISA",
		"schoolmgmt_desc_1": "Management platform for students, courses, and grades",
		"schoolmgmt_desc_2": "with admin dashboard and role-based access.",
		"schoolmgmt_status": "Status: Completed — 1st year project at ENSISA",
		"skills_title":   "Skills",
		"contact_title":  "Contact",
		"contact_reach":  "Feel free to reach out!",
		"contact_open":   "Open to opportunities in:",
		"contact_backend": "Backend Development",
		"contact_cyber":  "Cybersecurity",
		"contact_devops": "DevOps",
		"edu_title":      "Education",
		"ensisa_degree":  "Engineering Degree: Computer Science & Networks",
		"ensisa_school":  "ENSISA - National School of Engineers of South Alsace",
		"ensisa_period":  "Sept 2023 - Sept 2026",
		"ensisa_loc":     "Mulhouse, France",
		"ensisa_desc_1":  "Software development, databases, DevOps,",
		"ensisa_desc_2":  "networks & security, AI and deep learning.",
		"ensisa_desc_3":  "Hands-on projects in full-stack, cloud, and cybersecurity.",
		"cpge_degree":    "Preparatory Classes MPSI",
		"cpge_school":    "Carnot Prépas",
		"cpge_period":    "Sept 2021 - Jul 2023",
		"cpge_loc":       "Meknes, Morocco",
		"cpge_desc_1":    "Intensive program in Mathematics, Physics and",
		"cpge_desc_2":    "Engineering Sciences. Developed strong analytical",
		"cpge_desc_3":    "thinking and problem-solving skills.",
		"bac_degree":     "Baccalaureate - Science Math A",
		"bac_school":     "Lycée Ajana",
		"bac_period":     "2021",
		"bac_loc":        "Meknes, Morocco",
		"bac_desc_1":     "French Option. Strong foundation in mathematics",
		"bac_desc_2":     "and scientific reasoning.",
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
		"about_1":        "Ingénieur passionné par la conception de",
		"about_2":        "systèmes fiables et utiles — des APIs backend",
		"about_3":        "scalables aux solutions de cybersécurité.",
		"about_4":        "Animé par la curiosité et l'impact, avec un",
		"about_5":        "intérêt pour la sécurité et la qualité logicielle.",
		"about_6":        "Aussi passionné de F1, gaming et astrophysique.",
		"exp_title":      "Expériences",
		"exp_scroll":     "(↑↓ pour défiler)",
		"mission":        "Mission:",
		"stack":          "Stack:",
		"challenge":      "Défi:",
		"feedback":       "Avis:",
		"gw_role":        "Paris | Oct 2025 - Sept 2026",
		"gw_mission_1":   "Développement et maintien de GCap, la solution NDR",
		"gw_mission_2":   "(Network Detection & Response) de Gatewatcher, en",
		"gw_mission_3":   "parallèle de projets open-source de sécurité.",
		"gw_mission_4":   "Gestion d'infrastructure cloud (Proxmox, Nutanix,",
		"gw_mission_5":   "Azure, AWS) et création d'outillage interne.",
		"gw_challenge":   "Montée en compétence cybersécurité en partant de zéro.",
		"gw_feedback_1":  "Environnement stimulant avec une forte culture",
		"gw_feedback_2":  "d'ingénierie et une utilisation réfléchie de l'IA.",
		"gw_title":       "Ingénieur Cybersécurité",
		"etifak_role":    "Remote | Sept 2024 - Sept 2025",
		"etifak_title":   "Développeur Backend",
		"etifak_mission_1": "Développement backend pour marketplace B2B",
		"etifak_mission_2": "dans une startup. APIs REST, tests unitaires,",
		"etifak_mission_3": "et déploiement AWS.",
		"etifak_challenge": "Auto-apprentissage en tant que dev junior.",
		"etifak_feedback":  "Excellente communication d'équipe!",
		"suez_role":      "Paris | Juin - Août 2024",
		"suez_title":     "Data Engineer",
		"suez_mission_1": "Stage : Application intranet pour qualité des données.",
		"suez_mission_2": "Déduplication avec distance de Levenshtein.",
		"suez_feedback":  "Premier stage. Super dynamique Scrum!",
		"proj_title":     "Projets",
		"sls_desc_1":     "Plateforme communautaire proposant des setups gratuits",
		"sls_desc_2":     "pour Assetto Corsa Competizione. Projet full-stack avec",
		"sls_desc_3":     "API REST, gestion de base de données et frontend moderne.",
		"sls_status":     "Statut : En ligne & maintenu activement",
		"mealpass_desc_1": "Application mobile de gestion de distributions",
		"mealpass_desc_2": "alimentaires pour étudiants précaires. Système de",
		"mealpass_desc_3": "check-in par QR-code avec suivi en temps réel.",
		"mealpass_status": "Statut : Terminé — projet bénévole pour une association",
		"termfolio_desc_1": "Portfolio terminal construit avec Go et Bubble Tea.",
		"termfolio_desc_2": "TUI interactive avec support bilingue, accès SSH,",
		"termfolio_desc_3": "et rendu personnalisé avec des caractères box-drawing.",
		"termfolio_status":    "Statut : En ligne & maintenu — ",
		"termfolio_highlight": "vous le consultez en ce moment",
		"vectorart_desc_1": "Application de dessin vectoriel type Illustrator avec",
		"vectorart_desc_2": "outils de formes, calques et export.",
		"vectorart_status": "Statut : Terminé — projet 1ère année ENSISA",
		"discordclone_desc_1": "Plateforme de messagerie en temps réel inspirée de",
		"discordclone_desc_2": "Discord avec salons, authentification et chat en direct.",
		"discordclone_status": "Statut : Terminé — projet 1ère année ENSISA",
		"schoolmgmt_desc_1": "Plateforme de gestion d'étudiants, cours et notes",
		"schoolmgmt_desc_2": "avec tableau de bord admin et accès par rôles.",
		"schoolmgmt_status": "Statut : Terminé — projet 1ère année ENSISA",
		"skills_title":   "Compétences",
		"contact_title":  "Contact",
		"contact_reach":  "N'hésitez pas à me contacter!",
		"contact_open":   "Ouvert aux opportunités en:",
		"contact_backend": "Développement Backend",
		"contact_cyber":  "Cybersécurité",
		"contact_devops": "DevOps",
		"edu_title":      "Parcours_Académique",
		"ensisa_degree":  "Diplôme d'ingénieur : Informatique et Réseaux",
		"ensisa_school":  "ENSISA - École Nationale Supérieure d'Ingénieurs Sud-Alsace",
		"ensisa_period":  "Sept 2023 - Sept 2026",
		"ensisa_loc":     "Mulhouse, France",
		"ensisa_desc_1":  "Développement logiciel, bases de données, DevOps,",
		"ensisa_desc_2":  "réseaux et sécurité, IA et deep learning.",
		"ensisa_desc_3":  "Projets pratiques en full-stack, cloud et cybersécurité.",
		"cpge_degree":    "Classes Préparatoires MPSI",
		"cpge_school":    "Carnot Prépas",
		"cpge_period":    "Sept 2021 - Juil 2023",
		"cpge_loc":       "Meknès, Maroc",
		"cpge_desc_1":    "Programme intensif en Mathématiques, Physique et",
		"cpge_desc_2":    "Sciences de l'Ingénieur. Développement de la rigueur",
		"cpge_desc_3":    "analytique et des capacités de résolution de problèmes.",
		"bac_degree":     "Baccalauréat - Option Science Math A",
		"bac_school":     "Lycée Ajana",
		"bac_period":     "2021",
		"bac_loc":        "Meknès, Maroc",
		"bac_desc_1":     "Option Française. Solide formation en mathématiques",
		"bac_desc_2":     "et raisonnement scientifique.",
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
	scroll int
	lang   Lang

	// Welcome screen animation
	screen    Screen
	animPhase AnimPhase
	animPos   int
	animTicks int
	cursorOn  bool
}

// Tracks max scroll from the last render frame
var lastMaxScroll int

func initialModel() model {
	return model{
		width:    80,
		height:   24,
		cursor:   0,
		lang:     EN,
		cursorOn: true,
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
	return doTick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.screen {
		case WelcomeScreen:
			switch msg.String() {
			case "ctrl+c":
				return m, tea.Quit
			default:
				m.screen = PortfolioScreen
				return m, nil
			}
		case PortfolioScreen:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "left", "h":
				if m.cursor > 0 {
					m.cursor--
					m.scroll = 0
				}
			case "right", "l":
				if m.cursor < 3 {
					m.cursor++
					m.scroll = 0
				}
			case "up", "k":
				if m.scroll > 0 {
					m.scroll--
				}
			case "down", "j":
				if m.scroll < lastMaxScroll {
					m.scroll++
				}
			case "tab":
				if m.lang == EN {
					m.lang = FR
				} else {
					m.lang = EN
				}
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tickMsg:
		if m.screen == WelcomeScreen {
			advanceAnimation(&m)
			return m, doTick()
		}
	}

	return m, nil
}

func advanceAnimation(m *model) {
	m.animTicks++

	switch m.animPhase {
	case PhaseInitialBlink:
		if m.animTicks%blinkInterval == 0 {
			m.cursorOn = !m.cursorOn
		}
		if m.animTicks >= initialBlinkTime {
			m.animPhase = PhaseTypingDomain
			m.animTicks = 0
			m.animPos = 0
			m.cursorOn = true
		}

	case PhaseTypingDomain:
		if m.animPos < len(domainText) {
			m.animPos++
			m.cursorOn = true
		} else {
			m.animPhase = PhaseBlinkFirst
			m.animTicks = 0
			m.cursorOn = true
		}

	case PhaseBlinkFirst:
		if m.animTicks%blinkInterval == 0 {
			m.cursorOn = !m.cursorOn
		}
		if m.animTicks >= blinkDuration {
			m.animPhase = PhaseCursorHome
			m.animTicks = 0
			m.animPos = 0 // cursor jumps to start
			m.cursorOn = true
		}

	case PhaseCursorHome:
		// Pause after Home key — hands moving from Alt+Arrow to letter keys
		if m.animTicks >= handMovePause {
			m.animPhase = PhasePrepend
			m.animTicks = 0
			m.animPos = 0
			m.cursorOn = true
		}

	case PhasePrepend:
		if m.animPos < len(prependText) {
			m.animPos++
			m.cursorOn = true
		} else {
			m.animPhase = PhasePauseAfterPrepend
			m.animTicks = 0
			m.cursorOn = true
		}

	case PhasePauseAfterPrepend:
		// Pause after typing prefix — before cursor jumps to end
		if m.animTicks >= handMovePause {
			m.animPhase = PhaseBlinkSecond
			m.animTicks = 0
			m.cursorOn = true
		}

	case PhaseBlinkSecond:
		if m.animTicks%blinkInterval == 0 {
			m.cursorOn = !m.cursorOn
		}
		if m.animTicks >= blinkDuration {
			m.animPhase = PhaseCursorEnd
			m.animTicks = 0
			m.animPos = len(prependText) + len(domainText) // cursor jumps to end
			m.cursorOn = true
		}

	case PhaseCursorEnd:
		// Pause after End key — hands moving from Alt+Arrow to Backspace
		if m.animTicks >= handMovePause {
			m.animPhase = PhaseDelete
			m.animTicks = 0
			m.cursorOn = true
		}

	case PhaseDelete:
		if m.animPos > 0 {
			m.animPos -= deleteChars
			if m.animPos < 0 {
				m.animPos = 0
			}
		} else {
			m.animPhase = PhaseRestart
			m.animTicks = 0
		}

	case PhaseRestart:
		if m.animTicks >= restartPause {
			m.animPhase = PhaseTypingDomain
			m.animTicks = 0
			m.animPos = 0
			m.cursorOn = true
		}
	}
}

func (m model) View() string {
	if m.screen == WelcomeScreen {
		return m.welcomeView()
	}

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
	menuItems := []string{"Experience", m.t("edu_title"), "Projects", "Skills"}
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
	leftPanelWidth := innerWidth*30/100 + 1
	rightPanelWidth := innerWidth - leftPanelWidth - 1 // -1 for the middle │

	// Calculate navbar position (top right of overall view)
	navbarLen := lipgloss.Width(navbar)
	navbarRightPad := 1 // minimal padding from right edge
	navbarPadding := rightPanelWidth - navbarLen - 2 - navbarRightPad // -2 for ┤ and ├
	if navbarPadding < 0 {
		navbarPadding = 0
	}

	// Helper for clickable links (OSC 8 hyperlinks)
	link := func(url, text string) string {
		return fmt.Sprintf("\x1b]8;;%s\x07%s\x1b]8;;\x07", url, text)
	}

	// Site text box (top left)
	siteText := " " + link("https://termfolio.dev", "termfolio.dev") + " "
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
			titleStyle.Render("━━━ Gatewatcher ✕ "+m.t("gw_title")+" ━━━"),
			contentStyle.Render(m.t("gw_role")),
			mutedStyle.Render(m.t("mission"))+" "+contentStyle.Render(m.t("gw_mission_1")),
			contentStyle.Render(m.t("gw_mission_2")),
			contentStyle.Render(m.t("gw_mission_3")),
			contentStyle.Render(m.t("gw_mission_4")),
			contentStyle.Render(m.t("gw_mission_5")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Python, Ansible, Docker, Bash, Linux, CI/CD"),
			mutedStyle.Render(m.t("challenge"))+" "+contentStyle.Render(m.t("gw_challenge")),
			mutedStyle.Render(m.t("feedback"))+" "+contentStyle.Render(m.t("gw_feedback_1")),
			contentStyle.Render(m.t("gw_feedback_2")),
			"",
			titleStyle.Render("━━━ Etifak ✕ "+m.t("etifak_title")+" ━━━"),
			contentStyle.Render(m.t("etifak_role")),
			mutedStyle.Render(m.t("mission"))+" "+contentStyle.Render(m.t("etifak_mission_1")),
			contentStyle.Render(m.t("etifak_mission_2")),
			contentStyle.Render(m.t("etifak_mission_3")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Python, Django, PostgreSQL, Docker, AWS EC2"),
			mutedStyle.Render(m.t("challenge"))+" "+contentStyle.Render(m.t("etifak_challenge")),
			mutedStyle.Render(m.t("feedback"))+" "+contentStyle.Render(m.t("etifak_feedback")),
			"",
			titleStyle.Render("━━━ Suez Digital Solutions ✕ "+m.t("suez_title")+" ━━━"),
			contentStyle.Render(m.t("suez_role")),
			mutedStyle.Render(m.t("mission"))+" "+contentStyle.Render(m.t("suez_mission_1")),
			contentStyle.Render(m.t("suez_mission_2")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Python, Streamlit, Pandas, asyncio, Azure DevOps"),
			mutedStyle.Render(m.t("feedback"))+" "+contentStyle.Render(m.t("suez_feedback")),
		)

	case 1: // Education
		rightPanelContent = lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(m.t("edu_title"))+" "+mutedStyle.Render(m.t("exp_scroll")),
			"",
			titleStyle.Render("━━━ "+m.t("ensisa_degree")+" ━━━"),
			contentStyle.Render(m.t("ensisa_school")),
			mutedStyle.Render(m.t("ensisa_period")+" | "+m.t("ensisa_loc")),
			contentStyle.Render(m.t("ensisa_desc_1")),
			contentStyle.Render(m.t("ensisa_desc_2")),
			contentStyle.Render(m.t("ensisa_desc_3")),
			"",
			titleStyle.Render("━━━ "+m.t("cpge_degree")+" ━━━"),
			contentStyle.Render(m.t("cpge_school")),
			mutedStyle.Render(m.t("cpge_period")+" | "+m.t("cpge_loc")),
			contentStyle.Render(m.t("cpge_desc_1")),
			contentStyle.Render(m.t("cpge_desc_2")),
			contentStyle.Render(m.t("cpge_desc_3")),
			"",
			titleStyle.Render("━━━ "+m.t("bac_degree")+" ━━━"),
			contentStyle.Render(m.t("bac_school")),
			mutedStyle.Render(m.t("bac_period")+" | "+m.t("bac_loc")),
			contentStyle.Render(m.t("bac_desc_1")),
			contentStyle.Render(m.t("bac_desc_2")),
		)

	case 2: // Projects
		rightPanelContent = lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(m.t("proj_title"))+" "+mutedStyle.Render(m.t("exp_scroll")),
			"",
			titleStyle.Render("━━━ "+link("https://simplylovelysetups.com", "SimplyLovelySetups.com")+" ━━━"),
			mutedStyle.Render("Sept 2025 - Present"),
			contentStyle.Render(m.t("sls_desc_1")),
			contentStyle.Render(m.t("sls_desc_2")),
			contentStyle.Render(m.t("sls_desc_3")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("FastAPI, SQLAlchemy, MongoDB, NextJS, Docker"),
			mutedStyle.Render(m.t("sls_status")),
			"",
			titleStyle.Render("━━━ "+link("https://termfolio.dev", "termfolio.dev")+" ━━━"),
			mutedStyle.Render("Feb 2026 - Present"),
			contentStyle.Render(m.t("termfolio_desc_1")),
			contentStyle.Render(m.t("termfolio_desc_2")),
			contentStyle.Render(m.t("termfolio_desc_3")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Go, Bubble Tea, Lipgloss, SSH (Wish)"),
			mutedStyle.Render(m.t("termfolio_status"))+lipgloss.NewStyle().Background(primaryColor).Foreground(lipgloss.Color("#000000")).Bold(true).Render(" "+m.t("termfolio_highlight")+" "),
			"",
			titleStyle.Render("━━━ MealPass ━━━"),
			mutedStyle.Render("Mars 2025"),
			contentStyle.Render(m.t("mealpass_desc_1")),
			contentStyle.Render(m.t("mealpass_desc_2")),
			contentStyle.Render(m.t("mealpass_desc_3")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Django, PostgreSQL, Flutter, Dart, Docker"),
			mutedStyle.Render(m.t("mealpass_status")),
			"",
			titleStyle.Render("━━━ NeoShape ━━━"),
			mutedStyle.Render("2024"),
			contentStyle.Render(m.t("vectorart_desc_1")),
			contentStyle.Render(m.t("vectorart_desc_2")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Java, JavaFX"),
			mutedStyle.Render(m.t("vectorart_status")),
			"",
			titleStyle.Render("━━━ Nexus ━━━"),
			mutedStyle.Render("2024"),
			contentStyle.Render(m.t("discordclone_desc_1")),
			contentStyle.Render(m.t("discordclone_desc_2")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Python, Django, WebSocket, HTML/CSS"),
			mutedStyle.Render(m.t("discordclone_status")),
			"",
			titleStyle.Render("━━━ NewMoodle ━━━"),
			mutedStyle.Render("2024"),
			contentStyle.Render(m.t("schoolmgmt_desc_1")),
			contentStyle.Render(m.t("schoolmgmt_desc_2")),
			mutedStyle.Render(m.t("stack"))+" "+contentStyle.Render("Python, Django, PostgreSQL, HTML/CSS"),
			mutedStyle.Render(m.t("schoolmgmt_status")),
		)

	case 3: // Skills
		rightPanelContent = lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(m.t("skills_title"))+" "+mutedStyle.Render(m.t("exp_scroll")),
			"",
			titleStyle.Render("Languages")+"  "+contentStyle.Render("Python, Go, Rust, TypeScript, JavaScript, Java, C, F#"),
			titleStyle.Render("Backend")+"    "+contentStyle.Render("Django, DRF, FastAPI, PostgreSQL, MongoDB, SQLAlchemy"),
			titleStyle.Render("Frontend")+"   "+contentStyle.Render("React.js, NextJS, Flutter, HTML/CSS"),
			titleStyle.Render("DevOps")+"     "+contentStyle.Render("Docker, AWS, Azure, CI/CD, Linux, Ansible, Bash"),
			titleStyle.Render("Security")+"   "+contentStyle.Render("Suricata, Network Analysis, Selenium, Beautiful Soup"),
			titleStyle.Render("Data")+"       "+contentStyle.Render("Pandas, Streamlit, Azure ML"),
			titleStyle.Render("Tools")+"      "+contentStyle.Render("Git, Scrum/Agile, Code Review, REST APIs"),
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

	// Clamp scroll to valid range
	maxScroll := len(rightLines) - contentHeight
	if maxScroll < 0 {
		maxScroll = 0
	}
	lastMaxScroll = maxScroll
	if m.scroll > maxScroll {
		m.scroll = maxScroll
	}

	var framedContent string
	for i := 0; i < contentHeight; i++ {
		leftLine := ""
		rightLine := ""
		if i < len(leftLines) {
			leftLine = leftLines[i]
		}
		rightIdx := i + m.scroll
		if rightIdx < len(rightLines) {
			rightLine = rightLines[rightIdx]
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

// cursorBlock renders a character with inverted colors to simulate a terminal block cursor.
func cursorBlock(ch string) string {
	return lipgloss.NewStyle().
		Background(primaryColor).
		Foreground(lipgloss.Color("#000000")).
		Bold(true).
		Render(ch)
}

func (m model) welcomeView() string {
	width := m.width
	height := m.height
	if width < minWidth {
		width = minWidth
	}
	if height < minHeight {
		height = minHeight
	}

	// Cursor character
	cursor := " "
	if m.cursorOn {
		cursor = "█"
	}

	// Style the text
	animStyle := lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true)

	// Build display with cursor at the correct insertion point
	fullText := prependText + domainText
	var line string

	switch m.animPhase {
	case PhaseInitialBlink:
		line = animStyle.Render(cursor)
	case PhaseTypingDomain:
		line = animStyle.Render(domainText[:m.animPos] + cursor)
	case PhaseBlinkFirst:
		line = animStyle.Render(domainText + cursor)
	case PhaseCursorHome:
		if m.cursorOn {
			line = cursorBlock("t") + animStyle.Render(domainText[1:]+" ")
		} else {
			line = animStyle.Render(domainText + " ")
		}
	case PhasePrepend:
		if m.cursorOn {
			line = animStyle.Render(prependText[:m.animPos]) + cursorBlock("t") + animStyle.Render(domainText[1:]+" ")
		} else {
			line = animStyle.Render(prependText[:m.animPos] + domainText + " ")
		}
	case PhasePauseAfterPrepend:
		if m.cursorOn {
			line = animStyle.Render(prependText) + cursorBlock("t") + animStyle.Render(domainText[1:]+" ")
		} else {
			line = animStyle.Render(fullText + " ")
		}
	case PhaseBlinkSecond:
		line = animStyle.Render(fullText + cursor)
	case PhaseCursorEnd:
		// Cursor jumped to end, waiting before deleting
		line = animStyle.Render(fullText + cursor)
	case PhaseDelete:
		line = animStyle.Render(fullText[:m.animPos] + cursor)
	case PhaseRestart:
		line = animStyle.Render(cursor)
	}

	hint := mutedStyle.Render("Press any key to continue")

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		line,
		"",
		hint,
	)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
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
