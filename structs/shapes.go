package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, length float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.length
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r1 := rectangle{5, 5}
	c1 := circle{10}

	fmt.Println(r1, "area", r1.area())
	fmt.Println(c1, "area", c1.area())

	measure(r1)
	measure(c1)
}
