package main

import (
	"fmt"
)

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
