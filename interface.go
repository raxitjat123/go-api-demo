package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{radius: 6}
	rectangle := Rectangle{height: 6, width: 7}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))
}
