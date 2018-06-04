package main

/**
冒泡排序
1.最后一位和前一位比较，把小的值给前一位，这样先循环9次，此时第[0]位为最小值
2.进行第二轮，循环到第[1]位停止，此时第[1]位为第二小的值
3.i循环9次即可，每循环一次确定一位和后边值相比的最小值
 */
import (
	"fmt"
)

func bubbleSort(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := len - 1; j > i; j-- {
			if arr[j] < arr[j-1] {//最后以为和前一位比较，把小的值
				arr[j], arr[j-1] = arr[j-1], arr[j]//把小的值给前一位
			}
		}
	}
}

func main() {
	arr := []int{2, 78, 3, 6, 90, 33, 44, 0, -5, 99}
	bubbleSort(arr)
	fmt.Print(arr)
}
