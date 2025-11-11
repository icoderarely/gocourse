package basics

import (
	"fmt"
	"slices"
)

func main() {
	// similar to arays, but dont have sizes
	// var sliceName []type

	// declarations
	// var num []int
	// var num []int{1, 2, 3}
	// num := []int{1, 2, 3}

	// slice with a size
	// num := make([]int, 5)

	// slicing an array
	arr := [6]int{1, 2, 3, 4, 5}
	s := arr[1:4]
	s = append(s, 4, 7, 6)
	// fmt.Println(s)

	sCp := make([]int, len(s))
	copy(sCp, s)
	sCp[2] = 5
	fmt.Println("s1:", s, "s2:", sCp)

	if !slices.Equal(s, sCp) {
		fmt.Println("not equal slices")
	}

	// more used than arrays
	// underlying structure is arrays -> Every slice in Go internally points to an array in memory.
}
