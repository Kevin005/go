package main

import (
	"fmt"
	"time"
)

func main() {
	quitCh := make(chan bool, 1)

	go func() {
		for {
			select {
			case a, b := <-quitCh:
				fmt.Println(a, b)
			}
		}
	}()
	quitCh <- false
	time.Sleep(5 * time.Second)

}
