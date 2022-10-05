package main

import(
	"fmt"
)

func pointers(){
	fmt.Println("about pointers")
	i, j := 4, 100

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 8         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 10   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}