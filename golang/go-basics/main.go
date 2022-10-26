package main

import "fmt"

func main() {

	testArraysFunctions()

}

func testArraysFunctions() {
	var arr []int
	arr = append(arr, 2, 4, 6, 8)
	result := SumOfArray(arr)
	fmt.Printf("sum of input array is  %d\n ", result)
	arr2 := []int{1, 2, 3, 4}
	fmt.Printf("sum of input array2 is %d\n", SumOfArray(arr2))

	arr3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("number found? %t\n", LinearSearch(arr3, 11))

	sortedarray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	element := 20
	found := BinarySearch(sortedarray, element)
	fmt.Printf("element %d found in BinarySearch ? %t\n", element, found)
	arr2 = []int{1, -2, 3, 4, -4, 6, -14, 6, 2}
	largestSum := LargestSumSubArray(arr2)
	fmt.Printf("largest sum is %d\n", largestSum)
}
