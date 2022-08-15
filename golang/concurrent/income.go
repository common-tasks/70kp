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
	var mut sync.Mutex
	fmt.Printf("initial balance is %d", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "salary", Amount: 500},
		{Source: "FD", Amount: 5},
		{Source: "CC", Amount: 5},
	}
	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				mut.Lock()
				bankBalance = bankBalance + income.Amount
				mut.Unlock()
				fmt.Printf(" your income in week : %d is %d from %s ", week, bankBalance, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("final amount is %d",bankBalance)
}
