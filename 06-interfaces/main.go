package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// Has different logic englishBot
	return "Hello World!"
}

func (spanishBot) getGreeting() string {
	// Has different logic for spanishBot
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
