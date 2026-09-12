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
	var name string = "golang"
	var name1 = "golang"
	var age int32 = 30

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(name1)

	//shorthand syntax
	name3 := "name3"
	fmt.Println(name3)

	var myname string

	myname = "rohit joijode"

	fmt.Println(myname)

	var price float32 = 50.00

	var price1 = 509.33

	fmt.Println(price1)
	fmt.Println(price)
}
