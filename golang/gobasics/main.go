package main

import (
	"fmt"
	"runtime"
)

func main() {
	//switch statement

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
