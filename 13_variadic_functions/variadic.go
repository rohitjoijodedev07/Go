package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, nums := range nums {
		total = total + nums
	}

	return total

}

//interface any
// func sum1(nums ...interface{}) int {
// 	total := 0

// 	for _, nums := range nums {
// 		total = total + nums
// 	}

// 	return total

// }

func main() {
	fmt.Println(1, 2, 3, 5, 56, "hello")

	result := sum(3, 3, 4, 5, 5)

	fmt.Println(result)
	/////another way to call veriadic function in go
	nums := []int{3, 4, 5, 6}

	result1 := sum(nums...)

	fmt.Println(result1)
}
