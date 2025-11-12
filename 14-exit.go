package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Exit(): immediately terminates a program with a given status code
	// deferred functions wont be executed
	// 0 -> success; any other -> non success

	defer fmt.Println("defer func")

	fmt.Println("starting main()")

	// exit with status code
	os.Exit(1)

	// this will never be executed
	fmt.Println("end of main()")

	// use cases
	// 1. error handling; 2. termination conditions; 3. exit codes

	// best practices
	// 1. avoid deferred actions; 2. status codes; 3. avoid abusive use
}
