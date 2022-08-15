package main

import (
	"fmt"
	"strconv"
)

/*

About variables:

* variable declaration
* redeclaration and shadowing
* visibility
* Naming conventions
* Type conversions

*/

//declaring variables at package level
var b int = 1
var level int = 10

// UpperCaseVariable have global scope
var UpperCaseVariable int = 101

// Variables demo
func Variables() {
	fmt.Printf("%s\n", "about variables")
	blockLevelVariables()
	variablesConversion()
}

// block level variables
func blockLevelVariables() {

	fmt.Printf("%s", "block level variables declaration and assignment")

	var i int = 99 // variable i of type int

	j := 100 // compiler figures out what data type it is

	var k float32 // declare now
	k = 101.1     // assign a value in next statement

	var level int = 200 // shadowing of variable declared at package level

	var Somevar int = 10

	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v %T\n", k, k)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", level, level)
	fmt.Printf("%v\n", Somevar)

}

func variablesConversion() {

	fmt.Printf("variables conversion\n")

	var i int = 101

	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)

	fmt.Printf("%v, %T\n", j, j)

	//string conversion

	var a int = 5

	var str string = string(a)

	fmt.Printf("%v , %T\n", str, str) // this will just print the unicode character of 5

	str = strconv.Itoa(a)

	fmt.Printf("%v , %T\n", str, str) // this is actual string conversion

}