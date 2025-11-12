package basics

import "fmt"

func init() {
	fmt.Println("1. initializing package...")
}

func init() {
	fmt.Println("2. initializing package...")
}

func main() {
	// used to perform initialization tasks
	// initializing variables, performin set of ops, initializing states, registering components or configurations, db init
	// happens before main()
	// once per package
	// sequentially if there are many init functions

	fmt.Println("inside main()")

	// best practices
	// avoid side effects, init order, documentaion
}
