package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	loopWorker()
}

func loopWorker() {
	i := 0
	ticker := time.NewTicker(7 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// 执行我们想要的操作
			i++
			//doxx(i)
			fmt.Println(i)
		}
	}
}

func doxx(i int) {
	time.Sleep(7 * time.Second) // --- B
	fmt.Println("aaa", i, time.Now().Unix())
}
