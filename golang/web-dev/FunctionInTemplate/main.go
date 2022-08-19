package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Klass struct {
	StudentCount int
	ClassName    string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"gt": greet,
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func greet(s string) string {
	s = "hello " + s
	return s
}
func main() {
	var school []Klass = make([]Klass, 5)
	m1 := Klass{StudentCount: 7, ClassName: "montessori1"}
	m2 := Klass{StudentCount: 5, ClassName: "montessori2"}
	m3 := Klass{StudentCount: 10, ClassName: "montessori3"}
	school = append(school, m1)
	school = append(school, m2)
	school = append(school, m3)

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", school)
	if err != nil {
		log.Fatalln(err)
	}

}
