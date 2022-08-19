package main

import (
	"log"
	"os"
)

type GreatGuy struct {
	Name string
	Age  int
}

func passSliceOfStruct() {
	var greatGuys []GreatGuy = make([]GreatGuy, 5)
	guy1 := GreatGuy{Name: "MKG", Age: 100}
	greatGuys = append(greatGuys, guy1)
	guy2 := GreatGuy{Name: "JN", Age: 101}
	greatGuys = append(greatGuys, guy2)
	guy3 := GreatGuy{Name: "SCB", Age: 102}
	greatGuys = append(greatGuys, guy3)
	guy4 := GreatGuy{Name: "BS", Age: 103}
	greatGuys = append(greatGuys, guy4)
	guy5 := GreatGuy{Name: "CA", Age: 104}
	greatGuys = append(greatGuys, guy5)

	err := tpl.ExecuteTemplate(os.Stdout, "structslice.gohtml", greatGuys)
	if err != nil {
		log.Fatalln(err)
	}

}
