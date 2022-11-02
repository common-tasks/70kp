package goroutines

import (
	"fmt"
	"time"
)

func AboutChannels() {
	// ch := make(chan string)
	// go sendData(ch)
	// go getData(ch)
	// time.Sleep(1e9)
	// close(ch)

	ch2 := make(chan int,0)
	go pumpData(ch2)
	go serveData(ch2)
	time.Sleep(1e9)

	close(ch2)

}
func pumpData(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
func serveData(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch)
		fmt.Print(" ")
	}
}
func sendData(ch chan string) {
	ch <- "ranchi"
	ch <- "kokar"
	ch <- "kanke"

}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println(input)
	}

}
