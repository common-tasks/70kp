package dsa

import (
	"fmt"
)

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
		minIndex:=i
		if(a[i-1]<a[minIndex]){
			minIndex=i-1
		}
		if(a[i+1]<a[minIndex]){
			minIndex=i+1
		}
		if(minIndex!=a[i]){
			temp:=a[minIndex]
			a[minIndex]=a[i]
			a[i]=temp
		}


	}

	fmt.Println("wave form output")
	for _, v := range a {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println(" ")
}
