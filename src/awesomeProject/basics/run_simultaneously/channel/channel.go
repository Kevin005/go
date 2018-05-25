package channel

import "fmt"

func end() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1 ////这一行操作就会发生阻塞，因为前三行的放入数据的操作已经把channel填满了
}

func receive() {
	ch := make(chan int, 3)
	<-ch ////这一行会发生阻塞，因为channel才刚创建，是空的，没有东西可以取出
}

func channelType() {
	ch := make(chan int)     //无缓冲的channel，同等于make(chan int, 0)
	ch1 := make(chan int, 5) //一个缓冲区大小为5的channel
	fmt.Println(ch)
	fmt.Println(ch1)
}

func main() {
	receive()
}
