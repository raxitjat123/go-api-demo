package main

import (
	"fmt"
)

const Max int = 3

func main() {

	a := []int{3, 4, 5}
	var ptr [Max]*int

	for i := 0; i < Max; i++ {
		ptr[i] = &a[i]       //store address of array
		fmt.Println(*ptr[i]) // to get value write * which means value of that address
	}
}
