package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"math/rand"
)

// create a new type of desk which has slice of strings

type deck []string
var currentCards deck
//receiver function
func (cards deck) printDeck() {
	for i, card := range cards {
		fmt.Println(i,card)

	}

}

func subDeck(max int) deck {
 return currentCards[0:max]
}

func deal (d deck,max int) (deck,deck) {
	return d[0:max],d[max:]
}
// creates new deck

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"spades", "diamonds", "hearts", "club"}

	cardValues := []string{"ace", "two", "three", "four"}

	for _, suit := range cardSuits {

		for _, cardvalue := range cardValues {
			cards = append(cards, suit+" and "+cardvalue)
		}
	}
	currentCards = cards
	return cards
}

func (d deck)toString() string{

	var strSlice []string = []string(d)

	var finalString string = strings.Join(strSlice,",");
	
	return finalString
	
}

func (d deck)writeToFile(fileName string) error {	
	var content string = d.toString()
	var data [] byte = []byte(content)
	err := os.WriteFile(fileName, data, fs.ModeAppend)

    if err != nil {
        fmt.Println(err)
    }
	return err
}

func (d deck)readFromFile(filename string) deck{
    var content string
	var result deck
	data ,err:=os.ReadFile(filename)
	
	if(err != nil){
		fmt.Println(err)
		os.Exit(1)
	}

	content = string(data)
	dataslice :=strings.Split(content, ",")
	result = dataslice;

	return result

}

func (d deck) shuffle()  {

for i  := range d {
	newPosition:= rand.Intn(len(d)-1)
	d[i], d[newPosition] = d[newPosition],d[i]

}}
