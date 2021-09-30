package main

import "fmt"

type bot interface {
	getGreeting() string
}
type koreanBot struct{}
type japaneseBot struct{}

func main() {
	kb := koreanBot{}
	jb := japaneseBot{}

	printGreeting(kb)
	printGreeting(jb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (koreanBot) getGreeting() string {
	return "안녕하세요"
}

func (japaneseBot) getGreeting() string {
	return "こんにちは"
}
