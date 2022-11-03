package goroutines

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func Semaphore() {
	ch := make(chan string)
	go basicCompute(ch)

	doSomethingElse()
	var result string = <-ch
	fmt.Println("result of basicCompute is " + result)
	var arr []int
	var sortMax int=1000000000
	for i := 0; i < sortMax; i++ {
		var x int = rand.Intn(5000000-1) + 1
		arr = append(arr, x)
	}

	// display(arr)
	start:=time.Now()
	i:=len(arr)/2
	done:=make(chan bool)
	
	go sortSlice(arr[:i],done)
	go sortSlice(arr[i+1:],done)
	<-done
	<-done
	end:=time.Since(start)
	// display(arr)
	fmt.Printf("time taken in sort %f",end.Minutes())
}

// func display(a []int) {
// 	fmt.Println("in display function")
// 	for i := 0; i < len(a); i++ {
// 		fmt.Print(a[i])
// 		fmt.Print(" ")
// 	}
// }
func sortSlice(s []int,c chan bool) {
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	c<-true

}

func basicCompute(c chan string) {
	var sum int = 0
	start := time.Now()
	fmt.Println("start time " + start.GoString())
	for i := 0; i < 1000; i++ {
		sum = sum + i
	}
	elapsed := time.Since(start).Seconds()
	fmt.Println("time elapsed")
	fmt.Println(elapsed)

	c <- strconv.Itoa(sum)
}
func doSomethingElse() {
	var arr []int
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		arr = append(arr, i)
	}
	end := time.Since(start).Seconds()
	fmt.Printf("time taken in doSomethingElse %f of array length %d\n", end, len(arr))

}
