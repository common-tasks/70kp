package main

import (
	"fmt"
)

func main() {
	callBasics()
}

func callBasics() {
	var greet string = "hello world!"
	fmt.Print(greet)
	// fmt.Println([]byte(greet))

	cards := newDeck()
	// cards.printDeck()
	cards.shuffle()
	cards.printDeck()

	// cards.printDeck()

	// deal, remainingCards := deal(cards, 3)

	// deal.printDeck()
	// remainingCards.printDeck()

	// cd := newDeck();
	// // fmt.Println(cd.toString())
	// var filename string = "test.txt"
	// cd.writeToFile(filename)

	// content:=cd.readFromFile(filename)
	// fmt.Println("file read content ")
	// content.printDeck()



}
