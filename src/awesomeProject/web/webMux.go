package main

/**
实现了多路副有
 */
import (
	"net/http"
	"io"
)

func helloHandler1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "xiaoshi")
}

func helloHandler2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is /")

}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/xiaoshi", helloHandler1)
	mux.HandleFunc("/", helloHandler2)
	http.ListenAndServe(":8080", mux)
}
