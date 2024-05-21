package main

import (
	"fmt"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func GreetingPeople(language string) string {
	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	}

	return prefix
}
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return GreetingPeople(language) + name
}

func main() {
	fmt.Println("HI")
}
