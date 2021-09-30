package main

import "fmt"

type koreanBot struct{}
type japaneseBot struct{}

func main() {
	//kb := koreanBot{}
	jb := japaneseBot{}

	//printGreeting(kb)
	printGreeting(jb)
}

// func printGreeting(kb koreanBot) {
// 	fmt.Println(kb.getGreeting())
// }

func printGreeting(jb japaneseBot) {
	fmt.Println(jb.getGreeting())
}

func (koreanBot) getGreeting() string {
	return "안녕하세요"
}

func (japaneseBot) getGreeting() string {
	return "こんにちは"
}
