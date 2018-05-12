package main

import "net/http"

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.Handle("/", &helloHandler{})
	http.ListenAndServe(":9090", nil)
}
