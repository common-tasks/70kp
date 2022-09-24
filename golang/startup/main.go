package main

import 
(
	"fmt"
	"rsc.io/quote"
	
)

func main(){
	fmt.Println("hello world")
	fmt.Println(quote.Go())
}

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}