package _struct

import (
	"fmt"
	"log"
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

type StudentBoos struct {
	Books
	math_books    string
	english_books string
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "go language"
	Book2.title = "go language 2"
	printBook(Book1)

	StudentBoos := StudentBoos{
		Books{
			"",
			"",
			"",
			0,
		},
		"",
		"",
	}
	fmt.Println(StudentBoos.title)
	StudentBoos.Books = Books{}

	a := struct{}{}
	b := struct{}{}
	log.Println(a == b)
	log.Printf("%p,%p\n", &a, &b) //打印的指针地址竟然一样!
}

func printBook(book Books) {
	fmt.Println(book.title)
	fmt.Println(book.book_id)
	fmt.Println(book.author)
}
