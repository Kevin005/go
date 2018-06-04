package main

/**
对象初始化的三种方式
 */
import "fmt"

type TransINfo struct {

}

type Fragment interface {
	Exec(transInfo *TransINfo) error
}

type GetPodAction struct {

}

func (g GetPodAction) Exec(transInfo *TransINfo) error {
	fmt.Println("this is exec")
	return nil
}

func main(){
	//初始化方式1
	var fragment Fragment = new(GetPodAction)
	fragment.Exec(&TransINfo{})
	//初始化方式2
	var fragment1 Fragment = &GetPodAction{}
	fragment1.Exec(&TransINfo{})
	//初始化方式3
	var fragment2 Fragment = GetPodAction{}
	fragment2.Exec(&TransINfo{})
}
