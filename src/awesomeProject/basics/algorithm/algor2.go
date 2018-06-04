package main

/**
选择排序
1.用[0]位和后边所有位进行比较，如果小于[0]位的就和第[0]位替换数值，第一轮就会确保[0]位是最小值，
2.下来++判断第[1]位为第二最小值
 */
import "fmt"

func main() {
	arr := []int{2, 67, 33, 0, 45, 25, 77, 208, -8, -7}
	selectSort(arr)
	fmt.Println(arr)
}

func selectSort(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		var minInt int = i
		for j := i + 1; j < len; j++ {
			if arr[minInt] > arr[j] {
				arr[minInt],arr[j] = arr[j],arr[minInt]
			}
		}
	}
}
