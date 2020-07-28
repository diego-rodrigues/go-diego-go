package mapper

import (
	"github.com/rylans/getlang"
)

func Greet(s string) string {
	info := getlang.FromString(s)

	switch info.LanguageCode() {
	case "en":
		return "Hello!"
	case "de":
		return "Guten Tag"
	case "pt-PT":
		return "Ola"
	default:
		return "I don't know that language yet"
	}
}
