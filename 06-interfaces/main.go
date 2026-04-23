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

	printGreetingEB(eb)
	printGreetingSB(sb)
}

func (englishBot) getGreeting() string {
	// Has different logic englishBot
	return "Hello World!"
}

func (spanishBot) getGreeting() string {
	// Has different logic for spanishBot
	return "Hola!"
}

func printGreetingEB(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingSB(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
