package goroutines

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func AboutGoRoutines() {
	fmt.Printf("%s", "in main")
	var numOfCores = flag.Int("n",2,"number of CPU cores to use")
	flag.Parse()
	runtime.GOMAXPROCS(*numOfCores)
	
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9) // sleep works with a Duration in nanoseconds (ns) !
	fmt.Println("At the end of main()")

	wg:=new(sync.WaitGroup)
	wg.Add(2)
	go heavyMethod(wg)
	go extraheavyMethod(wg)
	wg.Wait()
	fmt.Println("all done!")


}
func longWait() {
	fmt.Println("start long wait")
	time.Sleep(5 * 1e9)
	fmt.Println("end of long wait")
}
func shortWait() {
	fmt.Println("start short wait")
	time.Sleep(2 * 1e9)
	fmt.Println("end short wait")
}

func heavyMethod(wg *sync.WaitGroup) {
	defer wg.Done()
}

func extraheavyMethod(wg *sync.WaitGroup) {
	defer wg.Done()
}
