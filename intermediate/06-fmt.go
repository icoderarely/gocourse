package intermediate

import "fmt"

// fmt package
// formatted input/ output functions, widely used for printing and formatting strings/numbers and other data types

// key functions
// 1. printing: Print(), Printf(), Println()
// 2. fromatting: Sprint(), Sprintf(), Sprintln()
// 3. scanning: Scan(), Scanf(), Scanln()
// 4. error: Errorf()

func main() {
	// formatting functions
	num := 45
	fmt.Println("----- fromatting function")
	// can also concatenate strings
	s := fmt.Sprint("Hello", "World!", 123, num) // only returns the resultant string, doesnt print anything
	fmt.Println(s)
	sln := fmt.Sprintln("Hello", "World!", 123, num) // adds space in between by default, and a new line at the end
	fmt.Print(sln)
	sf := fmt.Sprintf("Hello World %v", num)
	fmt.Printf("%v\n", sf)

	// scanning functions - through console
	fmt.Println("----- Scanning function")
	var name string
	var age int
	fmt.Println("enter name and age:")
	// fmt.Scan(&name, &age) // requires some data to be entered for it to proceed
	// fmt.Printf("name: %v of age: %v\n", name, age)

	// fmt.Scanln(&name, &age) // stops scanning at a new line
	// fmt.Printf("name: %v of age: %v\n", name, age)

	// fmt.Scanf("%s %d", &name, &age) // expects, there will be a string then a space and then an integer: follow the convention declared in the function
	fmt.Printf("name: %v of age: %v\n", name, age)

	// error function - returns string as a value
	fmt.Println("----- Error function")
	// if there are more than one input expected, you first handle error then move to the value part: value, err := func()
	age = 11
	err := checkAge(age)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("%d is less than 18", age)
	}
	return nil
}
