package main

import "fmt"

type shape interface {
	area() float64
	peri() float64
}

type Circle struct {
	r float64
}
type Rectangle struct {
	l, w float64
}

func (c Circle) area() float64 {
	return 3.14 * c.r * c.r
}

func (c Circle) peri() float64 {
	return 2 * 3.14 * c.r
}

func (r Rectangle) area() float64 {
	return r.l * r.w
}
func (r Rectangle) peri() float64 {
	return 2 * (r.l + r.w)
}
func main() {
	c := Circle{r: 5}
	r := Rectangle{l: 3, w: 5}

	var sh shape

	// Circle calculations
	sh = c
	fmt.Println("Circle Area:", sh.area())
	fmt.Println("Circle Perimeter:", sh.peri())

	// Rectangle calculations
	sh = r
	fmt.Println("Rectangle Area:", sh.area())
	fmt.Println("Rectangle Perimeter:", sh.peri())

}
