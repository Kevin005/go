package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan string)
	c := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		c <- "this is c"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		in <- "this is in"
	}()
	timeout := time.After(90 * time.Millisecond)
	select {
	case n := <- in:
		fmt.Println("received", n)
	case out := <- c:
		fmt.Println("sent", out)
	case <-timeout:
		fmt.Println("this is timeout")
	}
	fmt.Println("this is end")
}
