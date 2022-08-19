package main

import (
	"log"
	"os"
)

type games struct {
	Name    string
	Players int
}

func passStruct() bool {

	game := games{Name: "cricket", Players: 11}
	err := tpl.ExecuteTemplate(os.Stdout,"struct.gohtml",game)
	if(err!=nil){
		log.Fatalln(err)
	}
	return true
}
