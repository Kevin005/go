package main
//http://www.infoq.com/cn/news/2017/08/go-1-9-Type-Alias
import "fmt"

type MyInt1 int
type MyInt2 = int

func main() {
	var i int = 0
	//var i1 MyInt1 = i//编译报错，不能把int转为MyInt1类型
	var i2 MyInt2 = i
	fmt.Println(i2)
}
