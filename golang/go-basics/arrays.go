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
