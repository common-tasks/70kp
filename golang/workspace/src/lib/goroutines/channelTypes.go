package goroutines

import (
	"fmt"
	"time"
)

func BlockingChannel() {
	ch := make(chan string)
	go sendMsg(ch)

	go receiveMsg(ch)
	time.Sleep(time.Second)
	close(ch)

	select {}

}

func sendMsg(ch chan string) {
	fmt.Println("sending msg on the channel")
	ch <- "sending msg on the channel"
	fmt.Println("sent msg on the channel")
}
func receiveMsg(ch chan string) {
	var msg string = <-ch
	fmt.Println("received msg on the channel")
	fmt.Println(msg)

}
