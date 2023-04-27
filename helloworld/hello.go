package main

import (
	"fmt"
)

const (
	swedish            = "Swedish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	swedishHelloPrefix = "Hej "
	frenchHelloPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello(
		"world", "English",
	))
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case swedish:
		prefix = swedishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(input, lang string) string {
	if input == "" {
		input = "World"
	}

	return greetingPrefix(lang) + input
}
