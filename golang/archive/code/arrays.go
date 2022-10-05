package main

import (
	"fmt"
)

func aboutArrays(){
	println("about arrays")
	var a [2]string
	a[0] = "array element 1"
	a[1] = "array element 2"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func aboutslice(){
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}