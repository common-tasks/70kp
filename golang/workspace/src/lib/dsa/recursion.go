package dsa

import "fmt"

func RecursionMain() {
	fact := Factorial(6)
	fmt.Println("factorial is ", fact)
}
func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
func PrintBase16(){

}
