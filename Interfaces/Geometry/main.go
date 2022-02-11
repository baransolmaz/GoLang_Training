package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}
type shape interface {
	getArea() float64
}

func main() {
	s := square{5}
	t := triangle{3, 3}
	printArea(s)
	printArea(t)
}
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func printArea(s shape) {
	fmt.Printf("Area of Shape: %.3f\n", s.getArea())
}
