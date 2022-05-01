package messages

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Run() {
	var lang language.Tag
	//lang = language.MustParse("en-US")
	//lang = language.MustParse("en-GB")
	lang = language.MustParse("en-MY")
	//lang = language.MustParse("en-ID")

	p := message.NewPrinter(lang)
	p.Printf("hello %s\n", "John")
	var booksCount int
	//booksCount = 1
	booksCount = 100
	p.Printf("There are %d books\n", booksCount)
	fmt.Println("this is great")
}
