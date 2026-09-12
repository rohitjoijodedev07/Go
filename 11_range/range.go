// package range
package main

//import "fmt"

// In Go (Golang), fmt stands for Format.

// It is a standard Go package used for formatted I/O (input/output).

// // Common functions:

// fmt.Println() → prints with a newline
// fmt.Print() → prints without automatically adding a newline
// fmt.Printf() → prints using formatting
// fmt.Sprintf() → formats a string and returns it

//iterating over data structures
func main() {
	nums := []int{6, 7, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum := 0
	// for _, num := range nums {
	// 	sum = sum + num
	// }

	// fmt.Println(sum)

	// sum := 0
	// for i, num := range nums {
	// 	fmt.Println(num,i)
	// }

	//m := map[string]string{"fname": "john", "lname": "doe"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	//only for keys

	// for k, v := range m {
	// 	fmt.Println(k)
	// }

	//c means //unicode code point rune
	//i means // starting byte of rune
	//255 -> 1 byte , 2 byte
	// for i, c := range "golang" {
	// 	fmt.Println(i, string(c))
	// }

}
