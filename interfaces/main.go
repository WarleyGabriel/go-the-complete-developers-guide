package main

import "fmt"

type bot interface {
	getGreeting(int) string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting(1))
}

func (eb englishBot) getGreeting(i int) string {
	return fmt.Sprintf("Hello there! %d", i)
}

func (sb spanishBot) getGreeting(i int) string {
	return fmt.Sprintf("Hola! %d", i)
}
