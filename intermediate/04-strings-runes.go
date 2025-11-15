package intermediate

import "fmt"

// strings: sequence of bytes(unsinged int8 values), represents text, immutable
// array of chars(unicode), each alphabet is a rune

// runes: integer value that represnets a char
// rune is just an alias for int32, while byte represents a single 8-bit value (useful for ASCII or raw data), a rune represents a full Unicode character, which might need more than one byte.
// go encompasses chars from alot of languages across the globe - internationalization

// runes and chars
// similarities
// 1. represents char
// 2. storage size: occupy fixed
// differences
// 1. unicode support in go(4 byte) compare to char(1 byte in c - ascii)
// 2. type and size -> alias for int32, can represent wider range of text
// 3. encoding and decoding -> native support in go

func main() {
	msg := "Hello, Go!\t"
	rawMsg := `Hello\nGo!` // raw string literal, escape sequences dont work here
	fmt.Println(msg, rawMsg)

	// can use len()
	fmt.Println(len(msg), len(rawMsg))

	// can access char at a specific index returns byte value
	fmt.Println(msg[0]) // ascii value of 'H'

	// string concatenation
	greet := "hello,"
	name := "John"
	greetMsg := greet + " " + name
	fmt.Println(greetMsg)

	// comparisions is based on lexicography: ascii value comparisions
	// str1 < str2:dictionary based

	// string iteration
	for i, char := range greetMsg {
		fmt.Println("char:", string(char), "index:", i)
		// %d and all placeholders(c term), here in golang is called "format verbs"
	}

	// var ch rune = '🙂' // single quotes for runes, similar to char in other languages
	str := "🙂"
	fmt.Println([]byte(str), []rune(str))
}
