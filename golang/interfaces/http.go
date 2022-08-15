package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
type logWriter struct{}
func main() {
	resp, err := http.Get("http://www.google.com")

	if(err!=nil){
		fmt.Println("error happened " ,err)
		os.Exit(1)
	}
	var lw logWriter
    io.Copy(lw,resp.Body)


}

func (lw logWriter)Write(data []byte)(int,error){
	fmt.Println(string(data))
	fmt.Println("number of bytes",len(data))
	return len(data),nil
}