package basics

func main() {
	age := 23

	// if else
	// if age >= 18 {
	// 	fmt.Println("Adult")
	// } else if age < 14 {
	// 	fmt.Println("Universal")
	// } else {
	// 	fmt.Println("Universal Adult")
	// }

	// and -> &&; or -> ||

	// switch statement -> in go we dont need to use break statement at the end on cases
	// switch expression {
	// case value1:
	// 	// code to be executed
	// case valuen: ...
	// default: // if no case matches
	// }

	// to make the switch behave as default like in other languages and pass to the next case, we can use `fallthrough;`

	// 1. multiple conditions -> case "mon", "tue", "wed": ...code
	// 2. case with expression ->
	// 	a. boolean: switch { case num < 10: ...code }
	// 	b. expression: switch num { case 1: ...cdoe }
	// 3. switch with expression -> switch num / 2 {...code}
}
