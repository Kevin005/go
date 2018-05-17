package main

//程序中使用的是值传递, 所以两个值并没有实现交互，我们可以使用 引用传递 来实现交换效果
import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	/* 通过调用函数来交换值 */
	aaa := swap(a, b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
	fmt.Printf("交换后 b 的值 : %d\n", aaa)
}

/* 定义相互交换值的函数 */
func swap(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp;
}
