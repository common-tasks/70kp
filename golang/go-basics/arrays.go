package main


func SumOfArray(arr []int) int {

	var sum int
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}

func LinearSearch(arr []int,number int) bool {
	var isFound bool =false
	for i := 0; i < len(arr); i++ {
		if(arr[i]==number){
			isFound=true
			break
		}
	}
	return isFound
}
