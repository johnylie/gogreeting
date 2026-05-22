package gogreeting

import "time"

type Language string

const (
	EN Language = "en"
	ID Language = "id"
)

type Translation struct {
	Morning   string
	Afternoon string
	Evening   string
	Night     string
}

var defaultTranslations = map[Language]Translation{
	EN: {
		Morning:   "Good morning",
		Afternoon: "Good afternoon",
		Evening:   "Good evening",
		Night:     "Good night",
	},

	ID: {
		Morning:   "Selamat pagi",
		Afternoon: "Selamat siang",
		Evening:   "Selamat sore",
		Night:     "Selamat malam",
	},
}

func Greeting(
	name string,
	lang Language,
	custom *Translation,
) string {
	hour := time.Now().Hour()

	translations := defaultTranslations[lang]

	if custom != nil {
		if custom.Morning != "" {
			translations.Morning = custom.Morning
		}

		if custom.Afternoon != "" {
			translations.Afternoon = custom.Afternoon
		}

		if custom.Evening != "" {
			translations.Evening = custom.Evening
		}

		if custom.Night != "" {
			translations.Night = custom.Night
		}
	}

	var greeting string

	switch {
	case hour >= 5 && hour < 11:
		greeting = translations.Morning

	case hour >= 11 && hour < 15:
		greeting = translations.Afternoon

	case hour >= 15 && hour < 18:
		greeting = translations.Evening

	default:
		greeting = translations.Night
	}

	return greeting + ", " + name
}
