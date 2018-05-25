package channel

import "fmt"

func main_aFunction() {
	ch := make(chan int)
	go aFunction(ch)
	//ch <- 1
}

func aFunction(ch chan int) {
	fmt.Println("finish")
	<-ch
}

func main_bFunction() {
	ch := make(chan int)
	ch <- 1
	go bFunction(ch)
}
func bFunction(ch chan int) {
	fmt.Println("finish")
	<-ch
}

func main() {
	main_bFunction()
}
