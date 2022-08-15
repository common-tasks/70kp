package main

import (
	"fmt"
	"sync"
	// "time"
)
var msg string
var wgmsg sync.WaitGroup

func main() {
	raceConditionFunc()

	// time.Sleep(time.Second)
	waitGroupFunc()

}

func waitGroupFunc() {
	var wg sync.WaitGroup
	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	wg.Add(len(words))

	for i, word := range words {
		var data string = fmt.Sprintf("%d:%s", i, word)
		go print(data, &wg)
	}

	wg.Wait()
}

func raceConditionFunc() {
	msg = "hello world"
	var mt sync.Mutex

	wgmsg.Add(2)
	go updateMessage("hola",&mt,&wgmsg)
	go updateMessage("halo",&mt,&wgmsg)
	wgmsg.Wait()
	fmt.Println(msg)
}
func updateMessage(s string,mt *sync.Mutex,wg *sync.WaitGroup){
	defer wg.Done()
	mt.Lock()
	msg=s
	mt.Unlock()
}	

func print(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print(s)
	fmt.Print(" ")
}
