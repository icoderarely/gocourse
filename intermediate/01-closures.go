package intermediate

// Closures are functions that capture and remember variables from the scope in which they were created,
// even after that scope has ended. They allow functions to maintain state between calls without using global variables.
// Closures are commonly used for callbacks, data hiding, and creating function factories.

import "fmt"

func main() {
	sequence := adder()
	// currVal := sequence()
	// fmt.Println(currVal)
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())
	fmt.Println(sequence2())

	fmt.Println("---------------")

	subractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()
	// subractor: is the inner function that is returned
	// tailing `()` intermediately calls the outer function, thus returning the inside one to subractor: Immediately Invoked Function Expression (IIFE)
	// Outer function → creates the countdown variable and returns a function.
	// Inner function → uses countdown and performs subtraction.
	// subractor → refers to the inner function, the one that actually subtracts.

	fmt.Println(subractor(5))
	fmt.Println(subractor(10))
	fmt.Println(subractor(12))
}

func adder() func() int {
	i := 0
	fmt.Println("prev value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i", i)
		return i
	}
}
