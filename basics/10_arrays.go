package main

import "fmt"

func main() {
	// syntax: var arrayName [size]elementType
	// default values of int -> 0
	var nums [5]int
	fmt.Println(nums)

	nums[1] = 12
	fmt.Println(nums)

	alpha := [4]string{"a", "b", "c", "d"}
	fmt.Println("Fruits:", alpha)
	fmt.Println(alpha[2])

	// arrays are value type
	// passing arrays to function or assigning it to a variable creates a copy of the array

	// iterating array
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println("index:", i, "element:", nums[i])
	// }

	for index, value := range nums {
		fmt.Println("index:", index, "element:", value)
	}

	// for _, value := range nums -> `_` means we are discarding the value - here it is the index value
}
