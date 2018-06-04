package main

import "fmt"

type Student struct{
	Name string
}

func (s * Student) Set(name string){
	s.Name = name
}

func (s Student) Sett(name string){
	s.Name = name
	fmt.Println(s.Name)
}

func (s Student) Get()string{
	return s.Name
}

func main(){
	a := &Student{}
	a.Set("this is & name")
	fmt.Println(a.Get())

	b := Student{}
	b.Sett("this is ")
	fmt.Println(b.Get())

}