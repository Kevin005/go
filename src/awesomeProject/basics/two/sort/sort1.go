package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4,3,2,1,5,9,8,7,6}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("After reversed: ", a)
}