package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	readText()
	fmt.Println("tcp server listening")

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(connection)
	}
}

func readText() {
	s := "hello there \n I am here \n just a voicemail away \n set an appointment to meet me any time \n and I will cancel that"
	scanner := bufio.NewScanner(strings.NewReader(s))
	// scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func handle(connection net.Conn) {
	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		content := scanner.Text()
		fmt.Println(content)
	}
	defer connection.Close()

	fmt.Println("done")
}
