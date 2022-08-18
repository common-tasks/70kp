package main

import (
	"log"
	"os"
)

func composite() {

	slc := []string{"hello", "world", "I", "have", "arrived"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl3.gohtml", slc)

	if err != nil {
		log.Println(err)
	}
}
