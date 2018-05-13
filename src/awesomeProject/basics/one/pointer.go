package main

import "fmt"

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

func main() {
	var ptr *int
	fmt.Println("ptr的值为：%x\n", ptr)
}
