package main

import (
	"fmt"
)

type Book struct {
	ISBN  string
	Name  string
	Price int
	Author
}

type Author struct {
	FirstName string
	LastName  string
}

func main() {

	book1 := Book{
		ISBN:   "ISBN142563",
		Name:   "book1",
		Price:  250,
		Author: Author{FirstName: "Author-1 FirstName", LastName: "Author-1 LastName"},
	}

	book2 := Book{
		ISBN:   "ISBN142563",
		Name:   "book-2",
		Price:  450,
		Author: Author{FirstName: "Author-2 FirstName", LastName: "Author-2 LastName"},
	}
	fmt.Println(book1, book2)

}
