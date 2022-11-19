package concepts

import "fmt"

func RunConcepts() {
	TestDefer()
	ft := convertTemp(32)
	fmt.Println("temp in fahrenheit: ", ft)
	testString()
	point()
}
