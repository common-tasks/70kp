package goroutines

import (
	"fmt"
	"strconv"
	"time"
)

func Semaphore() {
	ch := make(chan string)
	go basicCompute(ch)
	var result string = <-ch
	fmt.Println("result of basicCompute is " + result)
}

func basicCompute(c chan string) {
	var sum int = 0
	start:=time.Now()
	fmt.Println("start time " + start.GoString())
	for i := 0; i < 100000; i++ {
		sum = sum + i
	}
	elapsed:=time.Since(start)
	fmt.Printf("time taken to do basic compute %d seconds \n",elapsed/1000)
	c <-strconv.Itoa(sum)
}