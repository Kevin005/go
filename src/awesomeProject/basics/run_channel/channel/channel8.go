package main

import (
	"fmt"
	sync2 "sync"
	"time"
)

func main() {
	ch := make(chan bool, 100000)
	sync := new(sync2.WaitGroup)
	var ch1 = make(chan bool, 1)
	go a(ch, sync, ch1)
	//go b(ch, sync)
	//go c(ch, sync)

	go waitBreak(ch1)
	for i := 0; i < 10; i++ {
		sync.Add(1)
		ch <- true
		time.Sleep(2 * time.Second)
	}

	sync.Wait()
	fmt.Println("countA", countA)
	fmt.Println("countB", countB)
	fmt.Println("countC", countC)
}

var countA = 0

func waitBreak(ch1 chan bool) {
	time.Sleep(5 * time.Second)
	ch1 <- true
}

func a(ch chan bool, sync *sync2.WaitGroup, ch1 chan bool) {
	for {
		select {
		case <-ch1:
			break
		case _, b := <-ch:
			countA ++
			sync.Done()
			fmt.Println("is a", b)
		}
		break
	}
}

var countB = 0

func b(ch chan bool, sync *sync2.WaitGroup) {
	for {
		select {
		case _, b := <-ch:
			countB ++
			sync.Done()
			fmt.Println("is b", b)
		}
	}
}

var countC = 0

func c(ch chan bool, sync *sync2.WaitGroup) {
	for {
		select {
		case _, b := <-ch:
			countC ++
			sync.Done()
			fmt.Println("is b", b)
		}
	}
}
