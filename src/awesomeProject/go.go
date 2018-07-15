package main

import "fmt"

func main(){
	slice := make([]int, 3, 5)
	fmt.Println("before:", slice)
	changeSliceMember(slice)
	fmt.Println("after:", slice)
}

func changeSliceMember(slice []int) {
	if len(slice) > 1 {
		slice[0] = 9
	}
	aaa := append(slice,3,4)
	aaa[4] = 5
	fmt.Println(aaa)
}