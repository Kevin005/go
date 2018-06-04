package main

import "fmt"

type MyMap map[string]string

func (s  MyMap) Set(name,value string){
	s[name] = value
}
func (s MyMap) Get(name string)string{
	return s[name]
}

func main() {
	aa := MyMap{}
	aa.Set("key","value")
	fmt.Println(aa.Get("key"))
}
