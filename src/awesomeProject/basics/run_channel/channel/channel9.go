package main

import (
	"time"
	"fmt"
)

func main() {
	conn := &Connetcion{}
	for i := 0; i < 22; i++ {
		ch := make(chan int, 1)
		ch <- i
		conn.action(ch, i)
	}
}

type Connetcion struct {}

func (conn *Connetcion) action(disconnectchan chan int, i int) {
	select {
	case index := <-disconnectchan:
		select {
		case <-time.After(time.Duration(i) * time.Second):
			fmt.Println(index)
		}
	}
}
