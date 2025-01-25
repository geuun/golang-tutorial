package _interfaces

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect1 struct {
	width, height float64
}
type circle1 struct {
	radius float64
}

func (r rect1) area() float64 {
	return r.width * r.height
}
func (r rect1) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle1) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle1) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle1); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func Main() {
	r := rect1{width: 3, height: 4}
	c := circle1{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
