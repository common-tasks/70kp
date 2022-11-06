package dsa

import (
	"fmt"
)

func ArrayMain() {

	// arr := []int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}
	// IndexArray(arr)
	result:=SmallestPositiveMissingNumber()
	fmt.Printf("smallest positive missing number is %d",result)

}
func SmallestPositiveMissingNumber() int {
	a := []int{8, 5, 6, 1, 9, 11, 2, 7, 3, 10}
	//sort the array
	for i := 0; i < len(a); i++ {
		for j := i+1; j < len(a); j++ {
			var temp int
			if(a[i]>a[j]){
				temp=a[i]
				a[i]=a[j]
				a[j]=temp

			}	
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i])
		fmt.Print(" ")
	}
	small:=a[0]
	for i := 1; i < len(a); i++ {
		small=small+1
		if(small!=a[i]){
			return small
		}
	}
	return -1
}
func IndexArray(arr []int) {
	ln := len(arr)
	for i := 0; i < ln; i++ {
		for j := i + 1; j < (ln - (i + 1)); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[j] = arr[i]
				arr[i] = temp
			}
		}
	}
	for i := 0; i < ln; i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
}

// for example: []int{1, 2, 3, 4, 5, 6}
func RotateArrayByK(arr []int, k int) {
	fmt.Printf("%s\n", "rotate an element by K positions")

	fmt.Printf("%s\n", "array elements before rotation")

	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
	for i := 0; i < len(arr)-k; i++ {
		var temp int = arr[i+k]
		arr[i+k] = arr[i]
		arr[i] = temp
	}

	fmt.Printf("%s", "array elements after rotation")
	fmt.Println("")
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println(" ")

}

// for example: []int{ 8, 1, 2, 3, 4, 5, 6, 4, 2 }
func ArrayWaveForm(a []int) {
	fmt.Println("wave form input")
	for _, v := range a {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println(" ")

	for i := 1; i < len(a); i = i + 2 {
		minIndex := i
		if a[i-1] < a[minIndex] {
			minIndex = i - 1
		}
		if a[i+1] < a[minIndex] {
			minIndex = i + 1
		}
		if minIndex != a[i] {
			temp := a[minIndex]
			a[minIndex] = a[i]
			a[i] = temp
		}

	}

	fmt.Println("wave form output")
	for _, v := range a {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println(" ")
}
