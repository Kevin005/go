package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         //这时候操作p会更改i的内存 point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         //此时&i的值已经改变 set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         //这时候操作p会直接更改j的内存  point to j
	*p = *p / 37   //此时&j的值已经改变 divide j through the pointer
	//fmt.Println(*p)
	fmt.Println(j) // see the new value of j
}

/**
42
21
73
 */