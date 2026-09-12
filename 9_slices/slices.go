package main

//import "fmt"

// In Go (Golang), fmt stands for Format.

// It is a standard Go package used for formatted I/O (input/output).

// // Common functions:

// fmt.Println() → prints with a newline
// fmt.Print() → prints without automatically adding a newline
// fmt.Printf() → prints using formatting
// fmt.Sprintf() → formats a string and returns it

//slice -> dynamic
//most used construct in go
// + useful methods

func main() {
	// uninitialized slice is nil

	// var nums []int

	// fmt.Println(nums)
	// fmt.Println(len(nums))

	// var nums = make([]int, 2, 5)
	// // var nums = make([]int, 0, 5) set 0 not append 2 index 0

	// //capacity --> maximum numbers of elements can fit
	// fmt.Println(cap(nums))
	// // fmt.Println(nums)
	// // fmt.Println(nums == nil)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// nums = append(nums, 8)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	//second way
	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// var nums = make([]int, 2, 5)
	// nums[0] = 3
	// nums[1] = 5
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	//copy function
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// nums = append(nums, 2)
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	///slice operator

	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:2]) //example from and to index
	// fmt.Println(nums[:1])

	// slices
	//var nums1 = []int{1, 2}
	//var nums2 = []int{1, 2}

	//fmt.Println(slices.Equal(nums1, nums2)) //this line compare indexs //result output true or false

	//2D slices also we created
	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)

}
