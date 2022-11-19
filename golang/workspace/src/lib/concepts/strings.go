package concepts

import (
	"fmt"
	"strings"
)

func testString() {
	var test string = "this is a sample string"
	r1 := strings.HasSuffix(test,"ring")
	fmt.Println("has suffix",r1)
	r2:=strings.HasPrefix(test,"thi")
	fmt.Println(r2)
	r2= strings.Contains(test,"is a")
	fmt.Println(r2)


}
