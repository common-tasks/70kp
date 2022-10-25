package main

import "fmt"

func main() {
	// newFunction()
	var arr []int
	arr = append(arr, 2, 4, 6, 8)
	result := SumOfArray(arr)
	fmt.Printf("sum of input array is  %d ", result)
	arr2:=[]int{1,2,3,4}
	fmt.Printf("sum of input array2 is %d",SumOfArray(arr2))

}

func newFunction() {
	var x int = 44
	fmt.Printf("%d %p\n", x, &x)
	var intPtr *int = &x
	*intPtr++
	fmt.Printf("%v\n", *intPtr)
	fmt.Printf("%v\n", x)
}
