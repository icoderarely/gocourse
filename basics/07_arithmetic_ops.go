package basics

import (
	"fmt"
	"math"
)

func main() {
	// Basic Arithmetic Operations
	// +, -, *, /, %
	// division between two integers is always in integer, result is truncated to zero.
	// to fix -> atleast one of the operands should be a floating point number => 22 / 7.0 or 22.0 / 7

	// Operator Precedence
	// () -> *, /, % -> +, -

	// Overflow
	// when the result of a computation exceeds the maximum value the data type can store
	// value wrapping to min value for signed integers and unexpected behaviours for unsigned integers
	fmt.Println("Overflow")
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)
	fmt.Println(maxInt + 1) // result: -9223372036854775808

	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)
	fmt.Println(uMaxInt + 1) // result: 0

	// Underflow
	// when the result is smaller than minimum value than a data type can hold
	fmt.Println("\nUndeflow")
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}

// why be mindful with overflow and underflow
// - program stability
// - data integrity
// - type safety

// mitigation stratergies
// - range checking
// - type conversion
// - error handling
