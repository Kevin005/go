package main

var (
	ch    = make(chan int, 20) //可以同时运行的routine数量为20
	count = 500
)

func main() {
	for i := 0; i < count; i++ {
		go fFunction(ch)
	}
	//知道了任务总量，可以像这样利用固定循环次数的循环检测所有的routine是否工作完毕
	for i := 0; i < count; i++ {
		<-ch
	}
}
func fFunction(ch chan int) {
	ch <- 1
}
