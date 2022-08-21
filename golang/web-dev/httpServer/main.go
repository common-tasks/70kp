package main

import (
	"html/template"
	"log"
	"net/http"
)

type myserver int

func (ms myserver) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "request.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("request.gohtml"))
}

func main() {
	var ms myserver
	http.ListenAndServe(":8080", ms)
}