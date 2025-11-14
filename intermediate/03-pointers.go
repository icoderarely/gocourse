package intermediate

import "fmt"

// Pointer is a variable that stores the memory address of another variable
// allows: modify value of a variable indirectly, pass large data structures efficiently between functions, manage memory directly for performance reasons

// pointer operations are limited to referencing(&) and referencing(*)
// there are no pointer arithmetic in golang like other languages

func main() {
	// declaration
	var ptr *int // ptr is pointer variable pointing to an integer variable
	var a int = 10
	ptr = &a
	fmt.Println(a, ptr, *ptr)

	modify(ptr)
	fmt.Println(a)
}

func modify(ptr *int) {
	*ptr++
}
