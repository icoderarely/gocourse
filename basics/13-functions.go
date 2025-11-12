package basics

import (
	"errors"
	"fmt"
)

func main() {
	// general function declaration
	// func <name>(parameter list) returnType {}
	// if more than one returnType then we can also use `()` to write multiple returnTypes

	// if return types are omitted, functionsr return default values for the return types specified

	// functions are first class objects/citizens

	// multiple return values by listing in parenthesis, with each return type having distinct type
	// func fnName(param1 type1, param2 type2...) (returnType1, returnType2) {
	// 	//codeblock
	// 	// return returnValue1, returnValue2...
	// }

	// fmt.Println(divide(24, 9))
	// q, r := divide(24, 9) -> also works to get those values and store in a variable before utilizing it or directly print

	// why?
	// 	error handling

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}

	// variadic functions: allows to accept a variable number of arguments
	// should be the last param
	// defined by elipsis ...
	// func fnName(param1 type1, param2 type2, param3 ...type3) returnType {}
	// type3 can have infinite number of params: 0 or thousand

	fmt.Println("Sum of 1 to 4 is:", sum(1, 2, 3, 4))

	// can unpack a slice and pass it as a variadic parameter
	nums := []int{1, 2, 4, 5}
	total := sum(nums...)
	fmt.Println(total)

	// defer: postpone the execution of a function, until surrounding function returns
	// returns normally or panics
	// always enclosed in another function
	// runs in last in first out manner
	process()

	// practical uses
	// 1. resource cleanup; 2. unlocaking mutexes; 3. loggin and tracing
	// best practices
	// 1. keep it short; 2. understand evalutation timing; 3. avoid complex control flow

	// panic: stops normal execution of flow immediately
	// panic(interface{}): any type of values means interface{}
	processPanic(0) // gives panic error + error stack
	// processPanic(-1)

	// panic only runs after defer functions are executed
	// err/nill > panic

	// recover: regain control of panicing go routine, only useful inside defer functions, manage behaviour of panic go routine to avoid abrupt termination
	fmt.Println("---------------")
	recoverProcess()
}

func recoverProcess() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("start process")
	panic("something went wrong")
	fmt.Println("end process")
}

func processPanic(input int) {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	if input < 0 {

		fmt.Println("before panic")
		panic("input < 0")
		fmt.Println("after panic") // anything after panic is unreachable, hence the error
	}
	fmt.Println("Processing input:", input)
}

func process() {
	defer fmt.Println("1. deferred statement executed")
	defer fmt.Println("2. deferred statement executed")
	defer fmt.Println("3. deferred statement executed")
	fmt.Println("normal execution statement")
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// another way to write this
// func divide(quotient int, remainder int) (int, int) {
// 	quotient = a / b
// 	remainder = a % b
// 	return
// }

func compare(a, b int) (string, error) {
	if a > b {
		return "a > b", nil
	} else if b > a {
		return "b > a", nil
	} else {
		return "", errors.New("unable to compare which is greator")
	}
}

// func something(a, b int, func someOp(int, int) int) func something(int, int) int {}
