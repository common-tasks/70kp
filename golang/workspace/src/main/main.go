package main

import (
	"fmt"
	"go-workspace/src/lib/dsa"
	"go-workspace/src/lib/goroutines"
	"go-workspace/src/lib/concepts"
)

func main() {
	fmt.Println("main method entry")
	callMethods()

}

func callMethods() {
	dsa.DSAProblems()

	goroutines.RunGoRoutines()
	
	concepts.RunConcepts()
}
