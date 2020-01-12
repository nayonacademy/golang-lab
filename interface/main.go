package main

import "fmt"

type englishBot struct{}
type spnashBot struct{}

type bot interface {
	getGreeting() string
}

func (englishBot) getGreeting() string {
	return "hello english"
}

func (spnashBot) getGreeting() string {
	return "hello spanish"
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func main() {
	english := englishBot{}
	spanish := spnashBot{}
	printGreeting(english)
	printGreeting(spanish)
}
