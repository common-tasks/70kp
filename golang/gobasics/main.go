package main

import (
	"fmt"
	"runtime"
)

type Greetings interface {
	Greet(msg string) string
}
type Mathematice interface {
	Add(number1 int, number2 int) int
}
type SayHello struct {
	Hmsg string
}

func (hello *SayHello) Greet(msg string) string {
	hello.Hmsg = "Halo"
	return hello.Hmsg + " " + msg
}
func (hello SayHello) PrintHello() {
	fmt.Println(hello.Hmsg)
}

type Calc struct {
}

func GreetAll(g Greetings, msg string) string {
	return g.Greet(msg)
}
func Addition(m Mathematice, number1 int, number2 int) int {
	return m.Add(number1, number2)
}
func (cal Calc) Add(number1 int, number2 int) int {

	return number1 + number2
}
func main() {

	// switchStatement()
	// deferStatement()
	// aboutArrays()
	// aboutSlices()
	// aboutMaps()
	aboutMethods()
}
func aboutMethods() {
	var calculator Calc
	result := Addition(calculator, 1, 2)
	fmt.Println(result)
	var hello SayHello
	hello.Hmsg = "Hola"
	hello.PrintHello()
	res := GreetAll(&hello, hello.Hmsg)
	hello.PrintHello()
	fmt.Println(res)

}
func aboutMaps() {

	list := make(map[string]string)
	list["one"] = "hello"
	list["two"] = "world"
	list["three"] = "arrived"
	for _, value := range list {
		fmt.Print(value)
	}
	fmt.Println()
	delete(list, "two")
	for _, v := range list {
		fmt.Print(v)
	}
	fmt.Println()
	elem, ok := list["one"]
	fmt.Printf("%v is present ? %v", elem, ok)
}
func aboutSlices() {
	sl := make([]string, 0)
	sl = append(sl, "hello")
	sl = append(sl, "world")
	sl = append(sl, "I am here")
	sl = append(sl, "finally")
	ln := len(sl)
	cp := cap(sl)
	fmt.Println(sl)
	fmt.Println(ln, cp)

}
func aboutArrays() {
	var a [6]string
	a[0] = "hello"
	a[1] = "world"
	a[2] = "call me anytime"
	a[3] = " and I will not pickup the call"
	fmt.Println(a)
}

func deferStatement() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

// switch statement
func switchStatement() {
	switch os := runtime.GOOS; os {
	case "windows":
		{
			fmt.Println("windows operating system")
		}
	case "linux":
		{
			fmt.Println("linux")
		}
	default:
		{
			fmt.Println(os)
		}

	}
}
