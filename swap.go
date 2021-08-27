package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	swap(&a, &b)

	fmt.Println(a, b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
