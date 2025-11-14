package intermediate

import "fmt"

// Recursion is a programming technique where a function solves a problem by calling itself with a smaller or simpler input. Each recursive call brings the problem closer to a base case, which stops further calls. It is useful for naturally repetitive or divide-and-conquer problems like trees, graphs, and mathematical sequences.

func main() {
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	// base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// recursive case: run till you reach base case
	fmt.Println(n)
	return n * factorial(n-1)
}
