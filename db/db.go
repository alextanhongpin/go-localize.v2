package db

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Run() {
	message.SetString(language.Dutch, "You have chosen to play %m.", "U heeft ervoor gekozen om %m te spelen.")
	message.SetString(language.Dutch, "basketball", "basketbal")
	message.SetString(language.Dutch, "hockey", "ijshockey")
	message.SetString(language.Dutch, "soccer", "voetbal")
	message.SetString(language.BritishEnglish, "soccer", "football")

	for _, sport := range []string{"soccer", "basketball", "hockey"} {
		for _, lang := range []string{"en", "en-GB", "nl"} {
			p := message.NewPrinter(language.Make(lang))
			fmt.Printf("%-6s %s\n", lang, p.Sprintf("You have chosen to play %m.", sport))
		}
		fmt.Println()
	}
}
