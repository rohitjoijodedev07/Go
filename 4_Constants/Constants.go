package main

import (
	"fmt"
)

//import "fmt"

// In Go (Golang), fmt stands for Format.

// It is a standard Go package used for formatted I/O (input/output).

// // Common functions:

// fmt.Println() → prints with a newline
// fmt.Print() → prints without automatically adding a newline
// fmt.Printf() → prints using formatting
// fmt.Sprintf() → formats a string and returns it

// Globle declaration
const age1 = 30

func main() {
	//local declaration
	const name = "go const values"
	const age = 30
	fmt.Println("")

	//multiple const
	const (
		port = 5000
		host = "localhost"
	)

	//not allowing this
	//port = 5500

	fmt.Println(port, host)
}
