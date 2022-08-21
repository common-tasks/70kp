package main

import (
	"fmt"
	"runtime"
)

func main() {

	// switchStatement()
	// deferStatement()
	// aboutArrays()
	aboutSlices()
}
func aboutSlices() {
	sl := make([]string, 0)
	sl = append(sl, "hello")
	sl = append(sl, "world")
	sl = append(sl, "I am here")
	sl = append(sl, "finally")
	ln := len(sl)
	cp := cap(sl)
	fmt.Println(sl)
	fmt.Println(ln, cp)

}
func aboutArrays() {
	var a [6]string
	a[0] = "hello"
	a[1] = "world"
	a[2] = "call me anytime"
	a[3] = " and I will not pickup the call"
	fmt.Println(a)
}

func deferStatement() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

// switch statement
func switchStatement() {
	switch os := runtime.GOOS; os {
	case "windows":
		{
			fmt.Println("windows operating system")
		}
	case "linux":
		{
			fmt.Println("linux")
		}
	default:
		{
			fmt.Println(os)
		}

	}
}
