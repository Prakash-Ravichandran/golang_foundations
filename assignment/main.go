package main

import "fmt"

type shape interface {
	printArea() 
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

func (s Square) printArea()  {
	fmt.Println("Area:", s.getArea())
}

func (t triangle) printArea()  {
	fmt.Println("Area:", t.getArea())
}

func main() {
     sq:= Square{sideLength: 5.3}
	 tr:= triangle{base: 7, height: 5}
	 sq.printArea()
	 tr.printArea()
}
