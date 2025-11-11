package main

// every go file belongs to a package and main is a special package which is executable and not a shared library
// when we run the program go looks for main package and executes main() first
// Package main = entry point of a Go executable.

import "fmt"

func main() {
	// the starting point of program execution
	// main() runs automatically when you execute the program

	fmt.Println("Hello World!")
	// fmt is a format package -> provides functions to formatted I/O similar to printf and scanf
	// we can use any language in unicode system
}

// Printing and Formatting: fmt offers various functions for printing and formatting output to the console or other output streams:
// 	- fmt.Println(): Prints a line to the console, adding a newline character at the end.
//   - fmt.Print(): Prints to the console without adding a newline.
//   - fmt.Printf(): Prints formatted output using format specifiers (e.g., %v for default value, %s for string, %d for integer).
//   - fmt.Sprintf(): Returns a formatted string without printing it.

// Scanning and Input: fmt also provides functions for reading input:
//   - fmt.Scanln(): Reads a line of input from the console.
//   - fmt.Scanf(): Reads formatted input using format specifiers.

// Error Handling: Functions like fmt.Errorf() are used to create formatted error messages.
