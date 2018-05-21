package main

import "fmt"

type S struct {
	i int
}

type I interface {
	Get() int
	Put(int)
}

func (p *S) Get() int {
	return p.i
}

func (p *S) Put(v int) {
	p.i = v
}

func f1(p I) {
	fmt.Println(p.Get())
	p.Put(888)
}

func f2(p interface{}) {
	switch t := p.(type) {
	case int:
		fmt.Println("this is number")
	case I:
		fmt.Println("this is I", t.Get())
	default:
		fmt.Println("unknow type")
	}
}

func add(a *S) {
	a.Put(999)
	fmt.Println(a.Get(), "in add func")
}

func addA(a S) {
	a.Put(2222)
	fmt.Println(a.Get(), "in add funcA")
}

func main() {
	var s S
	s.Put(333)
	fmt.Println(s.Get()) //333
	f1(&s)
	fmt.Println(s.Get()) //888
	f2(&s)               //this is I
	add(&s)              //999
	fmt.Println(s.Get()) //999
	addA(s)              //2222
	fmt.Println(s.Get()) //999
}
