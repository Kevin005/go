package basics

import "fmt"

type Animal interface {
	call()
}

type Dog struct {
}

type Rubby struct {
}

func (dog Dog) call() {
	fmt.Println("this is dog")
}

func (rubby Rubby) call() {
	fmt.Println("this is rubby")
}

func main2() {
	var animal Animal
	animal = new(Dog)
	animal.call()
	animal = new(Rubby)
	animal.call()
}
