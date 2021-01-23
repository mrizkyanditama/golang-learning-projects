package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func printArea(s shape) float64 {
	return s.getArea()
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func main() {
	tr := triangle{height: 10, base: 80}
	sq := square{sideLength: 40}

	fmt.Println(printArea(tr))
	fmt.Println(printArea(sq))
}
