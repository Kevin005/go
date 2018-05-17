package main

/**
实现了静态服务器
http://localhost:8080/basics/
 */
import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("/"))) //当前硬盘目录
	//http.ListenAndServe(":8080", http.FileServer(http.Dir("."))) //当前项目目录
}
