package main

import "fmt"

func main() {
	// newFunction()
	var arr []int
	arr = append(arr, 2, 4, 6, 8)
	result := SumOfArray(arr)
	fmt.Printf("sum of input array is  %d\n ", result)
	arr2 := []int{1, 2, 3, 4}
	fmt.Printf("sum of input array2 is %d\n", SumOfArray(arr2))

	arr3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("number found? %t\n", LinearSearch(arr3, 11))

	sortedarray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	element:=20
	found := BinarySearch(sortedarray, element)
	fmt.Printf("element %d found in BinarySearch ? %t\n", element,found)

}

func newFunction() {
	var x int = 44
	fmt.Printf("%d %p\n", x, &x)
	var intPtr *int = &x
	*intPtr++
	fmt.Printf("%v\n", *intPtr)
	fmt.Printf("%v\n", x)
}
