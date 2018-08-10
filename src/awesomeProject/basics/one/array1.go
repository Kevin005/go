package main

import "fmt"

func main() {

	bytes := [][]byte{[]byte("1"), []byte("2")}
	sum([]byte("3"), bytes...)
}

func sum(D []byte, args ...[]byte) {
	fmt.Println(D)
	fmt.Println(args, " ")
}
