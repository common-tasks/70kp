package main

import "fmt"

func main() {
	var x int =44
	fmt.Printf("%d %p\n",x,&x)
	var intPtr *int = &x;
	*intPtr++;
	fmt.Printf("%v\n",*intPtr)
	fmt.Printf("%v\n",x)

}