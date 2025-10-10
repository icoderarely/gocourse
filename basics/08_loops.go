package basics

import "fmt"

func main() {
	// syntax: for initialization; condition; post {...code}

	// on range
	// for i := 1; i <= 5; i++ {
	// fmt.Println(i)
	// }

	// iterate over collection
	// numbers := []int{1, 2, 3, 4, 5}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %v, Value: %v\n", index, value)
	// }

	// range
	// for i := range 10 {
	// 	fmt.Println(i + 1)
	// }

	// using for loop as while loop
	i := 1
	for i <= 5 {
		if i == 2 {
			fmt.Println("two")
			i++
			continue
		}
		fmt.Println(i)
		i++
	}
}
