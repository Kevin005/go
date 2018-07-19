package main

import (
	"fmt"
	"unsafe"
	"text/template"
)

type NodeInfo struct {
	Domain     string
	Nodemode   string
	IsValidity bool
}

type MyNodeInfo struct {
	Domain     string
	Nodemode   string
	IsValidity bool
}

func main() {
	nodeInfo := &NodeInfo{
		"123456",
		"123456",
		true,
	}
	//NodeInfo和MyNodeInfo之间进行强制类型转换
	myNodeInfo := &MyNodeInfo{
		"1",
		"1",
		false,
	}

	//nodeInfos := []*NodeInfo{
	//	{
	//		"123456",
	//		"123456",
	//		true,
	//	},
	//	{
	//		"123456",
	//		"123456",
	//		true,
	//	},
	//}
	myNodeInfo = (*MyNodeInfo)(unsafe.Pointer(nodeInfo))
	fmt.Println(myNodeInfo)    //说明已经从NodeInfo转为MyNodeInfo类型了
	myNodeInfo = &MyNodeInfo{} //说明已经从NodeInfo转为MyNodeInfo类型了
	fmt.Println(myNodeInfo)
}

//补充：实际使用中unsafe可用场景很少，稍微复杂一点的结构，比如struct,unsafe根本不能适用，正确的方法还是要靠 type assertion？？？
// MyTemplate 定义和 template.Template 只是形似
type MyTemplate struct {
	name       string
	parseTree  *unsafe.Pointer
	common     *unsafe.Pointer
	leftDelim  string
	rightDelim string
}

func main1() {
	t := template.New("Foo")
	p := (*MyTemplate)(unsafe.Pointer(t))
	p.name = "Bar" // 关键在这里，突破私有成员
	fmt.Println(p, t)
}
