package main

import "fmt"

//https://studygolang.com/articles/2192 常量计数器

const a = iota
const (
	b = iota
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
