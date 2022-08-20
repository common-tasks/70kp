package main

import (
	"fmt"
	"net/http"
)

type myhttp int

func (m myhttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var msg string = " my http server"
	fmt.Println(msg)
	fmt.Fprintln(w,msg)
	
}
func main() {
	
	var mhttp myhttp

	http.ListenAndServe("localhost:8080", mhttp)

}
