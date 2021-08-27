package main

import "fmt"

func fibonacci(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fibonacci(i-1) + fibonacci(i-2)
	}
}

func main() {

	num := 10
	for i := 0; i <= num; i++ {
		fmt.Println(fibonacci(i))
	}
}
