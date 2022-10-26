package main

func SumOfArray(arr []int) int {

	var sum int
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}

func LinearSearch(arr []int, number int) bool {
	var isFound bool = false
	for i := 0; i < len(arr); i++ {
		if arr[i] == number {
			isFound = true
			break
		}
	}
	return isFound
}
func BinarySearch(sarray []int, item int) bool {
	var isFound bool = false
	length := len(sarray)
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		if item == sarray[mid] {
			return true
		} else if item < sarray[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return isFound
}

// []int{1, -2, 3, 4, -4, 6, -14, 6, 2}
func LargestSumSubArray(array []int) int {

	var largestSum = 0
	var currentLargest = 0
	for i := 0; i < len(array); i++ {
		currentLargest = currentLargest + array[i]
		if currentLargest < 0 {
			currentLargest = 0
		}
		if currentLargest > largestSum {
			largestSum = currentLargest
		}

	}

	return largestSum
}
