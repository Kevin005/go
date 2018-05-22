package main

import "fmt"

func main() {
	//空数组就是一个切片
	balance := []int{1, 2, 3}
	balance = append(balance, 1, 2, 3, 4)
	fmt.Println(balance)

	var balance1 = []int{}
	balance1 = append(balance1, 1, 2)
	fmt.Println(balance1)
}
