package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {

	var bankBalance int
	fmt.Printf("initial balance is %d", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "salary", Amount: 500},
		{Source: "FD", Amount: 10},
		{Source: "CC", Amount: 5},
	}

	for i, income := range incomes {
		go func(i int, income Income) {

		}(i, income)
	}
}
