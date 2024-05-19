package main

import "fmt"

const englishGreeting = "Hello, "
const frenchGreeting = "Bonjour, "
const spanishGreeting = "Hola, "

const french = "French"
const spanish = "Spanish"

func Hello(s, language string) string {
	if s == "" {
		s = "世界"
	}

	greeting := getGreeting(language)

	return greeting + s
}

func getGreeting(language string) (greeting string) {
	switch language {
	case spanish:
		greeting = spanishGreeting
	case french:
		greeting = frenchGreeting
	default:
		greeting = englishGreeting
	}
	return greeting
}

func main() {
	fmt.Println(Hello("世界", ""))
}
