package intermediate

import "fmt"

// interfaces
type shape interface {
	area() int
}

type square struct {
	l int
}

func (s square) area() int {
	return s.l * s.l
}

func main() {
	s := square{l: 5}
	squareArea := s.area()
	fmt.Println(squareArea)
}
