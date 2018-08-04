package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool, 1)
	go abc(ch1)
	select {
	case <-ch1:
		fmt.Println("is success")
	case <-time.After(5 * time.Second):
		fmt.Println("time out")
	}
}

func abc(ch1 chan bool) {
	time.Sleep(10 * time.Second)
	ch1 <- true
}
