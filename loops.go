package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter number :")
	scanner.Scan()
	nubmer, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for x := 1; x <= int(nubmer); x++ {
		fmt.Println(x)
	}

}
