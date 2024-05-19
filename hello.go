package main

import "fmt"

const englishGreeting = "Hello"

func Hello(s string) string {
	if s == "" {
		s = "世界"
	}
	return englishGreeting + " " + s
}

func main() {
	fmt.Println(Hello("世界"))
}
