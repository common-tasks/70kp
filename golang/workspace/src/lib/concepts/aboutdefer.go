package concepts

import (
	"errors"
	"fmt"
)

func TestDefer() {
	func01()
}

func func01() {
	fmt.Println("func01 start")
	defer func02()

	fmt.Println("func01 end")
	_ = errors.New("my ownn error")
	fmt.Println("end statement in func01")
}
func func02() {
	fmt.Println("func02 start")
	fmt.Println("func02 end")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
