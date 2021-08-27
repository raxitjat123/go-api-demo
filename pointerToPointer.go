package main

import (
	"fmt"
)

func main() {
	var a int
	var ptr *int
	var pptr **int
	a = 3000
	ptr = &a
	pptr = &ptr

	fmt.Println("value of a= %d, ptr %d, pptr %d", a, *ptr, **pptr)
}
