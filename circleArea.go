package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func main() {
	c := Circle{x: 0, y: 0, radius: 5}
	fmt.Println("Circle area is %f", c.area())
}
