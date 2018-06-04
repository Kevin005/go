package main

import "fmt"

var bb = 100//声明 并负值，可以在还输外部
var cc int = 100//声明 并负值，可以在还输外部

func main() {
	//= 是赋值， := 是声明变量并赋值并且只能在函数内部
	var a int//先声明
	a  =100 //负值
	fmt.Println(a)

	var b = 100//声明 并负值，可以在还输外部
	fmt.Println(b)
	var c int = 100//声明 并负值，可以在还输外部
	fmt.Println(c)
	d := 100 //简写：声明 并负值，并且只能在函数内部
	fmt.Println(d)
}
