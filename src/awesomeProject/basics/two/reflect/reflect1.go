package main

import (
	"reflect"
	"log"
)

var i int = 1

func main(){
	type S struct {
		A string `json:"tag_a"`
	}
	s := S{}

	value := reflect.ValueOf(&s)

	value = reflect.Indirect(value)

	//获取结构体s的类型S
	vt := value.Type()

	//获取S中的A成员变量
	f, _ := vt.FieldByName("A")

	//获取成员变量A的db标签
	log.Println(f.Tag.Get("json"))
}
