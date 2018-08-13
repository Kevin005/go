package main

import (
	"fmt"
	"time"
)

func main() {
	waitc := make(chan struct{})
	a := make(chan bool, 1)
	go func() {
		time.Sleep(5 * time.Second)
		close(waitc)
	}()
	go func() {
		<-waitc
		a <- true
	}()
	count := 0
	bool := false
	for {
		select {
		case <-a:
			bool = true
			fmt.Println("this is end")
		case <-time.After(1 * time.Second):
			count ++
			fmt.Println(count)
		}
		if bool {
			break
		}
	}
}
