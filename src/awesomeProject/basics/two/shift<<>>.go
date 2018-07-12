package main

import "fmt"

func main() {
	x := 2
	y := 4
	//因为2代表二进制10，向左移动一位为100，表示01>>10>>11>>100 表示4
	//4同理
	fmt.Println(x << 1)//output:4
	fmt.Println(y >> 1)//output:2

}
