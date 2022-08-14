package main

import (
	"fmt"
	"sync"
	// "time"
)

func main() {
var wg sync.WaitGroup;
words := []string { "zero","one","two","three","four","five","six","seven","eight","nine",}
wg.Add(len(words))

for i, word := range words {
	var data string = fmt.Sprintf("%d:%s",i,word)
	go print(data,&wg)
}

// time.Sleep(time.Second)

wg.Wait()

}

func print(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print(s)
	fmt.Print(" ")
}