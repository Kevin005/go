package main

import "fmt"

type in interface {
	String()
}

type A struct {
}

func (a *A) String() {
	fmt.Println("this is A")
}

func (a *A) StringA() {
	fmt.Println("this is AA")
}

type B struct {
	str string
}

func (b *B) String() {
	fmt.Println("this is b")
}

func (b *B) stringB() {
	fmt.Println("this is bb")
}

func main() {
	var i in
	i = &A{}
	switch t := i.(type) {
	case *A:
		t.String()
		t.StringA()
	case *B:
		t.String()
		t.stringB()
	}
}
