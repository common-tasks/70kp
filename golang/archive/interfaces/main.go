package main

import "fmt"
type bot interface {
	getGreeting() string
}
type englishBot struct {}
type hindiBot struct{}

func main_prev() {
	fmt.Println("about interfaces")
	var eb englishBot
	printGreeting(eb)
	var hb hindiBot
	printGreeting(hb)
	

}

func (eb englishBot)getGreeting() string {
	return "hello there"
}

func(hb hindiBot) getGreeting() string {
	return "kaise ho bhai"
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}
