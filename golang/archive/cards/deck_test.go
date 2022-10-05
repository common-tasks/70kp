package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	fmt.Println("testing new deck feature")
	newCard := newDeck()

	var length int = len(newCard)
	if length != 16 {
		t.Error("length is not 16 but ", length)
	}

}
