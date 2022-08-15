package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}
	for {
		var receive string = <-channel
		go checkLink(receive,channel)
	}
	
}
func checkLink(link string, chn chan string) {
	// fmt.Println(" check link function is called ")
	time.Sleep(time.Minute)
	_, err := http.Get(link)
	if err != nil {
		msg := link + " is down "
		fmt.Println(msg)
		chn <- msg
	} else {
		msg := link + " is up "
		fmt.Println(msg)
		chn <- msg
	}
}
