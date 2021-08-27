package main

import (
	"fmt"
)

func main() {

	i := []int{1, 2, 3, 4, 5}
	j := i
	fmt.Println("before i value ", i)
	fmt.Println("before j value ", j)
	j[2] = 6
	fmt.Println("After i value ", i)
	fmt.Println("After j value ", j)

	// output
	// i := [...]int{1, 2, 3, 4, 5}
	// use copy of array

	// 	before i value  [1 2 3 4 5]
	// before j value  [1 2 3 4 5]
	// After i value  [1 2 3 4 5]
	// After j value  [1 2 6 4 5]

	// output
	// i := []int{1, 2, 3, 4, 5}
	// use address of array

	// 	before i value  [1 2 3 4 5]
	// before j value  [1 2 3 4 5]
	// After i value  [1 2 6 4 5]
	// After j value  [1 2 6 4 5]
}
