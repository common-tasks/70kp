package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("tcp server listening")
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		fmt.Println("tcp server called")
		io.WriteString(connection, "\n Hello from tcp server")
		fmt.Fprintln(connection, "\n how is it going")
		fmt.Fprintf(connection, "%v", "\n well I guess")

		connection.Close()
	}
}
