package main

import "fmt"

//import "fmt"

// In Go (Golang), fmt stands for Format.

// It is a standard Go package used for formatted I/O (input/output).

// // Common functions:

// fmt.Println() → prints with a newline
// fmt.Print() → prints without automatically adding a newline
// fmt.Printf() → prints using formatting
// fmt.Sprintf() → formats a string and returns it

func main() {
	//simple values
	fmt.Println(1 + 1)
	//string
	fmt.Println("Hello golang")
	//bool
	fmt.Println(true)
	fmt.Println(false)
	//floats
	fmt.Println(10.5)
	fmt.Println(7.0 / 3.0)
}
