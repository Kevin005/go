package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	close(ch) //如果执行了close()就立即关闭channel的话，下面的循环就不会有任何输出了
	for {
		data, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(data)
	}
}
