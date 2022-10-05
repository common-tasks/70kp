package main

/*
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

*/

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func (v vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func abs(v vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func scale(v *vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func aboutMethods() {
	fmt.Println("about methods")
	v := vertex{3, 4}
	fmt.Println(v.Abs())

	v1 := vertex{3, 4}
	v1.Scale(10)
	fmt.Println(v1.Abs())
}
