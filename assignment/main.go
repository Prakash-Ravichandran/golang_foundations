package main

import "fmt"

type shape interface {
	getArea() float64 
}

type Square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s triangle) getArea() float64 {
	return 0.5 * s.base * s.height
}


func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}


func main() {
     sq:= Square{sideLength: 5.3}
	 tr:= triangle{base: 7, height: 5}

	 printArea(sq)
	 printArea(tr)
}
