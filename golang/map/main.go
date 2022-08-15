package main

import (
	"fmt"
	
)

func main() {

	fmt.Println("------------about maps---------------")

	map1 := map[string]string{
		"red":"1",
		"blue":"2",
	}
	fmt.Print(map1)
	fmt.Println()
	var map2 map[string]string= make(map[string]string)
	map2["key1"]="value1"
	map2["key2"]="value2"
	map2["key3"]="value3"
	fmt.Print(map2)

	// delete(map2,"key3")
	fmt.Println()
	// fmt.Print(map2)
	// fmt.Scanf("f")
	printMap(map2)

	fmt.Println("------------that was all about maps---------------")
}

func printMap(list map[string]string){
	fmt.Println("printMap function")
	for key, value := range list {
		fmt.Print(key)
		fmt.Print(value)
		fmt.Println()
	}
}