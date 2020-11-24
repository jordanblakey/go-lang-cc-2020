package main

import (
	"fmt"
	"math"
)

// Interface defines a set of methods for a struct

// Shape interface definition
type Shape interface {
	area() float64
}

// Circle struct definition
type Circle struct {
	x, y, radius float64
}

// Rectangle struct definition
type Rectangle struct {
	w, h float64
}

func getArea(s Shape) float64 {
	return s.area()
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}

func main() {
	rect1 := Rectangle{w: 10, h: 20}
	circ1 := Circle{5, 10, 20}

	fmt.Println(rect1, circ1)
	fmt.Printf("Rect1 - w: %f h: %f\n", rect1.w, rect1.h)
	fmt.Printf("Circ1 - x: %f, y: %f, radius: %f\n", circ1.x, circ1.y, circ1.radius)

	fmt.Printf("Circle Area: %f\n", getArea(circ1))
	fmt.Printf("Rectangle Area: %f\n", getArea(rect1))
}
