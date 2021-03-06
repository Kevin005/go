package main
//闭包函数
import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := getSequence()
	fmt.Println(nextNumber())
}
