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

	greeting := englishGreeting
	switch language {
	case spanish:
		greeting = spanishGreeting
	case french:
		greeting = frenchGreeting
	}

	return greeting + s
}

func main() {
	fmt.Println(Hello("世界", ""))
}
