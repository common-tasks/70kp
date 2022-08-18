package main

import (
	"log"
	"os"
)

func passMap() {

	map01 := map[string]string{
		"name01": "james",
		"name02": "topno",
		"name03": "benjamin",
		"name04": "khalko",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl4.gohtml", map01)
	if err != nil {
		log.Fatalln(err)
	}

}
