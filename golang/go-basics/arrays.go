package main


func SumOfArray(arr []int) int {

	var sum int
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}
