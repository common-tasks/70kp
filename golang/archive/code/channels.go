package main

import "fmt"

var initial int = 100

func substract(s []int, c chan int) {

	for _, v := range s {
		initial -= v
	}
	c <- initial // send initial to c
}

func channelsGo() {

	fmt.Println("about channels")

	s := []int{10, 10, 10, 10, 10, 10}

	c := make(chan int)
	go substract(s[:len(s)], c)
	go substract(s[:len(s)], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
