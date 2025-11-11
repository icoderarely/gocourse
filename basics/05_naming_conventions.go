package basics

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase -> CalculateArea, UserInfo, NewHTTPRequest
	// used to name types on go: structs, interfaces, enums

	// snake_case -> user_id, first_name, https_request
	// commonly for variables, constants and files names

	// UPPERCASE
	// mostly for constants to stand out as they are immutable
	const MAXRETRIES = 5

	// mixedCase (camelCase) -> htmlDocument
	// variables
	employeeID := 1001

	// Use descriptive names and adhere to naming conventions (throughout the program maintain consistency)
	// avoid abbreviations
	// package names to be short, consise and lowercase
}
