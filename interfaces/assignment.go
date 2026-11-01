package interfaces

import "fmt"

type Shape interface {
	GetArea() float64
}

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

func (t Triangle) GetArea() float64 {
	return 0.5 * t.base * t.height
}

func (s Square) GetArea() float64 {
	return s.sideLength * s.sideLength
}

func PrintArea(sh Shape) {
	a := sh.GetArea()
	fmt.Printf("Area of the given shape is: %f", a)
}

func Example() {
	t := Triangle{base: 10, height: 10}
	s := Square{sideLength: 10}

	PrintArea(s)
	PrintArea(t)
}