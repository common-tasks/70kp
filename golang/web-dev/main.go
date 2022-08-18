package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl2.gohtml", "tpl3.gohtml","tpl4.gohtml"))
}
func main() {
	printSomething("hello", "world")

	singleVariable()

	composite()

	passMap()
}

func singleVariable() {
	err := tpl.Execute(os.Stdout, "tridev")
	if err != nil {
		log.Fatalln(err)
	}
}

func printSomething(s ...string) {
	for _, str := range s {
		fmt.Println(str)

	}

}
