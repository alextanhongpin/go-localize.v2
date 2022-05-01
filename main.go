package main

import (
	"github.com/alextanhongpin/go-localize/db"
	_ "github.com/alextanhongpin/go-localize/internal/translations"
	"github.com/alextanhongpin/go-localize/internal/translations/messages"
)

func main() {
	// Split them so that 'make generate' doesn't target both folders.
	messages.Run()

	// Example of localization loaded from db.
	db.Run()
}
