package basics

func main() {
	// with `var` -> var variablename type = value
	// if we not define a type, it will infer from the value but then value should be must
	// can be used for global variable declaration
	var name string = "Name"

	// using the sign `:=` -> has to have a value; type is inferred from the value
	// this can only be used within functions
	name1 := "Name1"

	// variables are block scoped
}
