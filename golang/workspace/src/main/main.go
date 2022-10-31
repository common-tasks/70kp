package main

import
(
	"go-workspace/src/lib/dsa"
	"fmt"
	"go-workspace/src/lib/goroutines"
)

func main(){
	fmt.Println("main method entry")
	callMethods()
	
}

func callMethods() {
	// rotateArrayProblem()
	// arrayWaveFormProblem()
	goroutines.AboutGoRoutines()
}

func rotateArrayProblem() {
	array := []int{1, 2, 3, 4, 5, 6}
	dsa.RotateArrayByK(array, 2)
}

func arrayWaveFormProblem() {
	array := []int{ 8, 1, 2, 3, 4, 5, 6, 4, 2 }
	dsa.ArrayWaveForm(array)

}