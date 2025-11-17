package intermediate

import "fmt"

// structs: defined using "type" and "struct" keyword, followed by "{}" containing a list of fields
// fields are defined by name and type
// composite data types, allows to group together different types of data in a single name, similar to classes, but doesnt have inheritance

type Person struct {
	fName string
	lName string
	age   int

	// embed on struct in another
	address Address

	// anonymous struct
	PhoneHomeCell
}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func (p Person) fullName() string {
	return p.fName + " " + p.lName
}

// modifying using pointers
func (p *Person) incrementAgeByOne() {
	p.age++
}

func main() {
	// initializations
	// if un-initialized, will take a default value of that type
	// struct literal
	p1 := Person{
		fName: "John",
		lName: "Doe",
		age:   23,
		address: Address{
			city:    "London",
			country: "UK",
		},
		// can access these directly without nesting into phoenePhoneHomeCell: PhoneHomeCell,
		PhoneHomeCell: PhoneHomeCell{
			home: "434535",
			cell: "432556",
		},
	}

	fmt.Println(p1.cell, p1.PhoneHomeCell, p1.PhoneHomeCell.home) // it does work, but there is a hint to remove the nesting

	p2 := Person{
		fName: "George",
	}

	p2.address.city = "New York"
	p2.address.country = "USA"

	// accessing fields of a struct -> using dot notation
	fmt.Printf("%v %v\n", p1.lName, p2.lName)

	// anonymous structs
	fmt.Println("----- anonymous functions")
	user := struct {
		username string
		password string
	}{
		username: "user",
		password: "pass",
	}
	fmt.Println(user.username, user.password)

	// methods
	fmt.Println("----- methods - value recievers")
	// structs are usually defined outisde of main function
	// in go you cannot define struct and use methods in same function
	fmt.Println(p1.fullName())

	fmt.Println("----- pointer recievers")
	// pointer recievers to modify struct fields within a method
	fmt.Println(p1)
	p1.incrementAgeByOne()
	fmt.Println(p1)
}
