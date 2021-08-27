package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Your age: ")
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	if age >= 19 {
		fmt.Println("You can ride alone!")
	} else if age >= 12 {
		fmt.Println("Must be ride with parents!!!!!!")
	} else {
		fmt.Println("You are child, Not allowed.")
	}
}
