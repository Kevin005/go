package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "go language"
	Book2.title = "go language 2"
	printBook(Book1)
}

func printBook(book Books) {
	fmt.Println(book.title)
	fmt.Println(book.book_id)
	fmt.Println(book.author)
}
