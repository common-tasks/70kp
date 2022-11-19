package concepts

import "fmt"

func point() {
	var x int = 5
	var y = &x
	fmt.Printf("address of %d is %p",x,y)
}