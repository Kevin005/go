package main

//func cfunction(ch chan int) {
//	fmt.Println("finish")
//	<-ch //goroutine执行完了就从channel取出一个数据
//}
//
//func main() {
//	ch := make(chan int, 10)
//	for i := 0; i < 1000; i++ {
//		//每当创建goroutine的时候就向channel中放入一个数据，如果里面已经有10个数据了，就会
//		//阻塞，由此我们将同时运行的goroutine的总数控制在<=10个的范围内
//		ch <- 1
//		go cfunction(ch)
//	}
//	// 这里只是示范个例子，当然，接下来应该有些更加周密的同步操作
//}

func dfunction(ch chan int) {
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	<-ch
}

func main() {
	//主routine的操作同上面那段代码
	ch := make(chan int, 5)
	for i := 0; i < 100; i++ {
		ch <- 1
		go dfunction(ch)
	}
	// 这段代码运行的结果为死锁,其实总结起来就一句话，"放得太快，取得太慢了
	//在使用带缓冲的channel时一定要注意放入与取出的速率问题
}
