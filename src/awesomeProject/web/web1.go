package main

import (
	"net/http"
	"io"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello acd")
}

func main() {

	// 通过 HandlerFunc 把函数转换成 Handler 接口的实现对象
	hh := http.HandlerFunc(helloHandler)
	http.Handle("/xiaoshi", hh)
	http.ListenAndServe(":8080", nil)
}
