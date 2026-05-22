package gogreeting

import "time"

func Selamat(name string) string {
	hour := time.Now().Hour()

	switch {
	case hour >= 5 && hour < 11:
		return "Selamat pagi, " + name
	case hour >= 11 && hour < 15:
		return "Selamat siang, " + name
	case hour >= 15 && hour < 18:
		return "Selamat sore, " + name
	default:
		return "Selamat malam, " + name
	}
}
