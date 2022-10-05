package main

import (
	"fmt"
	"math"
)

func forloop() {
	fmt.Println("about for loops")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func infiniteloop() {
	for {
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
