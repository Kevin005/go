package main

//https://www.cnblogs.com/jasonxuli/p/6802289.html
//https://gocn.io/m/question/1519 *和&符号有什么区别

import "fmt"

type demo struct {
	Value string
}

//&a 表示a的存储地址
func main1() {
	var a int = 0
	var ip *int      //声明指针变量ip
	ip = &a          //ip指针,指向a的存储地址
	fmt.Println(&a)  //a的存储地址
	fmt.Println(ip)  //ip指针指向a的存储地址
	fmt.Println(*ip) //访问ip指针的值
	fmt.Println(a)
}

func main2() {
	var ptr *int
	fmt.Println("ptr的值为：%x\n", ptr)
}

func main() {
	var pt *demo             //*用作指针类型声明
	pt = &demo{Value: "val"} //&取值地址
	fmt.Println(pt)
	d1 := *pt //*用作取地址指向的实际值
	fmt.Println(d1)
}
