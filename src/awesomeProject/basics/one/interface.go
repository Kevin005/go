package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("how are you,this is nokia")
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("how are you,this is iphone")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call();
	phone = new(IPhone)
	phone.call()
}
