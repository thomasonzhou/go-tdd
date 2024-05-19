package main

import "fmt"

const englishGreeting = "Hello, "
const frenchGreeting = "Bonjour, "

const french = "French"

func Hello(s, language string) string {
	if s == "" {
		s = "世界"
	}
	if language == french {
		return frenchGreeting + s
	}
	return englishGreeting + s
}

func main() {
	fmt.Println(Hello("世界", ""))
}
