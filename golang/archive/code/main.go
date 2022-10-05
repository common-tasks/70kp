package main

//go application consists of packages and main package is the executable
/*
Go programs are organized into packages.
A package is a collection of source files in the same directory that are compiled together
*/

import (
	"fmt"
	"greetings"
	"log"
)

// fmt package has string formatting functions

func main() { // main function is the entry point of the application

	printHelloWorld() // simple hello world

	Variables() // all about variables

	runbasicTypes() // about types

	callExternalPackage() // call external package

	//loop
	forloop()

	//pointers
	pointers()

	//structs
	aboutStructs()

	//arrays
	aboutArrays()

	//slice
	aboutslice()

	//maps
	aboutMaps()

	//about methods
	aboutMethods()

	//go routines
	runGo()

	// about channels
	channelsGo()

	//defer

	defer fmt.Println("i am deferred")

	fmt.Println("let me complete first")

	// fmt.Scanln() // wait for user input
}

func callExternalPackage() {

	var greetingmsg string = greetings.Hello("bhai")
	fmt.Printf("%s\n", greetingmsg)

	// with error handling

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello01("friend")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	// Request a greeting message.
	msg, erro := greetings.HelloRandom("brother")
	// If an error was returned, print it to the console and
	// exit the program.
	if erro != nil {
		log.Fatal(erro)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(msg)

	// A slice of names.
	names := []string{"cat", "dog", "spider"}
	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

}
