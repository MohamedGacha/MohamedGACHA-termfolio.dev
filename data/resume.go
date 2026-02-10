package data

// Resume data for Mohamed GACHA's termfolio

type Contact struct {
	Name     string
	Title    string
	Email    string
	Phone    string
	Location string
	LinkedIn string
	GitHub   string
}

type Experience struct {
	Title    string
	Company  string
	Period   string
	Location string
	Desc     string
	Tags     []string
}

type Project struct {
	Name string
	Desc string
	Date string
	Tags []string
	URL  string
}

type Education struct {
	Degree string
	School string
	Period string
	Desc   string
}

type Resume struct {
	Contact    Contact
	Experiences []Experience
	Projects   []Project
	Education  []Education
	Skills     []string
	Languages  map[string]string
	Interests  []string
}

func GetResume() Resume {
	return Resume{
		Contact: Contact{
			Name:     "Mohamed GACHA",
			Title:    "Ingénieur en informatique et réseaux",
			Email:    "simogacha@gmail.com",
			Phone:    "+33 6 05 76 45 74",
			Location: "France",
			LinkedIn: "mohamed-gacha",
			GitHub:   "MohamedGacha",
		},
		Experiences: []Experience{
			{
				Title:    "Alternance développeur Solution Cybersécurité",
				Company:  "Gatewatcher",
				Period:   "Oct 2025 - Sept 2026",
				Location: "Paris, La Défense",
				Desc:     "Développement et maintien de GCap (solution NDR) en parallèle de projets open-source de sécurité. Gestion d'infrastructure cloud et création d'outillage interne.",
				Tags:     []string{"Python3", "Git", "GitHub", "Suricata", "Docker", "CI/CD", "Bash", "Linux"},
			},
			{
				Title:    "Alternance développeur Back-end",
				Company:  "Etifak",
				Period:   "Sept 2024 - Sept 2025",
				Location: "Remote",
				Desc:     "Développement et déploiement d'une REST API pour une plateforme E-procurement B2B.",
				Tags:     []string{"Python3", "Git", "GitHub", "Django", "PostgreSQL", "JWT", "Docker", "ReactJS", "Swagger", "AWS EC2"},
			},
			{
				Title:    "Data Engineer",
				Company:  "Suez Digital Solutions",
				Period:   "Juin 2024 - Août 2024",
				Location: "Paris, La Défense",
				Desc:     "Développement d'une application web d'analyse, déduplication et mise en qualité de données.",
				Tags:     []string{"Python3", "Git", "Azure DevOps", "Streamlit", "Pandas", "asyncio", "threading", "requests"},
			},
		},
		Projects: []Project{
			{
				Name: "SimplyLovelySetups.com",
				Desc: "Site web gratuit de setups pour Assetto Corsa Competizione",
				Date: "Sept 2025 - En cours",
				Tags: []string{"Python3", "FastAPI", "SQLAlchemy", "Alembic", "MongoDB", "NextJS", "Docker"},
				URL:  "https://simplyLovelySetups.com",
			},
			{
				Name: "MealPass",
				Desc: "Application Mobile de gestion de distributions alimentaires pour étudiants précaires",
				Date: "Mars 2025",
				Tags: []string{"Python3", "Django", "PostgreSQL", "Flutter", "Dart", "Docker"},
				URL:  "",
			},
		},
		Education: []Education{
			{
				Degree: "Diplôme d'ingénieur: Informatique et Réseaux",
				School: "ENSISA - École Nationale Supérieure d'Ingénieurs Sud-Alsace - Mulhouse",
				Period: "Sept 2023 - Sept 2026",
				Desc:   "Développement logiciel, bases de données, DevOps, réseaux et sécurité, IA et deep learning.",
			},
			{
				Degree: "Classes préparatoires MPSI",
				School: "Carnot prépas | Meknes, Maroc",
				Period: "Sept 2021 - Juil 2023",
				Desc:   "Mathématiques, Physique et Sciences de l'Ingénieur",
			},
			{
				Degree: "Baccalauréat - Option Science Math A",
				School: "Lycée Ajana | Meknès, Maroc",
				Period: "2021",
				Desc:   "Option Française",
			},
		},
		Skills: []string{
			"Python3", "Django", "FastAPI", "Alembic", "SQLAlchemy",
			"PostgreSQL", "MongoDB", "JavaScript", "HTML5", "CSS3",
			"ReactJS", "Docker", "AWS", "Azure", "Git",
			"Java", "C", "Linux", "Bash", "F#", "C#",
			"UML", "WebSocket", "DRF", "Streamlit", "Pandas",
		},
		Languages: map[string]string{
			"Français": "Courant",
			"Anglais":  "Courant",
			"Arabe":    "Natale",
			"Espagnol": "En cours d'apprentissage",
		},
		Interests: []string{
			"Formule 1",
			"Jeux vidéo",
			"Astrophysique",
			"Aéronautique",
		},
	}
}
