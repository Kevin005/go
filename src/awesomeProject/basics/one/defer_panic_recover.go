package main
/**
http://www.cnblogs.com/liang1101/p/6842230.html
 */
import (
	"fmt"
	"log"
	"strconv"
)

/**
会倒叙执行
 */
func defer1() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	defer fmt.Println("d")
}

//捕获因未知输入导致的程序异常
func catch(nums ...int) int {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()
	return nums[1] * nums[2] * nums[3]
}

//主动抛出 panic，不推荐使用，可能会导致性能问题
func toFloat64(num string) (float64, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[w]", r)
		}
	}()
	if num == "" {
		panic("Param is null")
	}
	return strconv.ParseFloat(num, 10)
}

func main() {
	catch(1, 2)
	toFloat64("")
}
