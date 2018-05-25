package channel

import "fmt"

func gfunction(routineControl chan int, feedback chan string) {
	defer func() {
		<-routineControl
		feedback <- "finish"
	}()

	// do some process
	// ...
}

/**
如果任务的数量不固定
 */
func main() {
	var (
		routineCtl = make(chan int, 20)
		feedback   = make(chan string, 10000)

		msg      string
		allwork  int
		finished int
	)
	for i := 0; i < 100000000; i++ {
		routineCtl <- 1
		allwork++
		go gfunction(routineCtl, feedback)
	}

	for {
		msg = <-feedback
		if msg == "finish" {
			finished++
		}
		if finished == allwork {
			fmt.Println(allwork)
			break
		}
	}
}
