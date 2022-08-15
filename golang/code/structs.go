package main

import (
	"fmt"
)

//Vertex structure
type Vertex struct {
	X int
	Y int
}

func aboutStructs() {
	fmt.Println("about structure")
	fmt.Println(Vertex{1, 2})
}