package main

/**
实现了静态服务器
 */
import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
